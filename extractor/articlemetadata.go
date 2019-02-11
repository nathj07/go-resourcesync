package extractor

import (
	"encoding/json"
)

// CoreArticleWrapper is the top level struct for unmarshalling the CORE article data
type CoreArticleWrapper struct {
	Status string      `json:"status"`
	Data   CoreArticle `json:"data"`
}

// CoreArticle holds the article details
type CoreArticle struct {
	ID                 string            `json:"id"`
	Authors            []string          `json:"authors"`
	Contributors       []string          `json:"contributors"`
	DatePublished      string            `json:"datePublished"`
	Description        string            `json:"description"`
	Identifiers        []string          `json:"identifiers"`
	Language           CoreLanguage      `json:"language"`
	Publisher          string            `json:"publisher"`
	Relations          []string          `json:"relations"`
	Repositories       []CoreRepository  `json:"repositories"`
	RepositoryDocument CoreRepositoryDoc `json:"repositoryDocument"`
	Subjects           []string          `json:"subjects"`
	Title              string            `json:"title"`
	Topics             []string          `json:"topics"`
	Types              []string          `json:"types"`
	Year               int               `json:"year"`
	FulltextIdentifier string            `json:"fulltextIdentifier"`
	Oai                string            `json:"oai"`
	DownloadURL        string            `json:"downloadUrl"`
}

// CoreLanguage represents the language details supplied in the ResourceSync article JSON
type CoreLanguage struct {
	Code string `json:"code"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CoreRepository handles the repository provided details
type CoreRepository struct {
	ID                 string `json:"id"`
	OpenDoarID         int    `json:"openDoarId"`
	Name               string `json:"name"`
	URI                string `json:"uri"`
	URLHomepage        string `json:"urlHomepage"`
	URLOaipmh          string `json:"urlOaipmh"`
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

// CoreRepositoryDoc handles the repository document details
type CoreRepositoryDoc struct {
	PdfStatus       int   `json:"pdfStatus"`
	MetadataAdded   int64 `json:"metadataAdded"`
	MetadataUpdated int64 `json:"metadataUpdated"`
	DepositedDate   int64 `json:"depositedDate"`
}

type CoreExtractor struct{}

// ExtractArticle implements the Core extractor interface and returns a single Core Article ready for use in CIS
func (ce *CoreExtractor) ExtractArticle(rawData []byte) (*CoreArticleWrapper, error) {
	res := &CoreArticleWrapper{}
	err := json.Unmarshal(rawData, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

