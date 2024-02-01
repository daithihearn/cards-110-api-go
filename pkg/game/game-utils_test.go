package game

import "testing"

func TestGameUtils_validateNumberOfPlayers(t *testing.T) {
	tests := []struct {
		name           string
		playerIDs      []string
		expectingError bool
	}{
		{
			name:           "Valid number of players - 2",
			playerIDs:      []string{"1", "2"},
			expectingError: false,
		},
		{
			name:           "Valid number of players - 6",
			playerIDs:      []string{"1", "2", "3", "4", "5", "6"},
			expectingError: false,
		},
		{
			name:           "Invalid number of players - 1",
			playerIDs:      []string{"1"},
			expectingError: true,
		},
		{
			name:           "Invalid number of players - 7",
			playerIDs:      []string{"1", "2", "3", "4", "5", "6", "7"},
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateNumberOfPlayers(test.playerIDs)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			}
		})
	}
}
func TestGameUtils_ParseCall(t *testing.T) {
	tests := []struct {
		name           string
		callStr        string
		expectedCall   Call
		expectingError bool
	}{
		{
			name:         "Valid call - 0",
			callStr:      "0",
			expectedCall: Pass,
		},
		{
			name:         "Valid call - 10",
			callStr:      "10",
			expectedCall: Ten,
		},
		{
			name:         "Valid call - 15",
			callStr:      "15",
			expectedCall: Fifteen,
		},
		{
			name:         "Valid call - 20",
			callStr:      "20",
			expectedCall: Twenty,
		},
		{
			name:         "Valid call - 25",
			callStr:      "25",
			expectedCall: TwentyFive,
		},
		{
			name:         "Valid call - 30",
			callStr:      "30",
			expectedCall: Jink,
		},
		{
			name:           "Invalid call - 5",
			callStr:        "5",
			expectingError: true,
		},
		{
			name:           "Invalid call - 35",
			callStr:        "35",
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			call, err := ParseCall(test.callStr)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if call != test.expectedCall {
					t.Errorf("expected call to be %d, got %d", test.expectedCall, call)
				}
			}
		})
	}
}

func TestGameUtils_shuffle(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "Shuffle 2",
			input: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		},
		{
			name:  "Shuffle 6",
			input: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13"},
		},
		{
			name:  "Shuffle 0",
			input: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shuffled := shuffle(test.input)

			expectedSize := len(test.input)
			if len(shuffled) != expectedSize {
				t.Errorf("expected shuffled size to be %d, got %d", expectedSize, len(shuffled))
			}

			// Check that all the elements are in the shuffled slice
			for _, v := range test.input {
				found := false
				for _, s := range shuffled {
					if v == s {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected %s to be in shuffled slice", v)
				}
			}

			// Expect the order to be different if length > 8
			if len(test.input) > 8 {
				// Check that at least one element is in a different position
				found := false
				for i, v := range test.input {
					if v != shuffled[i] {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected at least one element to be in a different position")
				}
			}

		})
	}
}

func TestGameUtils_createPlayers(t *testing.T) {
	tests := []struct {
		name           string
		playerIDs      []string
		expectingError bool
	}{
		{
			name:           "Valid number of players - 2",
			playerIDs:      []string{"1", "2"},
			expectingError: false,
		},
		{
			name:           "Valid number of players - 6",
			playerIDs:      []string{"1", "2", "3", "4", "5", "6"},
			expectingError: false,
		},
		{
			name:           "Invalid number of players - 1",
			playerIDs:      []string{"1"},
			expectingError: true,
		},
		{
			name:           "Invalid number of players - 7",
			playerIDs:      []string{"1", "2", "3", "4", "5", "6", "7"},
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			players, err := createPlayers(test.playerIDs)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if len(players) != len(test.playerIDs) {
					t.Errorf("expected %d players, got %d", len(test.playerIDs), len(players))
				}
			}
		})
	}
}

func TestGameUtils_canRenage(t *testing.T) {
	tests := []struct {
		name           string
		leadCard       Card
		myTrumps       []Card
		expectedResult bool
	}{
		{
			name:           "Can renage - Jack",
			leadCard:       TwoHearts,
			myTrumps:       []Card{JackHearts},
			expectedResult: true,
		},
		{
			name:           "Can renage - Five",
			leadCard:       TwoHearts,
			myTrumps:       []Card{FiveHearts},
			expectedResult: true,
		},
		{
			name:           "Can renage - Ace Hearts",
			leadCard:       TwoHearts,
			myTrumps:       []Card{AceHearts},
			expectedResult: true,
		},
		{
			name:           "Can renage - Ace Hearts with different suit",
			leadCard:       TwoDiamonds,
			myTrumps:       []Card{AceHearts},
			expectedResult: true,
		},
		{
			name:           "Can renage - Joker",
			leadCard:       TenSpades,
			myTrumps:       []Card{Joker},
			expectedResult: true,
		},
		{
			name:           "Can't renage - Two Hearts",
			leadCard:       FiveHearts,
			myTrumps:       []Card{TwoHearts},
			expectedResult: false,
		},
		{
			name:           "Can't renage - ace of spades",
			leadCard:       FourSpades,
			myTrumps:       []Card{AceSpades},
			expectedResult: false,
		},
		{
			name:           "Can't renage - when higher value trump is lead",
			leadCard:       JackHearts,
			myTrumps:       []Card{Joker},
			expectedResult: false,
		},
		{
			name:           "Can't renage - when higher value trump is lead",
			leadCard:       FiveDiamonds,
			myTrumps:       []Card{JackDiamonds},
			expectedResult: false,
		},
		{
			name:           "Can't renage - when higher value trump is lead",
			leadCard:       Joker,
			myTrumps:       []Card{AceHearts},
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := canRenage(test.leadCard, test.myTrumps)
			if result != test.expectedResult {
				t.Errorf("expected %v, got %v", test.expectedResult, result)
			}
		})
	}
}

