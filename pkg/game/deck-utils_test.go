package game

import "testing"

func TestDeck_ShuffleCards(t *testing.T) {
	tests := []struct {
		name  string
		cards []CardName
	}{
		{
			name:  "Empty deck",
			cards: []CardName{},
		},
		{
			name:  "Single card",
			cards: []CardName{ACE_HEARTS},
		},
		{
			name:  "Multiple cards",
			cards: []CardName{ACE_HEARTS, TWO_HEARTS, THREE_HEARTS, FOUR_HEARTS, FIVE_HEARTS, SIX_HEARTS, SEVEN_HEARTS, EIGHT_HEARTS, NINE_HEARTS, TEN_HEARTS, JACK_HEARTS, QUEEN_HEARTS, KING_HEARTS, ACE_DIAMONDS, TWO_DIAMONDS, THREE_DIAMONDS, FOUR_DIAMONDS, FIVE_DIAMONDS, SIX_DIAMONDS, SEVEN_DIAMONDS, EIGHT_DIAMONDS, NINE_DIAMONDS, TEN_DIAMONDS, JACK_DIAMONDS, QUEEN_DIAMONDS, KING_DIAMONDS, ACE_CLUBS, TWO_CLUBS, THREE_CLUBS, FOUR_CLUBS, FIVE_CLUBS, SIX_CLUBS, SEVEN_CLUBS, EIGHT_CLUBS, NINE_CLUBS, TEN_CLUBS, JACK_CLUBS, QUEEN_CLUBS, KING_CLUBS, ACE_SPADES, TWO_SPADES, THREE_SPADES, FOUR_SPADES, FIVE_SPADES, SIX_SPADES, SEVEN_SPADES, EIGHT_SPADES, NINE_SPADES, TEN_SPADES, JACK_SPADES, QUEEN_SPADES, KING_SPADES, JOKER},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shuffled := ShuffleCards(test.cards)
			if len(shuffled) != len(test.cards) {
				t.Errorf("expected %d cards, got %d", len(test.cards), len(shuffled))
			}

			if !compare(shuffled, test.cards) {
				t.Errorf("expected cards to be the same, got %v", shuffled)
			}

			// Cards should not be in the same order as they started (unless the deck is empty or has a small number of cards)
			if len(test.cards) > 7 {
				var sameOrder = true
				for i, card := range test.cards {
					if card != shuffled[i] {
						sameOrder = false
						break
					}
				}
				if sameOrder {
					t.Errorf("expected cards to be shuffled, got %v", shuffled)
				}
			}
		})
	}
}

func TestDeck_DealCards(t *testing.T) {
	tests := []struct {
		name        string
		deck        []CardName
		numPlayers  int
		expectError bool
	}{
		{
			name:        "Empty deck",
			deck:        []CardName{},
			numPlayers:  2,
			expectError: true,
		},
		{
			name:        "Not enough cards",
			deck:        []CardName{ACE_HEARTS},
			numPlayers:  2,
			expectError: true,
		},
		{
			name:       "Multiple cards",
			deck:       []CardName{ACE_HEARTS, TWO_HEARTS, THREE_HEARTS, FOUR_HEARTS, FIVE_HEARTS, SIX_HEARTS, SEVEN_HEARTS, EIGHT_HEARTS, NINE_HEARTS, TEN_HEARTS, JACK_HEARTS, QUEEN_HEARTS, KING_HEARTS, ACE_DIAMONDS, TWO_DIAMONDS, THREE_DIAMONDS, FOUR_DIAMONDS, FIVE_DIAMONDS, SIX_DIAMONDS, SEVEN_DIAMONDS, EIGHT_DIAMONDS, NINE_DIAMONDS, TEN_DIAMONDS, JACK_DIAMONDS, QUEEN_DIAMONDS, KING_DIAMONDS, ACE_CLUBS, TWO_CLUBS, THREE_CLUBS, FOUR_CLUBS, FIVE_CLUBS, SIX_CLUBS, SEVEN_CLUBS, EIGHT_CLUBS, NINE_CLUBS, TEN_CLUBS, JACK_CLUBS, QUEEN_CLUBS, KING_CLUBS, ACE_SPADES, TWO_SPADES, THREE_SPADES, FOUR_SPADES, FIVE_SPADES, SIX_SPADES, SEVEN_SPADES, EIGHT_SPADES, NINE_SPADES, TEN_SPADES, JACK_SPADES, QUEEN_SPADES, KING_SPADES, JOKER},
			numPlayers: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			deck, hands, err := DealCards(test.deck, test.numPlayers)

			if test.expectError {
				if err == nil {
					t.Errorf("expected an error")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}

				// The output deck and hands should have all the same cards as the input deck
				var outputCards = deck
				for _, hand := range hands {
					outputCards = append(outputCards, hand...)
				}
				if !compare(outputCards, test.deck) {
					t.Errorf("expected cards to be the same, got %v", outputCards)
				}
			}
		})
	}
}

