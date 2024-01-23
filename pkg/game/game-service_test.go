package game

import (
	"cards-110-api/pkg/db"
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestGameService_Create(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name            string
		inputPlayerIDs  []string
		inputAdminID    string
		mockUpsertError *[]error
		expectingError  bool
	}{
		{
			name:            "simple create",
			inputPlayerIDs:  []string{"1", "2"},
			inputAdminID:    "1",
			mockUpsertError: &[]error{nil},
		},
		{
			name: "duplicate player IDs",
			inputPlayerIDs: []string{
				"1",
				"1",
			},
			inputAdminID:    "1",
			mockUpsertError: &[]error{nil},
			expectingError:  true,
		},
		{
			name:           "admin not in game",
			inputPlayerIDs: []string{"1", "2"},
			inputAdminID:   "3",
			expectingError: true,
		},
		{
			name: "error thrown",
			inputPlayerIDs: []string{
				"1",
				"2",
			},
			inputAdminID:    "1",
			mockUpsertError: &[]error{errors.New("failed to upsert")},
			expectingError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCol := &db.MockCollection[Game]{
				MockUpsertErr: test.mockUpsertError,
			}

			ds := &Service{
				Col: mockCol,
			}

			result, err := ds.Create(ctx, test.inputPlayerIDs, test.name, test.inputAdminID)

			if test.expectingError {
				if err == nil {
					t.Errorf("expected error %v, got %v", test.expectingError, err)
				}
			} else {
				if result.Name != test.name {
					t.Errorf("expected name %s, got %s", test.name, result.Name)
				}
				if result.AdminID != test.inputAdminID {
					t.Errorf("expected admin id %s, got %s", test.inputAdminID, result.AdminID)
				}
				if len(result.Players) != len(test.inputPlayerIDs) {
					t.Errorf("expected %d players, got %d", len(test.inputPlayerIDs), len(result.Players))
				}
				// Check that the players are in the game
				for _, playerID := range test.inputPlayerIDs {
					found := false
					for _, player := range result.Players {
						if player.ID == playerID {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("expected player %s to be in the game", playerID)
					}
				}
			}
		})
	}
}

func TestGameService_Get(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		mockResult     *[]Game
		mockExists     *[]bool
		mockError      *[]error
		expectedResult Game
		expectedExists bool
		expectingError bool
	}{
		{
			name:           "simple get",
			mockResult:     &[]Game{TwoPlayerGame()},
			mockExists:     &[]bool{true},
			mockError:      &[]error{nil},
			expectedResult: TwoPlayerGame(),
			expectedExists: true,
			expectingError: false,
		},
		{
			name: "error thrown",
			mockResult: &[]Game{
				{},
			},
			mockExists:     &[]bool{false},
			mockError:      &[]error{errors.New("something went wrong")},
			expectedResult: Game{},
			expectedExists: false,
			expectingError: true,
		},
		{
			name:           "not found",
			mockResult:     &[]Game{{}},
			mockExists:     &[]bool{false},
			mockError:      &[]error{nil},
			expectedResult: Game{},
			expectedExists: false,
			expectingError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCol := &db.MockCollection[Game]{
				MockFindOneResult: test.mockResult,
				MockFindOneExists: test.mockExists,
				MockFindOneErr:    test.mockError,
			}

			ds := &Service{
				Col: mockCol,
			}

			result, exists, err := ds.Get(ctx, "1")

			if test.expectingError {
				if err == nil {
					t.Errorf("expected error %v, got %v", test.expectingError, err)
				}
			} else {
				if !reflect.DeepEqual(result, test.expectedResult) {
					t.Errorf("expected result %v, got %v", test.expectedExists, exists)
				}
			}
			if exists != test.expectedExists {
				t.Errorf("expected exists %v, got %v", test.expectedExists, exists)
			}
		})
	}

}

