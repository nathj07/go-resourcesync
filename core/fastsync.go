package core

import (
	"encoding/json"
	"fmt"
	"strings"
)

// FSArticle is a struct representing the data returned within a resource dump/fast sync zip file.
// The structure is defined https://core.ac.uk/services/fastsync#structure
type FSArticle struct{
	DOI string `json:"doi"`
	CoreID string `json:"coreId"`
	Identifiers []string `json:"identifiers"` // additional identifiers, we get no further information on what they are
	Title string `json:"title"`
	Authors []string `json:"authors"`
	Enrichments FSArticleEnrichment `json:"enrichments"`
	Contributors []string `json:"contributors"`
	DatePublished string `json:"datePublished"`
	Abstract string `json:"abstract"`
	DownloadURL string `json:"downloadUrl"`
	FullTextIdentifier string `json:"fullTextIdentifier"`
	PDFHashValue string `json:"pdfHashValue"`
	Publisher string `json:"publisher"`
	RawRecordXML string `json:"rawRecordXML"`
	Journal FSJournal `json:"journals"`
	Language FSLanguage `json:"language"`
	Relations []string `json:"relations"`
	Year int `json:"year"`
	Topics []string `json:"topics"`
	Subjects []string `json:"subjects"`
	URLs []string `json:"urls"`
	FullText string `json:"fullText"`
}

// FSArticleEnrichment adds extra details to the article
type FSArticleEnrichment struct {
	References []string `json:"references"`
	DocType FSDocType `json:"documentType"`
}

// FSDocType details the type of document being handled and the confidence CORE has in it's accuracy
type FSDocType struct{
	Type string `json:"type"`
	Confidence float32 `json:"confidence"`
}

// FSLanguage holds the basic language string, the ISO 2-letter code and a CORE specific int value representing the language
type FSLanguage struct {
	Code string `json:"code"`
	Name string `json:"name"`
	ID int `json:"id"`
}

// FSJournal holds the journal title and lit os identifiers, typically ISSN
type FSJournal struct {
	Title string `json:"title"`
	Identifiers []string `json:"identifiers"`
}

// ExtractFSArticle is a convenience method around unmarshaling the article metadata returned in the CORE
// resourcedump. There is no fetching done here, and no further processing. Rather the Go struct is returned for
// further use by the consumer.
func (ce *Extractor) ExtractFSArticle(rawData []byte) (*FSArticle, error) {
	res := &FSArticle{}
	err := json.Unmarshal(rawData, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// String implements the Stringer interface to ensure consistent printing of the extracted fastsync article metadata.
// This intentionally does not print all the data, but rather the items considered pertinent
func (fs *FSArticle) String() string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "CORE ID: %s\n", fs.CoreID)
	if fs.Title != "" {
		fmt.Fprintf(sb, "Title: %s\n", fs.DownloadURL)
	}
	if len(fs.Authors) > 0{
		fmt.Fprintf(sb, "Authors: %s\n", strings.Join(fs.Authors, ","))
	}
	if fs.Publisher != "" {
		fmt.Fprintf(sb, "Published By: %s\n", fs.Publisher)
	}
	if fs.DownloadURL != "" {
		fmt.Fprintf(sb, "Download From: %s\n", fs.DownloadURL)
	}
	return strings.TrimSpace(sb.String())
}