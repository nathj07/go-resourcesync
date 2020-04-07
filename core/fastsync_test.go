package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractFSArticle(t *testing.T) {
	ce := &Extractor{}
	got, err := ce.ExtractFSArticle(testFSData)
	require.Nil(t, err)
	assert.Equal(t, expFSArticle, got)
}

func TestFSArticleString(t *testing.T) {
	assert.Equal(t, expFSArticleString, fmt.Sprintf("%s", expFSArticle))
}

//
// Test data and expected outputs
//

var expFSArticle = &FSArticle{
	DOI:    "",
	CoreID: "42138760",
	OAIID:  "oai:clok.uclan.ac.uk:14639",
	MAGID:  "2144009174",
	Identifiers: []string{
		"oai:clok.uclan.ac.uk:14639",
		"",
	},
	Title: "Research Report: 'Standing on my own two feet': Disadvantaged Teenagers, Intimate Partner Violence and Coercive Control",
	Authors: []string{
		"Wood, Marsha",
		"Barter, Christine",
		"Berridge, David",
	},
	Enrichments: FSArticleEnrichment{
		References: []FSReference{
			{
				ID:      36385704,
				Title:   "Couldn't Have Asked for More: Lessons Learned in Information Dissemination.Technologies for Primary Health Care (PRITECH), Occasional Operations Papers. Management Sciences for Health,Arlington,Virginia.",
				Authors: []string{},
				Date:    "1993",
				DOI:     "10.1056/test-doi",
				Raw:     "White, K (1993):We Couldn't Have Asked for More: Lessons Learned in Information Dissemination.Technologies for Primary Health Care (PRITECH), Occasional Operations Papers. Management Sciences for Health,Arlington,Virginia.",
			},
			{
				ID:      36385691,
				Title:   "Development,Africa Bureau, Office of Sustainable Development.Washington,",
				Authors: []string{},
				Date:    "1997",
				Raw:     "U.S Agency for International Development,Africa Bureau, Office of Sustainable Development.Washington, D.C. RAND (1997): Population Matters:A RAND Program of Policy-relevant Research Communication. RAND, Santa Monica, California, USA.",
			},
		},
		DocType: FSDocType{
			Type:       "research",
			Confidence: 1,
		},
		CitationCount: 2,
	},
	Contributors:       []string{},
	DatePublished:      "2010",
	Abstract:           "",
	DownloadURL:        "https://core.ac.uk/download/pdf/42138760.pdf",
	FullTextIdentifier: "",
	PDFHashValue:       "97a86466f1afdd62cd885ef03fc75c483bef767d",
	Publisher:          "NSPCC",
	RawRecordXML:       "<record><header><identifier>\n    \n    \n      oai:clok.uclan.ac.uk:14639</identifier><datestamp>\n      2017-10-11T13:40:45Z</datestamp><setSpec>\n      7374617475733D707562</setSpec><setSpec>\n      7375626A656374733D4C303030:4C353030</setSpec><setSpec>\n      74797065733D6D6F6E6F6772617068</setSpec></header><metadata><rioxx xmlns=\"http://www.rioxx.net/schema/v2.0/rioxx/\" xmlns:ali=\"http://ali.niso.org/2014/ali/1.0\" xmlns:dc=\"http://purl.org/dc/elements/1.1/\" xmlns:dcterms=\"http://purl.org/dc/terms/\" xmlns:rioxxterms=\"http://docs.rioxx.net/schema/v2.0/rioxxterms/\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:schemaLocation=\"http://www.rioxx.net/schema/v2.0/rioxx/ http://www.rioxx.net/schema/v2.0/rioxx/rioxx.xsd\" ><ali:free_to_read>\n    \n      </ali:free_to_read><ali:license_ref start_date=\"2010\" >http://creativecommons.org/licenses/by-nc-nd/3.0</ali:license_ref><dc:format>application/pdf</dc:format><dc:identifier>http://clok.uclan.ac.uk/14639/1/standing-own-two-feet-report.pdf</dc:identifier><dc:language>en</dc:language><dc:publisher>NSPCC</dc:publisher><dc:subject>L500</dc:subject><dc:title>Research Report: 'Standing on my own two feet': Disadvantaged Teenagers, Intimate Partner Violence and Coercive Control</dc:title><rioxxterms:author>Wood, Marsha</rioxxterms:author><rioxxterms:author>Barter, Christine</rioxxterms:author><rioxxterms:author>Berridge, David</rioxxterms:author><rioxxterms:publication_date>2010</rioxxterms:publication_date><rioxxterms:type>Monograph</rioxxterms:type><rioxxterms:version>VoR</rioxxterms:version></rioxx></metadata></record>",
	Journals: []FSJournal{
		{
			Title: "",
			Identifiers: []string{
				"issn:0143-7739",
				"0143-7739",
			},
		},
	},
	Language: FSLanguage{
		Code: "en",
		Name: "English",
		ID:   9,
	},
	Relations: []string{},
	Year:      2010,
	Topics: []string{
		"L500",
	},
	URLs: []string{
		"http://example.com/1",
		"http://www.example.com/2",
	},
	FullText: "RESEARCH REPORT\n'Standing on my own two feet':\nDisadvantaged Teenagers, Intimate Partner Violence  \nand Coercive Control\n",
	ISSN:     "1077-8306",
}

