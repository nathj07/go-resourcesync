package core

import (
	"encoding/json"
	"fmt"
	"github.com/nathj07/go-resourcesync/fetcher"
	"reflect"
	"strings"
)

// ArticleWrapper is the top level struct for unmarshalling the CORE article data
type ArticleWrapper struct {
	Status string  `json:"status"`
	Data   Article `json:"data"`
}

// Article holds the article details
type Article struct {
	ID                 string        `json:"id"`
	Authors            []string      `json:"authors"`
	Contributors       []string      `json:"contributors"`
	DatePublished      string        `json:"datePublished"`
	Description        string        `json:"description"`
	FullText           string        `json:"fullText"` // only returned if query params include 'fulltext=true'
	Identifiers        []string      `json:"identifiers"`
	Language           Language      `json:"language"`
	Publisher          string        `json:"publisher"`
	Relations          []string      `json:"relations"`
	Repositories       []Repository  `json:"repositories"`
	RepositoryDocument RepositoryDoc `json:"repositoryDocument"`
	Subjects           []string      `json:"subjects"`
	Title              string        `json:"title"`
	Topics             []string      `json:"topics"`
	Types              []string      `json:"types"`
	Year               int           `json:"year"`
	FullTextURLs       []string      `json:"fulltextUrls"` // only returned if the query params include 'urls=true'
	FullTextIdentifier string        `json:"fulltextIdentifier"`
	OAI                string        `json:"oai"`
	DownloadURL        string        `json:"downloadUrl"`
}

// Language represents the language details supplied in the ResourceSync article JSON
type Language struct {
	Code string `json:"code"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Repository handles the repository provided details
type Repository struct {
	ID                 string `json:"id"`
	OpenDoarID         int    `json:"openDoarId"`
	Name               string `json:"name"`
	URI                string `json:"uri"`
	URLHomepage        string `json:"urlHomepage"`
	URLOAIPMH          string `json:"urlOaipmh"`
	URIJournals        string `json:"uriJournals"`
	PhysicalName       string `json:"physicalName"`
	Source             string `json:"source"`
	Software           string `json:"software"`
	MetadataFormat     string `json:"metadataFormat"`
	Description        string `json:"description"`
	Journal            string `json:"journal"`
	RoarID             int    `json:"roarId"`
	PdfStatus          string `json:"pdfStatus"`
	NrUpdates          int    `json:"nrUpdates"`
	Disabled           bool   `json:"disabled"`
	LastUpdateTime     string `json:"lastUpdateTime"`
	RepositoryLocation string `json:"repositoryLocation"`
}

// RepositoryDoc handles the repository document details
type RepositoryDoc struct {
	PDFStatus       int   `json:"pdfStatus"`
	MetadataAdded   int64 `json:"metadataAdded"`
	MetadataUpdated int64 `json:"metadataUpdated"`
	DepositedDate   int64 `json:"depositedDate"`
}

type Extractor struct {
	Fetcher fetcher.RSFetcher
}

func (ce *Extractor) Process(target, apiKey string) (*ArticleWrapper, error) {
	data, status, err := ce.Fetcher.Fetch(target + "?apiKey=" + apiKey)
	if err != nil {
		return nil, fmt.Errorf("%d: %v", status, err)
	}
	return ce.ExtractArticle(data)
}

// ExtractArticle is a convenience method around unmarshalling the CORE article metadata
func (ce *Extractor) ExtractArticle(rawData []byte) (*ArticleWrapper, error) {
	res := &ArticleWrapper{}
	err := json.Unmarshal(rawData, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// String implements the Stringer interface to ensure consistent printing of the extracted article metadata.
// If fields are empty they are omitted from the output.
func (aw *ArticleWrapper) String() string {
	sb := &strings.Builder{}
	s := reflect.ValueOf(aw).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Fprintf(sb, "%s: %v\n", typeOfT.Field(i).Name, f.Interface())
	}
	return sb.String()
}
