package deck

// Suit represents the suit of the card.
type Suit string

// Define suits.
const (
	SuitHearts   Suit = "HEARTS"
	SuitDiamonds Suit = "DIAMONDS"
	SuitClubs    Suit = "CLUBS"
	SuitSpades   Suit = "SPADES"
	SuitWild     Suit = "WILD"
	SuitEmpty    Suit = "EMPTY"
)

// Card represents a card with a value, coldValue, suit, and renegable status.
type Card struct {
	Value     int
	ColdValue int
	Suit      Suit
	Renegable bool
}

// Define cards as constants.
var (
	EmptyCard     = Card{0, 0, SuitEmpty, false}
	TwoHearts     = Card{101, 2, SuitHearts, false}
	ThreeHearts   = Card{102, 3, SuitHearts, false}
	FourHearts    = Card{103, 4, SuitHearts, false}
	SixHearts     = Card{104, 6, SuitHearts, false}
	SevenHearts   = Card{105, 7, SuitHearts, false}
	EightHearts   = Card{106, 8, SuitHearts, false}
	NineHearts    = Card{107, 9, SuitHearts, false}
	TenHearts     = Card{108, 10, SuitHearts, false}
	QueenHearts   = Card{109, 12, SuitHearts, false}
	KingHearts    = Card{110, 13, SuitHearts, false}
	AceHearts     = Card{112, 0, SuitWild, true}
	JackHearts    = Card{114, 11, SuitHearts, true}
	FiveHearts    = Card{115, 5, SuitHearts, true}
	TwoDiamonds   = Card{101, 2, SuitDiamonds, false}
	ThreeDiamonds = Card{102, 3, SuitDiamonds, false}
	FourDiamonds  = Card{103, 4, SuitDiamonds, false}
	SixDiamonds   = Card{104, 6, SuitDiamonds, false}
	SevenDiamonds = Card{105, 7, SuitDiamonds, false}
	EightDiamonds = Card{106, 8, SuitDiamonds, false}
	NineDiamonds  = Card{107, 9, SuitDiamonds, false}
	TenDiamonds   = Card{108, 10, SuitDiamonds, false}
	QueenDiamonds = Card{109, 12, SuitDiamonds, false}
	KingDiamonds  = Card{110, 13, SuitDiamonds, false}
	AceDiamonds   = Card{111, 1, SuitDiamonds, false}
	JackDiamonds  = Card{114, 11, SuitDiamonds, true}
	FiveDiamonds  = Card{115, 5, SuitDiamonds, true}
	TenClubs      = Card{101, 1, SuitClubs, false}
	NineClubs     = Card{102, 2, SuitClubs, false}
	EightClubs    = Card{103, 3, SuitClubs, false}
	SevenClubs    = Card{104, 4, SuitClubs, false}
	SixClubs      = Card{105, 5, SuitClubs, false}
	FourClubs     = Card{106, 7, SuitClubs, false}
	ThreeClubs    = Card{107, 8, SuitClubs, false}
	TwoClubs      = Card{108, 9, SuitClubs, false}
	QueenClubs    = Card{109, 12, SuitClubs, false}
	KingClubs     = Card{110, 13, SuitClubs, false}
	AceClubs      = Card{111, 10, SuitClubs, false}
	JackClubs     = Card{114, 11, SuitClubs, true}
	FiveClubs     = Card{115, 6, SuitClubs, true}
	TenSpades     = Card{101, 1, SuitSpades, false}
	NineSpades    = Card{102, 2, SuitSpades, false}
	EightSpades   = Card{103, 3, SuitSpades, false}
	SevenSpades   = Card{104, 4, SuitSpades, false}
	SixSpades     = Card{105, 5, SuitSpades, false}
	FourSpades    = Card{106, 7, SuitSpades, false}
	ThreeSpades   = Card{107, 8, SuitSpades, false}
	TwoSpades     = Card{108, 9, SuitSpades, false}
	QueenSpades   = Card{109, 12, SuitSpades, false}
	KingSpades    = Card{110, 13, SuitSpades, false}
	AceSpades     = Card{111, 10, SuitSpades, false}
	JackSpades    = Card{114, 11, SuitSpades, true}
	FiveSpades    = Card{115, 6, SuitSpades, true}
	Joker         = Card{113, 0, SuitWild, true}
)

// Deck represents the structure of the deck document in MongoDB.
type Deck struct {
	ID    string `bson:"_id,omitempty"` // The bson tag specifies the field name in MongoDB.
	Cards []Card `bson:"cards"`         // Cards are represented as a slice of Card.
}
