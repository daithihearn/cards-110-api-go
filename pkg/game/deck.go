package game

// Suit represents the suit of the card.
type Suit string

const (
	Empty    Suit = "EMPTY"
	Clubs         = "Clubs"
	Diamonds      = "Diamonds"
	Hearts        = "Hearts"
	Spades        = "Spades"
	Wild          = "WILD"
)

type CardName string

const (
	EMPTY_CARD     CardName = "EMPTY"
	TWO_HEARTS              = "TWO_HEARTS"
	THREE_HEARTS            = "THREE_HEARTS"
	FOUR_HEARTS             = "FOUR_HEARTS"
	SIX_HEARTS              = "SIX_HEARTS"
	SEVEN_HEARTS            = "SEVEN_HEARTS"
	EIGHT_HEARTS            = "EIGHT_HEARTS"
	NINE_HEARTS             = "NINE_HEARTS"
	TEN_HEARTS              = "TEN_HEARTS"
	QUEEN_HEARTS            = "QUEEN_HEARTS"
	KING_HEARTS             = "KING_HEARTS"
	ACE_HEARTS              = "ACE_HEARTS"
	JACK_HEARTS             = "JACK_HEARTS"
	FIVE_HEARTS             = "FIVE_HEARTS"
	TWO_DIAMONDS            = "TWO_DIAMONDS"
	THREE_DIAMONDS          = "THREE_DIAMONDS"
	FOUR_DIAMONDS           = "FOUR_DIAMONDS"
	SIX_DIAMONDS            = "SIX_DIAMONDS"
	SEVEN_DIAMONDS          = "SEVEN_DIAMONDS"
	EIGHT_DIAMONDS          = "EIGHT_DIAMONDS"
	NINE_DIAMONDS           = "NINE_DIAMONDS"
	TEN_DIAMONDS            = "TEN_DIAMONDS"
	QUEEN_DIAMONDS          = "QUEEN_DIAMONDS"
	KING_DIAMONDS           = "KING_DIAMONDS"
	ACE_DIAMONDS            = "ACE_DIAMONDS"
	JACK_DIAMONDS           = "JACK_DIAMONDS"
	FIVE_DIAMONDS           = "FIVE_DIAMONDS"
	TEN_CLUBS               = "TEN_CLUBS"
	NINE_CLUBS              = "NINE_CLUBS"
	EIGHT_CLUBS             = "EIGHT_CLUBS"
	SEVEN_CLUBS             = "SEVEN_CLUBS"
	SIX_CLUBS               = "SIX_CLUBS"
	FOUR_CLUBS              = "FOUR_CLUBS"
	THREE_CLUBS             = "THREE_CLUBS"
	TWO_CLUBS               = "TWO_CLUBS"
	QUEEN_CLUBS             = "QUEEN_CLUBS"
	KING_CLUBS              = "KING_CLUBS"
	ACE_CLUBS               = "ACE_CLUBS"
	JACK_CLUBS              = "JACK_CLUBS"
	FIVE_CLUBS              = "FIVE_CLUBS"
	TEN_SPADES              = "TEN_SPADES"
	NINE_SPADES             = "NINE_SPADES"
	EIGHT_SPADES            = "EIGHT_SPADES"
	SEVEN_SPADES            = "SEVEN_SPADES"
	SIX_SPADES              = "SIX_SPADES"
	FOUR_SPADES             = "FOUR_SPADES"
	THREE_SPADES            = "THREE_SPADES"
	TWO_SPADES              = "TWO_SPADES"
	QUEEN_SPADES            = "QUEEN_SPADES"
	KING_SPADES             = "KING_SPADES"
	ACE_SPADES              = "ACE_SPADES"
	JACK_SPADES             = "JACK_SPADES"
	FIVE_SPADES             = "FIVE_SPADES"
	JOKER                   = "JOKER"
)

func (c CardName) Card() Card {
	return ParseCard(c)
}

// Card represents a card with a value, coldValue, suit, and renegable status.
type Card struct {
	Name      CardName
	Value     int
	ColdValue int
	Suit      Suit
	Renegable bool
}

