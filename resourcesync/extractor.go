package resourcesync

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Containing the data structures and associaterd functions

// ResourceListIndex defines the structure of a Resource List Index data set
type ResourceListIndex struct {
	XMLName  xml.Name   `xml:"sitemapindex"`
	RSLink   []RSLN     `xml:"ln"`
	RSMD     RSMD       `xml:"md"`
	IndexSet []IndexDef `xml:"sitemap"`
}

// String implements the stringer interface for ResourceListIndex ensuring consistent printing of values
func (rli *ResourceListIndex) String() string {
	str := fmt.Sprintf("ResourceListIndex\nNamespace: %s, Local: %s", rli.XMLName.Space, rli.XMLName.Local)
	if len(rli.RSLink) > 0 {
		linkTexts := []string{}
		str = fmt.Sprintf("%s\nTop Level LN", str)
		for _, ln := range rli.RSLink {
			linkTexts = append(linkTexts, ln.String())
		}
		str = fmt.Sprintf("%s\n%s", str, strings.Join(linkTexts, "\n"))
	}
	str = fmt.Sprintf("%s\nTop Level MD\n%s", str, rli.RSMD)
	if len(rli.IndexSet) > 0 {
		indices := []string{}
		str = fmt.Sprintf("%s\nIndex Set", str)
		for _, index := range rli.IndexSet {
			indices = append(indices, index.String())
		}
		str = fmt.Sprintf("%s\n%s", str, strings.Join(indices, "\n"))
	}
	return str
}

// ResourceList hods the data from a resource list
type ResourceList struct {
	XMLName xml.Name      `xml:"urlset"`
	RSLink  []RSLN        `xml:"ln"`
	RSMD    RSMD          `xml:"md"`
	URLSet  []ResourceURL `xml:"url"`
}

// String implements the stringer interface for ResourceList ensuring consistent printing of values
func (rl *ResourceList) String() string {
	str := fmt.Sprintf("ResourceList\nNamespace: %s, Local: %s", rl.XMLName.Space, rl.XMLName.Local)
	if len(rl.RSLink) > 0 {
		linkTexts := []string{}
		str = fmt.Sprintf("%s\nTop Level LN", str)
		for _, ln := range rl.RSLink {
			linkTexts = append(linkTexts, ln.String())
		}
		str = fmt.Sprintf("%s\n%s", str, strings.Join(linkTexts, "\n"))
	}
	str = fmt.Sprintf("%s\nTop Level MD\n%s", str, rl.RSMD)
	if len(rl.URLSet) > 0 {
		urls := []string{}
		str = fmt.Sprintf("%s\nIndex Set", str)
		for _, rURL := range rl.URLSet {
			urls = append(urls, rURL.String())
		}
		str = fmt.Sprintf("%s\n%s", str, strings.Join(urls, "\n"))
	}
	return str
}

// ResourceURL holds the data retrieved from the url tag set within a standard sitemap.xml
type ResourceURL struct {
	Loc        string `xml:"loc"`        // mandatory
	LastMod    string `xml:"lastmod"`    // optional
	ChangeFreq string `xml:"changefreq"` // optional
	RSMD       RSMD   `xml:"md"`         // optional
	RSLN       RSLN   `xml:"ln"`         // optional
}

// String implements the stringer interface for ResourceURL ensuring consistent printing of values
func (ru ResourceURL) String() string {
	return fmt.Sprintf("Loc: %s LastMod: %s ChangeFreq: %s RSMD: %s RSLN: %s",
		ru.Loc, ru.LastMod, ru.ChangeFreq, ru.RSMD.String(), ru.RSLN.String())
}

// IndexDef holds those items defined as making up the resource list index data set
type IndexDef struct {
	Loc     string `xml:"loc"`     // mandatory
	LastMod string `xml:"lastmod"` // optional
	RSMD    RSMD   `xml:"md"`      // optional
}

// String implements the stringer interface for IndexDef ensuring consistent printing of values
func (id IndexDef) String() string {
	return fmt.Sprintf("Loc: %s LastMod: %s RSMD: %s",
		id.Loc, id.LastMod, id.RSMD.String())
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
