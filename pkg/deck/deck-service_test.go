package deck

import (
	"cards-110-api/pkg/db"
	"context"
	"fmt"
	"testing"
)

var deck = Deck{
	ID: "1",
	Cards: []Card{
		TwoHearts,
	},
}

func TestGet(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		mockResult     *[]Deck
		mockExists     *[]bool
		mockError      *[]error
		expectedResult Deck
		expectedExists bool
		expectingError bool
	}{
		{
			name:           "success",
			mockResult:     &[]Deck{deck},
			mockExists:     &[]bool{true},
			mockError:      &[]error{nil},
			expectedResult: deck,
			expectedExists: true,
			expectingError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCol := &db.MockCollection[Deck]{
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
