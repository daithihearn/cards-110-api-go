package game

import "testing"

func TestGame_Me(t *testing.T) {
	tests := []struct {
		name           string
		game           Game
		playerID       string
		expectedPlayer Player
		expectingError bool
	}{
		{
			name:           "Player exists",
			game:           TwoPlayerGame(),
			playerID:       "1",
			expectedPlayer: TwoPlayerGame().Players[0],
		},
		{
			name:           "Player does not exist",
			game:           TwoPlayerGame(),
			playerID:       "3",
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			player, err := test.game.Me(test.playerID)
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if player.ID != test.expectedPlayer.ID {
					t.Errorf("expected player ID to be %s, got %s", test.expectedPlayer.ID, player.ID)
				}
				if player.Seat != test.expectedPlayer.Seat {
					t.Errorf("expected player seat to be %d, got %d", test.expectedPlayer.Seat, player.Seat)
				}
				if !compare(player.Cards, test.expectedPlayer.Cards) {
					t.Errorf("expected player cards to be %v, got %v", test.expectedPlayer.Cards, player.Cards)
				}
				if player.Call != test.expectedPlayer.Call {
					t.Errorf("expected player call to be %d, got %d", test.expectedPlayer.Call, player.Call)
				}
				if player.Score != test.expectedPlayer.Score {
					t.Errorf("expected player score to be %d, got %d", test.expectedPlayer.Score, player.Score)
				}
				if player.Rings != test.expectedPlayer.Rings {
					t.Errorf("expected player rings to be %d, got %d", test.expectedPlayer.Rings, player.Rings)
				}
				if player.TeamID != test.expectedPlayer.TeamID {
					t.Errorf("expected player team ID to be %s, got %s", test.expectedPlayer.TeamID, player.TeamID)
				}
				if player.Winner != test.expectedPlayer.Winner {
					t.Errorf("expected player winner to be %t, got %t", test.expectedPlayer.Winner, player.Winner)
				}
			}
		})
	}
}