var testFSData = []byte(`{
  "doi": null,
  "coreId": "42138760",
  "oai": "oai:clok.uclan.ac.uk:14639",
  "magID": "2144009174",
  "identifiers": [
    "oai:clok.uclan.ac.uk:14639",
    null
  ],
  "title": "Research Report: 'Standing on my own two feet': Disadvantaged Teenagers, Intimate Partner Violence and Coercive Control",
  "authors": [
    "Wood, Marsha",
    "Barter, Christine",
    "Berridge, David"
  ],
  "enrichments": {
    "references": [
      {
        "id": 36385704,
        "title": "Couldn't Have Asked for More: Lessons Learned in Information Dissemination.Technologies for Primary Health Care (PRITECH), Occasional Operations Papers. Management Sciences for Health,Arlington,Virginia.",
        "authors": [],
        "date": "1993",
        "doi": "10.1056/test-doi",
        "raw": "White, K (1993):We Couldn't Have Asked for More: Lessons Learned in Information Dissemination.Technologies for Primary Health Care (PRITECH), Occasional Operations Papers. Management Sciences for Health,Arlington,Virginia.",
        "cites": null
      },
      {
        "id": 36385691,
        "title": "Development,Africa Bureau, Office of Sustainable Development.Washington,",
        "authors": [],
        "date": "1997",
        "doi": null,
        "raw": "U.S Agency for International Development,Africa Bureau, Office of Sustainable Development.Washington, D.C. RAND (1997): Population Matters:A RAND Program of Policy-relevant Research Communication. RAND, Santa Monica, California, USA.",
        "cites": null
      }
    ],
    "documentType": {
      "type": "research",
      "confidence": 1
    },
    "citationCount": 2
  },
  "contributors": [],
  "datePublished": "2010",
  "abstract": null,
  "downloadUrl": "https://core.ac.uk/download/pdf/42138760.pdf",
  "fullTextIdentifier": null,
  "pdfHashValue": "97a86466f1afdd62cd885ef03fc75c483bef767d",
  "publisher": "NSPCC",
  "rawRecordXml": "<record><header><identifier>\n    \n    \n      oai:clok.uclan.ac.uk:14639</identifier><datestamp>\n      2017-10-11T13:40:45Z</datestamp><setSpec>\n      7374617475733D707562</setSpec><setSpec>\n      7375626A656374733D4C303030:4C353030</setSpec><setSpec>\n      74797065733D6D6F6E6F6772617068</setSpec></header><metadata><rioxx xmlns=\"http://www.rioxx.net/schema/v2.0/rioxx/\" xmlns:ali=\"http://ali.niso.org/2014/ali/1.0\" xmlns:dc=\"http://purl.org/dc/elements/1.1/\" xmlns:dcterms=\"http://purl.org/dc/terms/\" xmlns:rioxxterms=\"http://docs.rioxx.net/schema/v2.0/rioxxterms/\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:schemaLocation=\"http://www.rioxx.net/schema/v2.0/rioxx/ http://www.rioxx.net/schema/v2.0/rioxx/rioxx.xsd\" ><ali:free_to_read>\n    \n      </ali:free_to_read><ali:license_ref start_date=\"2010\" >http://creativecommons.org/licenses/by-nc-nd/3.0</ali:license_ref><dc:format>application/pdf</dc:format><dc:identifier>http://clok.uclan.ac.uk/14639/1/standing-own-two-feet-report.pdf</dc:identifier><dc:language>en</dc:language><dc:publisher>NSPCC</dc:publisher><dc:subject>L500</dc:subject><dc:title>Research Report: 'Standing on my own two feet': Disadvantaged Teenagers, Intimate Partner Violence and Coercive Control</dc:title><rioxxterms:author>Wood, Marsha</rioxxterms:author><rioxxterms:author>Barter, Christine</rioxxterms:author><rioxxterms:author>Berridge, David</rioxxterms:author><rioxxterms:publication_date>2010</rioxxterms:publication_date><rioxxterms:type>Monograph</rioxxterms:type><rioxxterms:version>VoR</rioxxterms:version></rioxx></metadata></record>",
  "journals": [
    {
      "title": null,
      "identifiers": [
        "issn:0143-7739",
        "0143-7739"
      ]
    }
  ],
  "language": {
    "code": "en",
    "id": 9,
    "name": "English"
  },
  "relations": [],
  "year": 2010,
  "topics": [
    "L500"
  ],
  "subject": [
    "L500"
  ],
  "urls": [
    "http://example.com/1",
    "http://www.example.com/2"
  ],
  "repositories": {
    "id": "18",
    "name": "CLoK",
    "openDoarId": 0,
    "roarId": 0
  },
  "fullText": "RESEARCH REPORT\n'Standing on my own two feet':\nDisadvantaged Teenagers, Intimate Partner Violence  \nand Coercive Control\n",
  "issn": "1077-8306"
}`)

var expFSArticleString = `CORE ID: 42138760
Title: https://core.ac.uk/download/pdf/42138760.pdf
Authors: Wood, Marsha,Barter, Christine,Berridge, David
Published By: NSPCC
Download From: https://core.ac.uk/download/pdf/42138760.pdf`
