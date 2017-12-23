package resourcesync

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetermineBaseType(t *testing.T) {
	rs := &ResourceSync{}
	type testData struct {
		tag      string
		testBody []byte
		expRType int
	}

	testTable := []testData{
		testData{
			tag:      "INDEX",
			testBody: testResourceListIndex,
			expRType: Index,
		},
		testData{
			tag:      "LIST",
			testBody: testResourceList,
			expRType: List,
		},
		testData{
			tag:      "LIST",
			testBody: testCapabilityList,
			expRType: List,
		},
		testData{
			tag:      "UNKNOWN",
			testBody: []byte(`<xml><unsupported>bad content</unsupported></xml>`),
			expRType: Unknown,
		},
	}
	for _, td := range testTable {
		if gotType := rs.determineBaseType(td.testBody); gotType != td.expRType {
			t.Errorf("%s unexpected type %d exp: %d", td.tag, gotType, td.expRType)
		}
	}
}

func TestParse(t *testing.T) {
	rs := &ResourceSync{}
	type testData struct {
		tag      string
		testBody []byte
		expRD    *ResourceData
	}

	testTable := []testData{
		testData{
			tag:      "INDEX",
			testBody: testResourceListIndex,
			expRD: &ResourceData{
				RType: Index,
				RLI: &ResourceListIndex{
					XMLName: xml.Name{
						Space: "http://www.sitemaps.org/schemas/sitemap/0.9",
						Local: "sitemapindex",
					},
					RSLink: []RSLN{
						RSLN{
							Rel:  "up",
							Href: "http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/capabilitylist.xml",
						},
					},
					RSMD: RSMD{
						Capability: "resourcelist",
						At:         "2017-05-16T13:55:36Z",
						Completed:  "2017-05-16T13:56:17Z",
					},
					IndexSet: []IndexDef{
						IndexDef{
							Loc: "\n\thttp://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist_0000.xml\n\t",
							RSMD: RSMD{
								At:        "2017-05-16T13:55:38Z",
								Completed: "2017-05-16T13:56:04Z",
							},
						},
						IndexDef{
							Loc: "\n\thttp://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist_0001.xml\n\t",
							RSMD: RSMD{
								At:        "2017-05-16T13:56:11Z",
								Completed: "2017-05-16T13:56:16Z",
							},
						},
					},
				},
			},
		},
		testData{
			tag:      "LIST",
			testBody: testResourceList,
			expRD: &ResourceData{
				RType: List,
				RL: &ResourceList{
					XMLName: xml.Name{
						Space: "http://www.sitemaps.org/schemas/sitemap/0.9",
						Local: "urlset",
					},
					RSLink: []RSLN{
						RSLN{
							Rel:  "up",
							Href: "http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/capabilitylist.xml",
						},
						RSLN{
							Rel:  "index",
							Href: "http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist-index.xml",
						},
					},
					RSMD: RSMD{
						Capability: "resourcelist",
						At:         "2017-05-16T13:56:11Z",
						Completed:  "2017-05-16T13:56:16Z",
					},
					URLSet: []ResourceURL{
						ResourceURL{
							Loc:     "\n\thttp://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/pdf/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9maW1tdS4yMDEyLjAwMTcwL3BkZg%3D%3D.pdf\n\t",
							LastMod: "2017-04-12T19:45:43Z",
							RSMD: RSMD{
								Hash:   "md5:d030c6d483b306029b0897630e67c550",
								Length: "360320",
								Type:   "application/pdf",
							},
							RSLN: RSLN{
								Rel:  "describedBy",
								Href: "http://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/metadata/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9maW1tdS4yMDEyLjAwMTcwL3BkZg%3D%3D.json",
							},
						},
						ResourceURL{
							Loc:     "\n\thttp://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/pdf/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9mbmV1ci4yMDE0LjAwMDgwL3BkZg%3D%3D.pdf\n\t",
							LastMod: "2017-04-13T05:26:04Z",
							RSMD: RSMD{
								Hash:   "md5:24012e8eb5e1aeb4659616019c6d743f",
								Length: "411256",
								Type:   "application/pdf",
							},
							RSLN: RSLN{
								Rel:  "describedBy",
								Href: "http://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/metadata/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9mbmV1ci4yMDE0LjAwMDgwL3BkZg%3D%3D.json",
							},
						},
					},
				},
			},
		},
		testData{
			tag:      "Capability",
			testBody: testCapabilityList,
			expRD: &ResourceData{
				RType: Capability,
				RL: &ResourceList{
					XMLName: xml.Name{
						Space: "http://www.sitemaps.org/schemas/sitemap/0.9",
						Local: "urlset",
					},
					RSLink: []RSLN{
						RSLN{
							Rel:  "up",
							Href: "http://publisher-connector.core.ac.uk/resourcesync/.well-known/resourcesync",
						},
					},
					RSMD: RSMD{
						Capability: "capabilitylist",
					},
					URLSet: []ResourceURL{
						ResourceURL{
							Loc: "\n\thttp://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist-index.xml\n\t",
							RSMD: RSMD{
								Capability: "resourcelist",
							},
						},
					},
				},
			},
		},
	}
	for _, td := range testTable {
		gotRD, err := rs.Parse(td.testBody)
		if err != nil {
			t.Errorf("%s Unexpected error in parse: %v", td.tag, err)
		}
		assert.Equal(t, td.expRD, gotRD)
	}

}

