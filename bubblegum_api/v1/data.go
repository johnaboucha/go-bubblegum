package bubblegum_api

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
)

type Card struct {
	CardID       int
	Year         string
	Manufacturer string
	Player       string
	Series       string
	Card_number  string
	Description  string
	Category     string
	Parallel     string
	Image        string
}

var cards = []Card{}
var categories = make(map[string]bool)

func GetGoogleSheet() {

	SHEET_ID := "1g1E_k1V1VuHXCwT0sZUEMjg-4pUmUzaOkSkCOO5PkFc"
	SHEET_NAME := "cards"
	url := "https://docs.google.com/spreadsheets/d/" + SHEET_ID + "/gviz/tq?tqx=out:csv&sheet=" + SHEET_NAME

	data, err := readCSVFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	for index, row := range data {
		// skip header
		if index == 0 {
			continue
		}
		card := Card{
			CardID:       index,
			Year:         row[0],
			Manufacturer: row[1],
			Player:       row[2],
			Series:       row[3],
			Card_number:  row[4],
			Description:  row[5],
			Category:     row[6],
			Parallel:     row[7],
			Image:        row[8],
		}
		cards = append(cards, card)
		categories[row[6]] = true
	}

}

func readCSVFromURL(url string) ([][]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Could not get CSV from URL: %v", err)
	}

	defer response.Body.Close()
	reader := csv.NewReader(response.Body)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Could not read CSV file: %v", err)
	}

	return data, nil
}