func TestGameUtils_isFollowing(t *testing.T) {
	tests := []struct {
		name           string
		myCard         CardName
		myCards        []CardName
		currentHand    Hand
		suit           Suit
		expectedResult bool
	}{
		{
			name:           "Can follow - Jack of Hearts",
			myCard:         JACK_HEARTS,
			myCards:        []CardName{JACK_HEARTS, FIVE_HEARTS, ACE_HEARTS},
			currentHand:    Hand{LeadOut: TWO_HEARTS},
			suit:           Hearts,
			expectedResult: true,
		},
		{
			name:           "Can follow - Ace of Hearts",
			myCard:         ACE_HEARTS,
			myCards:        []CardName{JACK_HEARTS, FIVE_HEARTS, ACE_HEARTS},
			currentHand:    Hand{LeadOut: TWO_HEARTS},
			suit:           Hearts,
			expectedResult: true,
		},
		{
			name:           "Can follow - Joker",
			myCard:         JOKER,
			myCards:        []CardName{JACK_HEARTS, FIVE_HEARTS, ACE_HEARTS, JOKER},
			currentHand:    Hand{LeadOut: TWO_HEARTS},
			suit:           Hearts,
			expectedResult: true,
		},
		{
			name:           "Can follow - Joker",
			myCard:         JOKER,
			myCards:        []CardName{JACK_HEARTS, FIVE_HEARTS, ACE_HEARTS, JOKER},
			currentHand:    Hand{LeadOut: TWO_HEARTS},
			suit:           Hearts,
			expectedResult: true,
		},
		{
			name:           "Can follow - I have no trumps",
			myCard:         JACK_HEARTS,
			myCards:        []CardName{JACK_HEARTS, FIVE_HEARTS, TWO_HEARTS},
			currentHand:    Hand{LeadOut: FIVE_DIAMONDS},
			suit:           Diamonds,
			expectedResult: true,
		},
		{
			name:           "Must follow a cold card",
			myCard:         JACK_HEARTS,
			myCards:        []CardName{JACK_HEARTS, FIVE_HEARTS, TWO_SPADES},
			currentHand:    Hand{LeadOut: FIVE_SPADES},
			suit:           Clubs,
			expectedResult: false,
		},
		{
			name:           "Following",
			myCard:         THREE_CLUBS,
			myCards:        []CardName{TWO_HEARTS, THREE_CLUBS, FOUR_DIAMONDS, FIVE_SPADES, SIX_HEARTS},
			currentHand:    Hand{LeadOut: ACE_CLUBS, PlayedCards: []PlayedCard{{PlayerID: "1", Card: ACE_CLUBS}}},
			suit:           Clubs,
			expectedResult: true,
		},
		{
			name:           "Not following",
			myCard:         FOUR_DIAMONDS,
			myCards:        []CardName{TWO_HEARTS, THREE_CLUBS, FOUR_DIAMONDS, FIVE_SPADES, SIX_HEARTS},
			currentHand:    Hand{LeadOut: ACE_CLUBS, PlayedCards: []PlayedCard{{PlayerID: "1", Card: ACE_CLUBS}}},
			suit:           Clubs,
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isFollowing(test.myCard, test.myCards, test.currentHand, test.suit)
			if result != test.expectedResult {
				t.Errorf("expected %v, got %v", test.expectedResult, result)
			}
		})
	}
}

func TestGameUtils_getActiveSuit(t *testing.T) {
	tests := []struct {
		name           string
		hand           Hand
		suit           Suit
		expectedResult Suit
		expectingError bool
	}{
		{
			name:           "Trump cards played",
			hand:           Hand{LeadOut: JACK_HEARTS, PlayedCards: []PlayedCard{{Card: JACK_HEARTS, PlayerID: "1"}, {Card: FIVE_HEARTS, PlayerID: "2"}, {Card: ACE_HEARTS, PlayerID: "3"}}},
			suit:           Hearts,
			expectedResult: Hearts,
		},
		{
			name:           "Joker played",
			hand:           Hand{LeadOut: JOKER, PlayedCards: []PlayedCard{{Card: JACK_HEARTS, PlayerID: "1"}, {Card: FIVE_HEARTS, PlayerID: "2"}, {Card: JOKER, PlayerID: "3"}}},
			suit:           Diamonds,
			expectedResult: Diamonds,
		},
		{
			name:           "Ace of hearts played",
			hand:           Hand{LeadOut: JACK_CLUBS, PlayedCards: []PlayedCard{{Card: JACK_CLUBS, PlayerID: "1"}, {Card: FIVE_HEARTS, PlayerID: "2"}, {Card: ACE_HEARTS, PlayerID: "3"}}},
			suit:           Diamonds,
			expectedResult: Diamonds,
		},
		{
			name:           "No trump cards played",
			hand:           Hand{LeadOut: JACK_HEARTS, PlayedCards: []PlayedCard{{Card: JACK_HEARTS, PlayerID: "1"}, {Card: FIVE_HEARTS, PlayerID: "2"}, {Card: ACE_DIAMONDS, PlayerID: "3"}}},
			suit:           Clubs,
			expectedResult: Hearts,
		},
		{
			name: "Suit not set",
			hand: Hand{LeadOut: JACK_HEARTS, PlayedCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: FIVE_HEARTS, PlayerID: "2"},
				{Card: ACE_DIAMONDS, PlayerID: "3"},
				{Card: TWO_CLUBS, PlayerID: "4"},
			}},
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := getActiveSuit(test.hand, test.suit)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if result != test.expectedResult {
					t.Errorf("expected %v, got %v", test.expectedResult, result)
				}
			}
		})
	}
}

