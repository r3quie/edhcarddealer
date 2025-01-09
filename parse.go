package edhcarddealer

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type BulkData struct {
	Data []struct {
		Type        string `json:"type"`
		DownloadURI string `json:"download_uri"`
	} `json:"data"`
}

// DownloadOracleCards downloads the oracle cards from scryfall and saves them to cache/oracle-cards.json
func DownloadOracleCards(path string) {
	out, err := os.Create(path)
	if err != nil {
		if e := os.Mkdir("cache", 0755); e != nil {
			panic(e)
		}
		out, err = os.Create(path)
		if err != nil {
			panic(err)
		}
	}
	defer out.Close()

	res, err := http.Get("https://api.scryfall.com/bulk-data")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bd, _ := io.ReadAll(res.Body)

	var bulk BulkData
	if err := json.Unmarshal(bd, &bulk); err != nil {
		log.Println(err)
	}
	for _, d := range bulk.Data {
		if d.Type == "oracle_cards" {
			res, err := http.Get(d.DownloadURI)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			io.Copy(out, res.Body)
		}
	}
}

func DownloadAllCards(path string) {
	// Create a temporary file for downloading
	tmpPath := path + ".tmp"
	out, err := os.Create(tmpPath)
	if err != nil {
		if e := os.Mkdir("cache", 0755); e != nil {
			panic(e)
		}
		out, err = os.Create(tmpPath)
		if err != nil {
			panic(err)
		}
	}
	defer out.Close()

	// Get bulk data info
	res, err := http.Get("https://api.scryfall.com/bulk-data")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bd, _ := io.ReadAll(res.Body)

	var bulk BulkData
	if err := json.Unmarshal(bd, &bulk); err != nil {
		log.Println(err)
		return
	}

	// Download the file
	for _, d := range bulk.Data {
		if d.Type == "all_cards" {
			res, err = http.Get(d.DownloadURI)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			// Use buffered copy
			buf := make([]byte, 1024*1024) // 1MB buffer
			if _, err := io.CopyBuffer(out, res.Body, buf); err != nil {
				log.Printf("Error copying data: %v", err)
				return
			}
			break
		}
	}

	// Close the file before moving it
	out.Close()

	// Move the temporary file to the final location
	if err := os.Rename(tmpPath, path); err != nil {
		log.Printf("Error moving file: %v", err)
		return
	}
}

func ParseCards[T Cards | CardsInfo](path string) T {
	// Open the file
	var f []byte
	for {
		ff, err := os.ReadFile(path)
		if err != nil {
			DownloadOracleCards(path)
			continue
		}
		f = ff
		break
	}
	var cs T

	if err := json.Unmarshal(f, &cs); err != nil {
		log.Println(err)
	}
	return cs
}

func ParseAllCards[T Cards | CardsInfo](path string) T {
	var cs T

	// Open the file
	f, err := os.Open(path)
	if err != nil {
		DownloadAllCards(path)
		f, err = os.Open(path)
		if err != nil {
			log.Printf("Error opening file after download: %v", err)
			return cs
		}
	}
	defer f.Close()

	// Create a decoder for streaming JSON parsing
	decoder := json.NewDecoder(f)

	// Parse the JSON array
	if err := decoder.Decode(&cs); err != nil {
		log.Printf("Error decoding JSON: %v", err)
	}

	return cs
}

var PathToCards = "cache/oracle-cards.json"

var PathToAllCards = "cache/all-cards.json"

var ParsedCards = ParseCards[Cards](PathToCards)

var ParsedAllCards = ParseAllCards[Cards](PathToAllCards)
var ParsedAllCardsInfo = ParseAllCards[CardsInfo](PathToAllCards)

var ParsedCardsInfo = ParseCards[CardsInfo](PathToCards)
