package game

import (
	"cards-110-api/pkg/cache"
	"cards-110-api/pkg/db"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type ServiceI interface {
	Create(ctx context.Context, playerIDs []string, name string, adminID string) (Game, error)
	Get(ctx context.Context, gameId string) (Game, bool, error)
	GetState(ctx context.Context, gameId string, playerId string) (State, bool, error)
	GetAll(ctx context.Context) ([]Game, error)
	Delete(ctx context.Context, gameId string, adminId string) error
	Call(ctx context.Context, gameId string, playerId string, call Call) (Game, error)
	SelectSuit(ctx context.Context, gameId string, playerId string, suit Suit, cards []CardName) (Game, error)
	Buy(ctx context.Context, gameId string, playerId string, cards []CardName) (Game, error)
	Play(ctx context.Context, gameId string, playerId string, card CardName) (Game, error)
}

type Service struct {
	Col   db.CollectionI[Game]
	Cache *cache.RedisCache[State]
}

func getCacheKey(gameId string, playerId string) string {
	return gameId + "-" + playerId
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

func (s *Service) GetState(ctx context.Context, gameId string, playerID string) (State, bool, error) {
	// Check the cache.
	state, found, err := s.Cache.Get(getCacheKey(gameId, playerID))
	if err != nil && found {
		return state, true, nil
	}

	// Get the game from the database.
	game, has, errG := s.Get(ctx, gameId)
	if errG != nil || !has {
		return State{}, has, errG
	}

	// Get the state for the player.
	state, err = game.GetState(playerID)
	if err != nil {
		return State{}, true, err
	}

	// Save the state to the cache.
	err = s.Cache.Set(getCacheKey(gameId, playerID), state, 10*time.Minute)
	if err != nil {
		log.Printf("Failed to save state to cache: %s", err)
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
func (s *Service) Call(ctx context.Context, gameId string, playerID string, call Call) (Game, error) {
	// Get the game from the database.
	game, has, err := s.Get(ctx, gameId)
	if err != nil {
		return Game{}, err
	}
	if !has {
		return Game{}, errors.New("game not found")
	}

	// Make the call.
	err = game.Call(playerID, call)
	if err != nil {
		return Game{}, err
	}

	// Save the game to the database.
	err = s.Col.UpdateOne(ctx, game, game.ID)
	if err != nil {
		return Game{}, err
	}

	// Save the state to the cache.
	state, err := game.GetState(playerID)
	if err != nil {
		return Game{}, err
	}

	err = s.Cache.Set(getCacheKey(gameId, playerID), state, 10*time.Minute)
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

	// Save the state to the cache.
	state, err := game.GetState(playerID)
	if err != nil {
		return Game{}, err
	}

	err = s.Cache.Set(getCacheKey(gameId, playerID), state, 10*time.Minute)
	if err != nil {
		return Game{}, err
	}

	return game, nil
}

// Buy cards
func (s *Service) Buy(ctx context.Context, gameId string, playerID string, cards []CardName) (Game, error) {
	// Get the game from the database.
	game, has, err := s.Get(ctx, gameId)
	if err != nil {
		return Game{}, err
	}
	if !has {
		return Game{}, errors.New("game not found")
	}

	// Buy the cards.
	err = game.Buy(playerID, cards)
	if err != nil {
		return Game{}, err
	}

	// Save the game to the database.
	err = s.Col.UpdateOne(ctx, game, game.ID)
	if err != nil {
		return Game{}, err
	}

	// Save the state to the cache.
	state, err := game.GetState(playerID)
	if err != nil {
		return Game{}, err
	}

	err = s.Cache.Set(getCacheKey(gameId, playerID), state, 10*time.Minute)
	if err != nil {
		return Game{}, err
	}

	return game, nil
}

// Play a card
func (s *Service) Play(ctx context.Context, gameId string, playerID string, card CardName) (Game, error) {
	// Get the game from the database.
	game, has, err := s.Get(ctx, gameId)
	if err != nil {
		return Game{}, err
	}
	if !has {
		return Game{}, errors.New("game not found")
	}

	// Play the card.
	err = game.Play(playerID, card)
	if err != nil {
		return Game{}, err
	}

	// Save the game to the database.
	err = s.Col.UpdateOne(ctx, game, game.ID)
	if err != nil {
		return Game{}, err
	}

	// Save the state to the cache.
	state, err := game.GetState(playerID)
	if err != nil {
		return Game{}, err
	}

	err = s.Cache.Set(getCacheKey(gameId, playerID), state, 10*time.Minute)
	if err != nil {
		return Game{}, err
	}

	// Invalidate the stats caches if the game is over
	if game.Status == Completed {
		for _, player := range game.Players {
			err = s.Cache.Delete("stats-" + player.ID)
			if err != nil {
				return Game{}, err
			}
		}
	}

	return game, nil
}