//
// Test Data
//

var testResourceListIndex = []byte(`<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:rs="http://www.openarchives.org/rs/terms/">
	<rs:ln href="http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/capabilitylist.xml" rel="up"/>
	<rs:md at="2017-05-16T13:55:36Z" capability="resourcelist" completed="2017-05-16T13:56:17Z"/>
	<sitemap>
	<loc>
	http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist_0000.xml
	</loc>
	<rs:md at="2017-05-16T13:55:38Z" completed="2017-05-16T13:56:04Z"/>
	</sitemap>
	<sitemap>
	<loc>
	http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist_0001.xml
	</loc>
	<rs:md at="2017-05-16T13:56:11Z" completed="2017-05-16T13:56:16Z"/>
	</sitemap>
	</sitemapindex>`)

var testResourceList = []byte(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:rs="http://www.openarchives.org/rs/terms/">
	<rs:ln href="http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/capabilitylist.xml" rel="up"/>
	<rs:ln href="http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist-index.xml" rel="index"/>
	<rs:md at="2017-05-16T13:56:11Z" capability="resourcelist" completed="2017-05-16T13:56:16Z"/>
	<url>
	<loc>
	http://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/pdf/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9maW1tdS4yMDEyLjAwMTcwL3BkZg%3D%3D.pdf
	</loc>
	<lastmod>2017-04-12T19:45:43Z</lastmod>
	<rs:md hash="md5:d030c6d483b306029b0897630e67c550" length="360320" type="application/pdf"/>
	<rs:ln href="http://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/metadata/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9maW1tdS4yMDEyLjAwMTcwL3BkZg%3D%3D.json" rel="describedBy"/>
	</url>
	<url>
	<loc>
	http://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/pdf/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9mbmV1ci4yMDE0LjAwMDgwL3BkZg%3D%3D.pdf
	</loc>
	<lastmod>2017-04-13T05:26:04Z</lastmod>
	<rs:md hash="md5:24012e8eb5e1aeb4659616019c6d743f" length="411256" type="application/pdf"/>
	<rs:ln href="http://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/metadata/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9mbmV1ci4yMDE0LjAwMDgwL3BkZg%3D%3D.json" rel="describedBy"/>
	</url>
	</urlset>`)

var testCapabilityList = []byte(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:rs="http://www.openarchives.org/rs/terms/">
	<rs:ln href="http://publisher-connector.core.ac.uk/resourcesync/.well-known/resourcesync" rel="up"/>
	<rs:md capability="capabilitylist"/>
	<url>
	<loc>
	http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist-index.xml
	</loc>
	<rs:md capability="resourcelist"/>
	</url>
	</urlset>`)
