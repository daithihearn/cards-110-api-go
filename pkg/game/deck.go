package game

// Suit represents the suit of the card.
type Suit string

const (
	EMPTY    Suit = "EMPTY"
	CLUBS         = "CLUBS"
	DIAMONDS      = "DIAMONDS"
	HEARTS        = "HEARTS"
	SPADES        = "SPADES"
	WILD          = "WILD"
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
	NAME      CardName
	Value     int
	ColdValue int
	Suit      Suit
	Renegable bool
}

// Define cards as constants.
var (
	EmptyCard     = Card{EMPTY_CARD, 0, 0, EMPTY, false}
	TwoHearts     = Card{NAME: TWO_HEARTS, Value: 2, ColdValue: 0, Suit: HEARTS, Renegable: false}
	ThreeHearts   = Card{THREE_HEARTS, 3, 0, HEARTS, false}
	FourHearts    = Card{FOUR_HEARTS, 4, 0, HEARTS, false}
	SixHearts     = Card{SIX_HEARTS, 6, 0, HEARTS, false}
	SevenHearts   = Card{SEVEN_HEARTS, 7, 0, HEARTS, false}
	EightHearts   = Card{EIGHT_HEARTS, 8, 0, HEARTS, false}
	NineHearts    = Card{NINE_HEARTS, 9, 0, HEARTS, false}
	TenHearts     = Card{TEN_HEARTS, 10, 0, HEARTS, false}
	QueenHearts   = Card{QUEEN_HEARTS, 12, 0, HEARTS, false}
	KingHearts    = Card{KING_HEARTS, 13, 0, HEARTS, false}
	AceHearts     = Card{ACE_HEARTS, 1, 0, HEARTS, false}
	JackHearts    = Card{JACK_HEARTS, 11, 0, HEARTS, true}
	FiveHearts    = Card{FIVE_HEARTS, 5, 0, HEARTS, true}
	TwoDiamonds   = Card{TWO_DIAMONDS, 2, 0, DIAMONDS, false}
	ThreeDiamonds = Card{THREE_DIAMONDS, 3, 0, DIAMONDS, false}
	FourDiamonds  = Card{FOUR_DIAMONDS, 4, 0, DIAMONDS, false}
	SixDiamonds   = Card{SIX_DIAMONDS, 6, 0, DIAMONDS, false}
	SevenDiamonds = Card{SEVEN_DIAMONDS, 7, 0, DIAMONDS, false}
	EightDiamonds = Card{EIGHT_DIAMONDS, 8, 0, DIAMONDS, false}
	NineDiamonds  = Card{NINE_DIAMONDS, 9, 0, DIAMONDS, false}
	TenDiamonds   = Card{TEN_DIAMONDS, 10, 0, DIAMONDS, false}
	QueenDiamonds = Card{QUEEN_DIAMONDS, 12, 0, DIAMONDS, false}
	KingDiamonds  = Card{KING_DIAMONDS, 13, 0, DIAMONDS, false}
	AceDiamonds   = Card{ACE_DIAMONDS, 1, 0, DIAMONDS, false}
	JackDiamonds  = Card{JACK_DIAMONDS, 11, 0, DIAMONDS, true}
	FiveDiamonds  = Card{FIVE_DIAMONDS, 5, 0, DIAMONDS, true}
	TenClubs      = Card{TEN_CLUBS, 10, 0, CLUBS, false}
	NineClubs     = Card{NINE_CLUBS, 9, 0, CLUBS, false}
	EightClubs    = Card{EIGHT_CLUBS, 8, 0, CLUBS, false}
	SevenClubs    = Card{SEVEN_CLUBS, 7, 0, CLUBS, false}
	SixClubs      = Card{SIX_CLUBS, 6, 0, CLUBS, false}
	FourClubs     = Card{FOUR_CLUBS, 4, 0, CLUBS, false}
	ThreeClubs    = Card{THREE_CLUBS, 3, 0, CLUBS, false}
	TwoClubs      = Card{TWO_CLUBS, 2, 0, CLUBS, false}
	QueenClubs    = Card{QUEEN_CLUBS, 12, 0, CLUBS, false}
	KingClubs     = Card{KING_CLUBS, 13, 0, CLUBS, false}
	AceClubs      = Card{ACE_CLUBS, 1, 0, CLUBS, false}
	JackClubs     = Card{JACK_CLUBS, 11, 0, CLUBS, true}
	FiveClubs     = Card{FIVE_CLUBS, 5, 0, CLUBS, true}
	TenSpades     = Card{TEN_SPADES, 10, 0, SPADES, false}
	NineSpades    = Card{NINE_SPADES, 9, 0, SPADES, false}
	EightSpades   = Card{EIGHT_SPADES, 8, 0, SPADES, false}
	SevenSpades   = Card{SEVEN_SPADES, 7, 0, SPADES, false}
	SixSpades     = Card{SIX_SPADES, 6, 0, SPADES, false}
	FourSpades    = Card{FOUR_SPADES, 4, 0, SPADES, false}
	ThreeSpades   = Card{THREE_SPADES, 3, 0, SPADES, false}
	TwoSpades     = Card{TWO_SPADES, 2, 0, SPADES, false}
	QueenSpades   = Card{QUEEN_SPADES, 12, 0, SPADES, false}
	KingSpades    = Card{KING_SPADES, 13, 0, SPADES, false}
	AceSpades     = Card{ACE_SPADES, 1, 0, SPADES, false}
	JackSpades    = Card{JACK_SPADES, 11, 0, SPADES, true}
	FiveSpades    = Card{FIVE_SPADES, 5, 0, SPADES, true}
	Joker         = Card{JOKER, 0, 0, WILD, true}
)

func (c Card) String() string {
	return string(c.NAME)
}
