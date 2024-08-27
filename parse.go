package edhcarddealer

import (
	"encoding/json"
	"log"
	"os"
)

func ParseCards(path string) Cards {
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

var PathToCards = "cache/oracle-cards.json"

var ParsedCards = ParseCards(PathToCards)
