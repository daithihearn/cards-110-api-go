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

func ParseCall(c string) (Call, error) {
	switch c {
	case "0":
		return Pass, nil
	case "10":
		return Ten, nil
	case "15":
		return Fifteen, nil
	case "20":
		return Twenty, nil
	case "25":
		return TwentyFive, nil
	case "30":
		return Jink, nil
	default:
		return 0, fmt.Errorf("invalid call")
	}
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
		PlayedCards:     make([]PlayedCard, 0),
	}

	// Create the round
	round := Round{
		Timestamp:      timestamp,
		Number:         1,
		DealerID:       dealerID,
		Status:         Calling,
		CurrentHand:    hand,
		CompletedHands: make([]Hand, 0),
	}

	return round, nil
}

func NewGame(playerIDs []string, name string, adminID string) (Game, error) {
	// Validate number of players is in the range 2-6
	err := validateNumberOfPlayers(playerIDs)
	if err != nil {
		return Game{}, err
	}

	// Verify the admin is in the list of players
	adminFound := false
	for _, playerID := range playerIDs {
		if playerID == adminID {
			adminFound = true
			break
		}
	}
	if !adminFound {
		return Game{}, errors.New("admin not found in players")
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
	deck, hands, err := DealCards(ShuffleCards(NewDeck()), len(players))
	if err != nil {
		return Game{}, err
	}
	var dummy []CardName
	for i, hand := range hands {
		if i >= len(players) {
			dummy = hand
			break
		}
		players[i].Cards = hand
	}

	// Create the game
	game := Game{
		ID:           "game-" + strconv.Itoa(rand.Intn(1000000)),
		Timestamp:    time.Now(),
		Name:         name,
		Status:       Active,
		AdminID:      adminID,
		Players:      players,
		Dummy:        dummy,
		CurrentRound: round,
		Deck:         deck,
	}

	return game, nil
}

// containsAllUnique checks if targetSlice contains all unique elements of referenceSlice.
func containsAllUnique(referenceSlice, targetSlice []CardName) bool {
	if len(targetSlice) > len(referenceSlice) {
		return false
	}

	existsInReferenceSlice := make(map[CardName]bool)
	for _, item := range referenceSlice {
		existsInReferenceSlice[item] = true
	}

	seenInTargetSlice := make(map[CardName]bool)
	for _, item := range targetSlice {
		if _, ok := existsInReferenceSlice[item]; !ok {
			return false // Element in targetSlice is not in referenceSlice
		}
		if _, seen := seenInTargetSlice[item]; seen {
			return false // Element is not unique in targetSlice
		}
		seenInTargetSlice[item] = true
	}
	return true
}

// compare checks if targetSlice and referenceSlice are equivalent
func compare(referenceSlice, targetSlice []CardName) bool {
	if len(referenceSlice) != len(targetSlice) {
		return false
	}
	return containsAllUnique(referenceSlice, targetSlice)
}
