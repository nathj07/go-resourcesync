package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unicode/utf8"
)

func main() {
	resp, err := http.Get("https://core.ac.uk/api-v2/articles/get/59329824?metadata=true&fulltext=true&citations=false&similar=false&duplicate=false&urls=false&faithfulMetadata=false&apiKey=OUAxP4mifI19JKNDznLw7B6CgEoZthXs")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("unexpected status code: %d", resp.StatusCode))
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	c := &CoreAPIResp{}
	json.Unmarshal(b, c)
	if c.Status != "OK" {
		panic(fmt.Errorf("unexpect API status: %s", c.Status))
	}
	fmt.Printf("Text Valid UTF8 (bytes): %t\n", utf8.Valid([]byte(c.Data.Fulltext)))
	fmt.Printf("Text Valid UTF8 (string): %t\n", utf8.ValidString(c.Data.Fulltext))

}

type CoreAPIResp struct {
	Status string   `json:"status"`
	Data   CoreData `json:"data"`
}

type CoreData struct {
	Fulltext string `json:"fullText"`
}
