package edhcarddealer

import (
	"math/rand"
	"slices"
)

type Card struct {
	Name          string `json:"name"`
	ProducedMana  []any  `json:"produced_mana"`
	ColorIdentity []any  `json:"color_identity"`
	OracleText    string `json:"oracle_text"`
	ImgUris       struct {
		Small  string `json:"small"`
		Normal string `json:"normal"`
		Large  string `json:"large"`
	} `json:"image_uris"`
	ScryUri  string `json:"scryfall_uri"`
	TypeLine string `json:"type_line"`
	ID       string `json:"id"`
}

type Cards []Card

type Deck struct {
	Cards         Cards
	ColorIdentity []any
}

// Riffle shuffles the deck
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

// Names returns the names of the cards
func (c Cards) Names() []string {
	n := make([]string, len(c))
	for i, x := range c {
		n[i] = x.Name
	}
	return n
}

func reverse(arr []Card) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// PutHandOnBottom puts the top 7 cards on the bottom of the deck
func (c *Cards) PutHandOnBottom() {
	*c = append((*c)[len(*c)-7:], (*c)[:len(*c)-7]...)
}

func (c *Cards) PileShuffle(n int) {
	piles := make([][]Card, n)

	// Distribute cards into piles
	for i := 0; i < len(*c); i++ {
		piles[i%n] = append(piles[i%n], (*c)[i])
	}
	for i := 0; i < n; i++ {
		reverse(piles[i])
	}
	shuffledDeck := make([]Card, 0, len(*c))
	for _, pile := range piles {
		shuffledDeck = append(shuffledDeck, pile...)
	}
	*c = shuffledDeck
}

// Contains checks if a card is in the slice
func (cs Cards) Contains(card Card) bool {
	for _, c := range cs {
		if c.Name == card.Name {
			return true
		}
	}
	return false
}

// DealHand deals a hand of 7 cards
func (deck Deck) DealHand() Cards {
	hand := make([]int, 7)
	h := make(Cards, 7)

	for i := range hand {
		card := rand.Intn(len(deck.Cards))
		for slices.Contains(hand, card) {
			card = rand.Intn(len(deck.Cards))
		}
		hand[i] = card
		h[i] = deck.Cards[card]
	}

	return h
}
