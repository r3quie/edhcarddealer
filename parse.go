package main

import (
	"math/rand"
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

var cards = parseCards("cache/oracle-cards.json")

type Cards []Card

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

func (c Cards) Names() []string {
	n := make([]string, len(c))
	for i, x := range c {
		n[i] = x.Name
	}
	return n
}

type Deck struct {
	Cards         Cards
	ColorIdentity []any
}

func (deck *Deck) PutHandOnBottom() {
	deck.Cards = append(deck.Cards[len(deck.Cards)-7:], deck.Cards[:len(deck.Cards)-7]...)
}

func (cs Cards) Contains(card Card) bool {
	for _, c := range cs {
		if c.Name == card.Name {
			return true
		}
	}
	return false
}

func (deck Deck) DealHand() Cards {
	hand := make(Cards, 7)

	for i := range hand {
		card := deck.Cards[rand.Intn(len(deck.Cards))]
		for hand.Contains(card) {
			card = deck.Cards[rand.Intn(len(deck.Cards))]
		}
		hand[i] = card
	}

	return hand
}
