package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Card struct {
	Name          string `json:"name"`
	ProducedMana  string `json:"produced_mana"`
	ColorIdentity string `json:"color_identity"`
	OracleText    string `json:"oracle_text"`
	Img           string `json:"image_uris"`
	ScryUri       string `json:"scryfall_uri"`
}

type Cards struct {
	Cards []Card `json:"cards"`
}

func ParseCards() {
	// Open the file
	jsonFile, err := os.Open("cache/oracle-cards.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var cards Cards

	if err := json.Unmarshal(byteValue, &cards); err != nil {
		fmt.Println(err)
	}

	for _, x := range cards.Cards {
		if strings.Contains(x.OracleText, "draw a card") {
			fmt.Println(x.Name)
		}
	}
}

func main() {
	ParseCards()
}
