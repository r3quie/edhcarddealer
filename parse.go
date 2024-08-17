package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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

var cards = parseCards()

/*
	type DCard struct {
		Card  Card
		Count int
	}

type Deck []DCard
*/
func parseCards() Cards {
	// Open the file
	jsonFile, err := os.Open("cache/oracle-cards.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var cs Cards

	if err := json.Unmarshal(byteValue, &cs); err != nil {
		log.Println(err)
	}
	return cs
}

func getValue(cardName string) (Card, error) {
	for _, card := range cards {
		if card.Name == cardName {
			return card, nil
		}
	}
	return Card{}, fmt.Errorf("Card not found")
}

func getDeck(input string) Cards {
	lines := strings.Split(input, "\r\n")

	var deck Cards

	re := regexp.MustCompile(`^(\d+)\s+(.*)$`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) == 3 {
			count := match[1]
			cardName := match[2]
			var countint int
			if x, err := strconv.Atoi(count); err != nil {
				countint = 1
			} else {
				countint = x
			}
			x, err := getValue(cardName)
			if err != nil {
				log.Println(err)
				deck = deck[:len(deck)-countint]
				continue
			}
			for k := 0; k < countint; k++ {
				deck = append(deck, x)
			}
		}
	}

	return deck
}

func simdeal(deck Cards) struct{ U, R, B, G, W, C, Non int } {
	hand := make(Cards, 7)
	for i := range hand {
		hand[i] = deck[rand.Intn(len(deck))]
	}

	blue, red, black, green, white, cless, non := 0, 0, 0, 0, 0, 0, 0
	for _, card := range hand {
		if card.ProducedMana == nil {
			non++
			continue
		}
		for _, color := range card.ProducedMana {
			switch color {
			case "U":
				blue++
			case "R":
				red++
			case "B":
				black++
			case "G":
				green++
			case "W":
				white++
			case "C":
				cless++
			}
		}
	}
	return struct {
		U   int
		R   int
		B   int
		G   int
		W   int
		C   int
		Non int
	}{blue, red, black, green, white, cless, non}
}

func simulate(deck Cards, n int) struct{ U, R, B, G, W, C, Non float64 } {

	total := struct{ U, R, B, G, W, C, Non int }{}
	avg_results := struct{ U, R, B, G, W, C, Non float64 }{}

	for i := 0; i < n; i++ {
		results := simdeal(deck)
		total.U += results.U
		total.R += results.R
		total.B += results.B
		total.G += results.G
		total.W += results.W
		total.C += results.C
		total.Non += results.Non
	}

	avg_results.U = float64(total.U) / float64(n)
	avg_results.R = float64(total.R) / float64(n)
	avg_results.B = float64(total.B) / float64(n)
	avg_results.G = float64(total.G) / float64(n)
	avg_results.W = float64(total.W) / float64(n)
	avg_results.C = float64(total.C) / float64(n)
	avg_results.Non = float64(total.Non) / float64(n)

	return avg_results
}

