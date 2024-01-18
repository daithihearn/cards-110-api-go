//go:build testutils

package game

import (
	"time"
)

func Dummy() []CardName {
	return []CardName{ACE_SPADES, KING_SPADES, QUEEN_SPADES, JACK_SPADES, JOKER}
}

func Player1() Player {
	return Player{
		ID:     "1",
		Seat:   1,
		Call:   0,
		Cards:  []CardName{ACE_HEARTS, KING_HEARTS, QUEEN_HEARTS, JACK_HEARTS, TEN_HEARTS},
		Bought: 0,
	}
}

func Player2() Player {
	return Player{
		ID:     "2",
		Seat:   2,
		Call:   0,
		Cards:  []CardName{ACE_DIAMONDS, KING_DIAMONDS, QUEEN_DIAMONDS, JACK_DIAMONDS, TEN_DIAMONDS},
		Bought: 0,
	}
}

func Player3() Player {
	return Player{
		ID:     "3",
		Seat:   3,
		Call:   0,
		Cards:  []CardName{ACE_CLUBS, KING_CLUBS, QUEEN_CLUBS, JACK_CLUBS, TEN_CLUBS},
		Bought: 0,
	}
}

func Player4() Player {
	return Player{
		ID:     "4",
		Seat:   4,
		Call:   0,
		Cards:  []CardName{NINE_CLUBS, EIGHT_CLUBS, SEVEN_CLUBS, SIX_CLUBS, FIVE_CLUBS},
		Bought: 0,
	}
}

func Player5() Player {
	return Player{
		ID:     "5",
		Seat:   5,
		Call:   0,
		Cards:  []CardName{NINE_DIAMONDS, EIGHT_DIAMONDS, SEVEN_DIAMONDS, SIX_DIAMONDS, FIVE_DIAMONDS},
		Bought: 0,
	}
}

func Player6() Player {
	return Player{
		ID:     "6",
		Seat:   6,
		Call:   0,
		Cards:  []CardName{NINE_HEARTS, EIGHT_HEARTS, SEVEN_HEARTS, SIX_HEARTS, FIVE_HEARTS},
		Bought: 0,
	}
}

func Player7() Player {
	return Player{
		ID:     "7",
		Seat:   7,
		Call:   0,
		Cards:  []CardName{NINE_SPADES, EIGHT_SPADES, SEVEN_SPADES, SIX_SPADES, FIVE_SPADES},
		Bought: 0,
	}
}

func OnePlayerGame() Game {
	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{Player1()},
		Dummy:     Dummy(),
		CurrentRound: Round{
			DealerID: "1",
			Status:   Calling,
			CurrentHand: Hand{
				CurrentPlayerID: "1",
			},
		},
		AdminID: "1",
	}
}

func TwoPlayerGame() Game {
	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{Player1(), Player2()},
		Dummy:     Dummy(),
		CurrentRound: Round{
			DealerID: "1",
			Status:   Calling,
			CurrentHand: Hand{
				CurrentPlayerID: "2",
			},
		},
		AdminID: "1",
	}
}

func ThreePlayerGame() Game {
	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{Player1(), Player2(), Player3()},
		Dummy:     Dummy(),
		CurrentRound: Round{
			DealerID: "1",
			Status:   Calling,
			CurrentHand: Hand{
				CurrentPlayerID: "2",
			},
		},
		AdminID: "1",
	}
}

func FourPlayerGame() Game {
	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{Player1(), Player2(), Player3(), Player4()},
		Dummy:     Dummy(),
		CurrentRound: Round{
			DealerID: "1",
			Status:   Calling,
			CurrentHand: Hand{
				CurrentPlayerID: "2",
			},
		},
		AdminID: "1",
	}
}

func FivePlayerGame() Game {
	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{Player1(), Player2(), Player3(), Player4(), Player5()},
		Dummy:     Dummy(),
		CurrentRound: Round{
			DealerID: "1",
			Status:   Calling,
			CurrentHand: Hand{
				CurrentPlayerID: "2",
			},
		},
		AdminID: "1",
	}
}

func SixPlayerGame() Game {
	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{Player1(), Player2(), Player3(), Player4(), Player5(), Player6()},
		Dummy:     Dummy(),
		CurrentRound: Round{
			DealerID: "1",
			Status:   Calling,
			CurrentHand: Hand{
				CurrentPlayerID: "5",
			},
		},
		AdminID: "1",
	}
}

func SevenPlayerGame() Game {
	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{Player1(), Player2(), Player3(), Player4(), Player5(), Player6(), Player7()},
		Dummy:     Dummy(),
		CurrentRound: Round{
			DealerID: "1",
			Status:   Calling,
			CurrentHand: Hand{
				CurrentPlayerID: "5",
			},
		},
		AdminID: "1",
	}
}

func CalledGame() Game {
	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{Player1(), Player2()},
		Dummy:     Dummy(),
		CurrentRound: Round{
			DealerID: "1",
			GoerID:   "2",
			Status:   Called,
			CurrentHand: Hand{
				CurrentPlayerID: "2",
			},
		},
		AdminID: "1",
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
