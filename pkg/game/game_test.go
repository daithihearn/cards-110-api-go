package game

import "testing"

type CallWrap struct {
	playerID             string
	call                 Call
	expectedNextPlayerID string
	expectingError       bool
	expectedRevision     int
}

func TestGame_Call(t *testing.T) {
	tests := []struct {
		name           string
		game           Game
		calls          []CallWrap
		expectedGoerID string
	}{
		{
			name: "simple call",
			game: TwoPlayerGame(),
			calls: []CallWrap{{
				playerID:             "2",
				call:                 Fifteen,
				expectedNextPlayerID: "1",
				expectedRevision:     1,
			}},
		},
		{
			name: "Completed game",
			game: CompletedGame(),
			calls: []CallWrap{{
				playerID:         "2",
				call:             Fifteen,
				expectingError:   true,
				expectedRevision: 0,
			}},
		},
		{
			name: "Game in called state",
			game: CalledGame(),
			calls: []CallWrap{{
				playerID:         "2",
				call:             Fifteen,
				expectingError:   true,
				expectedRevision: 0,
			}},
		},
		{
			name: "Invalid call 10 in a 2 player game",
			game: TwoPlayerGame(),
			calls: []CallWrap{{
				playerID:         "2",
				call:             Ten,
				expectingError:   true,
				expectedRevision: 0,
			}},
		},
		{
			name: "Valid 10 call in 6 player game",
			game: SixPlayerGame(),
			calls: []CallWrap{{
				playerID:             "5",
				call:                 Ten,
				expectedNextPlayerID: "6",
				expectedRevision:     1,
			}},
		},
		{
			name: "Call too low",
			game: TwoPlayerGame(),
			calls: []CallWrap{
				{
					playerID:             "2",
					call:                 Twenty,
					expectedNextPlayerID: "1",
					expectedRevision:     1,
				},
				{
					playerID:         "1",
					call:             Fifteen,
					expectingError:   true,
					expectedRevision: 1,
				},
			},
		},
		{
			name: "Dealer seeing call",
			game: TwoPlayerGame(),
			calls: []CallWrap{
				{
					playerID:             "2",
					call:                 Twenty,
					expectedNextPlayerID: "1",
					expectedRevision:     1,
				},
				{
					playerID:             "1",
					call:                 Twenty,
					expectedNextPlayerID: "2",
					expectedRevision:     2,
				},
			},
		},
		{
			name: "Dealer seeing call, dealer passes",
			game: TwoPlayerGame(),
			calls: []CallWrap{
				{
					playerID:             "2",
					call:                 Fifteen,
					expectedNextPlayerID: "1",
					expectedRevision:     1,
				},
				{
					playerID:             "1",
					call:                 Twenty,
					expectedNextPlayerID: "1",
					expectedRevision:     2,
				},
			},
			expectedGoerID: "1",
		},
		{
			name: "All players pass",
			game: TwoPlayerGame(),
			calls: []CallWrap{
				{
					playerID:             "2",
					call:                 Pass,
					expectedNextPlayerID: "1",
					expectedRevision:     1,
				},
				{
					playerID:             "1",
					call:                 Pass,
					expectedNextPlayerID: "1",
					expectedRevision:     2,
				},
			},
		},
		{
			name: "Player calls JINK in 6 player game",
			game: SixPlayerGame(),
			calls: []CallWrap{
				{
					playerID:             "5",
					call:                 Jink,
					expectedNextPlayerID: "1",
					expectedRevision:     1,
				},
			},
		},
		{
			name: "Dealer takes JINK in 6 player game",
			game: SixPlayerGame(),
			calls: []CallWrap{
				{
					playerID:             "5",
					call:                 Jink,
					expectedNextPlayerID: "1",
					expectedRevision:     1,
				},
				{
					playerID:             "1",
					call:                 Jink,
					expectedNextPlayerID: "1",
					expectedRevision:     2,
				},
			},
			expectedGoerID: "1",
		},
		{
			name: "Dealer lets JINK go in 6 player game",
			game: SixPlayerGame(),
			calls: []CallWrap{
				{
					playerID:             "5",
					call:                 Jink,
					expectedNextPlayerID: "1",
					expectedRevision:     1,
				},
				{
					playerID:             "1",
					call:                 Pass,
					expectedNextPlayerID: "5",
					expectedRevision:     2,
				},
			},
			expectedGoerID: "5",
		},
		{
			name: "Dealer takes caller, caller increases bet and dealer passes",
			game: TwoPlayerGame(),
			calls: []CallWrap{
				{
					playerID:             "2",
					call:                 Fifteen,
					expectedNextPlayerID: "1",
					expectedRevision:     1,
				},
				{
					playerID:             "1",
					call:                 Fifteen,
					expectedNextPlayerID: "2",
					expectedRevision:     2,
				},
				{
					playerID:             "2",
					call:                 Twenty,
					expectedNextPlayerID: "1",
					expectedRevision:     3,
				},
				{
					playerID:             "1",
					call:                 Pass,
					expectedNextPlayerID: "2",
					expectedRevision:     4,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, call := range test.calls {
				err := test.game.Call(call.playerID, call.call)
				if call.expectingError {
					if err == nil {
						t.Errorf("expected an error, got nil")
					}
				} else {
					var player Player
					for _, p := range test.game.Players {
						if p.ID == call.playerID {
							player = p
							break
						}
					}
					if player.Call != call.call {
						t.Errorf("expected player %s to call %d, got %d", call.playerID, call.call, test.game.Players[0].Call)
					}
					if test.game.CurrentRound.CurrentHand.CurrentPlayerID != call.expectedNextPlayerID {
						t.Errorf("expected player %s to be current player, got %s", call.expectedNextPlayerID, test.game.CurrentRound.CurrentHand.CurrentPlayerID)
					}
					// Check revision has been incremented
					if test.game.Revision != call.expectedRevision {
						t.Errorf("expected revision %d, got %d", call.expectedRevision, test.game.Revision)
					}
				}

			}
			if test.expectedGoerID != "" {
				if test.game.CurrentRound.GoerID != test.expectedGoerID {
					t.Errorf("expected goer to be %s, got %s", test.expectedGoerID, test.game.CurrentRound.GoerID)
				}
			}

		})
	}
}

func TestGame_MinKeep(t *testing.T) {
	tests := []struct {
		name            string
		game            Game
		expectedMinKeep int
		errorExpected   bool
	}{
		{
			name:            "2 player game",
			game:            TwoPlayerGame(),
			expectedMinKeep: 0,
		},
		{
			name:            "3 player game",
			game:            ThreePlayerGame(),
			expectedMinKeep: 0,
		},
		{
			name:            "4 player game",
			game:            FourPlayerGame(),
			expectedMinKeep: 0,
		},
		{
			name:            "5 player game",
			game:            FivePlayerGame(),
			expectedMinKeep: 1,
		},
		{
			name:            "6 player game",
			game:            SixPlayerGame(),
			expectedMinKeep: 2,
		},
		{
			name:          "Invalid number of players - 1",
			game:          OnePlayerGame(),
			errorExpected: true,
		},
		{
			name:          "Invalid number of players - 7",
			game:          SevenPlayerGame(),
			errorExpected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			minKeep, err := test.game.MinKeep()
			if test.errorExpected {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if minKeep != test.expectedMinKeep {
					t.Errorf("expected minKeep to be %d, got %d", test.expectedMinKeep, minKeep)
				}
			}
		})
	}
}

