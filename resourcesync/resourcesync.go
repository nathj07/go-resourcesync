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
	// Capability indicates this is a capability list
	Capability
	// ChangeList indicates this is a change list
	ChangeList
	// ChangeListIndex indicates this is a change list index
	ChangeListIndex
	// ResourceDumpManifest indicates this is the manifest data from a resourcedump
	ResourceDumpManifest
	// ChangeDumpManifest indicates this is the manifest data from a resourcedump
	ChangeDumpManifest
)

// These constants are correctly formatted strings that help to determine feed types
const (
	capabilityList = "capabilitylist"
	resourceList   = "resourcelist"
	changeList     = "changelist"
	// resourcedumpManifest is specific to the CORE fastsync. It is an xml file within the retrieved zip file that details
	// the relative local path for the unpacked items. Other than an extra attribute it conforms to the same schema as a changelist.
	resourcedumpManifest = "resourcedump-manifest"
	changedumpManifest   = "changedump-manifest"
)

// ErrUnsupportedFeedType is used when the feed type is not one of the supported set
var ErrUnsupportedFeedType = errors.New("unsupported feed type supplied")

// ResourceSync is the top level structure needed to interact with ResourceSync endpoints
type ResourceSync struct {
	Fetcher fetcher.RSFetcher
}

// ResourceData is the structure for holding the data returned from a ResoureceSync fetch.
// Only one of RL or RLI will be populated, the rType value will indicate which
type ResourceData struct {
	RL    *ResourceList
	RLI   *ResourceListIndex
	RType int // based on the type const above
}

// New is the simplest way to instantiate a ready to use ResourceSync object
func New(f fetcher.RSFetcher) *ResourceSync {
	return &ResourceSync{
		Fetcher: f,
	}
}

// Process takes the given target and fetches that page, parsing the data
// The returned structure indicate the type and then the relevant data can be found.
func (rs *ResourceSync) Process(baseTarget string) (*ResourceData, error) {
	data, status, err := rs.Fetcher.Fetch(baseTarget)
	if err != nil {
		return nil, fmt.Errorf("%d: %v", status, err)
	}
	rd, err := rs.Parse(data)
	if err != nil {
		return nil, err
	}
	return rd, nil
}

// Parse handles the unmarshaling of the feed data.
// The returned ResourceData will have one field populated and the RType value will indicate which.
func (rs *ResourceSync) Parse(feed []byte) (*ResourceData, error) {
	feedType := rs.determineBaseType(feed)
	switch feedType {
	case Index:
		return rs.parseIndexType(feed)
	case List:
		return rs.parseListType(feed)
	default:
		return nil, ErrUnsupportedFeedType
	}
}

func (rs *ResourceSync) parseIndexType(feed []byte) (*ResourceData, error) {
	rd := &ResourceData{
		RLI: &ResourceListIndex{},
		RL:  nil,
	}
	if err := xml.Unmarshal(feed, rd.RLI); err != nil {
		return nil, err
	}
	switch rd.RLI.RSMD.Capability {
	case changeList:
		rd.RType = ChangeListIndex
	case resourceList:
		rd.RType = Index
	default:
		return nil, ErrUnsupportedFeedType
	}
	return rd, nil
}

func (rs *ResourceSync) parseListType(feed []byte) (*ResourceData, error) {
	rd := &ResourceData{
		RLI: nil,
		RL:  &ResourceList{},
	}
	if err := xml.Unmarshal(feed, rd.RL); err != nil {
		return nil, err
	}
	switch rd.RL.RSMD.Capability {
	case resourceList:
		rd.RType = List
	case capabilityList:
		rd.RType = Capability
	case changeList:
		rd.RType = ChangeList
	case resourcedumpManifest:
		rd.RType = ResourceDumpManifest
	case changedumpManifest:
		rd.RType = ChangeDumpManifest
	default:
		return nil, ErrUnsupportedFeedType
	}
	return rd, nil
}

// determineBaseType simply establishes if the feed is <sitemapindex> or <urlset> at the top level.

// Any further determination is done under the parse method after the data has been unmarshalled.
func (rs *ResourceSync) determineBaseType(data []byte) int {
	if bytes.Contains(data, []byte("<sitemapindex")) {
		return Index
	}
	if bytes.Contains(data, []byte("<urlset")) {
		return List
	}
	return Unknown
}
