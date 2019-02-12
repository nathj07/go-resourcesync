package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nathj07/go-resourcesync/fetcher"
)

func TestProcess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	ce := &Extractor{Fetcher: &fetcher.BasicRSFetcher{}}
	_, err := ce.Process(server.URL, "normally_valid_key")
	require.NotNil(t, err)
}

func TestProcessFetchFail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(testArticleData))
	}))
	ce := &Extractor{Fetcher: &fetcher.BasicRSFetcher{}}
	article, err := ce.Process(server.URL, "normally_valid_key")
	if err != nil {
		t.Fatalf("Unexpected error from Process: %v", err)
	}
	assert.Equal(t, expArticleWrapper, article)
}

func TestArticleString(t *testing.T){
	assert.Equal(t, expArticleString, fmt.Sprintf("%s", expArticleWrapper))
}

//
// Test Data
//

// testArticle data does not include empty value fields, and the data has been anonymized
var testArticleData = []byte(`{
    "status": "OK",
    "data": {
        "id": "123456",
        "authors": [
            "Davies, Nathan"
        ],
        "contributors": [
            "Davies, Sarah",
            "Smith, John",
            "Brown, Emma"
        ],
        "datePublished": "2012",
        "description": "Keywords: journalist’s attitudes, media ethics, professional values, responsibility, accountability, transparency, self-regulation, Estonian journalism, MediaAct.",
        "identifiers": [
            "oai:dspace.ac:123456"
        ],
        "language": {
            "code": "et",
            "id": 11,
            "name": "Estonian"
        },
        "publisher": "Publishing House",
        "relations": [],
        "repositories": [
            {
                "id": "123",
                "openDoarId": 0,
                "name": "DSpace at Publishing House",
                "physicalName": "noname"
            }
        ],
        "repositoryDocument": {
            "pdfStatus": 1,
            "metadataAdded": 1377620298000,
            "metadataUpdated": 1547433963000,
            "depositedDate": 1369695600000
        },
        "subjects": [
            "Thesis"
        ],
        "title": "This is a long and complicated article",
        "topics": [
            "article",
            "complexity",
            "media"
        ],
        "year": 2012,
        "fulltextIdentifier": "https://example.com/download/pdf/123456.pdf",
        "oai": "oai:dspace.ac:123456",
        "downloadUrl": "https://example.com/download/pdf/123456.pdf"
    }
}`)

var expArticleWrapper = &ArticleWrapper{
	Status: "OK",
	Data: Article{
		ID:            "123456",
		Authors:       []string{"Davies, Nathan"},
		Contributors:  []string{"Davies, Sarah", "Smith, John", "Brown, Emma"},
		DatePublished: "2012",
		Description:   "Keywords: journalist’s attitudes, media ethics, professional values, responsibility, accountability, transparency, self-regulation, Estonian journalism, MediaAct.",
		Identifiers:   []string{"oai:dspace.ac:123456"},
		Language: Language{
			Code: "et",
			ID:   11,
			Name: "Estonian",
		},
		Publisher: "Publishing House",
		Relations: []string{},
		Repositories: []Repository{
			{
				ID:           "123",
				OpenDoarID:   0,
				Name:         "DSpace at Publishing House",
				PhysicalName: "noname",
			},
		},
		RepositoryDocument: RepositoryDoc{
			PDFStatus:       1,
			MetadataAdded:   1377620298000,
			MetadataUpdated: 1547433963000,
			DepositedDate:   1369695600000,
		},
		Subjects:           []string{"Thesis"},
		Title:              "This is a long and complicated article",
		Topics:             []string{"article", "complexity", "media"},
		Year:               2012,
		FullTextIdentifier: "https://example.com/download/pdf/123456.pdf",
		OAI:                "oai:dspace.ac:123456",
		DownloadURL:        "https://example.com/download/pdf/123456.pdf",
	},
}

var expArticleString = `Status: OK
Data: {123456 [Davies, Nathan] [Davies, Sarah Smith, John Brown, Emma] 2012 Keywords: journalist’s attitudes, media ethics, professional values, responsibility, accountability, transparency, self-regulation, Estonian journalism, MediaAct. [oai:dspace.ac:123456] {et 11 Estonian} Publishing House [] [{123 0 DSpace at Publishing House     noname      0  0 false  }] {1 1377620298000 1547433963000 1369695600000} [Thesis] This is a long and complicated article [article complexity media] [] 2012 https://example.com/download/pdf/123456.pdf oai:dspace.ac:123456 https://example.com/download/pdf/123456.pdf}
`