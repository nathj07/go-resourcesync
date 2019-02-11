package extractor

import (
	"encoding/xml"
	"fmt"
	"strings"
)

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
