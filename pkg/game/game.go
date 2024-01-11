package game

import (
	"fmt"
	"time"
)

type Status string

const (
	ACTIVE    Status = "ACTIVE"
	FINISHED         = "FINISHED"
	COMPLETED        = "COMPLETED"
	CANCELLED        = "CANCELLED"
)

type RoundStatus string

const (
	CALLING RoundStatus = "CALLING"
	CALLED              = "CALLED"
	BUYING              = "BUYING"
	PLAYING             = "PLAYING"
)

type Player struct {
	ID     string     `bson:"_id,omitempty" json:"id"`
	Seat   int        `bson:"seatNumber" json:"seatNumber"`
	Call   int        `bson:"call" json:"call"`
	Cards  []CardName `bson:"cards" json:"-"`
	Bought int        `bson:"cardsBought" json:"cardsBought"`
	Score  int        `bson:"score" json:"score"`
	Rings  int        `bson:"rings" json:"rings"`
	TeamID string     `bson:"teamId" json:"teamId"`
	Winner bool       `bson:"winner" json:"winner"`
}

type PlayedCard struct {
	PlayerID string   `bson:"playerId" json:"playerId"`
	Card     CardName `bson:"card" json:"card"`
}

type Hand struct {
	Timestamp       time.Time    `bson:"timestamp" json:"timestamp"`
	LeadOut         CardName     `bson:"leadOut" json:"leadOut"`
	CurrentPlayerID string       `bson:"currentPlayerId" json:"currentPlayerId"`
	PlayedCards     []PlayedCard `bson:"playedCards" json:"playedCards"`
}

type Round struct {
	Timestamp      time.Time   `bson:"timestamp" json:"timestamp"`
	Number         int         `bson:"number" json:"number"`
	DealerID       string      `bson:"dealerId" json:"dealerId"`
	GoerID         string      `bson:"goerId" json:"goerId"`
	Suit           Suit        `bson:"suit" json:"suit"`
	Status         RoundStatus `bson:"status" json:"status"`
	CurrentHand    Hand        `bson:"currentHand" json:"currentHand"`
	DealerSeeing   bool        `bson:"dealerSeeingCall" json:"dealerSeeingCall"`
	CompletedHands []Hand      `bson:"completedHands" json:"completedHands"`
}

type Game struct {
	ID           string     `bson:"_id,omitempty" json:"id"`
	AdminID      string     `bson:"adminId" json:"adminId"`
	Timestamp    time.Time  `bson:"timestamp" json:"timestamp"`
	Name         string     `bson:"name" json:"name"`
	Status       Status     `bson:"status" json:"status"`
	Players      []Player   `bson:"players" json:"players"`
	CurrentRound Round      `bson:"currentRound" json:"currentRound"`
	Completed    []Round    `bson:"completedRounds" json:"completedRounds"`
	Deck         []CardName `bson:"deck" json:"-"`
}

type GameState struct {
	ID           string     `json:"id"`
	Status       Status     `json:"status"`
	Me           Player     `json:"me"`
	IamSpectator bool       `json:"iamSpectator"`
	IsMyGo       bool       `json:"isMyGo"`
	IamGoer      bool       `json:"iamGoer"`
	IamDealer    bool       `json:"iamDealer"`
	IamAdmin     bool       `json:"iamAdmin"`
	MaxCall      int        `json:"maxCall"`
	Players      []Player   `json:"players"`
	Round        Round      `json:"round"`
	Cards        []CardName `json:"cards"`
}

func (g *Game) GetState(playerID string) (GameState, error) {
	// Get the current player
	var me Player
	for _, p := range g.Players {
		if p.ID == playerID {
			me = p
			break
		}
	}
	if me.ID == "" {
		return GameState{}, fmt.Errorf("player not found in game")
	}

	// 1. Find dummy
	var dummy Player
	if g.CurrentRound.GoerID == playerID {
		for _, p := range g.Players {
			if p.ID == "dummy" {
				dummy = p
				break
			}
		}
	}

	// 2. Get max call
	maxCall := -1
	for _, p := range g.Players {
		if p.Call > maxCall {
			maxCall = p.Call
		}
	}

	// 3. Add dummy if applicable
	iamGoer := g.CurrentRound.GoerID == playerID
	if iamGoer && g.CurrentRound.Status == CALLED && dummy.ID != "" {
		me.Cards = append(me.Cards, dummy.Cards...)
	}

	// 4. Return player's game state
	gameState := GameState{
		ID:           g.ID,
		Me:           me,
		IamSpectator: false,
		IsMyGo:       g.CurrentRound.CurrentHand.CurrentPlayerID == me.ID,
		IamGoer:      iamGoer,
		IamDealer:    g.CurrentRound.DealerID == me.ID,
		IamAdmin:     g.AdminID == me.ID,
		Cards:        me.Cards,
		Status:       g.Status,
		Round:        g.CurrentRound,
		MaxCall:      maxCall,
		Players:      make([]Player, 0),
	}

	for _, p := range g.Players {
		if p.ID != "dummy" {
			gameState.Players = append(gameState.Players, p)
		}
	}

	return gameState, nil
}

type PlayerStats struct {
	GameID    string    `bson:"gameId" json:"gameId"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	Winner    bool      `bson:"winner" json:"winner"`
	Score     int       `bson:"score" json:"score"`
	Rings     int       `bson:"rings" json:"rings"`
}

type CreateGameRequest struct {
	PlayerIDs []string `json:"players"`
	Name      string   `json:"name"`
}
