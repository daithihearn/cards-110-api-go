//go:build testutils

package game

import (
	"time"
)

func Player1() Player {
	return Player{
		ID:     "1",
		Seat:   1,
		Call:   0,
		Cards:  []CardName{},
		Bought: 0,
	}
}

func Player2() Player {
	return Player{
		ID:     "2",
		Seat:   2,
		Call:   0,
		Cards:  []CardName{},
		Bought: 0,
	}
}

func Player3() Player {
	return Player{
		ID:     "3",
		Seat:   3,
		Call:   0,
		Cards:  []CardName{},
		Bought: 0,
	}
}

func Player4() Player {
	return Player{
		ID:     "4",
		Seat:   4,
		Call:   0,
		Cards:  []CardName{},
		Bought: 0,
	}
}

func Player5() Player {
	return Player{
		ID:     "5",
		Seat:   5,
		Call:   0,
		Cards:  []CardName{},
		Bought: 0,
	}
}

func Player6() Player {
	return Player{
		ID:     "6",
		Seat:   6,
		Call:   0,
		Cards:  []CardName{},
		Bought: 0,
	}
}

func TwoPlayerGame() Game {
	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{Player1(), Player2()},
		AdminID:   "1",
	}
}

func CompletedGame() Game {
	return Game{
		ID:        "2",
		Name:      "Test Game",
		Status:    Completed,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{Player1(), Player2()},
		AdminID:   "1",
	}
}
