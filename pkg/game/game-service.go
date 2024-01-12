package game

import (
	"cards-110-api/pkg/db"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type ServiceI interface {
	Create(ctx context.Context, playerIDs []string, name string, adminID string) (Game, error)
	Get(ctx context.Context, gameId string) (Game, bool, error)
	GetAll(ctx context.Context) ([]Game, error)
}

type Service struct {
	Col db.CollectionI[Game]
}

// Create a new game.
func (s *Service) Create(ctx context.Context, playerIDs []string, name string, adminID string) (Game, error) {
	log.Printf("Creating new game (%s)", name)

	// Check for duplicate player IDs.
	uniquePlayerIDs := make(map[string]bool)
	for _, id := range playerIDs {
		uniquePlayerIDs[id] = true
	}
	if len(uniquePlayerIDs) != len(playerIDs) {
		return Game{}, errors.New("duplicate player IDs")
	}

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