func TestGame_SelectSuit(t *testing.T) {
	tests := []struct {
		name             string
		game             Game
		playerID         string
		suit             Suit
		cards            []CardName
		expectedSuit     Suit
		expectedStatus   RoundStatus
		expectedRevision int
		expectingError   bool
	}{
		{
			name:             "Valid selection - keep 1 from my hand and 1 from dummy",
			game:             CalledGame(),
			playerID:         "2",
			suit:             Diamonds,
			cards:            []CardName{ACE_DIAMONDS, KING_SPADES},
			expectedSuit:     Diamonds,
			expectedStatus:   Buying,
			expectedRevision: 1,
		},
		{
			name:             "Valid selection - not keeping any cards",
			game:             CalledGame(),
			playerID:         "2",
			suit:             Diamonds,
			cards:            []CardName{},
			expectedSuit:     Diamonds,
			expectedStatus:   Buying,
			expectedRevision: 1,
		},
		{
			name:             "Valid selection - keep 5 cards",
			game:             CalledGame(),
			playerID:         "2",
			suit:             Diamonds,
			cards:            []CardName{ACE_DIAMONDS, KING_SPADES, QUEEN_SPADES, JACK_SPADES, JOKER},
			expectedSuit:     Diamonds,
			expectedStatus:   Buying,
			expectedRevision: 1,
		},
		{
			name:           "Invalid player - not the goer",
			game:           CalledGame(),
			playerID:       "1",
			suit:           Hearts,
			cards:          []CardName{ACE_HEARTS},
			expectedSuit:   Hearts,
			expectedStatus: Buying,
			expectingError: true,
		},
		{
			name:           "Invalid state",
			game:           TwoPlayerGame(),
			playerID:       "1",
			suit:           Hearts,
			cards:          []CardName{ACE_HEARTS},
			expectedSuit:   Hearts,
			expectedStatus: Buying,
			expectingError: true,
		},
		{
			name:           "Invalid number of cards",
			game:           CalledGame(),
			playerID:       "2",
			suit:           Hearts,
			cards:          []CardName{ACE_SPADES, KING_SPADES, QUEEN_SPADES, JACK_SPADES, JOKER, ACE_DIAMONDS},
			expectedSuit:   Hearts,
			expectedStatus: Buying,
			expectingError: true,
		},
		{
			name:           "Invalid card - not in hand",
			game:           CalledGame(),
			playerID:       "2",
			suit:           Hearts,
			cards:          []CardName{FIVE_CLUBS},
			expectedSuit:   Hearts,
			expectedStatus: Buying,
			expectingError: true,
		},
		{
			name:           "Duplicate card",
			game:           CalledGame(),
			playerID:       "2",
			suit:           Hearts,
			cards:          []CardName{ACE_DIAMONDS, ACE_DIAMONDS},
			expectedSuit:   Hearts,
			expectedStatus: Buying,
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.game.SelectSuit(test.playerID, test.suit, test.cards)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if test.game.CurrentRound.Status != test.expectedStatus {
					t.Errorf("expected round status to be %s, got %s", test.expectedStatus, test.game.CurrentRound.Status)
				}
				if test.game.CurrentRound.Suit != test.expectedSuit {
					t.Errorf("expected suit to be %s, got %s", test.expectedSuit, test.game.CurrentRound.Suit)
				}
				if test.game.Revision != test.expectedRevision {
					t.Errorf("expected revision to be %d, got %d", test.expectedRevision, test.game.Revision)
				}
				// Check that he has all of retained the cards he selected
				state, err := test.game.GetState(test.playerID)
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if !containsAllUnique(state.Cards, test.cards) {
					t.Errorf("expected player to have all of the selected cards %v, got %v", test.cards, state.Cards)
				}
			}
		})
	}
}

func TestGame_ParseCall(t *testing.T) {
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