func TestGameUtils_findWinningCard(t *testing.T) {
	tests := []struct {
		name           string
		hand           Hand
		suit           Suit
		expectedResult PlayedCard
		expectingError bool
	}{
		{
			name:           "No cards played",
			hand:           Hand{LeadOut: JACK_HEARTS, PlayedCards: []PlayedCard{}},
			suit:           Hearts,
			expectedResult: PlayedCard{},
			expectingError: true,
		},
		{
			name:           "No trump cards played",
			hand:           Hand{LeadOut: JACK_HEARTS, PlayedCards: []PlayedCard{{Card: JACK_HEARTS, PlayerID: "1"}, {Card: FIVE_HEARTS, PlayerID: "2"}, {Card: ACE_SPADES, PlayerID: "3"}}},
			suit:           Diamonds,
			expectedResult: PlayedCard{Card: JACK_HEARTS, PlayerID: "1"},
		},
		{
			name:           "Trump cards played",
			hand:           Hand{LeadOut: JACK_HEARTS, PlayedCards: []PlayedCard{{Card: JACK_HEARTS, PlayerID: "1"}, {Card: FIVE_HEARTS, PlayerID: "2"}, {Card: ACE_HEARTS, PlayerID: "3"}}},
			suit:           Hearts,
			expectedResult: PlayedCard{Card: FIVE_HEARTS, PlayerID: "2"},
		},
		{
			name:           "Trump cards played - Joker",
			hand:           Hand{LeadOut: JACK_CLUBS, PlayedCards: []PlayedCard{{Card: JACK_CLUBS, PlayerID: "1"}, {Card: SIX_CLUBS, PlayerID: "2"}, {Card: JOKER, PlayerID: "3"}}},
			suit:           Hearts,
			expectedResult: PlayedCard{Card: JOKER, PlayerID: "3"},
		},
		{
			name:           "Trump cards played - Ace of hearts",
			hand:           Hand{LeadOut: JACK_CLUBS, PlayedCards: []PlayedCard{{Card: JACK_CLUBS, PlayerID: "1"}, {Card: SIX_CLUBS, PlayerID: "2"}, {Card: ACE_HEARTS, PlayerID: "3"}}},
			suit:           Hearts,
			expectedResult: PlayedCard{Card: ACE_HEARTS, PlayerID: "3"},
		},
		{
			name:           "Ten of trumps beats jack of cold suit",
			hand:           Hand{LeadOut: JACK_SPADES, PlayedCards: []PlayedCard{{Card: JACK_SPADES, PlayerID: "1"}, {Card: TEN_DIAMONDS, PlayerID: "2"}}},
			suit:           Diamonds,
			expectedResult: PlayedCard{Card: TEN_DIAMONDS, PlayerID: "2"},
		},
		{
			name:           "real world scenario",
			hand:           Hand{LeadOut: ACE_SPADES, PlayedCards: []PlayedCard{{Card: ACE_SPADES, PlayerID: "2"}, {Card: FIVE_SPADES, PlayerID: "1"}}},
			suit:           "SPADES",
			expectedResult: PlayedCard{Card: FIVE_SPADES, PlayerID: "1"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := findWinningCard(test.hand, test.suit)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if result != test.expectedResult {
					t.Errorf("expected %v, got %v", test.expectedResult, result)
				}
			}
		})
	}
}

