package extractor

import (
	"fmt"
	"strings"
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
)

// ResourceData is the structure for holding the data returned from a ResoureceSync fetch.
// Only one of RL or RLI will be populated, the rType value will indicate which
type ResourceData struct {
	RL    *ResourceList
	RLI   *ResourceListIndex
	RType int // based on the type const above
}

// RSLN is the namespaced ln values defined in the resourcesync protocol
type RSLN struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

// String implements the stringer interface for RSLN ensuring consistent printing of values
func (rsln RSLN) String() string {
	return fmt.Sprintf("Rel: %s HREF: %s", rsln.Rel, rsln.Href)
}

// RSMD is the namespaced md values defined in the resourcesync protocol.
// Not all values are present in all cases.
type RSMD struct {
	Capability string `xml:"capability,attr"`
	At         string `xml:"at,attr"`
	Completed  string `xml:"completed,attr"`
	// the next three are used on ResourceList only
	Hash   string `xml:"hash,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
	// the following are used on ChangeList only
	From     string `xml:"from,attr"`
	Until    string `xml:"until,attr"`
	Change   string `xml:"change,attr"`
	DateTime string `xml:"datetime,attr"`
}

// String implements the stringer interface for RSMD ensuring consistent printing of values
// As this structure holds data from different feed types only non empty strings are printed
func (rsmd RSMD) String() string {
	str := fmt.Sprintf("Capability: %s ", rsmd.Capability)
	if rsmd.At != "" {
		str = fmt.Sprintf("%sAt: %s ", str, rsmd.At)
	}
	if rsmd.Completed != "" {
		str = fmt.Sprintf("%sCompleted: %s ", str, rsmd.Completed)
	}
	if rsmd.Hash != "" {
		str = fmt.Sprintf("%sHash: %s ", str, rsmd.Hash)
	}
	if rsmd.Length != "" {
		str = fmt.Sprintf("%sLength: %s ", str, rsmd.Length)
	}
	if rsmd.Type != "" {
		str = fmt.Sprintf("%sType: %s ", str, rsmd.Type)
	}
	if rsmd.From != "" {
		str = fmt.Sprintf("%sFrom: %s ", str, rsmd.From)
	}
	if rsmd.Until != "" {
		str = fmt.Sprintf("%sUntil: %s ", str, rsmd.Until)
	}
	if rsmd.Change != "" {
		str = fmt.Sprintf("%sChange: %s ", str, rsmd.Change)
	}
	if rsmd.DateTime != "" {
		str = fmt.Sprintf("%sDateTime: %s ", str, rsmd.DateTime)
	}
	return strings.TrimSpace(str)
}
