package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

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
	// PLEASE TEST IF Sscanf FASTER THAN REGEXP
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
	hand := deck.DealHand()

	var result Result

	for _, card := range hand {
		if card.ProducedMana == nil {
			result.Non++
			continue
		}
		if strings.Contains(card.OracleText, "dd one mana of any color in your commander's color identity") {
			result.Count(deck.ColorIdentity)
			continue
		}
		result.Count(card.ProducedMana)
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