func main() {
	x := time.Now()
	deck := getDeck(`1 Aarakocra Sneak` + "\r" + `
1 Access Tunnel` + "\r" + `
1 Aether Tunnel` + "\r" + `
1 Ancestral Vision` + "\r" + `
1 Aqueous Form` + "\r" + `
1 As Foretold` + "\r" + `
1 Baleful Mastery` + "\r" + `
1 Blackblade Reforged` + "\r" + `
1 Blasphemous Act` + "\r" + `
1 Bojuka Bog` + "\r" + `
1 Braids, Conjurer Adept` + "\r" + `
1 Cascade Bluffs` + "\r" + `
1 Caves of Chaos Adventurer` + "\r" + `
1 Chaos Warp` + "\r" + `
1 Command Tower` + "\r" + `
1 Counterspell` + "\r" + `
1 Court of Ire` + "\r" + `
1 Creeping Bloodsucker` + "\r" + `
1 Crumbling Necropolis` + "\r" + `
1 Darkwater Catacombs` + "\r" + `
1 Defabricate` + "\r" + `
1 Descent into Avernus` + "\r" + `
1 Dimir Signet` + "\r" + `
1 Dragonmaster Outcast` + "\r" + `
1 Dragonskull Summit` + "\r" + `
1 Drowned Catacomb` + "\r" + `
1 Endless Evil` + "\r" + `
1 Exotic Orchard` + "\r" + `
1 Feed the Swarm` + "\r" + `
1 Feywild Caretaker` + "\r" + `
1 Gate to the Aether` + "\r" + `
1 Indulgent Tormentor` + "\r" + `
1 Infernal Grasp` + "\r" + `
1 Inspiring Refrain` + "\r" + `
6 Island` + "\r" + `
1 Izzet Signet` + "\r" + `
1 Kumena's Awakening` + "\r" + `
1 Lizard Blades` + "\r" + `
1 Mechanized Production` + "\r" + `
1 Midnight Clock` + "\r" + `
4 Mountain` + "\r" + `
1 Negate` + "\r" + `
1 Nightscape Familiar` + "\r" + `
1 Palace Siege` + "\r" + `
1 Passageway Seer` + "\r" + `
1 Path of Ancestry` + "\r" + `
1 Phyrexian Arena` + "\r" + `
1 Plargg and Nassari` + "\r" + `
1 Profane Tutor` + "\r" + `
1 Protection Racket` + "\r" + `
1 Rakdos Signet` + "\r" + `
1 Ravenloft Adventurer` + "\r" + `
1 Reliquary Tower` + "\r" + `
1 Replicating Ring` + "\r" + `
1 Rilsa Rael, Kingpin` + "\r" + `
1 Ring of Evos Isle` + "\r" + `
1 Ring of Valkas` + "\r" + `
1 Ring of Xathrid` + "\r" + `
1 Rogue's Passage` + "\r" + `
1 Rousing Refrain` + "\r" + `
1 Shivan Reef` + "\r" + `
1 Skyline Despot` + "\r" + `
1 Smoldering Marsh` + "\r" + `
1 Sphinx of the Second Sun` + "\r" + `
1 Star Whale` + "\r" + `
1 Stirring Bard` + "\r" + `
1 Stolen Strategy` + "\r" + `
1 Sulfur Falls` + "\r" + `
1 Sulfurous Springs` + "\r" + `
1 Sunken Hollow` + "\r" + `
2 Swamp` + "\r" + `
1 Swiftfoot Boots` + "\r" + `
1 Sword Coast Sailor` + "\r" + `
1 Talisman of Creativity` + "\r" + `
1 Talisman of Dominance` + "\r" + `
1 Talisman of Indulgence` + "\r" + `
1 Tavern Brawler` + "\r" + `
1 Temple of Deceit` + "\r" + `
1 Temple of Epiphany` + "\r" + `
1 Temple of Malice` + "\r" + `
1 Thassa, God of the Sea` + "\r" + `
1 The Ninth Doctor` + "\r" + `
1 The Tenth Doctor` + "\r" + `
1 Thopter Spy Network` + "\r" + `
1 Tomb of Horrors Adventurer` + "\r" + `
1 Twilight Prophet` + "\r" + `
1 Underground River` + "\r" + `
1 Vandalblast` + "\r" + `
1 Wheel of Fate` + "\r" + `
1 Whispersilk Cloak` + "\r" + `
1 Obeka, Splitter of Seconds`)

	result := simulate(deck, 1000000)
	fmt.Printf("Average Blue: %f\n", result.U)
	fmt.Printf("Average Red: %f\n", result.R)
	fmt.Printf("Average Black: %f\n", result.B)
	fmt.Printf("Average Green: %f\n", result.G)
	fmt.Printf("Average White: %f\n", result.W)
	fmt.Printf("Average Colorless: %f\n", result.C)
	fmt.Printf("Average Generating cards: %f\n", 7-result.Non)

	fmt.Println(time.Since(x))
}
