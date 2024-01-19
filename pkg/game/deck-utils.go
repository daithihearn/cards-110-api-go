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

func DealCards(deck []CardName, numPlayers int) ([]CardName, [][]CardName, error) {
	hands := make([][]CardName, numPlayers+1)
	// Deal the cards
	for i := 0; i < 5; i++ {
		for j := 0; j < numPlayers+1; j++ {
			if len(deck) == 0 {
				return nil, nil, fmt.Errorf("deck is empty")
			}
			hands[j] = append(hands[j], deck[0])
			deck = deck[1:]
		}
	}
	return deck, hands, nil
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
