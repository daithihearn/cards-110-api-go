package game

import (
	"time"
)

type Status string

const (
	Active    Status = "ACTIVE"
	Completed        = "COMPLETED"
)

type RoundStatus string

const (
	Calling RoundStatus = "CALLING"
	Called              = "CALLED"
	Buying              = "BUYING"
	Playing             = "PLAYING"
)

type Call int

const (
	Pass       Call = 0
	Ten             = 10
	Fifteen         = 15
	Twenty          = 20
	TwentyFive      = 25
	Jink            = 30
)

type Player struct {
	ID     string     `bson:"_id,omitempty" json:"id"`
	Seat   int        `bson:"seatNumber" json:"seatNumber"`
	Call   Call       `bson:"call" json:"call"`
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
	LeadOut         CardName     `bson:"leadOut" json:"leadOut,omitempty"`
	CurrentPlayerID string       `bson:"currentPlayerId" json:"currentPlayerId"`
	PlayedCards     []PlayedCard `bson:"playedCards" json:"playedCards"`
}

type Round struct {
	Timestamp      time.Time   `bson:"timestamp" json:"timestamp"`
	Number         int         `bson:"number" json:"number"`
	DealerID       string      `bson:"dealerId" json:"dealerId"`
	GoerID         string      `bson:"goerId" json:"goerId,omitempty"`
	Suit           Suit        `bson:"suit" json:"suit,omitempty"`
	Status         RoundStatus `bson:"status" json:"status"`
	CurrentHand    Hand        `bson:"currentHand" json:"currentHand"`
	DealerSeeing   bool        `bson:"dealerSeeingCall" json:"dealerSeeingCall"`
	CompletedHands []Hand      `bson:"completedHands" json:"completedHands"`
}

type Game struct {
	ID           string     `bson:"_id,omitempty" json:"id"`
	Revision     int        `bson:"revision" json:"revision"`
	AdminID      string     `bson:"adminId" json:"adminId"`
	Timestamp    time.Time  `bson:"timestamp" json:"timestamp"`
	Name         string     `bson:"name" json:"name"`
	Status       Status     `bson:"status" json:"status"`
	Players      []Player   `bson:"players" json:"players"`
	Dummy        []CardName `bson:"dummy" json:"-"`
	CurrentRound Round      `bson:"currentRound" json:"-"`
	Completed    []Round    `bson:"completedRounds" json:"-"`
	Deck         []CardName `bson:"deck" json:"-"`
}

type State struct {
	ID           string     `json:"id"`
	Revision     int        `bson:"revision" json:"revision"`
	Status       Status     `json:"status"`
	Me           Player     `json:"me"`
	IamSpectator bool       `json:"iamSpectator"`
	IsMyGo       bool       `json:"isMyGo"`
	IamGoer      bool       `json:"iamGoer"`
	IamDealer    bool       `json:"iamDealer"`
	IamAdmin     bool       `json:"iamAdmin"`
	MaxCall      Call       `json:"maxCall"`
	Players      []Player   `json:"players"`
	Round        Round      `json:"round"`
	PrevRound    Round      `json:"previousRound"`
	Cards        []CardName `json:"cards"`
}
