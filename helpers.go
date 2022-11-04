package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Cat struct {
	PermalinkName string
	PrettyName    string
}

func pullCategories(myHost string) []Cat {
	url := myHost + "/api/v1/categories/"

	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	cats := make(map[string][]string)
	json.Unmarshal(body, &cats)

	results := []Cat{}
	for _, cat := range cats["categories"] {
		pretty := strings.Replace(cat, "-", " ", -1)
		pretty = cases.Title(language.English).String(pretty)
		newCat := Cat{
			PermalinkName: cat,
			PrettyName:    pretty,
		}
		results = append(results, newCat)
	}
	return results
}
