package game

import (
	"fmt"
	"math/rand"
)

func ShuffleCards(cards []CardName) []CardName {
	shuffled := make([]CardName, len(cards))
	perm := rand.Perm(len(cards))

	for i, v := range perm {
		shuffled[v] = cards[i]
	}

	return shuffled
}

// DealCards deals the cards to the players and returns the remaining cards
// and the dummy hand.
func DealCards(deck []CardName, numPlayers int) ([]CardName, []CardName, [][]CardName, error) {
	if numPlayers < 2 || numPlayers > 6 {
		return nil, nil, nil, fmt.Errorf("invalid number of players")
	}
	dummy := make([]CardName, 5)
	hands := make([][]CardName, numPlayers)
	// Deal the cards
	for i := 0; i < 5; i++ {
		for j := 0; j < numPlayers+1; j++ {
			if len(deck) == 0 {
				return nil, nil, nil, fmt.Errorf("deck is empty")
			}
			if j == numPlayers {
				dummy[i] = deck[0]
			} else {
				hands[j] = append(hands[j], deck[0])
			}
			deck = deck[1:]
		}
	}
	return deck, dummy, hands, nil
}

func BuyCards(deck []CardName, cards []CardName) ([]CardName, []CardName, error) {
	for {
		if len(cards) == 5 {
			break
		}
		if len(deck) == 0 {
			return nil, nil, fmt.Errorf("deck is empty")
		}

		cards = append(cards, deck[0])
		deck = deck[1:]
	}
	return deck, cards, nil
}

func NewDeck() []CardName {
	return []CardName{
		TWO_HEARTS,
		THREE_HEARTS,
		FOUR_HEARTS,
		FIVE_HEARTS,
		SIX_HEARTS,
		SEVEN_HEARTS,
		EIGHT_HEARTS,
		NINE_HEARTS,
		TEN_HEARTS,
		JACK_HEARTS,
		QUEEN_HEARTS,
		KING_HEARTS,
		ACE_HEARTS,
		TWO_DIAMONDS,
		THREE_DIAMONDS,
		FOUR_DIAMONDS,
		FIVE_DIAMONDS,
		SIX_DIAMONDS,
		SEVEN_DIAMONDS,
		EIGHT_DIAMONDS,
		NINE_DIAMONDS,
		TEN_DIAMONDS,
		JACK_DIAMONDS,
		QUEEN_DIAMONDS,
		KING_DIAMONDS,
		ACE_DIAMONDS,
		TWO_CLUBS,
		THREE_CLUBS,
		FOUR_CLUBS,
		FIVE_CLUBS,
		SIX_CLUBS,
		SEVEN_CLUBS,
		EIGHT_CLUBS,
		NINE_CLUBS,
		TEN_CLUBS,
		JACK_CLUBS,
		QUEEN_CLUBS,
		KING_CLUBS,
		ACE_CLUBS,
		TWO_SPADES,
		THREE_SPADES,
		FOUR_SPADES,
		FIVE_SPADES,
		SIX_SPADES,
		SEVEN_SPADES,
		EIGHT_SPADES,
		NINE_SPADES,
		TEN_SPADES,
		JACK_SPADES,
		QUEEN_SPADES,
		KING_SPADES,
		ACE_SPADES,
		JOKER,
	}
}