func TestGameService_GetAll(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		mockResult     *[][]Game
		mockError      *[]error
		expectedResult []Game
		expectingError bool
	}{
		{
			name:           "simple get",
			mockResult:     &[][]Game{{TwoPlayerGame()}},
			mockError:      &[]error{nil},
			expectedResult: []Game{TwoPlayerGame()},
			expectingError: false,
		},
		{
			name:           "error thrown",
			mockResult:     &[][]Game{{}},
			mockError:      &[]error{errors.New("something went wrong")},
			expectedResult: []Game{},
			expectingError: true,
		},
		{
			name:           "no results should return empty array",
			mockResult:     &[][]Game{},
			mockError:      &[]error{nil},
			expectedResult: []Game{},
			expectingError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCol := &db.MockCollection[Game]{
				MockFindResult: test.mockResult,
				MockFindErr:    test.mockError,
			}

			ds := &Service{
				Col: mockCol,
			}

			result, err := ds.GetAll(ctx)

			if test.expectingError {
				if err == nil {
					t.Errorf("expected error %v, got %v", test.expectingError, err)
				}
			} else {
				if !reflect.DeepEqual(result, test.expectedResult) && len(test.expectedResult) != 0 {
					t.Errorf("expected result %v, got %v", test.expectedResult, result)
				}
			}
		})
	}
}

func TestGameService_GetState(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		gameID         string
		playerID       string
		mockGetResult  *[]Game
		mockGetExists  *[]bool
		mockGetError   *[]error
		expectedResult State
		expectingError bool
	}{
		{
			name:          "simple get state",
			gameID:        TwoPlayerGame().ID,
			playerID:      "2",
			mockGetResult: &[]Game{TwoPlayerGame()},
			mockGetExists: &[]bool{true},
			mockGetError:  &[]error{nil},
			expectedResult: State{
				ID:           TwoPlayerGame().ID,
				Revision:     TwoPlayerGame().Revision,
				Me:           TwoPlayerGame().Players[1],
				IamSpectator: false,
				IsMyGo:       false,
				IamGoer:      false,
				IamDealer:    false,
				IamAdmin:     false,
				Cards:        TwoPlayerGame().Players[1].Cards,
				Status:       TwoPlayerGame().Status,
				Round:        TwoPlayerGame().CurrentRound,
				MaxCall:      TwoPlayerGame().Players[0].Call,
				Players:      TwoPlayerGame().Players,
			},
			expectingError: false,
		},
		{
			name:           "game not found",
			gameID:         "1",
			playerID:       "2",
			mockGetResult:  &[]Game{{}},
			mockGetExists:  &[]bool{false},
			mockGetError:   &[]error{nil},
			expectingError: false,
		},
		{
			name:           "error thrown getting game",
			gameID:         "1",
			playerID:       "2",
			mockGetResult:  &[]Game{{}},
			mockGetExists:  &[]bool{false},
			mockGetError:   &[]error{errors.New("something went wrong")},
			expectingError: true,
		},
		{
			name:           "error thrown getting game - true exists",
			gameID:         "1",
			playerID:       "2",
			mockGetResult:  &[]Game{{}},
			mockGetExists:  &[]bool{true},
			mockGetError:   &[]error{errors.New("something went wrong")},
			expectingError: true,
		},
		{
			name:     "player not in the game",
			gameID:   TwoPlayerGame().ID,
			playerID: "3",
			mockGetResult: &[]Game{
				TwoPlayerGame(),
			},
			mockGetExists:  &[]bool{true},
			mockGetError:   &[]error{nil},
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCol := &db.MockCollection[Game]{
				MockFindOneResult: test.mockGetResult,
				MockFindOneExists: test.mockGetExists,
				MockFindOneErr:    test.mockGetError,
			}

			ds := &Service{
				Col: mockCol,
			}

			result, _, err := ds.GetState(ctx, test.gameID, test.playerID)

			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if result.ID != test.expectedResult.ID {
					t.Errorf("expected result %v, got %v", test.expectedResult.ID, result.ID)
				}

			}

		})
	}
}

