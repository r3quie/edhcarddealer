package edhcarddealer

func GetInfo(name string) CardInfo {
	for _, card := range ParsedCardsInfo {
		if card.Name == name {
			return card
		}
	}
	return CardInfo{}
}
