package bubblegum_api

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
)

type Card struct {
	CardID       int    `json:"id"`
	Year         string `json:"year"`
	Manufacturer string `json:"manufacturer"`
	Player       string `json:"player"`
	Series       string `json:"series"`
	Card_number  string `json:"card_number"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	Parallel     string `json:"parallel"`
	Image        string `json:"image"`
}

type Manufacturer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	YearFounded  string `json:"year_founded"`
	YearDefunct  string `json:"year_defuct"`
	Fate         string `json:"fate"`
	Headquarters string `json:"headquarters"`
	Website      string `json:"website"`
	Revenue      string `json:"revenue"`
	Employees    string `json:"employees"`
	Address      string `json:"address"`
}

type Player struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Position  string `json:"position"`
	BirthDate string `json:"birth_date"`
	DeathDate string `json:"death_date"`
	Throws    string `json:"throws"`
	Bats      string `json:"bats"`
	Height    string `json:"height"`
	Weight    string `json:"weight"`
}

type Team struct {
	ID              int    `json:"id"`
	Location        string `json:"location"`
	Team            string `json:"team"`
	League          string `json:"league"`
	LeagueLevel     string `json:"league_level"`
	YearEstablished string `json:"year_established"`
	YearDefunct     string `json:"year_defunct"`
}

var cards = []Card{}
var manufacturers = []Manufacturer{}
var players = []Player{}
var teams = []Team{}
var categories = make(map[string]bool)

func LoadData() {
	pullCards()
	pullManufacturers()
	pullPlayers()
	pullTeams()
}

func pullManufacturers() {
	SHEET_ID := "1g1E_k1V1VuHXCwT0sZUEMjg-4pUmUzaOkSkCOO5PkFc"
	SHEET_NAME := "manufacturers"
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
		manufacturer := Manufacturer{
			ID:           index,
			Name:         row[0],
			YearFounded:  row[1],
			YearDefunct:  row[2],
			Fate:         row[3],
			Headquarters: row[4],
			Website:      row[5],
			Revenue:      row[6],
			Employees:    row[7],
			Address:      row[8],
		}
		manufacturers = append(manufacturers, manufacturer)
	}
}

func pullTeams() {
	SHEET_ID := "1g1E_k1V1VuHXCwT0sZUEMjg-4pUmUzaOkSkCOO5PkFc"
	SHEET_NAME := "teams"
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
		team := Team{
			ID:              index,
			Location:        row[0],
			Team:            row[1],
			League:          row[2],
			LeagueLevel:     row[3],
			YearEstablished: row[4],
			YearDefunct:     row[5],
		}
		teams = append(teams, team)
	}
}

func pullPlayers() {
	SHEET_ID := "1g1E_k1V1VuHXCwT0sZUEMjg-4pUmUzaOkSkCOO5PkFc"
	SHEET_NAME := "players"
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
		player := Player{
			ID:        index,
			FirstName: row[0],
			LastName:  row[1],
			Position:  row[2],
			BirthDate: row[3],
			DeathDate: row[4],
			Throws:    row[5],
			Bats:      row[6],
			Height:    row[7],
			Weight:    row[8],
		}
		players = append(players, player)
	}
}

func pullCards() {

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