func TestGameUtils_findWinningCardsForRound(t *testing.T) {
	tests := []struct {
		name           string
		round          Round
		expectedResult []PlayedCard
		expectingError bool
	}{
		{
			name: "Round complete",
			round: Round{
				Suit: Hearts,
				CompletedHands: []Hand{
					{LeadOut: JACK_HEARTS, PlayedCards: []PlayedCard{{Card: JACK_HEARTS, PlayerID: "1"}, {Card: FIVE_HEARTS, PlayerID: "2"}, {Card: ACE_HEARTS, PlayerID: "3"}}},
					{LeadOut: TWO_SPADES, PlayedCards: []PlayedCard{{Card: TWO_SPADES, PlayerID: "1"}, {Card: SIX_SPADES, PlayerID: "2"}, {Card: ACE_SPADES, PlayerID: "3"}}},
					{LeadOut: SIX_DIAMONDS, PlayedCards: []PlayedCard{{Card: TWO_DIAMONDS, PlayerID: "1"}, {Card: SIX_DIAMONDS, PlayerID: "2"}, {Card: ACE_DIAMONDS, PlayerID: "3"}}},
					{LeadOut: ACE_CLUBS, PlayedCards: []PlayedCard{{Card: TWO_CLUBS, PlayerID: "1"}, {Card: SIX_CLUBS, PlayerID: "2"}, {Card: ACE_CLUBS, PlayerID: "3"}}},
					{LeadOut: SEVEN_HEARTS, PlayedCards: []PlayedCard{{Card: TWO_HEARTS, PlayerID: "1"}, {Card: SIX_HEARTS, PlayerID: "2"}, {Card: SEVEN_HEARTS, PlayerID: "3"}}},
				},
			},
			expectedResult: []PlayedCard{{Card: FIVE_HEARTS, PlayerID: "2"}, {Card: ACE_SPADES, PlayerID: "3"}, {Card: SIX_DIAMONDS, PlayerID: "2"}, {Card: ACE_CLUBS, PlayerID: "3"}, {Card: SEVEN_HEARTS, PlayerID: "3"}},
		},
		{
			name: "Round not complete",
			round: Round{
				Suit: Hearts,
				CompletedHands: []Hand{
					{LeadOut: JACK_HEARTS, PlayedCards: []PlayedCard{{Card: JACK_HEARTS, PlayerID: "1"}, {Card: FIVE_HEARTS, PlayerID: "2"}, {Card: ACE_HEARTS, PlayerID: "3"}}},
					{LeadOut: TWO_SPADES, PlayedCards: []PlayedCard{{Card: TWO_SPADES, PlayerID: "1"}, {Card: SIX_SPADES, PlayerID: "2"}, {Card: ACE_SPADES, PlayerID: "3"}}},
					{LeadOut: SIX_DIAMONDS, PlayedCards: []PlayedCard{{Card: TWO_DIAMONDS, PlayerID: "1"}, {Card: SIX_DIAMONDS, PlayerID: "2"}, {Card: ACE_DIAMONDS, PlayerID: "3"}}},
					{LeadOut: ACE_CLUBS, PlayedCards: []PlayedCard{{Card: TWO_CLUBS, PlayerID: "1"}, {Card: SIX_CLUBS, PlayerID: "2"}, {Card: ACE_CLUBS, PlayerID: "3"}}},
				},
			},
			expectingError: true,
		},
		{
			name: "No suit set",
			round: Round{CompletedHands: []Hand{
				{LeadOut: JACK_HEARTS, PlayedCards: []PlayedCard{{Card: JACK_HEARTS, PlayerID: "1"}, {Card: FIVE_HEARTS, PlayerID: "2"}, {Card: ACE_HEARTS, PlayerID: "3"}}},
				{LeadOut: TWO_SPADES, PlayedCards: []PlayedCard{{Card: TWO_SPADES, PlayerID: "1"}, {Card: SIX_SPADES, PlayerID: "2"}, {Card: ACE_SPADES, PlayerID: "3"}}},
				{LeadOut: SIX_DIAMONDS, PlayedCards: []PlayedCard{{Card: TWO_DIAMONDS, PlayerID: "1"}, {Card: SIX_DIAMONDS, PlayerID: "2"}, {Card: ACE_DIAMONDS, PlayerID: "3"}}},
				{LeadOut: ACE_CLUBS, PlayedCards: []PlayedCard{{Card: TWO_CLUBS, PlayerID: "1"}, {Card: SIX_CLUBS, PlayerID: "2"}, {Card: ACE_CLUBS, PlayerID: "3"}}},
				{LeadOut: SEVEN_HEARTS, PlayedCards: []PlayedCard{{Card: TWO_HEARTS, PlayerID: "1"}, {Card: SIX_HEARTS, PlayerID: "2"}, {Card: SEVEN_HEARTS, PlayerID: "3"}}},
			}},
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := findWinningCardsForRound(test.round)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if len(result) != len(test.expectedResult) {
					t.Errorf("expected %d cards, got %d", len(test.expectedResult), len(result))
				}
				// Check the cards are the same
				for i, v := range result {
					if v != test.expectedResult[i] {
						t.Errorf("expected %v, got %v", test.expectedResult[i], v)
					}
				}
			}
		})
	}
}

