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
		TeamID: "1",
		Seat:   1,
		Call:   0,
		Cards:  []CardName{ACE_HEARTS, KING_HEARTS, QUEEN_HEARTS, JACK_HEARTS, TEN_HEARTS},
		Bought: 0,
	}
}

func Player2() Player {
	return Player{
		ID:     "2",
		TeamID: "2",
		Seat:   2,
		Call:   0,
		Cards:  []CardName{ACE_DIAMONDS, KING_DIAMONDS, QUEEN_DIAMONDS, JACK_DIAMONDS, TEN_DIAMONDS},
		Bought: 0,
	}
}

func Player3() Player {
	return Player{
		ID:     "3",
		TeamID: "3",
		Seat:   3,
		Call:   0,
		Cards:  []CardName{ACE_CLUBS, KING_CLUBS, QUEEN_CLUBS, JACK_CLUBS, TEN_CLUBS},
		Bought: 0,
	}
}

func Player4() Player {
	return Player{
		ID:     "4",
		TeamID: "4",
		Seat:   4,
		Call:   0,
		Cards:  []CardName{NINE_CLUBS, EIGHT_CLUBS, SEVEN_CLUBS, SIX_CLUBS, FIVE_CLUBS},
		Bought: 0,
	}
}

func Player5() Player {
	return Player{
		ID:     "5",
		TeamID: "5",
		Seat:   5,
		Call:   0,
		Cards:  []CardName{NINE_DIAMONDS, EIGHT_DIAMONDS, SEVEN_DIAMONDS, SIX_DIAMONDS, FIVE_DIAMONDS},
		Bought: 0,
	}
}

func Player6() Player {
	return Player{
		ID:     "6",
		TeamID: "6",
		Seat:   6,
		Call:   0,
		Cards:  []CardName{NINE_HEARTS, EIGHT_HEARTS, SEVEN_HEARTS, SIX_HEARTS, FIVE_HEARTS},
		Bought: 0,
	}
}

func Player7() Player {
	return Player{
		ID:     "7",
		TeamID: "7",
		Seat:   7,
		Call:   0,
		Cards:  []CardName{NINE_SPADES, EIGHT_SPADES, SEVEN_SPADES, SIX_SPADES, FIVE_SPADES},
		Bought: 0,
	}
}

func PlayerCalled() Player {
	return Player{
		ID:     "PlayerCalled",
		TeamID: "PlayerCalled",
		Seat:   1,
		Call:   20,
		Cards:  []CardName{NINE_HEARTS, EIGHT_HEARTS, SEVEN_HEARTS, SIX_HEARTS, FOUR_SPADES},
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
	p4 := Player4()
	p4.TeamID = "1"
	p5 := Player5()
	p5.TeamID = "2"
	p6 := Player6()
	p6.TeamID = "3"

	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{Player1(), Player2(), Player3(), p4, p5, p6},
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

func CalledGameFivePlayers() Game {
	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{Player1(), Player2(), PlayerCalled(), Player4(), Player5()},
		Dummy:     Dummy(),
		CurrentRound: Round{
			DealerID: "1",
			GoerID:   "PlayerCalled",
			Status:   Called,
			CurrentHand: Hand{
				CurrentPlayerID: "PlayerCalled",
			},
		},
		AdminID: "1",
	}
}

func CalledGameThreePlayers() Game {
	game := CalledGameFivePlayers()
	game.Players = []Player{Player1(), Player2(), PlayerCalled()}
	return game
}

func BuyingGame(dealerId string) Game {
	deck := NewDeck()

	p1 := Player1()
	p1.Cards = []CardName{deck[0], deck[1], deck[2], deck[3], deck[4]}
	deck = deck[5:]
	p2 := Player2()
	p2.Cards = []CardName{deck[0], deck[1], deck[2], deck[3], deck[4]}
	deck = deck[5:]
	p3 := Player3()
	p3.Cards = []CardName{deck[0], deck[1], deck[2], deck[3], deck[4]}
	deck = deck[5:]
	pCalled := PlayerCalled()
	pCalled.Cards = []CardName{deck[0], deck[1], deck[2], deck[3], deck[4]}
	deck = deck[5:]
	dummy := []CardName{deck[0], deck[1], deck[2], deck[3], deck[4]}
	deck = deck[5:]

	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{p1, p2, p3, pCalled},
		CurrentRound: Round{
			DealerID: dealerId,
			GoerID:   "PlayerCalled",
			Status:   Buying,
			CurrentHand: Hand{
				CurrentPlayerID: "2",
			},
			Suit:   Hearts,
			Number: 1,
		},
		Dummy:   dummy,
		AdminID: "1",
		Deck:    deck,
	}
}

