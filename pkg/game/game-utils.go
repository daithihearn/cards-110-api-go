package game

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// validateNumberOfPlayers Validate the number of players is in the range 2-6
func validateNumberOfPlayers(playerIDs []string) error {
	// Validate number of players is in the range 2-6
	if len(playerIDs) < 2 || len(playerIDs) > 6 {
		return fmt.Errorf("invalid number of players (%d)", len(playerIDs))
	}
	return nil
}

// shuffle a slice of strings
func shuffle(input []string) []string {
	shuffled := make([]string, len(input))
	perm := rand.Perm(len(input))

	for i, v := range perm {
		shuffled[v] = input[i]
	}

	return shuffled
}

// createPlayers Create the players for a game.
// There can only be 2-6 players
// If there are 6 players we play 3 teams of 2
// All other games are played individually
func createPlayers(playerIDs []string) ([]Player, error) {
	// Validate number of players is in the range 2-6
	if err := validateNumberOfPlayers(playerIDs); err != nil {
		return nil, err
	}

	// Create the players
	players := make([]Player, len(playerIDs))

	for i, playerID := range playerIDs {
		teamID := i + 1
		if len(playerIDs) == 6 {
			// Assign teams for 6 players: 1&4, 2&5, 3&6
			teamID = (i % 3) + 1
		}
		players[i] = Player{
			ID:     playerID,
			Seat:   i + 1,
			TeamID: strconv.Itoa(teamID),
		}
	}

	return players, nil
}

func nextPlayer(players []Player, currentPlayerId string) (Player, error) {
	currentIndex := -1
	for i, player := range players {
		if player.ID == currentPlayerId {
			currentIndex = i
			break
		}
	}

	if currentIndex == -1 {
		return Player{}, errors.New(fmt.Sprintf("Can't find player %s", currentPlayerId))
	}

	nextIndex := (currentIndex + 1) % len(players)
	if players[nextIndex].ID == "dummy" {
		return nextPlayer(players, players[nextIndex].ID)
	}

	return players[nextIndex], nil
}

func createFirstRound(players []Player, dealerID string) (Round, error) {
	timestamp := time.Now()

	// Create the current hand
	currentPlayer, err := nextPlayer(players, dealerID)
	if err != nil {
		return Round{}, err
	}
	hand := Hand{
		Timestamp:       timestamp,
		CurrentPlayerID: currentPlayer.ID,
	}

	// Create the round
	round := Round{
		Timestamp:   timestamp,
		Number:      1,
		DealerID:    dealerID,
		Status:      CALLING,
		CurrentHand: hand,
	}

	return round, nil
}

func NewGame(playerIDs []string, name string, adminID string) (Game, error) {
	// Validate number of players is in the range 2-6
	err := validateNumberOfPlayers(playerIDs)
	if err != nil {
		return Game{}, err
	}

	// Randomise the order of the players
	shuffledPlayerIDs := shuffle(playerIDs)

	// Create the players
	players, err := createPlayers(shuffledPlayerIDs)
	if err != nil {
		return Game{}, err
	}

	// Assign a dealer and create the first round
	dealer := playerIDs[0]
	round, err := createFirstRound(players, dealer)

	// Deal the cards
	deck := ShuffleCards(NewDeck())
	deck, players = DealCards(deck, players)

	// Create the game
	game := Game{
		ID:           "game-" + strconv.Itoa(rand.Intn(1000000)),
		Timestamp:    time.Now(),
		Name:         name,
		Status:       ACTIVE,
		AdminID:      adminID,
		Players:      players,
		CurrentRound: round,
		Deck:         deck,
	}

	return game, nil
}