func TestGameUtils_checkForJink(t *testing.T) {
	tests := []struct {
		name           string
		winningCards   []PlayedCard
		players        []Player
		goerID         string
		expectedResult bool
		expectingError bool
	}{
		{
			name: "30 but called 20 - doubles",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: JACK_DIAMONDS, PlayerID: "1"},
				{Card: JACK_CLUBS, PlayerID: "4"},
				{Card: JACK_SPADES, PlayerID: "4"},
				{Card: JACK_HEARTS, PlayerID: "1"},
			},
			players: []Player{
				{ID: "1", TeamID: "1", Call: Twenty},
				{ID: "2", TeamID: "2"},
				{ID: "3", TeamID: "3"},
				{ID: "4", TeamID: "1"},
				{ID: "5", TeamID: "2"},
				{ID: "6", TeamID: "3"},
			},
			goerID:         "1",
			expectedResult: false,
		},
		{
			name: "Jink- doubles",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: JACK_DIAMONDS, PlayerID: "1"},
				{Card: JACK_CLUBS, PlayerID: "4"},
				{Card: JACK_SPADES, PlayerID: "4"},
				{Card: JACK_HEARTS, PlayerID: "1"},
			},
			players: []Player{
				{ID: "1", TeamID: "1", Call: Jink},
				{ID: "2", TeamID: "2"},
				{ID: "3", TeamID: "3"},
				{ID: "4", TeamID: "1", Call: Ten},
				{ID: "5", TeamID: "2", Call: Twenty},
				{ID: "6", TeamID: "3", Call: TwentyFive},
			},
			goerID:         "1",
			expectedResult: true,
		},
		{
			name: "Jink - two player game",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: JACK_DIAMONDS, PlayerID: "1"},
				{Card: JACK_CLUBS, PlayerID: "1"},
				{Card: JACK_SPADES, PlayerID: "1"},
				{Card: JACK_HEARTS, PlayerID: "1"},
			},
			players: []Player{
				{ID: "1", TeamID: "1", Call: Jink},
				{ID: "2", TeamID: "2"},
			},
			goerID:         "1",
			expectedResult: false,
		},
		{
			name: "30 but called 20 - two player game",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: JACK_DIAMONDS, PlayerID: "1"},
				{Card: JACK_CLUBS, PlayerID: "1"},
				{Card: JACK_SPADES, PlayerID: "1"},
				{Card: JACK_HEARTS, PlayerID: "1"},
			},
			players: []Player{
				{ID: "1", TeamID: "1", Call: Twenty},
				{ID: "2", TeamID: "2"},
			},
			goerID:         "1",
			expectedResult: false,
		},
		{
			name: "Jink - three player game",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: JACK_DIAMONDS, PlayerID: "1"},
				{Card: JACK_CLUBS, PlayerID: "1"},
				{Card: JACK_SPADES, PlayerID: "1"},
				{Card: JACK_HEARTS, PlayerID: "1"},
			},
			players: []Player{
				{ID: "1", TeamID: "1", Call: Jink},
				{ID: "2", TeamID: "2", Call: TwentyFive},
				{ID: "3", TeamID: "3"},
			},
			goerID:         "1",
			expectedResult: true,
		},
		{
			name: "30 but called 15 - three player game",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: JACK_DIAMONDS, PlayerID: "1"},
				{Card: JACK_CLUBS, PlayerID: "1"},
				{Card: JACK_SPADES, PlayerID: "1"},
				{Card: JACK_HEARTS, PlayerID: "1"},
			},
			players: []Player{
				{ID: "1", TeamID: "1", Call: Fifteen},
				{ID: "2", TeamID: "2"},
				{ID: "3", TeamID: "3"},
			},
			goerID:         "1",
			expectedResult: false,
		},
		{
			name: "Jink - four player game",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "4"},
				{Card: JACK_DIAMONDS, PlayerID: "4"},
				{Card: JACK_CLUBS, PlayerID: "4"},
				{Card: JACK_SPADES, PlayerID: "4"},
				{Card: JACK_HEARTS, PlayerID: "4"},
			},
			players: []Player{
				{ID: "1", TeamID: "1"},
				{ID: "2", TeamID: "2"},
				{ID: "3", TeamID: "3"},
				{ID: "4", TeamID: "4", Call: Jink},
			},
			goerID:         "4",
			expectedResult: true,
		},
		{
			name: "30 but called 25 - four player game",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "4"},
				{Card: JACK_DIAMONDS, PlayerID: "4"},
				{Card: JACK_CLUBS, PlayerID: "4"},
				{Card: JACK_SPADES, PlayerID: "4"},
				{Card: JACK_HEARTS, PlayerID: "4"},
			},
			players: []Player{
				{ID: "1", TeamID: "1", Call: TwentyFive},
				{ID: "2", TeamID: "2"},
				{ID: "3", TeamID: "3"},
				{ID: "4", TeamID: "4"},
			},
			goerID:         "4",
			expectedResult: false,
		},
		{
			name: "Jink - five player game",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "3"},
				{Card: JACK_DIAMONDS, PlayerID: "3"},
				{Card: JACK_CLUBS, PlayerID: "3"},
				{Card: JACK_SPADES, PlayerID: "3"},
				{Card: JACK_HEARTS, PlayerID: "3"},
			},
			players: []Player{
				{ID: "1", TeamID: "1"},
				{ID: "2", TeamID: "2"},
				{ID: "3", TeamID: "3", Call: Jink},
				{ID: "4", TeamID: "4"},
				{ID: "5", TeamID: "5"},
			},
			goerID:         "3",
			expectedResult: true,
		},
		{
			name: "30 but called 25 - five player game",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "3"},
				{Card: JACK_DIAMONDS, PlayerID: "3"},
				{Card: JACK_CLUBS, PlayerID: "3"},
				{Card: JACK_SPADES, PlayerID: "3"},
				{Card: JACK_HEARTS, PlayerID: "3"},
			},
			players: []Player{
				{ID: "1", TeamID: "1", Call: 25},
				{ID: "2", TeamID: "2"},
				{ID: "3", TeamID: "3"},
				{ID: "4", TeamID: "4"},
				{ID: "5", TeamID: "5"},
			},
			goerID:         "3",
			expectedResult: false,
		},
		{
			name: "Jink - no jink",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: JACK_DIAMONDS, PlayerID: "1"},
				{Card: JACK_CLUBS, PlayerID: "3"},
				{Card: JACK_SPADES, PlayerID: "2"},
				{Card: JACK_HEARTS, PlayerID: "1"},
			},
			players: []Player{
				{ID: "1", TeamID: "1"},
				{ID: "2", TeamID: "2"},
				{ID: "3", TeamID: "3"},
				{ID: "4", TeamID: "1"},
				{ID: "5", TeamID: "2"},
				{ID: "6", TeamID: "3"},
			},
			goerID:         "1",
			expectedResult: false,
		},
		{
			name: "invalid number of cards",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: JACK_DIAMONDS, PlayerID: "1"},
				{Card: JACK_CLUBS, PlayerID: "2"},
				{Card: JACK_SPADES, PlayerID: "2"},
			},
			players: []Player{
				{ID: "1", TeamID: "1", Call: Jink},
				{ID: "2", TeamID: "2"},
				{ID: "3", TeamID: "3"},
			},
			goerID:         "1",
			expectingError: true,
		},
		{
			name: "invalid player ID",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "invalid"},
				{Card: JACK_DIAMONDS, PlayerID: "1"},
				{Card: JACK_CLUBS, PlayerID: "2"},
				{Card: JACK_SPADES, PlayerID: "2"},
				{Card: JACK_HEARTS, PlayerID: "1"},
			},
			players: []Player{
				{ID: "1", TeamID: "1", Call: Jink},
				{ID: "2", TeamID: "2"},
				{ID: "3", TeamID: "3"},
			},
			goerID:         "1",
			expectingError: true,
		},
		{
			name: "invalid player ID",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: JACK_DIAMONDS, PlayerID: "invalid"},
				{Card: JACK_CLUBS, PlayerID: "2"},
				{Card: JACK_SPADES, PlayerID: "2"},
				{Card: JACK_HEARTS, PlayerID: "1"},
			},
			players: []Player{
				{ID: "1", TeamID: "1", Call: Jink},
				{ID: "2", TeamID: "2"},
				{ID: "3", TeamID: "3"},
			},
			goerID:         "1",
			expectingError: true,
		},
		{
			name: "jink but not for going team",
			winningCards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: JACK_DIAMONDS, PlayerID: "1"},
				{Card: JACK_CLUBS, PlayerID: "1"},
				{Card: JACK_SPADES, PlayerID: "1"},
				{Card: JACK_HEARTS, PlayerID: "1"},
			},
			players: []Player{
				{ID: "1", TeamID: "1"},
				{ID: "2", TeamID: "2"},
				{ID: "3", TeamID: "3"},
			},
			goerID:         "2",
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := checkForJink(test.winningCards, test.players, test.goerID)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if result != test.expectedResult {
					t.Errorf("expected %v, got %v", test.expectedResult, result)
				}
			}
		})
	}
}