func ParseCardName(s string) (CardName, error) {
	switch s {
	case "TWO_HEARTS":
		return TWO_HEARTS, nil
	case "THREE_HEARTS":
		return THREE_HEARTS, nil
	case "FOUR_HEARTS":
		return FOUR_HEARTS, nil
	case "SIX_HEARTS":
		return SIX_HEARTS, nil
	case "SEVEN_HEARTS":
		return SEVEN_HEARTS, nil
	case "EIGHT_HEARTS":
		return EIGHT_HEARTS, nil
	case "NINE_HEARTS":
		return NINE_HEARTS, nil
	case "TEN_HEARTS":
		return TEN_HEARTS, nil
	case "QUEEN_HEARTS":
		return QUEEN_HEARTS, nil
	case "KING_HEARTS":
		return KING_HEARTS, nil
	case "ACE_HEARTS":
		return ACE_HEARTS, nil
	case "JACK_HEARTS":
		return JACK_HEARTS, nil
	case "FIVE_HEARTS":
		return FIVE_HEARTS, nil
	case "TWO_DIAMONDS":
		return TWO_DIAMONDS, nil
	case "THREE_DIAMONDS":
		return THREE_DIAMONDS, nil
	case "FOUR_DIAMONDS":
		return FOUR_DIAMONDS, nil
	case "SIX_DIAMONDS":
		return SIX_DIAMONDS, nil
	case "SEVEN_DIAMONDS":
		return SEVEN_DIAMONDS, nil
	case "EIGHT_DIAMONDS":
		return EIGHT_DIAMONDS, nil
	case "NINE_DIAMONDS":
		return NINE_DIAMONDS, nil
	case "TEN_DIAMONDS":
		return TEN_DIAMONDS, nil
	case "QUEEN_DIAMONDS":
		return QUEEN_DIAMONDS, nil
	case "KING_DIAMONDS":
		return KING_DIAMONDS, nil
	case "ACE_DIAMONDS":
		return ACE_DIAMONDS, nil
	case "JACK_DIAMONDS":
		return JACK_DIAMONDS, nil
	case "FIVE_DIAMONDS":
		return FIVE_DIAMONDS, nil
	case "TEN_CLUBS":
		return TEN_CLUBS, nil
	case "NINE_CLUBS":
		return NINE_CLUBS, nil
	case "EIGHT_CLUBS":
		return EIGHT_CLUBS, nil
	case "SEVEN_CLUBS":
		return SEVEN_CLUBS, nil
	case "SIX_CLUBS":
		return SIX_CLUBS, nil
	case "FOUR_CLUBS":
		return FOUR_CLUBS, nil
	case "THREE_CLUBS":
		return THREE_CLUBS, nil
	case "TWO_CLUBS":
		return TWO_CLUBS, nil
	case "QUEEN_CLUBS":
		return QUEEN_CLUBS, nil
	case "KING_CLUBS":
		return KING_CLUBS, nil
	case "ACE_CLUBS":
		return ACE_CLUBS, nil
	case "JACK_CLUBS":
		return JACK_CLUBS, nil
	case "FIVE_CLUBS":
		return FIVE_CLUBS, nil
	case "TEN_SPADES":
		return TEN_SPADES, nil
	case "NINE_SPADES":
		return NINE_SPADES, nil
	case "EIGHT_SPADES":
		return EIGHT_SPADES, nil
	case "SEVEN_SPADES":
		return SEVEN_SPADES, nil
	case "SIX_SPADES":
		return SIX_SPADES, nil
	case "FOUR_SPADES":
		return FOUR_SPADES, nil
	case "THREE_SPADES":
		return THREE_SPADES, nil
	case "TWO_SPADES":
		return TWO_SPADES, nil
	case "QUEEN_SPADES":
		return QUEEN_SPADES, nil
	case "KING_SPADES":
		return KING_SPADES, nil
	case "ACE_SPADES":
		return ACE_SPADES, nil
	case "JACK_SPADES":
		return JACK_SPADES, nil
	case "FIVE_SPADES":
		return FIVE_SPADES, nil
	case "JOKER":
		return JOKER, nil
	default:
		return EMPTY_CARD, fmt.Errorf("invalid card name")
	}
}

func ParseCard(c CardName) Card {
	switch c {
	case TWO_HEARTS:
		return TwoHearts
	case THREE_HEARTS:
		return ThreeHearts
	case FOUR_HEARTS:
		return FourHearts
	case FIVE_HEARTS:
		return FiveHearts
	case SIX_HEARTS:
		return SixHearts
	case SEVEN_HEARTS:
		return SevenHearts
	case EIGHT_HEARTS:
		return EightHearts
	case NINE_HEARTS:
		return NineHearts
	case TEN_HEARTS:
		return TenHearts
	case JACK_HEARTS:
		return JackHearts
	case QUEEN_HEARTS:
		return QueenHearts
	case KING_HEARTS:
		return KingHearts
	case ACE_HEARTS:
		return AceHearts
	case TWO_DIAMONDS:
		return TwoDiamonds
	case THREE_DIAMONDS:
		return ThreeDiamonds
	case FOUR_DIAMONDS:
		return FourDiamonds
	case FIVE_DIAMONDS:
		return FiveDiamonds
	case SIX_DIAMONDS:
		return SixDiamonds
	case SEVEN_DIAMONDS:
		return SevenDiamonds
	case EIGHT_DIAMONDS:
		return EightDiamonds
	case NINE_DIAMONDS:
		return NineDiamonds
	case TEN_DIAMONDS:
		return TenDiamonds
	case JACK_DIAMONDS:
		return JackDiamonds
	case QUEEN_DIAMONDS:
		return QueenDiamonds
	case KING_DIAMONDS:
		return KingDiamonds
	case ACE_DIAMONDS:
		return AceDiamonds
	case TWO_CLUBS:
		return TwoClubs
	case THREE_CLUBS:
		return ThreeClubs
	case FOUR_CLUBS:
		return FourClubs
	case FIVE_CLUBS:
		return FiveClubs
	case SIX_CLUBS:
		return SixClubs
	case SEVEN_CLUBS:
		return SevenClubs
	case EIGHT_CLUBS:
		return EightClubs
	case NINE_CLUBS:
		return NineClubs
	case TEN_CLUBS:
		return TenClubs
	case JACK_CLUBS:
		return JackClubs
	case QUEEN_CLUBS:
		return QueenClubs
	case KING_CLUBS:
		return KingClubs
	case ACE_CLUBS:
		return AceClubs
	case TWO_SPADES:
		return TwoSpades
	case THREE_SPADES:
		return ThreeSpades
	case FOUR_SPADES:
		return FourSpades
	case FIVE_SPADES:
		return FiveSpades
	case SIX_SPADES:
		return SixSpades
	case SEVEN_SPADES:
		return SevenSpades
	case EIGHT_SPADES:
		return EightSpades
	case NINE_SPADES:
		return NineSpades
	case TEN_SPADES:
		return TenSpades
	case JACK_SPADES:
		return JackSpades
	case QUEEN_SPADES:
		return QueenSpades
	case KING_SPADES:
		return KingSpades
	case ACE_SPADES:
		return AceSpades
	case JOKER:
		return Joker
	default:
		return EmptyCard
	}
}
