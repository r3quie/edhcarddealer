package edhcarddealer

import (
	"bufio"
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
func DownloadOracleCards(path string, typ string) {
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
		if d.Type == typ {
			res, err := http.Get(d.DownloadURI)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			io.Copy(out, res.Body)
		}
	}
}

func ParseCards[T Cards | CardsInfo](path string) T {
	// Open the file
	var f []byte
	for {
		ff, err := os.ReadFile(path)
		if err != nil {
			DownloadOracleCards(path, "oracle_cards")
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

var PathToCards = "cache/oracle-cards.json"

var PathToAllCards = initAllCards()

var ParsedCards = ParseCards[Cards](PathToCards)

var ParsedCardsInfo = ParseCards[CardsInfo](PathToCards)

func initAllCards() string {
	_, err := os.Stat("cache/all_cards")
	if err != nil {
		os.Mkdir("cache/all_cards", 0755)
		ParseAllCards("cache/all_cards/")
	}
	return "cache/all_cards/"
}

func ParseAllCards(path string) {
	file, err := os.Open("cache/bulk-all-cards.json")
	if err != nil {
		DownloadOracleCards("cache/bulk-all-cards.json", "all_cards")
		file, err = os.Open("cache/bulk-all-cards.json")
	}
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const maxCapacity int = 1000000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for scanner.Scan() {
		var card Card
		if len(scanner.Text()) < 4 {
			continue
		}
		if err := json.Unmarshal(scanner.Bytes()[:len(scanner.Bytes())-1], &card); err != nil {
			log.Println("continueing", err.Error())
			continue
		}
		outFile, err := os.Create(path + card.ID + ".json")
		if err != nil {
			panic(err)
		}
		defer outFile.Close()
		b, _ := json.Marshal(card)
		outFile.Write(b)
	}

	if err := scanner.Err(); err != nil {
		return
	}
}
