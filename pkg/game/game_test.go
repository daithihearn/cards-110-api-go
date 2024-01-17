package game

import "testing"

type CallWrap struct {
	playerID             string
	call                 Call
	expectedNextPlayerID string
	expectingError       bool
}

func TestCall(t *testing.T) {
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
			}},
		},
		{
			name: "Completed game",
			game: CompletedGame(),
			calls: []CallWrap{{
				playerID:       "2",
				call:           Fifteen,
				expectingError: true,
			}},
		},
		{
			name: "Game in called state",
			game: CalledGame(),
			calls: []CallWrap{{
				playerID:       "2",
				call:           Fifteen,
				expectingError: true,
			}},
		},
		{
			name: "Invalid call 10 in a 2 player game",
			game: TwoPlayerGame(),
			calls: []CallWrap{{
				playerID:       "2",
				call:           Ten,
				expectingError: true,
			}},
		},
		{
			name: "Valid 10 call in 6 player game",
			game: SixPlayerGame(),
			calls: []CallWrap{{
				playerID:             "5",
				call:                 Ten,
				expectedNextPlayerID: "6",
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
				},
				{
					playerID:       "1",
					call:           Fifteen,
					expectingError: true,
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
				},
				{
					playerID:             "1",
					call:                 Twenty,
					expectedNextPlayerID: "2",
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
				},
				{
					playerID:             "1",
					call:                 Twenty,
					expectedNextPlayerID: "1",
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
				},
				{
					playerID:             "1",
					call:                 Pass,
					expectedNextPlayerID: "1",
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
				},
				{
					playerID:             "1",
					call:                 Jink,
					expectedNextPlayerID: "1",
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
				},
				{
					playerID:             "1",
					call:                 Pass,
					expectedNextPlayerID: "5",
				},
			},
			expectedGoerID: "5",
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
