package game

import (
	"cards-110-api/pkg/db"
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		inputPlayerIDs []string
		inputName      string
		inputAdminID   string
		mockError      *[]error
		expectingError bool
	}{
		{
			name:           "simple create",
			inputPlayerIDs: []string{"1", "2"},
			inputName:      "test",
			inputAdminID:   "1",
			mockError:      &[]error{nil},
		},
		{
			name: "duplicate player IDs",
			inputPlayerIDs: []string{
				"1",
				"1",
			},
			inputName:      "test",
			inputAdminID:   "1",
			mockError:      &[]error{nil},
			expectingError: true,
		},
		{
			name: "error thrown",
			inputPlayerIDs: []string{
				"1",
				"2",
			},
			inputName:      "test",
			inputAdminID:   "1",
			mockError:      &[]error{errors.New("failed to upsert")},
			expectingError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCol := &db.MockCollection[Game]{
				MockUpsertErr: test.mockError,
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
					t.Errorf("expected name %s, got %s", test.inputName, result.Name)
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

func TestGet(t *testing.T) {
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

func TestGetAll(t *testing.T) {
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

func TestCancel(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name            string
		gameToCancel    string
		adminID         string
		mockGetResult   *[]Game
		mockGetExists   *[]bool
		mockGetError    *[]error
		mockUpdateError *[]error
		expectedResult  Game
		expectingError  bool
	}{
		{
			name:            "simple cancel",
			gameToCancel:    TwoPlayerGame().ID,
			adminID:         "1",
			mockGetResult:   &[]Game{TwoPlayerGame()},
			mockGetExists:   &[]bool{true},
			mockGetError:    &[]error{nil},
			mockUpdateError: &[]error{nil},
			expectedResult:  TwoPlayerGame(),
			expectingError:  false,
		},
		{
			name:         "error thrown",
			gameToCancel: TwoPlayerGame().ID,
			adminID:      "1",
			mockGetResult: &[]Game{
				{},
			},
			mockGetExists:   &[]bool{false},
			mockGetError:    &[]error{errors.New("something went wrong")},
			mockUpdateError: &[]error{nil},
			expectedResult:  Game{},
			expectingError:  true,
		},
		{
			name:           "not found",
			gameToCancel:   TwoPlayerGame().ID,
			adminID:        "1",
			mockGetResult:  &[]Game{{}},
			mockGetExists:  &[]bool{false},
			mockGetError:   &[]error{nil},
			expectedResult: Game{},
			expectingError: true,
		},
		{
			name:         "update error",
			gameToCancel: TwoPlayerGame().ID,
			adminID:      "1",
			mockGetResult: &[]Game{
				TwoPlayerGame(),
			},
			mockGetExists:   &[]bool{true},
			mockGetError:    &[]error{nil},
			mockUpdateError: &[]error{errors.New("something went wrong")},
			expectedResult:  Game{},
			expectingError:  true,
		},
		{
			name:         "not admin",
			gameToCancel: TwoPlayerGame().ID,
			adminID:      "2",
			mockGetResult: &[]Game{
				TwoPlayerGame(),
			},
			mockGetExists:   &[]bool{true},
			mockGetError:    &[]error{nil},
			mockUpdateError: &[]error{nil},
			expectedResult:  Game{},
			expectingError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCol := &db.MockCollection[Game]{
				MockFindOneResult: test.mockGetResult,
				MockFindOneExists: test.mockGetExists,
				MockFindOneErr:    test.mockGetError,
				MockUpdateOneErr:  test.mockUpdateError,
			}

			ds := &Service{
				Col: mockCol,
			}

			result, err := ds.Cancel(ctx, test.gameToCancel, test.adminID)

			if test.expectingError {
				if err == nil {
					t.Errorf("expected error %v, got %v", test.expectingError, err)
				}
			} else {
				if result.Status != CANCELLED {
					t.Errorf("expected result %v, got %v", test.expectedResult, result)
				}

				// Check all fields are the same except status
				if result.ID != test.expectedResult.ID {
					t.Errorf("expected result %v, got %v", test.expectedResult, result)
				}
				if !reflect.DeepEqual(result.Players, test.expectedResult.Players) {
					t.Errorf("expected result %v, got %v", test.expectedResult, result)
				}
				if !reflect.DeepEqual(result.CurrentRound, test.expectedResult.CurrentRound) {
					t.Errorf("expected result %v, got %v", test.expectedResult, result)
				}
				if result.AdminID != test.expectedResult.AdminID {
					t.Errorf("expected result %v, got %v", test.expectedResult, result)
				}
				if result.Name != test.expectedResult.Name {
					t.Errorf("expected result %v, got %v", test.expectedResult, result)
				}
			}
		})
	}
}
