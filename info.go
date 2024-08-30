package edhcarddealer

// Returns the legalities of a card in CardsInfo
func (cc CardsInfo) Legalities(name string) any {
	for _, Card := range ParsedCardsInfo {
		if Card.Name == name {
			return Card.Legalities
		}
	}
	return nil
}

// Returns the multiverse IDs of a card in CardsInfo
func (cc CardsInfo) MultiverseIds(name string) []any {
	for _, Card := range ParsedCardsInfo {
		if Card.Name == name {
			return Card.MultiverseIds
		}
	}
	return nil
}

// Returns the colors of a card in CardsInfo
func (cc CardsInfo) Colors(name string) []any {
	for _, Card := range ParsedCardsInfo {
		if Card.Name == name {
			return Card.Colors
		}
	}
	return nil
}

// Returns the color identity of a card in CardsInfo
func (cc CardsInfo) ColorIdentity(name string) []any {
	for _, Card := range ParsedCardsInfo {
		if Card.Name == name {
			return Card.ColorIdentity
		}
	}
	return nil
}

// Returns the keywords of a card in CardsInfo
func (cc CardsInfo) Keywords(name string) []any {
	for _, Card := range ParsedCardsInfo {
		if Card.Name == name {
			return Card.Keywords
		}
	}
	return nil
}

// Returns the games of a card in CardsInfo
func (cc CardsInfo) Games(name string) []string {
	for _, Card := range ParsedCardsInfo {
		if Card.Name == name {
			return Card.Games
		}
	}
	return nil
}

// Returns the finishes of a card in CardsInfo
func (cc CardsInfo) Finishes(name string) []string {
	for _, Card := range ParsedCardsInfo {
		if Card.Name == name {
			return Card.Finishes
		}
	}
	return nil
}

// Returns the prices of a card in CardsInfo
func (cc CardsInfo) Prices(name string) any {
	for _, Card := range ParsedCardsInfo {
		if Card.Name == name {
			return Card.Prices
		}
	}
	return nil
}

// Returns the related URIs of a card in CardsInfo
func (cc CardsInfo) RelatedUris(name string) any {
	for _, Card := range ParsedCardsInfo {
		if Card.Name == name {
			return Card.RelatedUris
		}
	}
	return nil
}

// Returns the purchase URIs of a card in CardsInfo
func (cc CardsInfo) PurchaseUris(name string) any {
	for _, Card := range ParsedCardsInfo {
		if Card.Name == name {
			return Card.PurchaseUris
		}
	}
	return nil
}

/*
// Returns the card faces of a card in CardsInfo
func (cc CardsInfo) CardFaces(name string) []any {
    for _, Card := range ParsedCardsInfo {
        if Card.Name == name {
            return Card.CardFaces
        }
    }
    return nil
}
*/
