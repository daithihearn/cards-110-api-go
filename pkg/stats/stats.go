package stats

import "time"

type PlayerStats struct {
	GameID    string    `bson:"gameId" json:"gameId"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	Winner    bool      `bson:"winner" json:"winner"`
	Score     int       `bson:"score" json:"score"`
	Rings     int       `bson:"rings" json:"rings"`
}
