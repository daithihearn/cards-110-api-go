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
	GetState(ctx context.Context, gameId string, playerId string) (State, bool, error)
	GetAll(ctx context.Context) ([]Game, error)
	Delete(ctx context.Context, gameId string, adminId string) error
	Call(ctx context.Context, gameId string, playerId string, call Call) (Game, error)
	SelectSuit(ctx context.Context, id string, id2 string, suit Suit, cards []CardName) (Game, error)
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

func (s *Service) GetState(ctx context.Context, gameId string, playerId string) (State, bool, error) {
	// Get the game from the database.
	game, has, err := s.Get(ctx, gameId)
	if err != nil || !has {
		return State{}, has, err
	}

	// Get the state for the player.
	state, err := game.GetState(playerId)
	if err != nil {
		return State{}, true, err
	}

	return state, true, nil
}

// GetAll Get all games.
func (s *Service) GetAll(ctx context.Context) ([]Game, error) {
	return s.Col.Find(ctx, bson.M{})
}

// Delete a game.
func (s *Service) Delete(ctx context.Context, gameId string, adminId string) error {
	// Get the game from the database.
	game, has, err := s.Get(ctx, gameId)
	if err != nil {
		return err
	}
	if !has {
		return errors.New("game not found")
	}

	// Check correct admin
	if game.AdminID != adminId {
		return errors.New("not admin")
	}

	// Can only remove a game that is in an active state
	if game.Status != Active {
		return errors.New("can only delete games that are in an active state")
	}

	// Delete the game from the database.
	return s.Col.DeleteOne(ctx, game.ID)
}

// Call make a call
func (s *Service) Call(ctx context.Context, gameId string, playerId string, call Call) (Game, error) {
	// Get the game from the database.
	game, has, err := s.Get(ctx, gameId)
	if err != nil {
		return Game{}, err
	}
	if !has {
		return Game{}, errors.New("game not found")
	}

	// Make the call.
	err = game.Call(playerId, call)
	if err != nil {
		return Game{}, err
	}

	// Save the game to the database.
	err = s.Col.UpdateOne(ctx, game, game.ID)
	if err != nil {
		return Game{}, err
	}
	return game, nil
}

// SelectSuit select a suit
func (s *Service) SelectSuit(ctx context.Context, gameId string, playerID string, suit Suit, cards []CardName) (Game, error) {
	// Get the game from the database.
	game, has, err := s.Get(ctx, gameId)
	if err != nil {
		return Game{}, err
	}
	if !has {
		return Game{}, errors.New("game not found")
	}

	// Select the suit.
	err = game.SelectSuit(playerID, suit, cards)
	if err != nil {
		return Game{}, err
	}

	// Save the game to the database.
	err = s.Col.UpdateOne(ctx, game, game.ID)
	if err != nil {
		return Game{}, err
	}
	return game, nil
}
