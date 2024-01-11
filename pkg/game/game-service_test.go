package game

import (
	"cards-110-api/pkg/db"
	"context"
	"fmt"
	"testing"
	"time"
)

var game = Game{
	ID:        "1",
	Name:      "Test Game",
	Status:    ACTIVE,
	Timestamp: time.Now(),
	Players: []Player{
		{
			ID:     "1",
			Seat:   1,
			Call:   0,
			Cards:  []CardName{},
			Bought: 0,
			Score:  0,
			Rings:  0,
			TeamID: "1",
			Winner: false,
		},
		{
			ID:     "2",
			Seat:   2,
			Call:   0,
			Cards:  []CardName{},
			Bought: 0,
			Score:  0,
			Rings:  0,
			TeamID: "2",
			Winner: false,
		},
	},
	AdminID: "1",
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
			name:           "success",
			mockResult:     &[]Game{game},
			mockExists:     &[]bool{true},
			mockError:      &[]error{nil},
			expectedResult: game,
			expectedExists: true,
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

			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", test.expectedResult) {
				t.Errorf("expected result %v, got %v", test.expectedResult, result)
			}
			if exists != test.expectedExists {
				t.Errorf("expected exists %v, got %v", test.expectedExists, exists)
			}
			if (err != nil) != test.expectingError {
				t.Errorf("expected error %v, got %v", test.expectingError, err)
			}
		})
	}

}