func TestGame_GetState(t *testing.T) {
	tests := []struct {
		name          string
		game          Game
		playerID      string
		expectedState State
	}{
		{
			name:     "Player exists",
			game:     TwoPlayerGame(),
			playerID: "1",
			expectedState: State{
				ID:       TwoPlayerGame().ID,
				Revision: TwoPlayerGame().Revision,
				Me:       TwoPlayerGame().Players[0], Cards: TwoPlayerGame().Players[0].Cards,
				IamDealer: true, IamGoer: false,
				IamSpectator: false, IsMyGo: false,
				Status:  TwoPlayerGame().Status,
				MaxCall: 0,
				Players: TwoPlayerGame().Players,
				Round:   TwoPlayerGame().CurrentRound,
			},
		},
		{
			name:     "Player does not exist so is a spectator",
			game:     TwoPlayerGame(),
			playerID: "3",
			expectedState: State{
				ID:           TwoPlayerGame().ID,
				Revision:     TwoPlayerGame().Revision,
				IamSpectator: true,
				Status:       TwoPlayerGame().Status,
				MaxCall:      0,
				Players:      TwoPlayerGame().Players,
				Round:        TwoPlayerGame().CurrentRound,
			},
		},
		{
			name:     "Game in Called state",
			game:     CalledGameFivePlayers(),
			playerID: "2",
			expectedState: State{
				ID:           CalledGameFivePlayers().ID,
				Revision:     CalledGameFivePlayers().Revision,
				Me:           CalledGameFivePlayers().Players[1],
				Cards:        CalledGameFivePlayers().Players[1].Cards,
				IamDealer:    false,
				IamGoer:      false,
				IamSpectator: false,
				IsMyGo:       false,
				Status:       CalledGameFivePlayers().Status,
				MaxCall:      20,
				Players:      CalledGameFivePlayers().Players,
				Round:        CalledGameFivePlayers().CurrentRound,
			},
		},
		{
			name:     "Game in Buying state",
			game:     BuyingGame("3"),
			playerID: "2",
			expectedState: State{
				ID:           BuyingGame("3").ID,
				Revision:     BuyingGame("3").Revision,
				Me:           BuyingGame("3").Players[1],
				Cards:        BuyingGame("3").Players[1].Cards,
				IamDealer:    false,
				IamGoer:      false,
				IamSpectator: false,
				IsMyGo:       true,
				Status:       BuyingGame("3").Status,
				MaxCall:      20,
				Players:      BuyingGame("3").Players,
				Round:        BuyingGame("3").CurrentRound,
			},
		},
		{
			name:     "Game in Playing state",
			game:     PlayingGame_RoundStart("3"),
			playerID: "PlayerCalled",
			expectedState: State{
				ID:           PlayingGame_RoundStart("3").ID,
				Revision:     PlayingGame_RoundStart("3").Revision,
				Me:           PlayingGame_RoundStart("3").Players[3],
				Cards:        PlayingGame_RoundStart("3").Players[3].Cards,
				IamDealer:    false,
				IamGoer:      true,
				IamSpectator: false,
				IsMyGo:       false,
				Status:       PlayingGame_RoundStart("3").Status,
				MaxCall:      20,
				Players:      PlayingGame_RoundStart("3").Players,
				Round:        PlayingGame_RoundStart("3").CurrentRound,
			},
		},
		{
			name:     "Game with completed rounds",
			game:     GameWithCompletedRounds(),
			playerID: "1",
			expectedState: State{
				ID:           GameWithCompletedRounds().ID,
				Revision:     GameWithCompletedRounds().Revision,
				Me:           GameWithCompletedRounds().Players[0],
				Cards:        GameWithCompletedRounds().Players[0].Cards,
				IamDealer:    true,
				IamGoer:      true,
				IamSpectator: false,
				IsMyGo:       false,
				Status:       GameWithCompletedRounds().Status,
				MaxCall:      0,
				Players:      GameWithCompletedRounds().Players,
				Round:        GameWithCompletedRounds().CurrentRound,
				PrevRound:    GameWithCompletedRounds().Completed[len(GameWithCompletedRounds().Completed)-1],
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			state := test.game.GetState(test.playerID)

			if state.ID != test.expectedState.ID {
				t.Errorf("expected game ID to be %s, got %s", test.expectedState.ID, state.ID)
			}
			if state.Revision != test.expectedState.Revision {
				t.Errorf("expected game revision to be %d, got %d", test.expectedState.Revision, state.Revision)
			}
			if state.Status != test.expectedState.Status {
				t.Errorf("expected game status to be %s, got %s", test.expectedState.Status, state.Status)
			}
			if state.Me.ID != test.expectedState.Me.ID {
				t.Errorf("expected player ID to be %s, got %s", test.expectedState.Me.ID, state.Me.ID)
			}
			if state.Me.Seat != test.expectedState.Me.Seat {
				t.Errorf("expected player seat to be %d, got %d", test.expectedState.Me.Seat, state.Me.Seat)
			}
			if !compare(state.Me.Cards, test.expectedState.Me.Cards) {
				t.Errorf("expected player cards to be %v, got %v", test.expectedState.Me.Cards, state.Me.Cards)
			}
			if state.Me.Call != test.expectedState.Me.Call {
				t.Errorf("expected player call to be %d, got %d", test.expectedState.Me.Call, state.Me.Call)
			}
			if state.Me.Score != test.expectedState.Me.Score {
				t.Errorf("expected player score to be %d, got %d", test.expectedState.Me.Score, state.Me.Score)
			}
			if state.Me.Rings != test.expectedState.Me.Rings {
				t.Errorf("expected player rings to be %d, got %d", test.expectedState.Me.Rings, state.Me.Rings)
			}
			if state.Me.TeamID != test.expectedState.Me.TeamID {
				t.Errorf("expected player team ID to be %s, got %s", test.expectedState.Me.TeamID, state.Me.TeamID)
			}
			if state.Me.Winner != test.expectedState.Me.Winner {
				t.Errorf("expected player winner to be %t, got %t", test.expectedState.Me.Winner, state.Me.Winner)
			}
			if state.IamDealer != test.expectedState.IamDealer {
				t.Errorf("expected IamDealer to be %t, got %t", test.expectedState.IamDealer, state.IamDealer)
			}
			if state.IamGoer != test.expectedState.IamGoer {
				t.Errorf("expected IamGoer to be %t, got %t", test.expectedState.IamGoer, state.IamGoer)
			}
			if state.IamSpectator != test.expectedState.IamSpectator {
				t.Errorf("expected IamSpectator to be %t, got %t", test.expectedState.IamSpectator, state.IamSpectator)
			}
			if state.IsMyGo != test.expectedState.IsMyGo {
				t.Errorf("expected IsMyGo to be %t, got %t", test.expectedState.IsMyGo, state.IsMyGo)
			}
			if state.MaxCall != test.expectedState.MaxCall {
				t.Errorf("expected MaxCall to be %d, got %d", test.expectedState.MaxCall, state.MaxCall)
			}
			if state.Round.Number != test.expectedState.Round.Number {
				t.Errorf("expected Round Number to be %d, got %d", test.expectedState.Round.Number, state.Round.Number)
			}
			if state.PrevRound.Number != test.expectedState.PrevRound.Number {
				t.Errorf("expected PrevRound Number to be %d, got %d", test.expectedState.PrevRound.Number, state.PrevRound.Number)
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

func TestGame_completeHand(t *testing.T) {
	tests := []struct {
		name           string
		game           Game
		nextPlayer     string
		completedHands int
		expectingError bool
	}{
		{
			name:           "No cards played",
			game:           PlayingGame_RoundStart("2"),
			expectingError: true,
		},
		{
			name:           "First hand completed",
			game:           PlayingGame_Hand1Complete("2"),
			completedHands: 1,
			nextPlayer:     "3",
		},
		{
			name:           "Final hand completed",
			game:           PlayingGame_FinalHandComplete("1"),
			completedHands: 5,
			nextPlayer:     "3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.game.completeHand()
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if test.game.CurrentRound.CurrentHand.CurrentPlayerID != test.nextPlayer {
					t.Errorf("expected current player to be %s, got %s", test.nextPlayer, test.game.CurrentRound.CurrentHand.CurrentPlayerID)
				}
				if len(test.game.CurrentRound.CurrentHand.PlayedCards) != 0 {
					t.Errorf("expected played cards to be empty, got %v", test.game.CurrentRound.CurrentHand.PlayedCards)
				}
				if len(test.game.CurrentRound.CompletedHands) != test.completedHands {
					t.Errorf("expected completed hands to be %d, got %d", test.completedHands, len(test.game.CurrentRound.CompletedHands))
				}
			}
		})
	}
}

func TestGame_completeRound(t *testing.T) {
	tests := []struct {
		name                string
		game                Game
		expectedRoundNumber int
		nextDealer          string
		nextPlayer          string
		expectingError      bool
	}{
		{
			name:           "Round not complete",
			game:           PlayingGame_RoundStart("2"),
			expectingError: true,
		},
		{
			name:                "Round complete",
			game:                PlayingGame_AllHandsComplete("1"),
			expectedRoundNumber: 2,
			nextDealer:          "2",
			nextPlayer:          "3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.game.completeRound()
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if test.game.Status != Active {
					t.Errorf("expected game status to be %s, got %s", Active, test.game.Status)
				}
				if len(test.game.CurrentRound.CompletedHands) != 0 {
					t.Errorf("expected completed hands to be 0, got %d", len(test.game.CurrentRound.CompletedHands))
				}
				if test.game.CurrentRound.Number != test.expectedRoundNumber {
					t.Errorf("expected round number to be %d, got %d", test.expectedRoundNumber, test.game.CurrentRound.Number)
				}
				// All players should have 5 cards and no call
				for _, player := range test.game.Players {
					if len(player.Cards) != 5 {
						t.Errorf("expected player %s to have 5 cards, got %d", player.ID, len(player.Cards))
					}
					if player.Call != Pass {
						t.Errorf("expected player %s to have call %d, got %d", player.ID, Pass, player.Call)
					}
				}
				// Dummy should have 5 cards
				if len(test.game.Dummy) != 5 {
					t.Errorf("expected dummy to have 5 cards, got %d", len(test.game.Dummy))
				}
				// Round status should be calling
				if test.game.CurrentRound.Status != Calling {
					t.Errorf("expected round status to be %s, got %s", Calling, test.game.CurrentRound.Status)
				}
				// Dealer should have changed
				if test.game.CurrentRound.DealerID != test.nextDealer {
					t.Errorf("expected dealer to be %s, got %s", test.nextDealer, test.game.CurrentRound.DealerID)
				}
				// Current player should have changed
				if test.game.CurrentRound.CurrentHand.CurrentPlayerID != test.nextPlayer {
					t.Errorf("expected current player to be %s, got %s", test.nextPlayer, test.game.CurrentRound.CurrentHand.CurrentPlayerID)
				}
			}
		})
	}

}

func TestGame_completeGame(t *testing.T) {
	tests := []struct {
		name           string
		game           Game
		winningTeam    string
		expectingError bool
	}{
		{
			name:        "Game complete",
			game:        CompletedGame(),
			winningTeam: "1",
		},
		{
			name:           "Game not complete - calling",
			game:           TwoPlayerGame(),
			expectingError: true,
		},
		{
			name:           "Game not complete - buying",
			game:           BuyingGame("1"),
			expectingError: true,
		},
		{
			name:           "Game not complete - playing",
			game:           PlayingGame_RoundStart("1"),
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.game.completeGame()
			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if test.game.Status != Completed {
					t.Errorf("expected game status to be %s, got %s", Completed, test.game.Status)
				}
				for _, player := range test.game.Players {
					if player.TeamID == test.winningTeam {
						if !player.Winner {
							t.Errorf("expected player %s to be a winner", player.ID)
						}
					} else {
						if player.Winner {
							t.Errorf("expected player %s to not be a winner", player.ID)
						}
					}
				}
			}
		})

	}
}

