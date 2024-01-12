package stats

import (
	"cards-110-api/pkg/db"
	"cards-110-api/pkg/game"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type ServiceI interface {
	GetStats(ctx context.Context, playerId string) ([]game.PlayerStats, error)
}

type Service struct {
	Col db.CollectionI[game.Game]
}

// GetStats Get the stats for a player.
func (s *Service) GetStats(ctx context.Context, playerId string) ([]game.PlayerStats, error) {

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "status", Value: "FINISHED"}, {Key: "players._id", Value: playerId}}}},
		{{Key: "$unwind", Value: "$players"}},
		{{Key: "$match", Value: bson.D{{Key: "players._id", Value: playerId}}}},
		{{Key: "$project", Value: bson.D{
			{Key: "gameId", Value: "$_id"},
			{Key: "timestamp", Value: "$timestamp"},
			{Key: "winner", Value: "$players.winner"},
			{Key: "score", Value: "$players.score"},
			{Key: "rings", Value: "$players.rings"},
		}}},
	}

	cursor, err := s.Col.Aggregate(ctx, pipeline)
	if err != nil {
		return []game.PlayerStats{}, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cursor, ctx)

	// Return an empty slice if there are no results.
	if cursor.RemainingBatchLength() == 0 {
		return []game.PlayerStats{}, nil
	}

	// Iterate over the cursor and decode each result.
	var results []game.PlayerStats
	for cursor.Next(ctx) {
		var result bson.M
		if err = cursor.Decode(&result); err != nil {
			// Log the detailed error
			log.Printf("Error decoding cursor result: %v", err)
			return []game.PlayerStats{}, err
		}

		// Map the result to a PlayerStats struct.
		playerStats := game.PlayerStats{}

		// Safely assert types
		if gameId, ok := result["gameId"].(string); ok {
			playerStats.GameID = gameId
		} else {
			// Handle missing or invalid gameId
			return []game.PlayerStats{}, fmt.Errorf("failed to decode gameID")
		}

		if timestamp, ok := result["timestamp"].(primitive.DateTime); ok {
			// Convert primitive.DateTime to time.Time
			playerStats.Timestamp = time.Unix(int64(timestamp)/1000, 0)
		} else {
			// Handle missing or invalid timestamp
			return []game.PlayerStats{}, fmt.Errorf("timestamp is not a valid DateTime")
		}

		if winner, ok := result["winner"].(bool); ok {
			playerStats.Winner = winner
		} else {
			return []game.PlayerStats{}, fmt.Errorf("winner is not a bool")
		}

		if score, ok := result["score"].(int32); ok {
			playerStats.Score = int(score)
		} else {
			return []game.PlayerStats{}, fmt.Errorf("score is not an int")
		}

		if rings, ok := result["rings"].(int32); ok {
			playerStats.Rings = int(rings)
		} else {
			return []game.PlayerStats{}, fmt.Errorf("rings is not an int")
		}

		results = append(results, playerStats)

	}

	return results, nil
}
