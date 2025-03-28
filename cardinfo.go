package edhcarddealer

type CardsInfo []CardInfo

type CardInfo struct {
	Object        string `json:"object"`
	ID            string `json:"id"`
	OracleID      string `json:"oracle_id"`
	MultiverseIds []any  `json:"multiverse_ids"`
	TcgplayerID   int    `json:"tcgplayer_id,omitempty"`
	Name          string `json:"name"`
	Lang          string `json:"lang"`
	ReleasedAt    string `json:"released_at"`
	URI           string `json:"uri"`
	ScryfallURI   string `json:"scryfall_uri"`
	Layout        string `json:"layout"`
	HighresImage  bool   `json:"highres_image"`
	ImageStatus   string `json:"image_status"`
	ImageUris     struct {
		Small  string `json:"small"`
		Normal string `json:"normal"`
		Large  string `json:"large"`
	} `json:"image_uris,omitempty"`
	ManaCost      string  `json:"mana_cost,omitempty"`
	Cmc           float64 `json:"cmc"`
	TypeLine      string  `json:"type_line"`
	OracleText    string  `json:"oracle_text,omitempty"`
	Colors        []any   `json:"colors,omitempty"`
	ColorIdentity []any   `json:"color_identity"`
	Keywords      []any   `json:"keywords"`
	Legalities    struct {
		Standard        string `json:"standard"`
		Future          string `json:"future"`
		Historic        string `json:"historic"`
		Timeless        string `json:"timeless"`
		Gladiator       string `json:"gladiator"`
		Pioneer         string `json:"pioneer"`
		Explorer        string `json:"explorer"`
		Modern          string `json:"modern"`
		Legacy          string `json:"legacy"`
		Pauper          string `json:"pauper"`
		Vintage         string `json:"vintage"`
		Penny           string `json:"penny"`
		Commander       string `json:"commander"`
		Oathbreaker     string `json:"oathbreaker"`
		Standardbrawl   string `json:"standardbrawl"`
		Brawl           string `json:"brawl"`
		Alchemy         string `json:"alchemy"`
		Paupercommander string `json:"paupercommander"`
		Duel            string `json:"duel"`
		Oldschool       string `json:"oldschool"`
		Premodern       string `json:"premodern"`
		Predh           string `json:"predh"`
	} `json:"legalities"`
	Games           []string `json:"games"`
	Reserved        bool     `json:"reserved"`
	Foil            bool     `json:"foil"`
	Nonfoil         bool     `json:"nonfoil"`
	Finishes        []string `json:"finishes"`
	Oversized       bool     `json:"oversized"`
	Promo           bool     `json:"promo"`
	Reprint         bool     `json:"reprint"`
	Variation       bool     `json:"variation"`
	SetID           string   `json:"set_id"`
	Set             string   `json:"set"`
	SetName         string   `json:"set_name"`
	SetType         string   `json:"set_type"`
	SetURI          string   `json:"set_uri"`
	SetSearchURI    string   `json:"set_search_uri"`
	ScryfallSetURI  string   `json:"scryfall_set_uri"`
	RulingsURI      string   `json:"rulings_uri"`
	PrintsSearchURI string   `json:"prints_search_uri"`
	CollectorNumber string   `json:"collector_number"`
	Digital         bool     `json:"digital"`
	Rarity          string   `json:"rarity"`
	FlavorText      string   `json:"flavor_text,omitempty"`
	CardBackID      string   `json:"card_back_id,omitempty"`
	Artist          string   `json:"artist"`
	ArtistIds       []string `json:"artist_ids"`
	IllustrationID  string   `json:"illustration_id,omitempty"`
	BorderColor     string   `json:"border_color"`
	Frame           string   `json:"frame"`
	FullArt         bool     `json:"full_art"`
	Textless        bool     `json:"textless"`
	Booster         bool     `json:"booster"`
	StorySpotlight  bool     `json:"story_spotlight"`
	EdhrecRank      int      `json:"edhrec_rank,omitempty"`
	Prices          struct {
		Usd       any `json:"usd"`
		UsdFoil   any `json:"usd_foil"`
		UsdEtched any `json:"usd_etched"`
		Eur       any `json:"eur"`
		EurFoil   any `json:"eur_foil"`
		Tix       any `json:"tix"`
	} `json:"prices"`
	RelatedUris struct {
		TcgplayerInfiniteArticles string `json:"tcgplayer_infinite_articles"`
		TcgplayerInfiniteDecks    string `json:"tcgplayer_infinite_decks"`
		Edhrec                    string `json:"edhrec"`
	} `json:"related_uris,omitempty"`
	PurchaseUris struct {
		Tcgplayer   string `json:"tcgplayer"`
		Cardmarket  string `json:"cardmarket"`
		Cardhoarder string `json:"cardhoarder"`
	} `json:"purchase_uris,omitempty"`
	MtgoID       int `json:"mtgo_id,omitempty"`
	MtgoFoilID   int `json:"mtgo_foil_id,omitempty"`
	CardmarketID int `json:"cardmarket_id,omitempty"`
	CardFaces    []struct {
		Object         string `json:"object"`
		Name           string `json:"name"`
		ManaCost       string `json:"mana_cost"`
		TypeLine       string `json:"type_line"`
		OracleText     string `json:"oracle_text"`
		Colors         []any  `json:"colors"`
		Artist         string `json:"artist"`
		ArtistID       string `json:"artist_id"`
		IllustrationID string `json:"illustration_id"`
		ImageUris      struct {
			Small  string `json:"small"`
			Normal string `json:"normal"`
			Large  string `json:"large"`
		} `json:"image_uris"`
	} `json:"card_faces,omitempty"`
}
