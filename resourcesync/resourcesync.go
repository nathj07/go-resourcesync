package resourcesync

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/nathj07/go-resourcesync/fetcher"
)

// resource type constant to define the type of data being processed
const (
	// Unknown indicates we cannot determine the type of feed we are processing
	Unknown = iota
	// List indicates a ResourceList
	List
	// Index indicates a ResourceListIndex
	Index
)

// ErrUnsupportedFeedType is used when the feed type is not one of the supported set
var ErrUnsupportedFeedType = errors.New("unsupported feed type supplied")

// ResourceSync is the top level structure needed to interact with ResourceSync endpoints
type ResourceSync struct {
	Fetcher   fetcher.RSFetcher
	indexChan chan IndexDef
	listChan  chan ResourceURL
}

// ResourceData is the structure for holding the data returned from a ResoureceSync fetch.
// Only one of rl or rli will be populated, the rType value will indicate which
type ResourceData struct {
	RL    *ResourceList
	RLI   *ResourceListIndex
	RType int // based on the type const above
}

// New is the simplest way to instantiate a ready to use ResourceSync object
func New(f fetcher.RSFetcher) *ResourceSync {
	return &ResourceSync{
		Fetcher:   f,
		indexChan: make(chan IndexDef),
		listChan:  make(chan ResourceURL),
	}
}

// Process takes the given target and fetches that page, pasring the data
// The returned structure indicate the type and then the relevant data can be found.
func (rs *ResourceSync) Process(baseTarget string) (*ResourceData, error) {
	data, status, err := rs.Fetcher.Fetch(baseTarget)
	if err != nil {
		return nil, fmt.Errorf("%d: %v", status, err)
	}
	rd, err := rs.Parse(data)
	if err != nil {
		return nil, fmt.Errorf("Parse failed: %v", err)
	}
	return rd, nil
}

// Parse handles the unmarshaling of the feed data.
// The returned ResourceData will have one field populated and the RType value will indicate which.
func (rs *ResourceSync) Parse(feed []byte) (*ResourceData, error) {
	feedType := rs.determineType(feed)
	rd := &ResourceData{
		RL:  &ResourceList{},
		RLI: &ResourceListIndex{},
	}
	switch feedType {
	case Index:
		if err := xml.Unmarshal(feed, rd.RLI); err != nil {
			return nil, err
		}
		rd.RType = Index
	case List:
		if err := xml.Unmarshal(feed, rd.RL); err != nil {
			return nil, err
		}
		rd.RType = List
	default:
		return nil, ErrUnsupportedFeedType
	}
	return rd, nil
}

func (rs *ResourceSync) determineType(data []byte) int {
	if bytes.Contains(data, []byte("<sitemapindex")) {
		return Index
	}
	if bytes.Contains(data, []byte("<urlset")) {
		return List
	}
	return Unknown
}

// TODO:
// Function to do it all in one go; a `Process` function that does fetching and then parsing
// I think that for now the client will be very simple and it will be up to the caller to
// determine how far to follow the chain.
// In the future this could be changed with a slight variation/addition to the public API

//
// Data Structures
//

// ResourceListIndex defines the structure of a Resource List Index data set
type ResourceListIndex struct {
	XMLName  xml.Name   `xml:"sitemapindex"`
	RSLink   []RSLN     `xml:"ln"`
	RSMD     RSMD       `xml:"md"`
	IndexSet []IndexDef `xml:"sitemap"`
}

// ResourceList hods the data from a resource list
type ResourceList struct {
	XMLName xml.Name      `xml:"urlset"`
	RSLink  []RSLN        `xml:"ln"`
	RSMD    RSMD          `xml:"md"`
	URLSet  []ResourceURL `xml:"url"`
}

// ResourceURL holds the data retrieved from the url tag set within a standard sitemap.xml
type ResourceURL struct {
	Loc        string `xml:"loc"`        // mandatory
	LastMod    string `xml:"lastmod"`    // optional
	ChangeFreq string `xml:"changefreq"` // optional
	RSMD       RSMD   `xml:"md"`         // optional
	RSLN       RSLN   `xml:"ln"`         // optional
}

// IndexDef holds those items defined as making up the resource list index data set
type IndexDef struct {
	Loc     string `xml:"loc"`     // mandatory
	LastMod string `xml:"lastmod"` // optional
	RSMD    RSMD   `xml:"md"`      // optional
}

// RSLN is the namespaced ln values defined in the resourcesync protocol
type RSLN struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

// RSMD is the namespaced md values defined in the resourcesync protocol
type RSMD struct {
	Capability string `xml:"capability,attr"`
	At         string `xml:"at,attr"`
	Completed  string `xml:"completed,attr"`
	// the next three are used on ResourceList only
	Hash   string `xml:"hash,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}