func TestGameUtils_getTeamID(t *testing.T) {
	tests := []struct {
		name           string
		playerID       string
		players        []Player
		expectedResult string
		expectingError bool
	}{
		{
			name:           "Team 1",
			playerID:       "1",
			players:        []Player{{ID: "1", TeamID: "1"}},
			expectedResult: "1",
		},
		{
			name:           "Team 2",
			playerID:       "2",
			players:        []Player{{ID: "1", TeamID: "1"}, {ID: "2", TeamID: "2"}},
			expectedResult: "2",
		},
		{
			name:           "Team 3",
			playerID:       "3",
			players:        []Player{{ID: "1", TeamID: "1"}, {ID: "2", TeamID: "2"}, {ID: "3", TeamID: "3"}},
			expectedResult: "3",
		},
		{
			name:           "Team 4",
			playerID:       "4",
			players:        []Player{{ID: "1", TeamID: "1"}, {ID: "2", TeamID: "2"}, {ID: "3", TeamID: "3"}, {ID: "4", TeamID: "4"}},
			expectedResult: "4",
		},
		{
			name:           "Team 5",
			playerID:       "5",
			players:        []Player{{ID: "1", TeamID: "1"}, {ID: "2", TeamID: "2"}, {ID: "3", TeamID: "3"}, {ID: "4", TeamID: "4"}, {ID: "5", TeamID: "5"}},
			expectedResult: "5",
		},
		{
			name:           "Team 6",
			playerID:       "6",
			players:        []Player{{ID: "1", TeamID: "1"}, {ID: "2", TeamID: "2"}, {ID: "3", TeamID: "3"}, {ID: "4", TeamID: "1"}, {ID: "5", TeamID: "2"}, {ID: "6", TeamID: "3"}},
			expectedResult: "3",
		},
		{
			name:           "Invalid player ID",
			playerID:       "invalid",
			players:        []Player{{ID: "1", TeamID: "1"}, {ID: "2", TeamID: "2"}, {ID: "3", TeamID: "3"}, {ID: "4", TeamID: "1"}, {ID: "5", TeamID: "2"}, {ID: "6", TeamID: "3"}},
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := getTeamID(test.playerID, test.players)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if result != test.expectedResult {
					t.Errorf("expected %v, got %v", test.expectedResult, result)
				}
			}
		})
	}
}

