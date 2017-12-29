package resourcesync

import (
	"fmt"
	"testing"
)

// TestResourceListIndexString ensures the string formatting of a full ResourceListIndex is as expected.
// This will cover the other data structures that form part of the ResourceListIndex definition.
func TestResourceListIndexString(t *testing.T) {
	exp := `ResourceListIndex
Namespace: http://www.sitemaps.org/schemas/sitemap/0.9, Local: sitemapindex
Top Level LN
Rel: up HREF: http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/capabilitylist.xml
Top Level MD
Capability: resourcelist At: 2017-05-16T13:55:36Z Completed: 2017-05-16T13:56:17Z
Index Set
Loc: 
	http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist_0000.xml
	 LastMod:  RSMD: Capability:  At: 2017-05-16T13:55:38Z Completed: 2017-05-16T13:56:04Z
Loc: 
	http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist_0001.xml
	 LastMod:  RSMD: Capability:  At: 2017-05-16T13:56:11Z Completed: 2017-05-16T13:56:16Z`
	got := fmt.Sprintf("%s", expIndexRD.RLI)
	if got != exp {
		t.Errorf("Unexpected response printing ResourceListIndex:\nGot:\n%s\nExp:\n%s", got, exp)
	}
}

// TestResourceListString ensures the string formatting of a full ResourceList is as expected.
// This will cover the other data structures that form part of the ResourceList definition.
func TestResourceListString(t *testing.T) {
	exp := `ResourceList
Namespace: http://www.sitemaps.org/schemas/sitemap/0.9, Local: urlset
Top Level LN
Rel: up HREF: http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/capabilitylist.xml
Rel: index HREF: http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist-index.xml
Top Level MD
Capability: resourcelist At: 2017-05-16T13:56:11Z Completed: 2017-05-16T13:56:16Z
Index Set
Loc: 
	http://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/pdf/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9maW1tdS4yMDEyLjAwMTcwL3BkZg%3D%3D.pdf
	 LastMod: 2017-04-12T19:45:43Z ChangeFreq:  RSMD: Capability:  Hash: md5:d030c6d483b306029b0897630e67c550 Length: 360320 Type: application/pdf RSLN: Rel: describedBy HREF: http://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/metadata/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9maW1tdS4yMDEyLjAwMTcwL3BkZg%3D%3D.json
Loc: 
	http://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/pdf/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9mbmV1ci4yMDE0LjAwMDgwL3BkZg%3D%3D.pdf
	 LastMod: 2017-04-13T05:26:04Z ChangeFreq:  RSMD: Capability:  Hash: md5:24012e8eb5e1aeb4659616019c6d743f Length: 411256 Type: application/pdf RSLN: Rel: describedBy HREF: http://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/metadata/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9mbmV1ci4yMDE0LjAwMDgwL3BkZg%3D%3D.json`
	// Using Sprintf here ensures the formatting is correct but also that Stringer is properly implemented
	got := fmt.Sprintf("%s", expListRD.RL)
	if got != exp {
		t.Errorf("Unexpected response printing ResourceList:\nGot:\n%s\nExp:\n%s", got, exp)
	}
}
