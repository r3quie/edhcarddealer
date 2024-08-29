package edhcarddealer

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// takes a card name and returns the card struct from the cards slice

type InputString struct {
	Input string
	Typ   string
}

func GetCard(cardName string) (Card, error) {
	for _, card := range ParsedCards {
		if card.Name == cardName {
			return card, nil
		}
	}
	return Card{}, fmt.Errorf("Card not found")
}

// takes a string of cards for MTGO and returns Cards

func InputToLines(i InputString) []string {
	var lines []string

	if i.Typ == "urlencoded" {
		if ii, err := url.QueryUnescape(i.Input); err == nil {
			i.Input = ii
		}
	}

	if strings.Contains(i.Input, "\r\n") {
		lines = strings.Split(i.Input, "\r\n")
	} else if !strings.Contains(i.Input, "\n") {
		lines = regexp.MustCompile(`(\d+\s+\w+\s+\w+)`).FindAllString(i.Input, -1)
		log.Println(lines)
	} else {
		lines = strings.Split(i.Input, "\n")
	}
	return lines
}

func GetDeck(decklist InputString) Deck {
	lines := InputToLines(decklist)
	var deck Deck
	// PLEASE TEST IF Sscanf FASTER THAN REGEXP
	re := regexp.MustCompile(`^(\d+)\s+(.*)$`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) == 3 {
			count := match[1]
			cardName := match[2]

			countint, _ := strconv.Atoi(count)

			x, err := GetCard(cardName)
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

func Simdeal(deck Deck) Result {
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

func Simulate(decklist InputString, n int) Results {

	deck := GetDeck(decklist)

	if n > 1000000 {
		n = 1000000
	}
	if n < 1 {
		n = 1
	}

	var total Result

	for i := 0; i < n; i++ {
		total.Add(Simdeal(deck))
	}

	return total.Average(n)
}