func PlayingGame_RoundStart(dealerId string) Game {
	game := BuyingGame(dealerId)
	game.CurrentRound.Status = Playing
	game.CurrentRound.CurrentHand.CurrentPlayerID = "1"
	game.Dummy = []CardName{}
	return game
}

func PlayingGame_Hand1Complete(dealerId string) Game {
	game := PlayingGame_RoundStart(dealerId)

	playedCards := make([]PlayedCard, 0)

	// Play the first card from each player's hand.
	for i, player := range game.Players {
		playedCards = append(playedCards, PlayedCard{
			PlayerID: player.ID,
			Card:     player.Cards[0],
		})
		game.Players[i].Cards = player.Cards[1:]
	}

	game.CurrentRound.Suit = Hearts
	game.CurrentRound.CurrentHand.PlayedCards = playedCards

	return game
}

func PlayingGame_FinalHandComplete(dealerId string) Game {
	game := PlayingGame_RoundStart(dealerId)

	playedCards := make([]PlayedCard, 0)

	// Play the first card from each player's hand.
	for i, player := range game.Players {

		playedCards = append(playedCards, PlayedCard{
			PlayerID: player.ID,
			Card:     player.Cards[0],
		})
		game.Players[i].Cards = player.Cards[1:]
	}
	game.CurrentRound.CurrentHand.PlayedCards = playedCards

	// Populate the completed hands from the rest of the cards
	// loop 4 times
	for i := 0; i < 4; i++ {
		cards := make([]PlayedCard, 0)
		for j, player := range game.Players {
			cards = append(cards, PlayedCard{
				PlayerID: player.ID,
				Card:     player.Cards[0],
			})
			game.Players[j].Cards = player.Cards[1:]
		}
		game.CurrentRound.CompletedHands = append(game.CurrentRound.CompletedHands, Hand{
			PlayedCards: cards,
			LeadOut:     cards[0].Card,
		})
	}

	game.CurrentRound.Suit = Hearts

	return game
}

func PlayingGame_AllHandsComplete(dealerId string) Game {
	game := PlayingGame_RoundStart(dealerId)

	game.CurrentRound.CurrentHand.PlayedCards = make([]PlayedCard, 0)

	// Populate the completed hands from the rest of the cards
	// loop 4 times
	for i := 0; i < 5; i++ {
		cards := make([]PlayedCard, 0)
		for j, player := range game.Players {
			cards = append(cards, PlayedCard{
				PlayerID: player.ID,
				Card:     player.Cards[0],
			})
			game.Players[j].Cards = player.Cards[1:]
		}
		game.CurrentRound.CompletedHands = append(game.CurrentRound.CompletedHands, Hand{
			PlayedCards: cards,
			LeadOut:     cards[0].Card,
		})
	}

	game.CurrentRound.Suit = Hearts

	return game
}

