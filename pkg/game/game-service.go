package game

import (
	"cards-110-api/pkg/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type ServiceI interface {
	Create(ctx context.Context, playerIDs []string, name string, adminID string) (Game, error)
	Get(ctx context.Context, gameId string) (Game, bool, error)
	GetAll(ctx context.Context) ([]Game, error)
	GetStats(ctx context.Context, playerId string) ([]PlayerStats, error)
}

type Service struct {
	Col db.CollectionI[Game]
}

// Create a new game.
func (s *Service) Create(ctx context.Context, playerIDs []string, name string, adminID string) (Game, error) {
	log.Printf("Creating new game (%s)", name)

	// Create a new game.
	game, err := NewGame(playerIDs, name, adminID)
	if err != nil {
		return Game{}, err
	}

	// Save the game to the database.
	err = s.Col.Upsert(ctx, game, game.ID)
	if err != nil {
		return Game{}, err
	}

	return game, nil
}

// Get a game by ID.
func (s *Service) Get(ctx context.Context, gameId string) (Game, bool, error) {
	return s.Col.FindOne(ctx, bson.M{"_id": gameId})
}

// GetAll Get all games.
func (s *Service) GetAll(ctx context.Context) ([]Game, error) {
	return s.Col.Find(ctx, bson.M{})
}

// GetStats Get the stats for a player.
func (s *Service) GetStats(ctx context.Context, playerId string) ([]PlayerStats, error) {

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

	return results, nil
}