func TestGame_applyScores(t *testing.T) {
	tests := []struct {
		name           string
		game           Game
		expectedScores []Player
		expectingError bool
	}{
		{
			name: "Caller doesn't make contract",
			game: PlayingGame_AllHandsComplete("1"),
			expectedScores: []Player{
				{
					ID:     "1",
					Score:  10,
					Rings:  0,
					TeamID: "1",
				},
				{
					ID:     "2",
					Score:  5,
					Rings:  0,
					TeamID: "2",
				},
				{
					ID:     "3",
					Score:  15,
					Rings:  0,
					TeamID: "3",
				},
				{
					ID:     "PlayerCalled",
					Score:  -20,
					Rings:  1,
					TeamID: "PlayerCalled",
				},
			},
		},
		{
			name: "Caller doesn't make contract - doubles",
			game: PlayingGame_DoesntMakeContract_Doubles(),
			expectedScores: []Player{
				{
					ID:     "1",
					Score:  15,
					Rings:  0,
					TeamID: "1",
				},
				{
					ID:     "2",
					Score:  0,
					Rings:  0,
					TeamID: "2",
				},
				{
					ID:     "3",
					Score:  -25,
					Rings:  1,
					TeamID: "3",
				},
				{
					ID:     "4",
					Score:  15,
					Rings:  0,
					TeamID: "1",
				},
				{
					ID:     "5",
					Score:  0,
					Rings:  0,
					TeamID: "2",
				},
				{
					ID:     "6",
					Score:  -25,
					Rings:  1,
					TeamID: "3",
				},
			},
		},
		{
			name: "Caller makes JINK",
			game: PlayingGame_Jink(),
			expectedScores: []Player{
				{ID: "1"},
				{ID: "2"},
				{ID: "3", Score: 60},
			},
		},
		{
			name: "Caller makes Jink - doubles",
			game: PlayingGame_Jink_Doubles(),
			expectedScores: []Player{
				{ID: "1"},
				{ID: "2"},
				{ID: "3", Score: 60},
				{ID: "4"},
				{ID: "5"},
				{ID: "6", Score: 60},
			},
		},
		{
			name: "Caller make 30 but didn't call jink",
			game: PlayingGame_Thirty(),
			expectedScores: []Player{
				{ID: "1"},
				{ID: "2"},
				{ID: "3", Score: 30},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.game.applyScores()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			for _, player := range test.game.Players {
				var expectedScore int
				var expectedRings int
				for _, expectedPlayer := range test.expectedScores {
					if expectedPlayer.ID == player.ID {
						expectedScore = expectedPlayer.Score
						expectedRings = expectedPlayer.Rings
					}
				}
				if player.Score != expectedScore {
					t.Errorf("expected player %s score to be %d, got %d", player.ID, expectedScore, player.Score)
				}
				if player.Rings != expectedRings {
					t.Errorf("expected player %s rings to be %d, got %d", player.ID, expectedRings, player.Rings)
				}
			}
		})
	}
}

