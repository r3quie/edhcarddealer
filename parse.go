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
		panic(err)
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

func ParseCards(path string) Cards {
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
	var cs Cards

	if err := json.Unmarshal(f, &cs); err != nil {
		log.Println(err)
	}
	return cs
}

var PathToCards = "cache/oracle-cards.json"

var ParsedCards = ParseCards(PathToCards)
