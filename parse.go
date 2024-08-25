package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// PLEASE TEST IF Sscanf FASTER THAN REGEXP

type Card struct {
	Name          string `json:"name"`
	ProducedMana  []any  `json:"produced_mana"`
	ColorIdentity []any  `json:"color_identity"`
	OracleText    string `json:"oracle_text"`
	ImgUris       struct {
		Small      string `json:"small"`
		Normal     string `json:"normal"`
		Large      string `json:"large"`
		Png        string `json:"png"`
		ArtCrop    string `json:"art_crop"`
		BorderCrop string `json:"border_crop"`
	} `json:"image_uris"`
	ScryUri string `json:"scryfall_uri"`
}

type Cards []Card

type Result struct {
	U, R, B, G, W, C, Non int
}

type Results struct {
	U, R, B, G, W, C, Non float64
}

var cards = parseCards("cache/oracle-cards.json")

// Add adds values of r2 to r
func (r *Result) Add(r2 Result) {
	r.U += r2.U
	r.R += r2.R
	r.B += r2.B
	r.G += r2.G
	r.W += r2.W
	r.C += r2.C
	r.Non += r2.Non
}

type Deck struct {
	Cards         Cards
	ColorIdentity []any
}

// Returns the average of the Result
func (r Result) Average(n int) Results {
	return Results{
		U:   float64(r.U) / float64(n),
		R:   float64(r.R) / float64(n),
		B:   float64(r.B) / float64(n),
		G:   float64(r.G) / float64(n),
		W:   float64(r.W) / float64(n),
		C:   float64(r.C) / float64(n),
		Non: float64(r.Non) / float64(n),
	}
}

func (r Results) String() string {
	return fmt.Sprintf("U: %.2f\nR: %.2f\nB: %.2f\nG: %.2f\nW: %.2f\nC: %.2f\nNon: %.2f\n", r.U, r.R, r.B, r.G, r.W, r.C, r.Non)
}

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

// takes a card name and returns the card struct from the cards slice
func getCard(cardName string) (Card, error) {
	for _, card := range cards {
		if card.Name == cardName {
			return card, nil
		}
	}
	return Card{}, fmt.Errorf("Card not found")
}

// takes a string of cards for MTGO and returns Cards
func getDeck(input string) Deck {

	lines := strings.Split(input, "\r\n")

	var deck Deck

	re := regexp.MustCompile(`^(\d+)\s+(.*)$`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) == 3 {
			count := match[1]
			cardName := match[2]

			countint, _ := strconv.Atoi(count)

			x, err := getCard(cardName)
			if err != nil {
				log.Println(err)
				continue
			}
			if len(deck.ColorIdentity) < len(x.ColorIdentity) {
				deck.ColorIdentity = x.ColorIdentity
			}
			for k := 0; k < countint; k++ {
				deck.Cards = append(deck.Cards, x)
			}
		}
	}

	return deck
}

func simdeal(deck Deck) Result {
	hand := make(Cards, 7)
	for i := range hand {
		hand[i] = deck.Cards[rand.Intn(len(deck.Cards))]
	}

	var result Result

	count := func(ProducedMana []any) {
		for _, color := range ProducedMana {
			switch color {
			case "U":
				result.U++
			case "R":
				result.R++
			case "B":
				result.B++
			case "G":
				result.G++
			case "W":
				result.W++
			case "C":
				result.C++
			}
		}
	}

	for _, card := range hand {
		if card.ProducedMana == nil {
			result.Non++
			continue
		}
		if strings.Contains(card.OracleText, "dd one mana of any color in your commander's color identity") {
			count(deck.ColorIdentity)
			continue
		}
		count(card.ProducedMana)
	}
	return result
}

func simulate(decklist string, n int) Results {

	deck := getDeck(decklist)

	if n > 1000000 {
		n = 1000000
	}
	if n < 1 {
		n = 1
	}

	var total Result

	for i := 0; i < n; i++ {
		total.Add(simdeal(deck))
	}

	avg_results := total.Average(n)

	return avg_results
}

func (c Cards) Riffle() {

	// using append 'cause you gotta love slices
	first := append(Cards{}, c[len(c)/2:]...)
	second := append(Cards{}, c[:len(c)/2]...)

	for i := 0; i < len(c); i++ {
		if i%2 == 0 {
			c[i] = first[i/2]
		} else {
			c[i] = second[i/2]
		}
	}
}

func (deck *Deck) PutHandOnBottom() {
	deck.Cards = append(deck.Cards[len(deck.Cards)-7:], deck.Cards[:len(deck.Cards)-7]...)
}

func (c Cards) Names() []string {
	n := make([]string, len(c))
	for i, x := range c {
		n[i] = x.Name
	}
	return n
}