func PlayingGame_Jink() Game {
	p1 := Player1()
	p1.Cards = []CardName{}
	p2 := Player2()
	p2.Cards = []CardName{}
	p2.Call = Twenty
	p3 := Player3()
	p3.Cards = []CardName{}
	p3.Call = Jink

	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{p1, p2, p3},
		CurrentRound: Round{
			DealerID: "1",
			GoerID:   "3",
			Status:   Buying,
			CurrentHand: Hand{
				CurrentPlayerID: "2",
			},
			CompletedHands: []Hand{
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "3", Card: ACE_CLUBS}, {PlayerID: "1", Card: TEN_HEARTS}, {PlayerID: "2", Card: TEN_DIAMONDS},
					},
				},
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "3", Card: FIVE_CLUBS}, {PlayerID: "1", Card: KING_CLUBS}, {PlayerID: "2", Card: ACE_DIAMONDS},
					},
				},
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "3", Card: JACK_CLUBS}, {PlayerID: "1", Card: KING_HEARTS}, {PlayerID: "2", Card: KING_DIAMONDS},
					},
				},
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "3", Card: JOKER}, {PlayerID: "1", Card: QUEEN_HEARTS}, {PlayerID: "2", Card: QUEEN_DIAMONDS},
					},
				},
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "3", Card: ACE_HEARTS}, {PlayerID: "1", Card: JACK_HEARTS}, {PlayerID: "2", Card: JACK_DIAMONDS},
					},
				},
			},
			Suit:   Clubs,
			Number: 1,
		},
		Dummy:   Dummy(),
		AdminID: "1",
		Deck:    NewDeck(),
	}
}

func PlayingGame_Jink_Doubles() Game {
	p1 := Player1()
	p1.Cards = []CardName{}
	p2 := Player2()
	p2.Cards = []CardName{}
	p2.Call = Twenty
	p3 := Player3()
	p3.Cards = []CardName{}
	p3.Call = Jink
	p4 := Player4()
	p4.Cards = []CardName{}
	p4.TeamID = "1"
	p5 := Player5()
	p5.Cards = []CardName{}
	p5.TeamID = "2"
	p6 := Player6()
	p6.Cards = []CardName{}
	p6.TeamID = "3"

	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{p1, p2, p3, p4, p5, p6},
		CurrentRound: Round{
			DealerID: "1",
			GoerID:   "3",
			Status:   Playing,
			CurrentHand: Hand{
				CurrentPlayerID: "2",
			},
			CompletedHands: []Hand{
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "1", Card: TEN_HEARTS}, {PlayerID: "2", Card: TEN_DIAMONDS}, {PlayerID: "3", Card: ACE_CLUBS}, {PlayerID: "4", Card: TEN_CLUBS}, {PlayerID: "5", Card: TEN_SPADES}, {PlayerID: "6", Card: TEN_DIAMONDS},
					},
				},
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "1", Card: KING_CLUBS}, {PlayerID: "2", Card: ACE_DIAMONDS}, {PlayerID: "3", Card: FIVE_CLUBS}, {PlayerID: "4", Card: ACE_CLUBS}, {PlayerID: "5", Card: ACE_SPADES}, {PlayerID: "6", Card: ACE_DIAMONDS},
					},
				},
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "1", Card: KING_HEARTS}, {PlayerID: "2", Card: KING_DIAMONDS}, {PlayerID: "3", Card: JACK_CLUBS}, {PlayerID: "4", Card: SIX_DIAMONDS}, {PlayerID: "5", Card: KING_SPADES}, {PlayerID: "6", Card: KING_DIAMONDS},
					},
				},
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "1", Card: QUEEN_HEARTS}, {PlayerID: "2", Card: TWO_HEARTS}, {PlayerID: "3", Card: QUEEN_DIAMONDS}, {PlayerID: "4", Card: QUEEN_CLUBS}, {PlayerID: "5", Card: QUEEN_SPADES}, {PlayerID: "6", Card: JOKER},
					},
				},
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "1", Card: JACK_HEARTS}, {PlayerID: "2", Card: JACK_DIAMONDS}, {PlayerID: "3", Card: ACE_HEARTS}, {PlayerID: "4", Card: THREE_DIAMONDS}, {PlayerID: "5", Card: JACK_SPADES}, {PlayerID: "6", Card: JACK_DIAMONDS},
					},
				},
			},
			Suit:   Clubs,
			Number: 1,
		},
		Dummy:   Dummy(),
		AdminID: "1",
		Deck:    NewDeck(),
	}
}