func TestGame_isGameOver(t *testing.T) {
	tests := []struct {
		name           string
		game           Game
		expectedResult bool
	}{
		{
			name:           "Game not over",
			game:           TwoPlayerGame(),
			expectedResult: false,
		},
		{
			name:           "Game over",
			game:           CompletedGame(),
			expectedResult: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.game.isGameOver()
			if result != test.expectedResult {
				t.Errorf("expected result to be %t, got %t", test.expectedResult, result)
			}
		})
	}
}

func TestGame_validateCaller(t *testing.T) {
	tests := []struct {
		name           string
		game           Game
		playerID       string
		desiredStatus  RoundStatus
		expectingError bool
	}{
		{
			name:           "Player is caller",
			game:           PlayingGame_RoundStart("1"),
			playerID:       "1",
			desiredStatus:  Playing,
			expectingError: false,
		},
		{
			name:           "Player is not caller",
			game:           PlayingGame_RoundStart("1"),
			playerID:       "2",
			desiredStatus:  Playing,
			expectingError: true,
		},
		{
			name:           "Player is caller - wrong status",
			game:           PlayingGame_RoundStart("1"),
			playerID:       "PlayerCalled",
			desiredStatus:  Calling,
			expectingError: true,
		},
		{
			name:           "Invalid playerID",
			game:           PlayingGame_RoundStart("1"),
			desiredStatus:  Playing,
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.game.validateCaller(test.playerID, test.desiredStatus)
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
				expectedRevision:     3,
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
			game: CalledGameFivePlayers(),
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
					expectedRevision:     3,
				},
				{
					playerID:         "1",
					call:             Fifteen,
					expectingError:   true,
					expectedRevision: 3,
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
					expectedRevision:     3,
				},
				{
					playerID:             "1",
					call:                 Twenty,
					expectedNextPlayerID: "2",
					expectedRevision:     4,
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
					expectedRevision:     3,
				},
				{
					playerID:             "1",
					call:                 Twenty,
					expectedNextPlayerID: "1",
					expectedRevision:     4,
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
					expectedRevision:     3,
				},
				{
					playerID:             "1",
					call:                 Pass,
					expectedNextPlayerID: "1",
					expectedRevision:     4,
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
					expectedRevision:     3,
				},
				{
					playerID:             "1",
					call:                 Fifteen,
					expectedNextPlayerID: "2",
					expectedRevision:     4,
				},
				{
					playerID:             "2",
					call:                 Twenty,
					expectedNextPlayerID: "1",
					expectedRevision:     5,
				},
				{
					playerID:             "1",
					call:                 Pass,
					expectedNextPlayerID: "2",
					expectedRevision:     6,
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
					if err != nil {
						t.Errorf("expected no error, got %v", err)
					}
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

func TestGame_SelectSuit(t *testing.T) {
	tests := []struct {
		name                 string
		game                 Game
		playerID             string
		suit                 Suit
		cards                []CardName
		expectedNextPlayerID string
		expectedStatus       RoundStatus
		expectedRevision     int
		expectingError       bool
	}{
		{
			name:                 "Valid selection - keep 1 from my hand and 1 from dummy",
			game:                 CalledGameFivePlayers(),
			playerID:             "PlayerCalled",
			suit:                 Hearts,
			cards:                []CardName{NINE_HEARTS, JOKER},
			expectedNextPlayerID: "2",
			expectedStatus:       Buying,
			expectedRevision:     1,
		},
		{
			name:                 "Valid selection - not keeping any cards",
			game:                 CalledGameThreePlayers(),
			playerID:             "PlayerCalled",
			suit:                 Diamonds,
			cards:                []CardName{},
			expectedNextPlayerID: "2",
			expectedStatus:       Buying,
			expectedRevision:     1,
		},
		{
			name:                 "Valid selection - keep 5 cards",
			game:                 CalledGameFivePlayers(),
			playerID:             "PlayerCalled",
			suit:                 Hearts,
			cards:                []CardName{NINE_HEARTS, EIGHT_HEARTS, SEVEN_HEARTS, SIX_HEARTS, JOKER},
			expectedStatus:       Buying,
			expectedNextPlayerID: "2",
			expectedRevision:     1,
		},
		{
			name:           "Invalid player - not the goer",
			game:           CalledGameFivePlayers(),
			playerID:       "1",
			suit:           Hearts,
			cards:          []CardName{ACE_HEARTS},
			expectingError: true,
		},
		{
			name:           "Invalid number of cards - too many",
			game:           CalledGameFivePlayers(),
			playerID:       "PlayerCalled",
			suit:           Hearts,
			cards:          []CardName{NINE_HEARTS, EIGHT_HEARTS, SEVEN_HEARTS, SIX_HEARTS, FOUR_SPADES, JOKER},
			expectingError: true,
		},
		{
			name:           "Invalid number of cards - too few",
			game:           CalledGameFivePlayers(),
			playerID:       "PlayerCalled",
			suit:           Hearts,
			cards:          []CardName{},
			expectingError: true,
		},
		{
			name:           "Invalid card - not in hand",
			game:           CalledGameFivePlayers(),
			playerID:       "PlayerCalled",
			suit:           Hearts,
			cards:          []CardName{FIVE_CLUBS},
			expectingError: true,
		},
		{
			name:           "Duplicate card",
			game:           CalledGameFivePlayers(),
			playerID:       "PlayerCalled",
			suit:           Hearts,
			cards:          []CardName{JOKER, JOKER},
			expectingError: true,
		},
		{
			name:           "Invalid suit",
			game:           CalledGameFivePlayers(),
			playerID:       "PlayerCalled",
			suit:           "invalid",
			cards:          []CardName{JOKER},
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
				if test.game.CurrentRound.Suit != test.suit {
					t.Errorf("expected suit to be %s, got %s", test.suit, test.game.CurrentRound.Suit)
				}
				if test.game.Revision != test.expectedRevision {
					t.Errorf("expected revision to be %d, got %d", test.expectedRevision, test.game.Revision)
				}
				// Check that he has all of retained the cards he selected
				state := test.game.GetState(test.playerID)
				if !compare(state.Cards, test.cards) {
					t.Errorf("expected player to have all of the selected cards %v, got %v", test.cards, state.Cards)
				}
				if test.expectedNextPlayerID != "" {
					if test.game.CurrentRound.CurrentHand.CurrentPlayerID != test.expectedNextPlayerID {
						t.Errorf("expected next player to be %s, got %s", test.expectedNextPlayerID, test.game.CurrentRound.CurrentHand.CurrentPlayerID)
					}
				}
			}
		})
	}
}

func TestGame_Buy(t *testing.T) {
	tests := []struct {
		name                 string
		game                 Game
		playerID             string
		cards                []CardName
		expectedStatus       RoundStatus
		expectedRevision     int
		expectedNextPlayerID string
		expectingError       bool
	}{
		{
			name:                 "Valid selection - keep 1 from my hand",
			game:                 BuyingGame("3"),
			playerID:             "2",
			cards:                []CardName{SEVEN_HEARTS},
			expectedStatus:       Buying,
			expectedRevision:     1,
			expectedNextPlayerID: "3",
		},
		{
			name:                 "Valid selection - keep all 5",
			game:                 BuyingGame("1"),
			playerID:             "2",
			cards:                []CardName{SEVEN_HEARTS, EIGHT_HEARTS, NINE_HEARTS, TEN_HEARTS, JACK_HEARTS},
			expectedStatus:       Buying,
			expectedRevision:     1,
			expectedNextPlayerID: "3",
		},
		{
			name:                 "dealer buying should cause status change",
			game:                 BuyingGame("2"),
			playerID:             "2",
			cards:                []CardName{SEVEN_HEARTS, EIGHT_HEARTS, NINE_HEARTS, TEN_HEARTS, JACK_HEARTS},
			expectedStatus:       Playing,
			expectedRevision:     1,
			expectedNextPlayerID: "1",
		},
		{
			name:             "Not current player",
			game:             BuyingGame("3"),
			playerID:         "1",
			cards:            []CardName{TWO_HEARTS},
			expectedRevision: 0,
			expectingError:   true,
		},
		{
			name:             "Valid selection - not keeping any cards",
			game:             BuyingGame("3"),
			playerID:         "2",
			cards:            []CardName{},
			expectedStatus:   Buying,
			expectedRevision: 1,
		},
		{
			name:             "Invalid state",
			game:             TwoPlayerGame(),
			playerID:         "1",
			cards:            []CardName{ACE_HEARTS},
			expectedStatus:   Buying,
			expectingError:   true,
			expectedRevision: 2,
		},
		{
			name:             "Invalid number of cards",
			game:             BuyingGame("3"),
			playerID:         "2",
			cards:            []CardName{ACE_SPADES, KING_SPADES, QUEEN_SPADES, JACK_SPADES, JOKER, ACE_DIAMONDS},
			expectedStatus:   Buying,
			expectingError:   true,
			expectedRevision: 0,
		},
		{
			name:             "Invalid card - not in hand",
			game:             BuyingGame("3"),
			playerID:         "2",
			cards:            []CardName{FIVE_CLUBS},
			expectedStatus:   Buying,
			expectingError:   true,
			expectedRevision: 0,
		},
		{
			name:             "Duplicate card",
			game:             BuyingGame("3"),
			playerID:         "2",
			cards:            []CardName{ACE_DIAMONDS, ACE_DIAMONDS},
			expectedStatus:   Buying,
			expectingError:   true,
			expectedRevision: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.game.Buy(test.playerID, test.cards)
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
				// Check that he has all of retained the cards he selected
				state := test.game.GetState(test.playerID)
				if len(state.Cards) != 5 {
					t.Errorf("expected player to have 5 cards, got %d", len(state.Cards))
				}
				if !containsAllUnique(state.Cards, test.cards) {
					t.Errorf("expected player to have all of the selected cards %v, got %v", test.cards, state.Cards)
				}
				if test.expectedNextPlayerID != "" {
					if test.game.CurrentRound.CurrentHand.CurrentPlayerID != test.expectedNextPlayerID {
						t.Errorf("expected next player to be %s, got %s", test.expectedNextPlayerID, test.game.CurrentRound.CurrentHand.CurrentPlayerID)
					}
				}
			}
			if test.game.Revision != test.expectedRevision {
				t.Errorf("expected revision to be %d, got %d", test.expectedRevision, test.game.Revision)
			}
		})
	}
}

