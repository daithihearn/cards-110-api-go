package game

import (
	"cards-110-api/pkg/cache"
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
		name               string
		gameID             string
		playerID           string
		mockGetResult      *[]Game
		mockGetExists      *[]bool
		mockGetError       *[]error
		mockGetCacheResult *[]State
		mockGetCacheExists *[]bool
		mockGetCacheError  *[]error
		mockSetCacheError  *[]error
		expectedResult     State
		expectedExists     bool
		expectingError     bool
	}{
		{
			name:               "simple get state",
			gameID:             TwoPlayerGame().ID,
			playerID:           "2",
			mockGetResult:      &[]Game{TwoPlayerGame()},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockGetCacheResult: &[]State{},
			mockGetCacheExists: &[]bool{false},
			mockGetCacheError:  &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectedResult: State{
				ID:           TwoPlayerGame().ID,
				Revision:     TwoPlayerGame().Revision,
				Me:           TwoPlayerGame().Players[1],
				IamSpectator: false,
				IsMyGo:       true,
				IamGoer:      false,
				IamDealer:    false,
				IamAdmin:     false,
				Cards:        TwoPlayerGame().Players[1].Cards,
				Status:       TwoPlayerGame().Status,
				Round:        TwoPlayerGame().CurrentRound,
				MaxCall:      TwoPlayerGame().Players[0].Call,
				Players:      TwoPlayerGame().Players,
			},
			expectedExists: true,
			expectingError: false,
		},
		{
			name:               "game not found",
			gameID:             "1",
			playerID:           "2",
			mockGetResult:      &[]Game{{}},
			mockGetExists:      &[]bool{false},
			mockGetError:       &[]error{nil},
			mockGetCacheResult: &[]State{},
			mockGetCacheExists: &[]bool{false},
			mockGetCacheError:  &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectedExists:     false,
			expectingError:     false,
		},
		{
			name:               "error thrown getting game",
			gameID:             "1",
			playerID:           "2",
			mockGetResult:      &[]Game{{}},
			mockGetExists:      &[]bool{false},
			mockGetError:       &[]error{errors.New("something went wrong")},
			mockGetCacheResult: &[]State{},
			mockGetCacheExists: &[]bool{false},
			mockGetCacheError:  &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectedExists:     false,
			expectingError:     true,
		},
		{
			name:               "error thrown getting game - true exists",
			gameID:             "1",
			playerID:           "2",
			mockGetResult:      &[]Game{{}},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{errors.New("something went wrong")},
			mockGetCacheResult: &[]State{},
			mockGetCacheExists: &[]bool{false},
			mockGetCacheError:  &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectedExists:     false,
			expectingError:     true,
		},
		{
			name:     "player not in the game",
			gameID:   TwoPlayerGame().ID,
			playerID: "3",
			mockGetResult: &[]Game{
				TwoPlayerGame(),
			},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockGetCacheResult: &[]State{},
			mockGetCacheExists: &[]bool{false},
			mockGetCacheError:  &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectedExists:     false,
			expectingError:     true,
		},
		{
			name: "successful cache hit",
			mockGetCacheResult: &[]State{{
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
			}},
			mockGetCacheExists: &[]bool{true},
			mockGetCacheError:  &[]error{nil},
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
			expectedExists: true,
		},
		{
			name:               "error writing to cache should return an error",
			gameID:             TwoPlayerGame().ID,
			playerID:           "2",
			mockGetResult:      &[]Game{TwoPlayerGame()},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockGetCacheResult: &[]State{},
			mockGetCacheExists: &[]bool{false},
			mockGetCacheError:  &[]error{nil},
			mockSetCacheError:  &[]error{errors.New("failed to write to cache")},
			expectedResult: State{
				ID:           TwoPlayerGame().ID,
				Revision:     TwoPlayerGame().Revision,
				Me:           TwoPlayerGame().Players[1],
				IamSpectator: false,
				IsMyGo:       true,
				IamGoer:      false,
				IamDealer:    false,
				IamAdmin:     false,
				Cards:        TwoPlayerGame().Players[1].Cards,
				Status:       TwoPlayerGame().Status,
				Round:        TwoPlayerGame().CurrentRound,
				MaxCall:      TwoPlayerGame().Players[0].Call,
				Players:      TwoPlayerGame().Players,
			},
			expectedExists: true,
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

			mockCache := &cache.MockCache[State]{
				MockGetErr:    test.mockGetCacheError,
				MockGetExists: test.mockGetCacheExists,
				MockGetResult: test.mockGetCacheResult,
				MockSetErr:    test.mockSetCacheError,
			}

			ds := &Service{
				Col:   mockCol,
				Cache: mockCache,
			}

			result, exists, err := ds.GetState(ctx, test.gameID, test.playerID)

			if test.expectingError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
			} else {
				if test.expectedExists != exists {
					t.Errorf("expected exists %v, got %v", test.expectedExists, exists)
				}
				// Check each field of the state
				if result.ID != test.expectedResult.ID {
					t.Errorf("expected id %s, got %s", test.expectedResult.ID, result.ID)
				}
				if result.Revision != test.expectedResult.Revision {
					t.Errorf("expected revision %d, got %d", test.expectedResult.Revision, result.Revision)
				}
				if result.Me.ID != test.expectedResult.Me.ID {
					t.Errorf("expected me id %s, got %s", test.expectedResult.Me.ID, result.Me.ID)
				}
				if result.IamSpectator != test.expectedResult.IamSpectator {
					t.Errorf("expected iam spectator %v, got %v", test.expectedResult.IamSpectator, result.IamSpectator)
				}
				if result.IsMyGo != test.expectedResult.IsMyGo {
					t.Errorf("expected is my go %v, got %v", test.expectedResult.IsMyGo, result.IsMyGo)
				}
				if result.IamGoer != test.expectedResult.IamGoer {
					t.Errorf("expected iam goer %v, got %v", test.expectedResult.IamGoer, result.IamGoer)
				}
				if result.IamDealer != test.expectedResult.IamDealer {
					t.Errorf("expected iam dealer %v, got %v", test.expectedResult.IamDealer, result.IamDealer)
				}
				if result.IamAdmin != test.expectedResult.IamAdmin {
					t.Errorf("expected iam admin %v, got %v", test.expectedResult.IamAdmin, result.IamAdmin)
				}
				if !reflect.DeepEqual(result.Cards, test.expectedResult.Cards) {
					t.Errorf("expected cards %v, got %v", test.expectedResult.Cards, result.Cards)
				}
				if result.Status != test.expectedResult.Status {
					t.Errorf("expected status %v, got %v", test.expectedResult.Status, result.Status)
				}
				if !reflect.DeepEqual(result.Round, test.expectedResult.Round) {
					t.Errorf("expected round %v, got %v", test.expectedResult.Round, result.Round)
				}
				if result.MaxCall != test.expectedResult.MaxCall {
					t.Errorf("expected max call %v, got %v", test.expectedResult.MaxCall, result.MaxCall)
				}
				if !reflect.DeepEqual(result.Players, test.expectedResult.Players) {
					t.Errorf("expected players %v, got %v", test.expectedResult.Players, result.Players)
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
		mockSetCacheError  *[]error
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
			mockSetCacheError:  &[]error{nil},
			expectedRevision:   3,
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
		{
			name:               "error writing to cache should return an error",
			gameID:             TwoPlayerGame().ID,
			playerID:           "2",
			call:               Jink,
			mockGetResult:      &[]Game{TwoPlayerGame()},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{errors.New("failed to write to cache")},
			expectedRevision:   2,
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

			mockCache := &cache.MockCache[State]{
				MockSetErr: test.mockSetCacheError,
			}

			ds := &Service{
				Col:   mockCol,
				Cache: mockCache,
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
		mockSetCacheError  *[]error
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
			mockSetCacheError:  &[]error{nil},
			expectedRevision:   1,
		},
		{
			name:               "game not found",
			gameID:             "1",
			playerID:           "2",
			suit:               Clubs,
			cards:              []CardName{ACE_CLUBS},
			mockGetResult:      &[]Game{{}},
			mockGetExists:      &[]bool{false},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectingError:     true,
		},
		{
			name:               "error thrown getting game",
			gameID:             "1",
			playerID:           "2",
			suit:               Clubs,
			cards:              []CardName{ACE_CLUBS},
			mockGetResult:      &[]Game{{}},
			mockGetExists:      &[]bool{false},
			mockGetError:       &[]error{errors.New("something went wrong")},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectingError:     true,
		},
		{
			name:               "error thrown getting game - true exists",
			gameID:             "1",
			playerID:           "2",
			suit:               Clubs,
			cards:              []CardName{ACE_CLUBS},
			mockGetResult:      &[]Game{{}},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{errors.New("something went wrong")},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectingError:     true,
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
		{
			name:               "error writing to cache should return an error",
			gameID:             CalledGameFivePlayers().ID,
			playerID:           "PlayerCalled",
			suit:               Hearts,
			cards:              []CardName{NINE_HEARTS, EIGHT_HEARTS, SEVEN_HEARTS, SIX_HEARTS},
			mockGetResult:      &[]Game{CalledGameFivePlayers()},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{errors.New("failed to write to cache")},
			expectedRevision:   1,
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

			mockCache := &cache.MockCache[State]{
				MockSetErr: test.mockSetCacheError,
			}

			ds := &Service{
				Col:   mockCol,
				Cache: mockCache,
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

func TestGameService_Buy(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name               string
		gameID             string
		playerID           string
		cards              []CardName
		mockGetResult      *[]Game
		mockGetExists      *[]bool
		mockGetError       *[]error
		mockUpdateOneError *[]error
		mockSetCacheError  *[]error
		expectingError     bool
		expectedRevision   int
	}{
		{
			name:               "simple buy",
			gameID:             "1",
			playerID:           "2",
			cards:              []CardName{SEVEN_HEARTS, EIGHT_HEARTS, NINE_HEARTS},
			mockGetResult:      &[]Game{BuyingGame("1")},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectedRevision:   1,
		},
		{
			name:               "not keeping any cards",
			gameID:             "1",
			playerID:           "2",
			cards:              []CardName{},
			mockGetResult:      &[]Game{BuyingGame("1")},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectedRevision:   1,
		},
		{
			name:               "game not found",
			gameID:             "1",
			playerID:           "2",
			cards:              []CardName{ACE_DIAMONDS, KING_DIAMONDS, QUEEN_DIAMONDS},
			mockGetResult:      &[]Game{},
			mockGetExists:      &[]bool{false},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectingError:     true,
		},
		{
			name:               "error thrown getting game",
			gameID:             "1",
			playerID:           "2",
			cards:              []CardName{ACE_DIAMONDS, KING_DIAMONDS, QUEEN_DIAMONDS},
			mockGetResult:      &[]Game{},
			mockGetExists:      &[]bool{false},
			mockGetError:       &[]error{errors.New("something went wrong")},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectingError:     true,
		},
		{
			name:               "error writing to cache should return an error",
			gameID:             "1",
			playerID:           "2",
			cards:              []CardName{SEVEN_HEARTS, EIGHT_HEARTS, NINE_HEARTS},
			mockGetResult:      &[]Game{BuyingGame("1")},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{errors.New("failed to write to cache")},
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

			mockCache := &cache.MockCache[State]{
				MockSetErr: test.mockSetCacheError,
			}

			ds := &Service{
				Col:   mockCol,
				Cache: mockCache,
			}

			game, err := ds.Buy(ctx, test.gameID, test.playerID, test.cards)

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
				// Check cards have been bought
				if player.ID == "" {
					t.Errorf("Player not found")
				}

				if len(player.Cards) != 5 {
					t.Errorf("expected 5 cards, got %d", len(player.Cards))
				}
				// Check the cards are still in the players hand
				if !containsAllUnique(player.Cards, test.cards) {
					t.Errorf("expected cards %v, got %v", test.cards, player.Cards)
				}

				// Check revision has been incremented
				if game.Revision != test.expectedRevision {
					t.Errorf("expected revision %d, got %d", test.expectedRevision, game.Revision)
				}
			}
		})
	}
}

func TestGameService_Play(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name               string
		gameID             string
		playerID           string
		card               CardName
		mockGetResult      *[]Game
		mockGetExists      *[]bool
		mockGetError       *[]error
		mockUpdateOneError *[]error
		mockSetCacheError  *[]error
		expectingError     bool
		expectedRevision   int
	}{
		{
			name:               "simple play",
			gameID:             "1",
			playerID:           "1",
			card:               TWO_HEARTS,
			mockGetResult:      &[]Game{PlayingGame_RoundStart("1")},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectedRevision:   1,
		},
		{
			name:               "playing a card that is not in the players hand",
			gameID:             "1",
			playerID:           "1",
			card:               ACE_HEARTS,
			mockGetResult:      &[]Game{PlayingGame_RoundStart("1")},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectingError:     true,
		},
		{
			name:               "game not found",
			gameID:             "1",
			playerID:           "1",
			card:               TWO_HEARTS,
			mockGetResult:      &[]Game{},
			mockGetExists:      &[]bool{false},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectingError:     true,
		},
		{
			name:               "error thrown getting game",
			gameID:             "1",
			playerID:           "1",
			card:               TWO_HEARTS,
			mockGetResult:      &[]Game{},
			mockGetExists:      &[]bool{false},
			mockGetError:       &[]error{errors.New("something went wrong")},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{nil},
			expectingError:     true,
		},
		{
			name:               "error writing to cache should return an error",
			gameID:             "1",
			playerID:           "1",
			card:               TWO_HEARTS,
			mockGetResult:      &[]Game{PlayingGame_RoundStart("1")},
			mockGetExists:      &[]bool{true},
			mockGetError:       &[]error{nil},
			mockUpdateOneError: &[]error{nil},
			mockSetCacheError:  &[]error{errors.New("failed to write to cache")},
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

			mockCache := &cache.MockCache[State]{
				MockSetErr: test.mockSetCacheError,
			}

			ds := &Service{
				Col:   mockCol,
				Cache: mockCache,
			}

			game, err := ds.Play(ctx, test.gameID, test.playerID, test.card)

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

				// Check the card has been played
				if len(player.Cards) != 4 {
					t.Errorf("expected 4 cards, got %d", len(player.Cards))
				}

				// Ensure the played card is not in the hand anymore
				if contains(player.Cards, test.card) {
					t.Errorf("expected card %v to be removed from hand", test.card)
				}

				// Check revision has been incremented
				if game.Revision != test.expectedRevision {
					t.Errorf("expected revision %d, got %d", test.expectedRevision, game.Revision)
				}
			}
		})
	}
}