func TestGameService_Delete(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name               string
		gameToCancel       string
		adminID            string
		mockGetResult      *[]Game
		mockGetExists      *[]bool
		mockGetError       *[]error
		mockDeleteOneError *[]error
		expectingError     bool
	}{
		{
			name:               "simple cancel",
			gameToCancel:       TwoPlayerGame().ID,
			adminID:            "1",
			mockGetResult:      &[]Game{TwoPlayerGame()},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockDeleteOneError: &[]error{nil},
			expectingError:     false,
		},
		{
			name:         "error thrown",
			gameToCancel: TwoPlayerGame().ID,
			adminID:      "1",
			mockGetResult: &[]Game{
				{},
			},
			mockGetExists:      &[]bool{false},
			mockGetError:       &[]error{errors.New("something went wrong")},
			mockDeleteOneError: &[]error{nil},
			expectingError:     true,
		},
		{
			name:               "not found",
			gameToCancel:       TwoPlayerGame().ID,
			adminID:            "1",
			mockGetResult:      &[]Game{{}},
			mockGetExists:      &[]bool{false},
			mockGetError:       &[]error{nil},
			mockDeleteOneError: &[]error{nil},
			expectingError:     true,
		},
		{
			name:         "update error",
			gameToCancel: TwoPlayerGame().ID,
			adminID:      "1",
			mockGetResult: &[]Game{
				TwoPlayerGame(),
			},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockDeleteOneError: &[]error{errors.New("something went wrong")},
			expectingError:     true,
		},
		{
			name:         "not admin",
			gameToCancel: TwoPlayerGame().ID,
			adminID:      "2",
			mockGetResult: &[]Game{
				TwoPlayerGame(),
			},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockDeleteOneError: &[]error{nil},
			expectingError:     true,
		},
		{
			name:         "Game completed",
			gameToCancel: CompletedGame().ID,
			adminID:      "1",
			mockGetResult: &[]Game{
				CompletedGame(),
			},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockDeleteOneError: &[]error{nil},
			expectingError:     true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCol := &db.MockCollection[Game]{
				MockFindOneResult: test.mockGetResult,
				MockFindOneExists: test.mockGetExists,
				MockFindOneErr:    test.mockGetError,
				MockDeleteOneErr:  test.mockDeleteOneError,
			}

			ds := &Service{
				Col: mockCol,
			}

			err := ds.Delete(ctx, test.gameToCancel, test.adminID)

			if test.expectingError && err == nil {
				t.Errorf("expected error %v, got %v", test.expectingError, err)
			}
		})
	}
}

func TestGameService_Call(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name               string
		gameID             string
		playerID           string
		call               Call
		mockGetResult      *[]Game
		mockGetExists      *[]bool
		mockGetError       *[]error
		mockUpdateOneError *[]error
		expectingError     bool
		expectedRevision   int
	}{
		{
			name:               "simple call",
			gameID:             TwoPlayerGame().ID,
			playerID:           "2",
			call:               Jink,
			mockGetResult:      &[]Game{TwoPlayerGame()},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{nil},
			expectedRevision:   1,
		},
		{
			name:           "game not found",
			gameID:         "1",
			playerID:       "2",
			call:           Jink,
			mockGetResult:  &[]Game{{}},
			mockGetExists:  &[]bool{false},
			mockGetError:   &[]error{nil},
			expectingError: true,
		},
		{
			name:           "error thrown getting game",
			gameID:         "1",
			playerID:       "2",
			call:           Jink,
			mockGetResult:  &[]Game{{}},
			mockGetExists:  &[]bool{false},
			mockGetError:   &[]error{errors.New("something went wrong")},
			expectingError: true,
		},
		{
			name:           "error thrown getting game - true exists",
			gameID:         "1",
			playerID:       "2",
			call:           Jink,
			mockGetResult:  &[]Game{{}},
			mockGetExists:  &[]bool{true},
			mockGetError:   &[]error{errors.New("something went wrong")},
			expectingError: true,
		},
		{
			name:               "error thrown updating game",
			gameID:             TwoPlayerGame().ID,
			playerID:           "2",
			call:               Jink,
			mockGetResult:      &[]Game{TwoPlayerGame()},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{errors.New("something went wrong")},
			expectingError:     true,
		},
		{
			name:           "not the players turn",
			gameID:         TwoPlayerGame().ID,
			playerID:       "1",
			call:           Jink,
			mockGetResult:  &[]Game{TwoPlayerGame()},
			mockGetExists:  &[]bool{true},
			mockGetError:   &[]error{nil},
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCol := &db.MockCollection[Game]{
				MockFindOneResult: test.mockGetResult,
				MockFindOneExists: test.mockGetExists,
				MockFindOneErr:    test.mockGetError,
				MockUpdateOneErr:  test.mockUpdateOneError,
			}

			ds := &Service{
				Col: mockCol,
			}

			game, err := ds.Call(ctx, test.gameID, test.playerID, test.call)

			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				var player Player
				for _, p := range game.Players {
					if p.ID == test.playerID {
						player = p
						break
					}
				}
				// Check call has been made
				if player.ID == "" {
					t.Errorf("Player not found")
				}

				if player.Call != test.call {
					t.Errorf("expected call %v, got %v", test.call, player.Call)
				}
				// Check revision has been incremented
				if game.Revision != test.expectedRevision {
					t.Errorf("expected revision %d, got %d", test.expectedRevision, game.Revision)
				}
			}
		})
	}
}