func TestGame_Play(t *testing.T) {
	tests := []struct {
		name               string
		game               Game
		playerID           string
		card               CardName
		expectedStatus     RoundStatus
		expectedNextPlayer string
		expectingError     bool
	}{
		{
			name:               "Valid play - first card",
			game:               PlayingGame_RoundStart("1"),
			playerID:           "1",
			card:               TWO_HEARTS,
			expectedStatus:     Playing,
			expectedNextPlayer: "2",
		},
		{
			name:           "Try to play a card that isn't in your hand",
			game:           PlayingGame_RoundStart("1"),
			playerID:       "1",
			card:           ACE_SPADES,
			expectedStatus: Playing,
			expectingError: true,
		},
		{
			name:               "Player 1 - wins trick",
			game:               PlayingGame_RoundStart_FirstCardPlayed(),
			playerID:           "2",
			card:               THREE_CLUBS,
			expectedStatus:     Playing,
			expectedNextPlayer: "1",
		},
		{
			name:               "Player 1 - wins trick - different seating arrangement",
			game:               PlayingGame_WinHand(),
			playerID:           "player2",
			card:               THREE_SPADES,
			expectedStatus:     Playing,
			expectedNextPlayer: "player1",
		},
		{
			name:               "Player 2 - wins trick",
			game:               PlayingGame_RoundStart_FirstCardPlayed(),
			playerID:           "2",
			card:               FIVE_CLUBS,
			expectedStatus:     Playing,
			expectedNextPlayer: "2",
		},
		{
			name:           "Not following suit",
			game:           PlayingGame_RoundStart_FirstCardPlayed(),
			playerID:       "2",
			card:           FOUR_DIAMONDS,
			expectedStatus: Playing,
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.game.Play(test.playerID, test.card)
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
				if test.game.CurrentRound.CurrentHand.CurrentPlayerID != test.expectedNextPlayer {
					t.Errorf("expected next player to be %s, got %s", test.expectedNextPlayer, test.game.CurrentRound.CurrentHand.CurrentPlayerID)
				}
				// Check that he has all of retained the cards he selected
				state := test.game.GetState(test.playerID)
				if contains(state.Cards, test.card) {
					t.Errorf("expected player to not have played card %s, got %v", test.card, state.Cards)
				}

				if state.Round.CurrentHand.CurrentPlayerID != test.expectedNextPlayer {
					t.Errorf("expected next player to be %s, got %s", test.expectedNextPlayer, state.Round.CurrentHand.CurrentPlayerID)
				}
			}
		})
	}
}
