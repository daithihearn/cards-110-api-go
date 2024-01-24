package stats

import (
	"cards-110-api/pkg/cache"
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
	GetStats(ctx context.Context, playerId string) ([]PlayerStats, error)
}

type Service struct {
	Col   db.CollectionI[game.Game]
	Cache *cache.RedisCache[[]PlayerStats]
}

func getCacheKey(playerId string) string {
	return "stats-" + playerId
}

// GetStats Get the stats for a player.
func (s *Service) GetStats(ctx context.Context, playerID string) ([]PlayerStats, error) {
	// Check the cache.
	stats, found, err := s.Cache.Get(getCacheKey(playerID))
	if err == nil && found {
		return stats, nil
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "status", Value: "FINISHED"}, {Key: "players._id", Value: playerID}}}},
		{{Key: "$unwind", Value: "$players"}},
		{{Key: "$match", Value: bson.D{{Key: "players._id", Value: playerID}}}},
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
		return []PlayerStats{}, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cursor, ctx)

	// Return an empty slice if there are no results.
	if cursor.RemainingBatchLength() == 0 {
		return []PlayerStats{}, nil
	}

	// Iterate over the cursor and decode each result.
	var results []PlayerStats
	for cursor.Next(ctx) {
		var result bson.M
		if err = cursor.Decode(&result); err != nil {
			// Log the detailed error
			log.Printf("Error decoding cursor result: %v", err)
			return []PlayerStats{}, err
		}

		// Map the result to a PlayerStats struct.
		playerStats := PlayerStats{}

		// Safely assert types
		if gameId, ok := result["gameId"].(string); ok {
			playerStats.GameID = gameId
		} else {
			// Handle missing or invalid gameId
			return []PlayerStats{}, fmt.Errorf("failed to decode gameID")
		}

		if timestamp, ok := result["timestamp"].(primitive.DateTime); ok {
			// Convert primitive.DateTime to time.Time
			playerStats.Timestamp = time.Unix(int64(timestamp)/1000, 0)
		} else {
			// Handle missing or invalid timestamp
			return []PlayerStats{}, fmt.Errorf("timestamp is not a valid DateTime")
		}

		if winner, ok := result["winner"].(bool); ok {
			playerStats.Winner = winner
		} else {
			return []PlayerStats{}, fmt.Errorf("winner is not a bool")
		}

		if score, ok := result["score"].(int32); ok {
			playerStats.Score = int(score)
		} else {
			return []PlayerStats{}, fmt.Errorf("score is not an int")
		}

		if rings, ok := result["rings"].(int32); ok {
			playerStats.Rings = int(rings)
		} else {
			return []PlayerStats{}, fmt.Errorf("rings is not an int")
		}

		results = append(results, playerStats)

	}

	// Save the result to the cache.
	err = s.Cache.Set(getCacheKey(playerID), results, 2*time.Minute)
	if err != nil {
		log.Printf("Failed to save state to cache: %s", err)
	}

	return results, nil
}