func TestGameService_SelectSuit(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name               string
		gameID             string
		playerID           string
		suit               Suit
		cards              []CardName
		mockGetResult      *[]Game
		mockGetExists      *[]bool
		mockGetError       *[]error
		mockUpdateOneError *[]error
		expectingError     bool
		expectedRevision   int
	}{
		{
			name:               "simple select suit",
			gameID:             CalledGameFivePlayers().ID,
			playerID:           "PlayerCalled",
			suit:               Hearts,
			cards:              []CardName{NINE_HEARTS, EIGHT_HEARTS, SEVEN_HEARTS, SIX_HEARTS},
			mockGetResult:      &[]Game{CalledGameFivePlayers()},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{nil},
			expectedRevision:   1,
		},
		{
			name:           "game not found",
			gameID:         "1",
			playerID:       "2",
			suit:           Clubs,
			cards:          []CardName{ACE_CLUBS},
			mockGetResult:  &[]Game{{}},
			mockGetExists:  &[]bool{false},
			mockGetError:   &[]error{nil},
			expectingError: true,
		},
		{
			name:           "error thrown getting game",
			gameID:         "1",
			playerID:       "2",
			suit:           Clubs,
			cards:          []CardName{ACE_CLUBS},
			mockGetResult:  &[]Game{{}},
			mockGetExists:  &[]bool{false},
			mockGetError:   &[]error{errors.New("something went wrong")},
			expectingError: true,
		},
		{
			name:           "error thrown getting game - true exists",
			gameID:         "1",
			playerID:       "2",
			suit:           Clubs,
			cards:          []CardName{ACE_CLUBS},
			mockGetResult:  &[]Game{{}},
			mockGetExists:  &[]bool{true},
			mockGetError:   &[]error{errors.New("something went wrong")},
			expectingError: true,
		},
		{
			name:               "error thrown updating game",
			gameID:             CalledGameFivePlayers().ID,
			playerID:           "2",
			suit:               Clubs,
			cards:              []CardName{KING_DIAMONDS},
			mockGetResult:      &[]Game{CalledGameFivePlayers()},
			mockGetError:       &[]error{nil},
			mockGetExists:      &[]bool{true},
			mockUpdateOneError: &[]error{errors.New("something went wrong")},
			expectingError:     true,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			mockCol := &db.MockCollection[Game]{
				MockFindOneResult: test.mockGetResult,
				MockFindOneExists: test.mockGetExists,
				MockFindOneErr:    test.mockGetError,
				MockUpdateOneErr:  test.mockUpdateOneError,
			}

			ds := &Service{
				Col: mockCol,
			}

			game, err := ds.SelectSuit(ctx, test.gameID, test.playerID, test.suit, test.cards)

			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				var player Player
				for _, p := range game.Players {
					if p.ID == test.playerID {
						player = p
						break
					}
				}
				// Check suit has been selected
				if player.ID == "" {
					t.Errorf("Player not found")
				}

				if game.CurrentRound.Suit != test.suit {
					t.Errorf("expected suit %v, got %v", test.suit, game.CurrentRound.Suit)
				}
				// Check revision has been incremented
				if game.Revision != test.expectedRevision {
					t.Errorf("expected revision %d, got %d", test.expectedRevision, game.Revision)
				}
			}
		})
	}
}
