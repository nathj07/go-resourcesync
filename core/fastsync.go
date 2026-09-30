package core

import (
	"encoding/json"
	"fmt"
	"strings"
)

// FSArticle is a struct representing the data returned within a resource dump/fast sync zip file.
// The structure is defined https://core.ac.uk/services/fastsync#structure
type FSArticle struct {
	DOI                string              `json:"doi"`
	CoreID             string              `json:"coreId"`
	OAIID              string              `json:"oai"`
	MAGID              string              `json:"magId"`
	Identifiers        []string            `json:"identifiers"` // additional identifiers, we get no further information on what they are
	Title              string              `json:"title"`
	Authors            []string            `json:"authors"`
	Enrichments        FSArticleEnrichment `json:"enrichments"`
	Contributors       []string            `json:"contributors"`
	DatePublished      string              `json:"datePublished"`
	Abstract           string              `json:"abstract"`
	DownloadURL        string              `json:"downloadUrl"`
	FullTextIdentifier string              `json:"fullTextIdentifier"`
	PDFHashValue       string              `json:"pdfHashValue"`
	Publisher          string              `json:"publisher"`
	RawRecordXML       string              `json:"rawRecordXml"`
	Journals           []FSJournal         `json:"journals"`
	Language           FSLanguage          `json:"language"`
	Relations          []string            `json:"relations"`
	Year               int                 `json:"year"`
	Topics             []string            `json:"topics"`
	Subjects           []string            `json:"subjects"`
	URLs               []string            `json:"urls"`
	FullText           string              `json:"fullText"`
	ISSN               string              `json:"issn"`
}

// FSArticleEnrichment adds extra details to the article
type FSArticleEnrichment struct {
	References    []FSReference `json:"references"`
	DocType       FSDocType     `json:"documentType"`
	CitationCount int           `json:"citationCount"`
}

// FSReference holds the enrichment reference data
type FSReference struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Date    string   `json:"date"`
	DOI     string   `json:"doi"`
	Raw     string   `json:"raw"`
	Cites   []int    `json:"cites"`
}

// FSDocType details the type(s) of document being handled and the confidence CORE has in its accuracy.
// CORE returns "type" as either a single string or an array of strings; both are normalised into Type.
type FSDocType struct {
	Type       []string `json:"type"`
	Confidence float32  `json:"confidence"`
}

// UnmarshalJSON allows documentType.type to be decoded whether CORE returns it as a
// single JSON string, an array of strings, or null. Everything is normalised into Type.
func (d *FSDocType) UnmarshalJSON(data []byte) error {
	// Capture type as raw so we can inspect its shape after decoding the rest.
	var aux struct {
		Type       json.RawMessage `json:"type"`
		Confidence float32         `json:"confidence"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	d.Confidence = aux.Confidence

	if len(aux.Type) == 0 || string(aux.Type) == "null" {
		d.Type = nil
		return nil
	}

	// An array of strings decodes directly; a scalar string will fail here and fall through.
	var arr []string
	if err := json.Unmarshal(aux.Type, &arr); err == nil {
		d.Type = arr
		return nil
	}

	var single string
	if err := json.Unmarshal(aux.Type, &single); err != nil {
		return err
	}
	d.Type = []string{single}
	return nil
}

// FSLanguage holds the basic language string, the ISO 2-letter code and a CORE specific int value representing the language
type FSLanguage struct {
	Code string `json:"code"`
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// FSJournal holds the journal title and lit os identifiers, typically ISSN
type FSJournal struct {
	Title       string   `json:"title"`
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
		fmt.Fprintf(sb, "Title: %s\n", fs.Title)
	}
	if len(fs.Authors) > 0 {
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
