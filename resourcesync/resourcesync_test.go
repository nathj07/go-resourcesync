package resourcesync

import (
	"encoding/xml"
	"fmt"
	"github.com/nathj07/go-resourcesync/extractor"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nathj07/go-resourcesync/fetcher"

	"github.com/stretchr/testify/assert"
)

func TestDetermineBaseType(t *testing.T) {
	rs := &ResourceSync{}


	testTable := []struct {
		tag      string
		testBody []byte
		expRType int
	}{
		{
			tag:      "INDEX",
			testBody: testResourceListIndex,
			expRType: Index,
		},
		{
			tag:      "LIST",
			testBody: testResourceList,
			expRType: List,
		},
		{
			tag:      "LIST",
			testBody: testCapabilityList,
			expRType: List,
		},
		{
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

	testTable := []struct {
		tag      string
		testBody []byte
		expRD    *ResourceData
	}{
		{
			tag:      "INDEX",
			testBody: testResourceListIndex,
			expRD:    expIndexRD,
		},
		{
			tag:      "LIST",
			testBody: testResourceList,
			expRD:    expListRD,
		},
		{
			tag:      "CAPABILITY",
			testBody: testCapabilityList,
			expRD:    expCapabilityRD,
		},
		{
			tag:      "CHANGELIST",
			testBody: testChangeList,
			expRD:    expChangeListRD,
		},
		{
			tag:      "CHANGELISTINDEX",
			testBody: testChangeListIndex,
			expRD:    expChangeListIndexRD,
		},
	}
	for _, td := range testTable {
		gotRD, err := rs.Parse(td.testBody)
		if err != nil {
			t.Errorf("%s Unexpected error in parse: %v", td.tag, err)
		}
		assert.Equal(t, td.expRD, gotRD, td.tag)
	}
}

func TestProcess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(testResourceListIndex))
	}))
	rs := New(&fetcher.BasicRSFetcher{})
	rd, err := rs.Process(server.URL)
	if err != nil {
		t.Fatalf("Unexpected error from Process: %v", err)
	}
	assert.Equal(t, expIndexRD, rd)
}

func TestProcessUnsupportedType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(testUnsupported))
	}))
	rs := New(&fetcher.BasicRSFetcher{})
	rd, err := rs.Process(server.URL)
	if err != ErrUnsupportedFeedType {
		t.Fatalf("Unexpected error: %v; exp: %v.", err, ErrUnsupportedFeedType)
	}
	if rd != nil {
		t.Errorf("Unexpected data returned %v", rd)
	}
}

