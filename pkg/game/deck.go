package game

// Suit represents the suit of the card.
type Suit string

const (
	Empty    Suit = "EMPTY"
	Clubs         = "CLUBS"
	Diamonds      = "DIAMONDS"
	Hearts        = "HEARTS"
	Spades        = "SPADES"
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
	EmptyCard     = Card{EMPTY_CARD, 0, 0, Empty, false}
	TwoHearts     = Card{Name: TWO_HEARTS, Value: 2, ColdValue: 0, Suit: Hearts, Renegable: false}
	ThreeHearts   = Card{THREE_HEARTS, 3, 0, Hearts, false}
	FourHearts    = Card{FOUR_HEARTS, 4, 0, Hearts, false}
	SixHearts     = Card{SIX_HEARTS, 6, 0, Hearts, false}
	SevenHearts   = Card{SEVEN_HEARTS, 7, 0, Hearts, false}
	EightHearts   = Card{EIGHT_HEARTS, 8, 0, Hearts, false}
	NineHearts    = Card{NINE_HEARTS, 9, 0, Hearts, false}
	TenHearts     = Card{TEN_HEARTS, 10, 0, Hearts, false}
	QueenHearts   = Card{QUEEN_HEARTS, 12, 0, Hearts, false}
	KingHearts    = Card{KING_HEARTS, 13, 0, Hearts, false}
	AceHearts     = Card{ACE_HEARTS, 1, 0, Hearts, false}
	JackHearts    = Card{JACK_HEARTS, 11, 0, Hearts, true}
	FiveHearts    = Card{FIVE_HEARTS, 5, 0, Hearts, true}
	TwoDiamonds   = Card{TWO_DIAMONDS, 2, 0, Diamonds, false}
	ThreeDiamonds = Card{THREE_DIAMONDS, 3, 0, Diamonds, false}
	FourDiamonds  = Card{FOUR_DIAMONDS, 4, 0, Diamonds, false}
	SixDiamonds   = Card{SIX_DIAMONDS, 6, 0, Diamonds, false}
	SevenDiamonds = Card{SEVEN_DIAMONDS, 7, 0, Diamonds, false}
	EightDiamonds = Card{EIGHT_DIAMONDS, 8, 0, Diamonds, false}
	NineDiamonds  = Card{NINE_DIAMONDS, 9, 0, Diamonds, false}
	TenDiamonds   = Card{TEN_DIAMONDS, 10, 0, Diamonds, false}
	QueenDiamonds = Card{QUEEN_DIAMONDS, 12, 0, Diamonds, false}
	KingDiamonds  = Card{KING_DIAMONDS, 13, 0, Diamonds, false}
	AceDiamonds   = Card{ACE_DIAMONDS, 1, 0, Diamonds, false}
	JackDiamonds  = Card{JACK_DIAMONDS, 11, 0, Diamonds, true}
	FiveDiamonds  = Card{FIVE_DIAMONDS, 5, 0, Diamonds, true}
	TenClubs      = Card{TEN_CLUBS, 10, 0, Clubs, false}
	NineClubs     = Card{NINE_CLUBS, 9, 0, Clubs, false}
	EightClubs    = Card{EIGHT_CLUBS, 8, 0, Clubs, false}
	SevenClubs    = Card{SEVEN_CLUBS, 7, 0, Clubs, false}
	SixClubs      = Card{SIX_CLUBS, 6, 0, Clubs, false}
	FourClubs     = Card{FOUR_CLUBS, 4, 0, Clubs, false}
	ThreeClubs    = Card{THREE_CLUBS, 3, 0, Clubs, false}
	TwoClubs      = Card{TWO_CLUBS, 2, 0, Clubs, false}
	QueenClubs    = Card{QUEEN_CLUBS, 12, 0, Clubs, false}
	KingClubs     = Card{KING_CLUBS, 13, 0, Clubs, false}
	AceClubs      = Card{ACE_CLUBS, 1, 0, Clubs, false}
	JackClubs     = Card{JACK_CLUBS, 11, 0, Clubs, true}
	FiveClubs     = Card{FIVE_CLUBS, 5, 0, Clubs, true}
	TenSpades     = Card{TEN_SPADES, 10, 0, Spades, false}
	NineSpades    = Card{NINE_SPADES, 9, 0, Spades, false}
	EightSpades   = Card{EIGHT_SPADES, 8, 0, Spades, false}
	SevenSpades   = Card{SEVEN_SPADES, 7, 0, Spades, false}
	SixSpades     = Card{SIX_SPADES, 6, 0, Spades, false}
	FourSpades    = Card{FOUR_SPADES, 4, 0, Spades, false}
	ThreeSpades   = Card{THREE_SPADES, 3, 0, Spades, false}
	TwoSpades     = Card{TWO_SPADES, 2, 0, Spades, false}
	QueenSpades   = Card{QUEEN_SPADES, 12, 0, Spades, false}
	KingSpades    = Card{KING_SPADES, 13, 0, Spades, false}
	AceSpades     = Card{ACE_SPADES, 1, 0, Spades, false}
	JackSpades    = Card{JACK_SPADES, 11, 0, Spades, true}
	FiveSpades    = Card{FIVE_SPADES, 5, 0, Spades, true}
	Joker         = Card{JOKER, 0, 0, Wild, true}
)

func (c Card) String() string {
	return string(c.Name)
}
