package game

import "testing"

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