// Define cards as constants.
var (
	EmptyCard     = Card{Name: EMPTY_CARD, Value: 0, ColdValue: 0, Suit: Empty, Renegable: false}
	TwoHearts     = Card{Name: TWO_HEARTS, Value: 101, ColdValue: 2, Suit: Hearts, Renegable: false}
	ThreeHearts   = Card{Name: THREE_HEARTS, Value: 102, ColdValue: 3, Suit: Hearts, Renegable: false}
	FourHearts    = Card{Name: FOUR_HEARTS, Value: 103, ColdValue: 4, Suit: Hearts, Renegable: false}
	FiveHearts    = Card{Name: FIVE_HEARTS, Value: 115, ColdValue: 5, Suit: Hearts, Renegable: true}
	SixHearts     = Card{Name: SIX_HEARTS, Value: 104, ColdValue: 6, Suit: Hearts, Renegable: false}
	SevenHearts   = Card{Name: SEVEN_HEARTS, Value: 105, ColdValue: 7, Suit: Hearts, Renegable: false}
	EightHearts   = Card{Name: EIGHT_HEARTS, Value: 106, ColdValue: 8, Suit: Hearts, Renegable: false}
	NineHearts    = Card{Name: NINE_HEARTS, Value: 107, ColdValue: 9, Suit: Hearts, Renegable: false}
	TenHearts     = Card{Name: TEN_HEARTS, Value: 108, ColdValue: 10, Suit: Hearts, Renegable: false}
	JackHearts    = Card{Name: JACK_HEARTS, Value: 114, ColdValue: 11, Suit: Hearts, Renegable: true}
	QueenHearts   = Card{Name: QUEEN_HEARTS, Value: 109, ColdValue: 12, Suit: Hearts, Renegable: false}
	KingHearts    = Card{Name: KING_HEARTS, Value: 110, ColdValue: 13, Suit: Hearts, Renegable: false}
	AceHearts     = Card{Name: ACE_HEARTS, Value: 112, ColdValue: 0, Suit: Wild, Renegable: true}
	TwoDiamonds   = Card{Name: TWO_DIAMONDS, Value: 101, ColdValue: 2, Suit: Diamonds, Renegable: false}
	ThreeDiamonds = Card{Name: THREE_DIAMONDS, Value: 102, ColdValue: 3, Suit: Diamonds, Renegable: false}
	FourDiamonds  = Card{Name: FOUR_DIAMONDS, Value: 103, ColdValue: 4, Suit: Diamonds, Renegable: false}
	FiveDiamonds  = Card{Name: FIVE_DIAMONDS, Value: 115, ColdValue: 5, Suit: Diamonds, Renegable: true}
	SixDiamonds   = Card{Name: SIX_DIAMONDS, Value: 104, ColdValue: 6, Suit: Diamonds, Renegable: false}
	SevenDiamonds = Card{Name: SEVEN_DIAMONDS, Value: 105, ColdValue: 7, Suit: Diamonds, Renegable: false}
	EightDiamonds = Card{Name: EIGHT_DIAMONDS, Value: 106, ColdValue: 8, Suit: Diamonds, Renegable: false}
	NineDiamonds  = Card{Name: NINE_DIAMONDS, Value: 107, ColdValue: 9, Suit: Diamonds, Renegable: false}
	TenDiamonds   = Card{Name: TEN_DIAMONDS, Value: 108, ColdValue: 10, Suit: Diamonds, Renegable: false}
	JackDiamonds  = Card{Name: JACK_DIAMONDS, Value: 114, ColdValue: 11, Suit: Diamonds, Renegable: true}
	QueenDiamonds = Card{Name: QUEEN_DIAMONDS, Value: 109, ColdValue: 12, Suit: Diamonds, Renegable: false}
	KingDiamonds  = Card{Name: KING_DIAMONDS, Value: 110, ColdValue: 13, Suit: Diamonds, Renegable: false}
	AceDiamonds   = Card{Name: ACE_DIAMONDS, Value: 111, ColdValue: 1, Suit: Diamonds, Renegable: false}
	TwoClubs      = Card{Name: TWO_CLUBS, Value: 108, ColdValue: 9, Suit: Clubs, Renegable: false}
	ThreeClubs    = Card{Name: THREE_CLUBS, Value: 107, ColdValue: 8, Suit: Clubs, Renegable: false}
	FourClubs     = Card{Name: FOUR_CLUBS, Value: 106, ColdValue: 7, Suit: Clubs, Renegable: false}
	FiveClubs     = Card{Name: FIVE_CLUBS, Value: 115, ColdValue: 6, Suit: Clubs, Renegable: true}
	SixClubs      = Card{Name: SIX_CLUBS, Value: 105, ColdValue: 5, Suit: Clubs, Renegable: false}
	SevenClubs    = Card{Name: SEVEN_CLUBS, Value: 104, ColdValue: 4, Suit: Clubs, Renegable: false}
	EightClubs    = Card{Name: EIGHT_CLUBS, Value: 103, ColdValue: 3, Suit: Clubs, Renegable: false}
	NineClubs     = Card{Name: NINE_CLUBS, Value: 102, ColdValue: 2, Suit: Clubs, Renegable: false}
	TenClubs      = Card{Name: TEN_CLUBS, Value: 101, ColdValue: 1, Suit: Clubs, Renegable: false}
	JackClubs     = Card{Name: JACK_CLUBS, Value: 114, ColdValue: 11, Suit: Clubs, Renegable: true}
	QueenClubs    = Card{Name: QUEEN_CLUBS, Value: 109, ColdValue: 12, Suit: Clubs, Renegable: false}
	KingClubs     = Card{Name: KING_CLUBS, Value: 110, ColdValue: 13, Suit: Clubs, Renegable: false}
	AceClubs      = Card{Name: ACE_CLUBS, Value: 111, ColdValue: 10, Suit: Clubs, Renegable: false}
	TwoSpades     = Card{Name: TWO_SPADES, Value: 108, ColdValue: 9, Suit: Spades, Renegable: false}
	ThreeSpades   = Card{Name: THREE_SPADES, Value: 107, ColdValue: 8, Suit: Spades, Renegable: false}
	FourSpades    = Card{Name: FOUR_SPADES, Value: 106, ColdValue: 7, Suit: Spades, Renegable: false}
	FiveSpades    = Card{Name: FIVE_SPADES, Value: 115, ColdValue: 6, Suit: Spades, Renegable: true}
	SixSpades     = Card{Name: SIX_SPADES, Value: 105, ColdValue: 5, Suit: Spades, Renegable: false}
	SevenSpades   = Card{Name: SEVEN_SPADES, Value: 104, ColdValue: 4, Suit: Spades, Renegable: false}
	EightSpades   = Card{Name: EIGHT_SPADES, Value: 103, ColdValue: 3, Suit: Spades, Renegable: false}
	NineSpades    = Card{Name: NINE_SPADES, Value: 102, ColdValue: 2, Suit: Spades, Renegable: false}
	TenSpades     = Card{Name: TEN_SPADES, Value: 101, ColdValue: 1, Suit: Spades, Renegable: false}
	JackSpades    = Card{Name: JACK_SPADES, Value: 114, ColdValue: 11, Suit: Spades, Renegable: true}
	QueenSpades   = Card{Name: QUEEN_SPADES, Value: 109, ColdValue: 12, Suit: Spades, Renegable: false}
	KingSpades    = Card{Name: KING_SPADES, Value: 110, ColdValue: 13, Suit: Spades, Renegable: false}
	AceSpades     = Card{Name: ACE_SPADES, Value: 111, ColdValue: 10, Suit: Spades, Renegable: false}
	Joker         = Card{Name: JOKER, Value: 113, ColdValue: 0, Suit: Wild, Renegable: true}
)

func (c Card) String() string {
	return string(c.Name)
}