func TestGameUtils_findWinningTeam(t *testing.T) {
	tests := []struct {
		name           string
		players        []Player
		round          Round
		expectedResult string
		expectingError bool
	}{
		{
			name: "Team 1 wins",
			players: []Player{
				{ID: "1", TeamID: "1", Score: 110},
				{ID: "2", TeamID: "2", Score: 0},
			},
			round:          Round{},
			expectedResult: "1",
		},
		{
			name: "Team 2 wins",
			players: []Player{
				{ID: "1", TeamID: "1", Score: 0},
				{ID: "2", TeamID: "2", Score: 110},
				{ID: "3", TeamID: "3", Score: 100},
				{ID: "4", TeamID: "4", Score: 105},
			},
			expectedResult: "2",
		},
		{
			name: "Two teams over 110 - the goer is one of them",
			players: []Player{
				{ID: "1", TeamID: "1", Score: 110},
				{ID: "2", TeamID: "2", Score: 0},
				{ID: "3", TeamID: "3", Score: 120},
			},
			round:          Round{GoerID: "1"},
			expectedResult: "1",
		},
		{
			name: "Two teams over 110 - the goer is not one of them",
			players: []Player{
				{ID: "1", TeamID: "1", Score: 110},
				{ID: "2", TeamID: "2", Score: 0},
				{ID: "3", TeamID: "3", Score: 120},
			},
			round: Round{GoerID: "2", Suit: Clubs, CompletedHands: []Hand{
				{LeadOut: JACK_HEARTS, PlayedCards: []PlayedCard{{Card: JACK_HEARTS, PlayerID: "1"}, {Card: TWO_SPADES, PlayerID: "2"}, {Card: ACE_HEARTS, PlayerID: "3"}}},
				{LeadOut: TWO_SPADES, PlayedCards: []PlayedCard{{Card: SEVEN_CLUBS, PlayerID: "1"}, {Card: THREE_SPADES, PlayerID: "2"}, {Card: SEVEN_CLUBS, PlayerID: "3"}}},
				{LeadOut: THREE_CLUBS, PlayedCards: []PlayedCard{{Card: JOKER, PlayerID: "1"}, {Card: FOUR_SPADES, PlayerID: "2"}, {Card: THREE_CLUBS, PlayerID: "3"}}},
				{LeadOut: TWO_CLUBS, PlayedCards: []PlayedCard{{Card: ACE_CLUBS, PlayerID: "1"}, {Card: TWO_CLUBS, PlayerID: "2"}, {Card: FIVE_HEARTS, PlayerID: "3"}}},
				{LeadOut: JACK_CLUBS, PlayedCards: []PlayedCard{{Card: JACK_CLUBS, PlayerID: "1"}, {Card: SIX_SPADES, PlayerID: "2"}, {Card: ACE_SPADES, PlayerID: "3"}}},
			}},
			expectedResult: "3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := findWinningTeam(test.players, test.round)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if result != test.expectedResult {
					t.Errorf("expected %v, got %v", test.expectedResult, result)
				}
			}
		})
	}
}

func TestGameUtils_getTeamsOver110(t *testing.T) {
	tests := []struct {
		name           string
		players        []Player
		expectedResult map[string]bool
	}{
		{
			name: "1 team over 110",
			players: []Player{
				{ID: "1", TeamID: "1", Score: 110},
				{ID: "2", TeamID: "2", Score: 0},
				{ID: "3", TeamID: "3", Score: 100},
			},
			expectedResult: map[string]bool{"1": true},
		},
		{
			name: "2 teams over 110",
			players: []Player{
				{ID: "1", TeamID: "1", Score: 110},
				{ID: "2", TeamID: "2", Score: 0},
				{ID: "3", TeamID: "3", Score: 120},
			},
			expectedResult: map[string]bool{"1": true, "3": true},
		},
		{
			name: "no teams over 110",
			players: []Player{
				{ID: "1", TeamID: "1", Score: 100},
				{ID: "2", TeamID: "2", Score: 0},
				{ID: "3", TeamID: "3", Score: 105},
			},
			expectedResult: map[string]bool{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := getTeamsOver110(test.players)
			if len(result) != len(test.expectedResult) {
				t.Errorf("expected %d teams, got %d", len(test.expectedResult), len(result))
			}
			for k, v := range result {
				if test.expectedResult[k] != v {
					t.Errorf("expected %v, got %v", test.expectedResult[k], v)
				}
			}
		})
	}
}

