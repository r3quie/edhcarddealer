package main

import (
	"encoding/json"
	"log"
	"os"
)

func parseCards(path string) Cards {
	// Open the file
	f, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var cs Cards

	if err := json.Unmarshal(f, &cs); err != nil {
		log.Println(err)
	}
	return cs
}

var pathToCards = "cache/oracle-cards.json"

var cards = parseCards(pathToCards)