func TestProcessUnsupportedXML(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<xml><unsupported>bad content</unsupported></xml>")
	}))
	rs := New(&fetcher.BasicRSFetcher{})
	rd, err := rs.Process(server.URL)
	if err != ErrUnsupportedFeedType {
		t.Fatalf("Unexpected error: %v; exp: %v.", err, ErrUnsupportedFeedType)
	}
	if rd != nil {
		t.Errorf("Unexpected data returned %v", rd)
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

var expIndexRD = &ResourceData{
	RType: Index,
	RLI: &extractor.ResourceListIndex{
		XMLName: xml.Name{
			Space: "http://www.sitemaps.org/schemas/sitemap/0.9",
			Local: "sitemapindex",
		},
		RSLink: []extractor.RSLN{
			extractor.RSLN{
				Rel:  "up",
				Href: "http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/capabilitylist.xml",
			},
		},
		RSMD: extractor.RSMD{
			Capability: "resourcelist",
			At:         "2017-05-16T13:55:36Z",
			Completed:  "2017-05-16T13:56:17Z",
		},
		IndexSet: []extractor.IndexDef{
			extractor.IndexDef{
				Loc: "\n\thttp://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist_0000.xml\n\t",
				RSMD: extractor.RSMD{
					At:        "2017-05-16T13:55:38Z",
					Completed: "2017-05-16T13:56:04Z",
				},
			},
			extractor.IndexDef{
				Loc: "\n\thttp://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist_0001.xml\n\t",
				RSMD: extractor.RSMD{
					At:        "2017-05-16T13:56:11Z",
					Completed: "2017-05-16T13:56:16Z",
				},
			},
		},
	},
}

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

var expListRD *ResourceData = &extractor.ResourceData{
	RType: List,
	RL: &extractor.ResourceList{
		XMLName: xml.Name{
			Space: "http://www.sitemaps.org/schemas/sitemap/0.9",
			Local: "urlset",
		},
		RSLink: []extractor.RSLN{
			extractor.RSLN{
				Rel:  "up",
				Href: "http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/capabilitylist.xml",
			},
			extractor.RSLN{
				Rel:  "index",
				Href: "http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist-index.xml",
			},
		},
		RSMD: extractor.RSMD{
			Capability: "resourcelist",
			At:         "2017-05-16T13:56:11Z",
			Completed:  "2017-05-16T13:56:16Z",
		},
		URLSet: []extractor.ResourceURL{
			extractor.ResourceURL{
				Loc:     "\n\thttp://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/pdf/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9maW1tdS4yMDEyLjAwMTcwL3BkZg%3D%3D.pdf\n\t",
				LastMod: "2017-04-12T19:45:43Z",
				RSMD: extractor.RSMD{
					Hash:   "md5:d030c6d483b306029b0897630e67c550",
					Length: "360320",
					Type:   "application/pdf",
				},
				RSLN: extractor.RSLN{
					Rel:  "describedBy",
					Href: "http://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/metadata/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9maW1tdS4yMDEyLjAwMTcwL3BkZg%3D%3D.json",
				},
			},
			extractor.ResourceURL{
				Loc:     "\n\thttp://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/pdf/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9mbmV1ci4yMDE0LjAwMDgwL3BkZg%3D%3D.pdf\n\t",
				LastMod: "2017-04-13T05:26:04Z",
				RSMD: extractor.RSMD{
					Hash:   "md5:24012e8eb5e1aeb4659616019c6d743f",
					Length: "411256",
					Type:   "application/pdf",
				},
				RSLN: extractor.RSLN{
					Rel:  "describedBy",
					Href: "http://publisher-connector.core.ac.uk/resourcesync/data/Frontiers/metadata/000/aHR0cDovL2pvdXJuYWwuZnJvbnRpZXJzaW4ub3JnL2FydGljbGUvMTAuMzM4OS9mbmV1ci4yMDE0LjAwMDgwL3BkZg%3D%3D.json",
				},
			},
		},
	},
}

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

var expCapabilityRD = &ResourceData{
	RType: Capability,
	RL: &extractor.ResourceList{
		XMLName: xml.Name{
			Space: "http://www.sitemaps.org/schemas/sitemap/0.9",
			Local: "urlset",
		},
		RSLink: []extractor.RSLN{
			extractor.RSLN{
				Rel:  "up",
				Href: "http://publisher-connector.core.ac.uk/resourcesync/.well-known/resourcesync",
			},
		},
		RSMD: extractor.RSMD{
			Capability: "capabilitylist",
		},
		URLSet: []extractor.ResourceURL{
			extractor.ResourceURL{
				Loc: "\n\thttp://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist-index.xml\n\t",
				RSMD: extractor.RSMD{
					Capability: "resourcelist",
				},
			},
		},
	},
}

var testUnsupported = []byte(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:rs="http://www.openarchives.org/rs/terms/">
	<rs:ln href="http://publisher-connector.core.ac.uk/resourcesync/.well-known/resourcesync" rel="up"/>
	<rs:md capability="unsupported"/>
	<url>
	<loc>
	http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/pdf/resourcelist-index.xml
	</loc>
	<rs:md capability="resourcelist"/>
	</url>
	</urlset>`)

var testChangeList = []byte(`<?xml version="1.0" encoding="UTF-8"?>
	<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
			xmlns:rs="http://www.openarchives.org/rs/terms/">
	  <rs:ln rel="up"
			 href="http://example.com/dataset1/capabilitylist.xml"/>
	  <rs:ln rel="index"
			 href="http://example.com/dataset1/changelist.xml"/>
	  <rs:md capability="changelist"
			 from="2013-01-02T00:00:00Z"
			 until="2013-01-03T00:00:00Z"/>
	  <url>
		  <loc>http://example.com/res7.html</loc>
		  <rs:md change="created" datetime="2013-01-02T12:00:00Z"/>
	  </url>
	  <url>
		  <loc>http://example.com/res9.pdf</loc>
		  <rs:md change="updated" datetime="2013-01-02T13:00:00Z"/>
	  </url>
	  <url>
		  <loc>http://example.com/res5.tiff</loc>
		  <rs:md change="deleted" datetime="2013-01-02T19:00:00Z"/>
	  </url>
	  <url>
		  <loc>http://example.com/res7.html</loc>
		  <rs:md change="updated" datetime="2013-01-02T20:00:00Z"/>
	  </url>
	</urlset>`)

var expChangeListRD = &extractor.ResourceData{
	RType: ChangeList,
	RL: &extractor.ResourceList{
		XMLName: xml.Name{
			Space: "http://www.sitemaps.org/schemas/sitemap/0.9",
			Local: "urlset",
		},
		RSLink: []extractor.RSLN{
			extractor.RSLN{
				Rel:  "up",
				Href: "http://example.com/dataset1/capabilitylist.xml",
			},
			extractor.RSLN{
				Rel:  "index",
				Href: "http://example.com/dataset1/changelist.xml",
			},
		},
		RSMD: extractor.RSMD{
			Capability: "changelist",
			From:       "2013-01-02T00:00:00Z",
			Until:      "2013-01-03T00:00:00Z",
		},
		URLSet: []extractor.ResourceURL{
			extractor.ResourceURL{
				Loc: "http://example.com/res7.html",
				RSMD: extractor.RSMD{
					Change:   "created",
					DateTime: "2013-01-02T12:00:00Z",
				},
			},
			extractor.ResourceURL{
				Loc: "http://example.com/res9.pdf",
				RSMD: extractor.RSMD{
					Change:   "updated",
					DateTime: "2013-01-02T13:00:00Z",
				},
			},
			extractor.ResourceURL{
				Loc: "http://example.com/res5.tiff",
				RSMD: extractor.RSMD{
					Change:   "deleted",
					DateTime: "2013-01-02T19:00:00Z",
				},
			},
			extractor.ResourceURL{
				Loc: "http://example.com/res7.html",
				RSMD: extractor.RSMD{
					Change:   "updated",
					DateTime: "2013-01-02T20:00:00Z",
				},
			},
		},
	},
}

var testChangeListIndex = []byte(`<?xml version="1.0" encoding="UTF-8"?>
	<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
			xmlns:rs="http://www.openarchives.org/rs/terms/">
		<rs:ln rel="up"
				href="http://example.com/dataset1/capabilitylist.xml"/>
		<rs:md capability="changelist"
				from="2013-01-01T00:00:00Z"/>
		<sitemap>
			<loc>http://example.com/20130101-changelist.xml</loc>
			<rs:md from="2013-01-01T00:00:00Z" 
					until="2013-01-02T00:00:00Z"/>
		</sitemap>
		<sitemap>
			<loc>http://example.com/20130102-changelist.xml</loc>
			<rs:md from="2013-01-02T00:00:00Z" 
					until="2013-01-03T00:00:00Z"/>
		</sitemap>
		<sitemap>
			<loc>http://example.com/20130103-changelist.xml</loc>
			<rs:md from="2013-01-03T00:00:00Z"/>
		</sitemap>
	</sitemapindex>`)

var expChangeListIndexRD = &extractor.ResourceData{
	RType: ChangeListIndex,
	RLI: &extractor.ResourceListIndex{
		XMLName: xml.Name{
			Space: "http://www.sitemaps.org/schemas/sitemap/0.9",
			Local: "sitemapindex",
		},
		RSLink: []extractor.RSLN{
			extractor.RSLN{
				Rel:  "up",
				Href: "http://example.com/dataset1/capabilitylist.xml",
			},
		},
		RSMD: extractor.RSMD{
			Capability: "changelist",
			From:       "2013-01-01T00:00:00Z",
		},
		IndexSet: []extractor.IndexDef{
			extractor.IndexDef{
				Loc: "http://example.com/20130101-changelist.xml",
				RSMD: extractor.RSMD{
					From:  "2013-01-01T00:00:00Z",
					Until: "2013-01-02T00:00:00Z",
				},
			},
			extractor.IndexDef{
				Loc: "http://example.com/20130102-changelist.xml",
				RSMD: extractor.RSMD{
					From:  "2013-01-02T00:00:00Z",
					Until: "2013-01-03T00:00:00Z",
				},
			},
			extractor.IndexDef{
				Loc: "http://example.com/20130103-changelist.xml",
				RSMD: extractor.RSMD{
					From: "2013-01-03T00:00:00Z",
				},
			},
		},
	},
}