func TestGameUtils_findBestTrump(t *testing.T) {
	tests := []struct {
		name           string
		cards          []PlayedCard
		suit           Suit
		expectedResult PlayedCard
		expectTrump    bool
		expectingError bool
	}{
		{
			name: "Jack of Hearts",
			cards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: JACK_DIAMONDS, PlayerID: "2"},
				{Card: JACK_CLUBS, PlayerID: "3"},
			},
			suit:           Hearts,
			expectedResult: PlayedCard{Card: JACK_HEARTS, PlayerID: "1"},
			expectTrump:    true,
		},
		{
			name: "Joker",
			cards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: ACE_HEARTS, PlayerID: "2"},
				{Card: JOKER, PlayerID: "3"},
			},
			suit:           Diamonds,
			expectedResult: PlayedCard{Card: JOKER, PlayerID: "3"},
			expectTrump:    true,
		},
		{
			name: "Ace of Hearts",
			cards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: ACE_HEARTS, PlayerID: "2"},
				{Card: ACE_CLUBS, PlayerID: "3"},
			},
			suit:           Clubs,
			expectedResult: PlayedCard{Card: ACE_HEARTS, PlayerID: "2"},
			expectTrump:    true,
		},
		{
			name: "No trump cards",
			cards: []PlayedCard{
				{Card: JACK_HEARTS, PlayerID: "1"},
				{Card: TWO_SPADES, PlayerID: "2"},
				{Card: ACE_CLUBS, PlayerID: "3"},
			},
			suit:        Diamonds,
			expectTrump: false,
		},
		{
			name:           "No cards played",
			cards:          []PlayedCard{},
			suit:           Clubs,
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, trump, err := findBestTrump(test.cards, test.suit)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if result != test.expectedResult {
					t.Errorf("expected %v, got %v", test.expectedResult, result)
				}
				if trump != test.expectTrump {
					t.Errorf("expected %v, got %v", test.expectTrump, trump)
				}
			}
		})
	}
}

func TestGameUtils_findPlayer(t *testing.T) {
	tests := []struct {
		name           string
		playerID       string
		players        []Player
		expectedResult Player
		expectingError bool
	}{
		{
			name:           "Player 1",
			playerID:       "1",
			players:        []Player{{ID: "1", TeamID: "1"}, {ID: "2", TeamID: "2"}},
			expectedResult: Player{ID: "1", TeamID: "1"},
		},
		{
			name:           "Player 2",
			playerID:       "2",
			players:        []Player{{ID: "1", TeamID: "1"}, {ID: "2", TeamID: "2"}},
			expectedResult: Player{ID: "2", TeamID: "2"},
		},
		{
			name:           "Player not found",
			playerID:       "3",
			players:        []Player{{ID: "1", TeamID: "1"}, {ID: "2", TeamID: "2"}},
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := findPlayer(test.playerID, test.players)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if result.ID != test.expectedResult.ID {
					t.Errorf("expected %v, got %v", test.expectedResult, result)
				}
			}
		})
	}
}

func TestGameUtils_CalculateScores(t *testing.T) {
	tests := []struct {
		name           string
		winningCards   []PlayedCard
		players        []Player
		suit           Suit
		expectedResult map[string]int
		expectingError bool
	}{
		{
			name:         "Add up scores",
			winningCards: []PlayedCard{{Card: JACK_HEARTS, PlayerID: "1"}, {Card: JACK_DIAMONDS, PlayerID: "2"}, {Card: JACK_CLUBS, PlayerID: "4"}, {Card: JACK_SPADES, PlayerID: "4"}, {Card: FIVE_HEARTS, PlayerID: "1"}},
			players: []Player{
				{ID: "1", TeamID: "1", Score: 0},
				{ID: "2", TeamID: "2", Score: 0},
				{ID: "3", TeamID: "3", Score: 0},
				{ID: "4", TeamID: "4", Score: 0},
			},
			suit: Hearts,
			expectedResult: map[string]int{
				"1": 15,
				"2": 5,
				"4": 10,
			},
		},
		{
			name:         "Team game",
			winningCards: []PlayedCard{{Card: JACK_HEARTS, PlayerID: "1"}, {Card: JACK_DIAMONDS, PlayerID: "2"}, {Card: JACK_CLUBS, PlayerID: "4"}, {Card: JACK_SPADES, PlayerID: "4"}, {Card: FIVE_HEARTS, PlayerID: "1"}},
			players: []Player{
				{ID: "1", TeamID: "1", Score: 0},
				{ID: "2", TeamID: "2", Score: 0},
				{ID: "3", TeamID: "3", Score: 0},
				{ID: "4", TeamID: "1", Score: 0},
				{ID: "5", TeamID: "2", Score: 0},
				{ID: "6", TeamID: "3", Score: 0},
			},
			suit: Hearts,
			expectedResult: map[string]int{
				"1": 25,
				"2": 5,
			},
		},
		{
			name:         "No trump cards",
			winningCards: []PlayedCard{{Card: JACK_HEARTS, PlayerID: "1"}, {Card: JACK_DIAMONDS, PlayerID: "2"}, {Card: JACK_CLUBS, PlayerID: "4"}, {Card: TWO_CLUBS, PlayerID: "4"}, {Card: FIVE_HEARTS, PlayerID: "1"}},
			players: []Player{
				{ID: "1", TeamID: "1", Score: 0},
				{ID: "2", TeamID: "2", Score: 0},
				{ID: "3", TeamID: "3", Score: 0},
				{ID: "4", TeamID: "4", Score: 0},
			},
			suit: Spades,
			expectedResult: map[string]int{
				"1": 10,
				"2": 5,
				"4": 10,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := calculateScores(test.winningCards, test.players, test.suit)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}

				for k, v := range result {
					if test.expectedResult[k] != v {
						t.Errorf("expected %v, got %v", test.expectedResult[k], v)
					}
				}
			}
		})
	}

}
