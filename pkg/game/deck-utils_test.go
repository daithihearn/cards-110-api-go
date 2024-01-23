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
			name:       "Full deck two players",
			deck:       NewDeck(),
			numPlayers: 2,
		},
		{
			name:       "Full deck three players",
			deck:       NewDeck(),
			numPlayers: 3,
		},
		{
			name:       "Full deck four players",
			deck:       NewDeck(),
			numPlayers: 4,
		},
		{
			name:       "Full deck five players",
			deck:       NewDeck(),
			numPlayers: 5,
		},
		{
			name:       "Full deck six players",
			deck:       NewDeck(),
			numPlayers: 6,
		},
		{
			name:        "Full deck seven players",
			deck:        NewDeck(),
			numPlayers:  7,
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			deck, dummy, hands, err := DealCards(test.deck, test.numPlayers)

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
				outputCards = append(outputCards, dummy...)
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

func TestDeck_ParseCardName(t *testing.T) {
	tests := []struct {
		name        string
		cardName    string
		expectError bool
		expected    CardName
	}{
		{
			name:        "Empty string",
			cardName:    "",
			expectError: true,
		},
		{
			name:        "Invalid card name",
			cardName:    "INVALID",
			expectError: true,
		},
		{
			name:        "ACE_HEARTS",
			cardName:    "ACE_HEARTS",
			expectError: false,
			expected:    ACE_HEARTS,
		},
		{
			name:        "TWO_HEARTS",
			cardName:    "TWO_HEARTS",
			expectError: false,
			expected:    TWO_HEARTS,
		},
		{
			name:        "THREE_HEARTS",
			cardName:    "THREE_HEARTS",
			expectError: false,
			expected:    THREE_HEARTS,
		},
		{
			name:        "FOUR_HEARTS",
			cardName:    "FOUR_HEARTS",
			expectError: false,
			expected:    FOUR_HEARTS,
		},
		{
			name:        "FIVE_HEARTS",
			cardName:    "FIVE_HEARTS",
			expectError: false,
			expected:    FIVE_HEARTS,
		},
		{
			name:        "SIX_HEARTS",
			cardName:    "SIX_HEARTS",
			expectError: false,
			expected:    SIX_HEARTS,
		},
		{
			name:        "SEVEN_HEARTS",
			cardName:    "SEVEN_HEARTS",
			expectError: false,
			expected:    SEVEN_HEARTS,
		},
		{
			name:        "EIGHT_HEARTS",
			cardName:    "EIGHT_HEARTS",
			expectError: false,
			expected:    EIGHT_HEARTS,
		},
		{
			name:        "NINE_HEARTS",
			cardName:    "NINE_HEARTS",
			expectError: false,
			expected:    NINE_HEARTS,
		},
		{
			name:        "TEN_HEARTS",
			cardName:    "TEN_HEARTS",
			expectError: false,
			expected:    TEN_HEARTS,
		},
		{
			name:        "JACK_HEARTS",
			cardName:    "JACK_HEARTS",
			expectError: false,
			expected:    JACK_HEARTS,
		},
		{
			name:        "QUEEN_HEARTS",
			cardName:    "QUEEN_HEARTS",
			expectError: false,
			expected:    QUEEN_HEARTS,
		},
		{
			name:        "KING_HEARTS",
			cardName:    "KING_HEARTS",
			expectError: false,
			expected:    KING_HEARTS,
		},
		{
			name:        "ACE_DIAMONDS",
			cardName:    "ACE_DIAMONDS",
			expectError: false,
			expected:    ACE_DIAMONDS,
		},
		{
			name:        "TWO_DIAMONDS",
			cardName:    "TWO_DIAMONDS",
			expectError: false,
			expected:    TWO_DIAMONDS,
		},
		{
			name:        "THREE_DIAMONDS",
			cardName:    "THREE_DIAMONDS",
			expectError: false,
			expected:    THREE_DIAMONDS,
		},
		{
			name:        "FOUR_DIAMONDS",
			cardName:    "FOUR_DIAMONDS",
			expectError: false,
			expected:    FOUR_DIAMONDS,
		},
		{
			name:        "FIVE_DIAMONDS",
			cardName:    "FIVE_DIAMONDS",
			expectError: false,
			expected:    FIVE_DIAMONDS,
		},
		{
			name:        "SIX_DIAMONDS",
			cardName:    "SIX_DIAMONDS",
			expectError: false,
			expected:    SIX_DIAMONDS,
		},
		{
			name:        "SEVEN_DIAMONDS",
			cardName:    "SEVEN_DIAMONDS",
			expectError: false,
			expected:    SEVEN_DIAMONDS,
		},
		{
			name:        "EIGHT_DIAMONDS",
			cardName:    "EIGHT_DIAMONDS",
			expectError: false,
			expected:    EIGHT_DIAMONDS,
		},
		{
			name:        "NINE_DIAMONDS",
			cardName:    "NINE_DIAMONDS",
			expectError: false,
			expected:    NINE_DIAMONDS,
		},
		{
			name:        "TEN_DIAMONDS",
			cardName:    "TEN_DIAMONDS",
			expectError: false,
			expected:    TEN_DIAMONDS,
		},
		{
			name:        "JACK_DIAMONDS",
			cardName:    "JACK_DIAMONDS",
			expectError: false,
			expected:    JACK_DIAMONDS,
		},
		{
			name:        "QUEEN_DIAMONDS",
			cardName:    "QUEEN_DIAMONDS",
			expectError: false,
			expected:    QUEEN_DIAMONDS,
		},
		{
			name:        "KING_DIAMONDS",
			cardName:    "KING_DIAMONDS",
			expectError: false,
			expected:    KING_DIAMONDS,
		},
		{
			name:        "ACE_CLUBS",
			cardName:    "ACE_CLUBS",
			expectError: false,
			expected:    ACE_CLUBS,
		},
		{
			name:        "TWO_CLUBS",
			cardName:    "TWO_CLUBS",
			expectError: false,
			expected:    TWO_CLUBS,
		},
		{
			name:        "THREE_CLUBS",
			cardName:    "THREE_CLUBS",
			expectError: false,
			expected:    THREE_CLUBS,
		},
		{
			name:        "FOUR_CLUBS",
			cardName:    "FOUR_CLUBS",
			expectError: false,
			expected:    FOUR_CLUBS,
		},
		{
			name:        "FIVE_CLUBS",
			cardName:    "FIVE_CLUBS",
			expectError: false,
			expected:    FIVE_CLUBS,
		},
		{
			name:        "SIX_CLUBS",
			cardName:    "SIX_CLUBS",
			expectError: false,
			expected:    SIX_CLUBS,
		},
		{
			name:        "SEVEN_CLUBS",
			cardName:    "SEVEN_CLUBS",
			expectError: false,
			expected:    SEVEN_CLUBS,
		},
		{
			name:        "EIGHT_CLUBS",
			cardName:    "EIGHT_CLUBS",
			expectError: false,
			expected:    EIGHT_CLUBS,
		},
		{
			name:        "NINE_CLUBS",
			cardName:    "NINE_CLUBS",
			expectError: false,
			expected:    NINE_CLUBS,
		},
		{
			name:        "TEN_CLUBS",
			cardName:    "TEN_CLUBS",
			expectError: false,
			expected:    TEN_CLUBS,
		},
		{
			name:        "JACK_CLUBS",
			cardName:    "JACK_CLUBS",
			expectError: false,
			expected:    JACK_CLUBS,
		},
		{
			name:        "QUEEN_CLUBS",
			cardName:    "QUEEN_CLUBS",
			expectError: false,
			expected:    QUEEN_CLUBS,
		},
		{
			name:        "KING_CLUBS",
			cardName:    "KING_CLUBS",
			expectError: false,
			expected:    KING_CLUBS,
		},
		{
			name:        "ACE_SPADES",
			cardName:    "ACE_SPADES",
			expectError: false,
			expected:    ACE_SPADES,
		},
		{
			name:        "TWO_SPADES",
			cardName:    "TWO_SPADES",
			expectError: false,
			expected:    TWO_SPADES,
		},
		{
			name:        "THREE_SPADES",
			cardName:    "THREE_SPADES",
			expectError: false,
			expected:    THREE_SPADES,
		},
		{
			name:        "FOUR_SPADES",
			cardName:    "FOUR_SPADES",
			expectError: false,
			expected:    FOUR_SPADES,
		},
		{
			name:        "FIVE_SPADES",
			cardName:    "FIVE_SPADES",
			expectError: false,
			expected:    FIVE_SPADES,
		},
		{
			name:        "SIX_SPADES",
			cardName:    "SIX_SPADES",
			expectError: false,
			expected:    SIX_SPADES,
		},
		{
			name:        "SEVEN_SPADES",
			cardName:    "SEVEN_SPADES",
			expectError: false,
			expected:    SEVEN_SPADES,
		},
		{
			name:        "EIGHT_SPADES",
			cardName:    "EIGHT_SPADES",
			expectError: false,
			expected:    EIGHT_SPADES,
		},
		{
			name:        "NINE_SPADES",
			cardName:    "NINE_SPADES",
			expectError: false,
			expected:    NINE_SPADES,
		},
		{
			name:        "TEN_SPADES",
			cardName:    "TEN_SPADES",
			expectError: false,
			expected:    TEN_SPADES,
		},
		{
			name:        "JACK_SPADES",
			cardName:    "JACK_SPADES",
			expectError: false,
			expected:    JACK_SPADES,
		},
		{
			name:        "QUEEN_SPADES",
			cardName:    "QUEEN_SPADES",
			expectError: false,
			expected:    QUEEN_SPADES,
		},
		{
			name:        "KING_SPADES",
			cardName:    "KING_SPADES",
			expectError: false,
			expected:    KING_SPADES,
		},
		{
			name:        "JOKER",
			cardName:    "JOKER",
			expectError: false,
			expected:    JOKER,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cardName, err := ParseCardName(test.cardName)

			if test.expectError {
				if err == nil {
					t.Errorf("expected an error")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}

				if cardName != test.expected {
					t.Errorf("expected card name to be %s, got %s", test.expected, cardName)
				}
			}
		})
	}
}

