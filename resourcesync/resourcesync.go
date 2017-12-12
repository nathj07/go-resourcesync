package resourcesync

import (
	"bytes"
	"encoding/xml"
	"errors"

	"github.com/nathj07/go-resourcesync/fetcher"
)

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
	Fetcher fetcher.RSFetcher
}

// New is the simplest way to instantiate a ready to use ResourceSync object
func New(f fetcher.RSFetcher) *ResourceSync {
	return &ResourceSync{
		Fetcher: f,
	}
}

// Parse handles the unmarshaling of the feed data.
// TODO: MAKE THIS MORE GENERAL TO SERVE INDEX AND LIST OR HAVE DIFFERENT FUNCTIONS.
// THINK ABOUT NICEST API FOR USER OF LIBRARY
func (rs *ResourceSync) Parse(feed []byte) (*ResourceList, error) {
	feedType := rs.determineType(feed)
	switch feedType {
	case Index:
		return nil, ErrUnsupportedFeedType
	case List:
		rl := &ResourceList{}
		if err := xml.Unmarshal(feed, rl); err != nil {
			return nil, err
		}
		return rl, nil
	default:
		return nil, ErrUnsupportedFeedType
	}

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
// Function to retrieve data
// Function to parse data - this should be exported and take a []byte to parse returning an exported struct
// Consider not returning the raw dta structure but a simplified parsed data structure
// Function to do it all in one go?
// I think that for now the client will be very simple and it will be up to the caller to
// determine how far to follow the chain.
// In the future this could be changed with a slight variation/addition to the public API

//
// Data Structures
//

// ResourceListIndex defines the structure of a Resource List Index data set
type ResourceListIndex struct {
	XMLName     xml.Name   `xml:"sitemapindex"`
	RSLink      []RSLN     `xml:"ln"`
	RSMD        RSMD       `xml:"md"`
	ResourceSet []IndexDef `xml:"sitemap"`
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
}

// IndexDef holds those items defined as making up the resource list index data set
type IndexDef struct {
	Loc     string `xml:"loc"`     // mandatory
	LastMod string `xml:"lastmod"` // optional
	RSMD    RSMD   `xml:"rs:md"`   // optional
}

// RSLN is the namespaced ln values defined in the resourcesync protocol
type RSLN struct {
	Rel  string `xml:"rel"`
	Href string `xml:"href"`
}

// RSMD is the namespaced md values defined in the resourcesync protocol
type RSMD struct {
	Capability string `xml:"capability,attr"`
	At         string `xml:"at,attr"`
	Completed  string `xml:"completed,attr"`
	// the next three are used on ResourceList only
	Hash   string `xml:"hash"`
	Length string `xml:"length"`
	Type   string `xml:"type"`
}