func PlayingGame_DoesntMakeContract_Doubles() Game {
	p1 := Player1()
	p1.Cards = []CardName{}
	p2 := Player2()
	p2.Cards = []CardName{}
	p2.Call = Twenty
	p3 := Player3()
	p3.Cards = []CardName{}
	p3.Call = TwentyFive
	p4 := Player4()
	p4.Cards = []CardName{}
	p4.TeamID = "1"
	p5 := Player5()
	p5.Cards = []CardName{}
	p5.TeamID = "2"
	p6 := Player6()
	p6.Cards = []CardName{}
	p6.TeamID = "3"

	return Game{
		ID:        "1",
		Name:      "Test Game",
		Status:    Active,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{p1, p2, p3, p4, p5, p6},
		CurrentRound: Round{
			DealerID: "1",
			GoerID:   "3",
			Status:   Playing,
			CurrentHand: Hand{
				CurrentPlayerID: "2",
			},
			CompletedHands: []Hand{
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "1", Card: TEN_HEARTS}, {PlayerID: "2", Card: TEN_DIAMONDS}, {PlayerID: "3", Card: ACE_CLUBS}, {PlayerID: "4", Card: TEN_CLUBS}, {PlayerID: "5", Card: TEN_SPADES}, {PlayerID: "6", Card: TEN_DIAMONDS},
					},
				},
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "1", Card: FIVE_CLUBS}, {PlayerID: "2", Card: ACE_DIAMONDS}, {PlayerID: "3", Card: SIX_CLUBS}, {PlayerID: "4", Card: ACE_CLUBS}, {PlayerID: "5", Card: ACE_SPADES}, {PlayerID: "6", Card: ACE_DIAMONDS},
					},
				},
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "1", Card: KING_HEARTS}, {PlayerID: "2", Card: KING_DIAMONDS}, {PlayerID: "3", Card: KING_CLUBS}, {PlayerID: "4", Card: JACK_CLUBS}, {PlayerID: "5", Card: KING_SPADES}, {PlayerID: "6", Card: KING_DIAMONDS},
					},
				},
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "1", Card: QUEEN_HEARTS}, {PlayerID: "2", Card: TWO_HEARTS}, {PlayerID: "3", Card: QUEEN_DIAMONDS}, {PlayerID: "4", Card: QUEEN_CLUBS}, {PlayerID: "5", Card: QUEEN_SPADES}, {PlayerID: "6", Card: JOKER},
					},
				},
				{
					PlayedCards: []PlayedCard{
						{PlayerID: "1", Card: JACK_HEARTS}, {PlayerID: "2", Card: JACK_DIAMONDS}, {PlayerID: "3", Card: ACE_HEARTS}, {PlayerID: "4", Card: THREE_DIAMONDS}, {PlayerID: "5", Card: JACK_SPADES}, {PlayerID: "6", Card: JACK_DIAMONDS},
					},
				},
			},
			Suit:   Clubs,
			Number: 1,
		},
		Dummy:   Dummy(),
		AdminID: "1",
		Deck:    NewDeck(),
	}
}

func PlayingGame_Thirty() Game {
	game := PlayingGame_Jink()
	game.Players[2].Call = TwentyFive
	return game
}

func CompletedGame() Game {
	p1 := Player1()
	p1.Score = 110
	p2 := Player2()
	p2.Score = 90

	return Game{
		ID:        "2",
		Name:      "Test Game",
		Status:    Completed,
		Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Players:   []Player{p1, p2},
		AdminID:   "1",
	}
}