func TestDeck_parseCard(t *testing.T) {
	tests := []struct {
		name     string
		cardName CardName
		expected Card
	}{
		{
			name:     "ACE_HEARTS",
			cardName: ACE_HEARTS,
			expected: AceHearts,
		},
		{
			name:     "TWO_HEARTS",
			cardName: TWO_HEARTS,
			expected: TwoHearts,
		},
		{
			name:     "THREE_HEARTS",
			cardName: THREE_HEARTS,
			expected: ThreeHearts,
		},
		{
			name:     "FOUR_HEARTS",
			cardName: FOUR_HEARTS,
			expected: FourHearts,
		},
		{
			name:     "FIVE_HEARTS",
			cardName: FIVE_HEARTS,
			expected: FiveHearts,
		},
		{
			name:     "SIX_HEARTS",
			cardName: SIX_HEARTS,
			expected: SixHearts,
		},
		{
			name:     "SEVEN_HEARTS",
			cardName: SEVEN_HEARTS,
			expected: SevenHearts,
		},
		{
			name:     "EIGHT_HEARTS",
			cardName: EIGHT_HEARTS,
			expected: EightHearts,
		},
		{
			name:     "NINE_HEARTS",
			cardName: NINE_HEARTS,
			expected: NineHearts,
		},
		{
			name:     "TEN_HEARTS",
			cardName: TEN_HEARTS,
			expected: TenHearts,
		},
		{
			name:     "JACK_HEARTS",
			cardName: JACK_HEARTS,
			expected: JackHearts,
		},
		{
			name:     "QUEEN_HEARTS",
			cardName: QUEEN_HEARTS,
			expected: QueenHearts,
		},
		{
			name:     "KING_HEARTS",
			cardName: KING_HEARTS,
			expected: KingHearts,
		},
		{
			name:     "ACE_DIAMONDS",
			cardName: ACE_DIAMONDS,
			expected: AceDiamonds,
		},
		{
			name:     "TWO_DIAMONDS",
			cardName: TWO_DIAMONDS,
			expected: TwoDiamonds,
		},
		{
			name:     "THREE_DIAMONDS",
			cardName: THREE_DIAMONDS,
			expected: ThreeDiamonds,
		},
		{
			name:     "FOUR_DIAMONDS",
			cardName: FOUR_DIAMONDS,
			expected: FourDiamonds,
		},
		{
			name:     "FIVE_DIAMONDS",
			cardName: FIVE_DIAMONDS,
			expected: FiveDiamonds,
		},
		{
			name:     "SIX_DIAMONDS",
			cardName: SIX_DIAMONDS,
			expected: SixDiamonds,
		},
		{
			name:     "SEVEN_DIAMONDS",
			cardName: SEVEN_DIAMONDS,
			expected: SevenDiamonds,
		},
		{
			name:     "EIGHT_DIAMONDS",
			cardName: EIGHT_DIAMONDS,
			expected: EightDiamonds,
		},
		{
			name:     "NINE_DIAMONDS",
			cardName: NINE_DIAMONDS,
			expected: NineDiamonds,
		},
		{
			name:     "TEN_DIAMONDS",
			cardName: TEN_DIAMONDS,
			expected: TenDiamonds,
		},
		{
			name:     "JACK_DIAMONDS",
			cardName: JACK_DIAMONDS,
			expected: JackDiamonds,
		},
		{
			name:     "QUEEN_DIAMONDS",
			cardName: QUEEN_DIAMONDS,
			expected: QueenDiamonds,
		},
		{
			name:     "KING_DIAMONDS",
			cardName: KING_DIAMONDS,
			expected: KingDiamonds,
		},
		{
			name:     "ACE_CLUBS",
			cardName: ACE_CLUBS,
			expected: AceClubs,
		},
		{
			name:     "TWO_CLUBS",
			cardName: TWO_CLUBS,
			expected: TwoClubs,
		},
		{
			name:     "THREE_CLUBS",
			cardName: THREE_CLUBS,
			expected: ThreeClubs,
		},
		{
			name:     "FOUR_CLUBS",
			cardName: FOUR_CLUBS,
			expected: FourClubs,
		},
		{
			name:     "FIVE_CLUBS",
			cardName: FIVE_CLUBS,
			expected: FiveClubs,
		},
		{
			name:     "SIX_CLUBS",
			cardName: SIX_CLUBS,
			expected: SixClubs,
		},
		{
			name:     "SEVEN_CLUBS",
			cardName: SEVEN_CLUBS,
			expected: SevenClubs,
		},
		{
			name:     "EIGHT_CLUBS",
			cardName: EIGHT_CLUBS,
			expected: EightClubs,
		},
		{
			name:     "NINE_CLUBS",
			cardName: NINE_CLUBS,
			expected: NineClubs,
		},
		{
			name:     "TEN_CLUBS",
			cardName: TEN_CLUBS,
			expected: TenClubs,
		},
		{
			name:     "JACK_CLUBS",
			cardName: JACK_CLUBS,
			expected: JackClubs,
		},
		{
			name:     "QUEEN_CLUBS",
			cardName: QUEEN_CLUBS,
			expected: QueenClubs,
		},
		{
			name:     "KING_CLUBS",
			cardName: KING_CLUBS,
			expected: KingClubs,
		},
		{
			name:     "ACE_SPADES",
			cardName: ACE_SPADES,
			expected: AceSpades,
		},
		{
			name:     "TWO_SPADES",
			cardName: TWO_SPADES,
			expected: TwoSpades,
		},
		{
			name:     "THREE_SPADES",
			cardName: THREE_SPADES,
			expected: ThreeSpades,
		},
		{
			name:     "FOUR_SPADES",
			cardName: FOUR_SPADES,
			expected: FourSpades,
		},
		{
			name:     "FIVE_SPADES",
			cardName: FIVE_SPADES,
			expected: FiveSpades,
		},
		{
			name:     "SIX_SPADES",
			cardName: SIX_SPADES,
			expected: SixSpades,
		},
		{
			name:     "SEVEN_SPADES",
			cardName: SEVEN_SPADES,
			expected: SevenSpades,
		},
		{
			name:     "EIGHT_SPADES",
			cardName: EIGHT_SPADES,
			expected: EightSpades,
		},
		{
			name:     "NINE_SPADES",
			cardName: NINE_SPADES,
			expected: NineSpades,
		},
		{
			name:     "TEN_SPADES",
			cardName: TEN_SPADES,
			expected: TenSpades,
		},
		{
			name:     "JACK_SPADES",
			cardName: JACK_SPADES,
			expected: JackSpades,
		},
		{
			name:     "QUEEN_SPADES",
			cardName: QUEEN_SPADES,
			expected: QueenSpades,
		},
		{
			name:     "KING_SPADES",
			cardName: KING_SPADES,
			expected: KingSpades,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			card := test.cardName.Card()
			if card != test.expected {
				t.Errorf("expected card to be %v, got %v", test.expected, card)
			}
		})
	}

}
