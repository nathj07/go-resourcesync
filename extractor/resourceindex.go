package extractor

import (
	"encoding/xml"
	"fmt"
	"strings"
)

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