func TestDeck_BuyCards(t *testing.T) {
	tests := []struct {
		name          string
		deck          []CardName
		cards         []CardName
		expectError   bool
		expectedDeck  []CardName
		expectedCards []CardName
	}{
		{
			name:        "Empty deck",
			deck:        []CardName{},
			cards:       []CardName{},
			expectError: true,
		},
		{
			name:        "Not enough cards",
			deck:        []CardName{ACE_HEARTS},
			cards:       []CardName{},
			expectError: true,
		},
		{
			name:        "Not enough cards",
			deck:        []CardName{ACE_DIAMONDS},
			cards:       []CardName{ACE_HEARTS},
			expectError: true,
		},
		{
			name:          "Enough cards",
			deck:          []CardName{ACE_DIAMONDS},
			cards:         []CardName{ACE_HEARTS, TWO_HEARTS, THREE_HEARTS, FOUR_HEARTS},
			expectedDeck:  []CardName{},
			expectedCards: []CardName{ACE_HEARTS, TWO_HEARTS, THREE_HEARTS, FOUR_HEARTS, ACE_DIAMONDS},
		},
		{
			name:          "Loads of cards",
			deck:          []CardName{ACE_HEARTS, TWO_HEARTS, THREE_HEARTS, FOUR_HEARTS, FIVE_HEARTS, SIX_HEARTS, SEVEN_HEARTS, EIGHT_HEARTS, NINE_HEARTS, TEN_HEARTS, JACK_HEARTS, QUEEN_HEARTS, KING_HEARTS, ACE_DIAMONDS, TWO_DIAMONDS, THREE_DIAMONDS, FOUR_DIAMONDS, FIVE_DIAMONDS, SIX_DIAMONDS, SEVEN_DIAMONDS, EIGHT_DIAMONDS, NINE_DIAMONDS, TEN_DIAMONDS, JACK_DIAMONDS, QUEEN_DIAMONDS, KING_DIAMONDS, ACE_CLUBS, TWO_CLUBS, THREE_CLUBS, FOUR_CLUBS, FIVE_CLUBS, SIX_CLUBS, SEVEN_CLUBS, EIGHT_CLUBS, NINE_CLUBS, TEN_CLUBS, JACK_CLUBS, QUEEN_CLUBS, KING_CLUBS, ACE_SPADES, TWO_SPADES, THREE_SPADES, FOUR_SPADES, FIVE_SPADES, SIX_SPADES, SEVEN_SPADES, EIGHT_SPADES, NINE_SPADES, TEN_SPADES, JACK_SPADES, QUEEN_SPADES, KING_SPADES, JOKER},
			cards:         []CardName{},
			expectedDeck:  []CardName{SIX_HEARTS, SEVEN_HEARTS, EIGHT_HEARTS, NINE_HEARTS, TEN_HEARTS, JACK_HEARTS, QUEEN_HEARTS, KING_HEARTS, ACE_DIAMONDS, TWO_DIAMONDS, THREE_DIAMONDS, FOUR_DIAMONDS, FIVE_DIAMONDS, SIX_DIAMONDS, SEVEN_DIAMONDS, EIGHT_DIAMONDS, NINE_DIAMONDS, TEN_DIAMONDS, JACK_DIAMONDS, QUEEN_DIAMONDS, KING_DIAMONDS, ACE_CLUBS, TWO_CLUBS, THREE_CLUBS, FOUR_CLUBS, FIVE_CLUBS, SIX_CLUBS, SEVEN_CLUBS, EIGHT_CLUBS, NINE_CLUBS, TEN_CLUBS, JACK_CLUBS, QUEEN_CLUBS, KING_CLUBS, ACE_SPADES, TWO_SPADES, THREE_SPADES, FOUR_SPADES, FIVE_SPADES, SIX_SPADES, SEVEN_SPADES, EIGHT_SPADES, NINE_SPADES, TEN_SPADES, JACK_SPADES, QUEEN_SPADES, KING_SPADES, JOKER},
			expectedCards: []CardName{ACE_HEARTS, TWO_HEARTS, THREE_HEARTS, FOUR_HEARTS, FIVE_HEARTS},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			deck, cards, err := BuyCards(test.deck, test.cards)

			if test.expectError {
				if err == nil {
					t.Errorf("expected an error")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}

				if !compare(deck, test.expectedDeck) {
					t.Errorf("expected deck to be %v, got %v", test.expectedDeck, deck)
				}
				if !compare(cards, test.expectedCards) {
					t.Errorf("expected cards to be %v, got %v", test.expectedCards, cards)
				}
			}
		})
	}
}

func TestDeck_NewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck) != 53 {
		t.Errorf("expected 53 cards, got %d", len(deck))
	}
}
