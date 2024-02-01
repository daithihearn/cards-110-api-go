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
	deck, dummy, hands, err := DealCards(ShuffleCards(NewDeck()), len(players))
	if err != nil {
		return Game{}, err
	}

	for i, hand := range hands {
		players[i].Cards = hand
	}

	// Create the game
	game := Game{
		ID:           strconv.Itoa(rand.Intn(10000000)),
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

func contains(referenceSlice []CardName, target CardName) bool {
	for _, item := range referenceSlice {
		if item == target {
			return true
		}
	}
	return false
}

// compare checks if targetSlice and referenceSlice are equivalent
func compare(referenceSlice, targetSlice []CardName) bool {
	if len(referenceSlice) != len(targetSlice) {
		return false
	}
	return containsAllUnique(referenceSlice, targetSlice)
}

func canRenage(leadOut Card, myTrumps []Card) bool {
	for _, trump := range myTrumps {
		if !trump.Renegable || trump.Value <= leadOut.Value {
			return false
		}
	}
	return true
}

func isFollowing(myCard CardName, myCards []CardName, currentHand Hand, suit Suit) bool {
	mySuit := myCard.Card().Suit
	leadOut := currentHand.LeadOut.Card()
	trumpLead := leadOut.Suit == suit || leadOut.Suit == Wild

	if trumpLead {
		var myTrumps []Card
		for _, card := range myCards {
			s := card.Card().Suit
			if s == suit || s == Wild {
				myTrumps = append(myTrumps, card.Card())
			}
		}

		return len(myTrumps) == 0 ||
			mySuit == suit ||
			mySuit == Wild ||
			canRenage(leadOut, myTrumps)
	}

	var mySuitedCards []Card
	for _, card := range myCards {
		if card.Card().Suit == leadOut.Suit {
			mySuitedCards = append(mySuitedCards, card.Card())
		}
	}

	return len(mySuitedCards) == 0 ||
		mySuit == suit ||
		mySuit == Wild ||
		mySuit == leadOut.Suit
}

// getActiveSuit Was a suit or wild card played? If not set the lead out card as the suit
func getActiveSuit(hand Hand, suit Suit) (Suit, error) {
	if suit == "" {
		return "", errors.New("suit not set")
	}
	// If a trump was played, the active suit is the trump suit
	for _, playedCard := range hand.PlayedCards {
		if playedCard.Card.Card().Suit == suit || playedCard.Card.Card().Suit == Wild {
			return suit, nil
		}
	}

	return hand.LeadOut.Card().Suit, nil
}

func findWinningCard(hand Hand, suit Suit) (PlayedCard, error) {
	if len(hand.PlayedCards) == 0 {
		return PlayedCard{}, errors.New("no cards played")
	}
	if !suit.isValid() {
		return PlayedCard{}, errors.New("invalid suit")
	}

	// Get active suit
	activeSuit, err := getActiveSuit(hand, suit)
	if err != nil {
		return PlayedCard{}, err
	}

	var winningCard PlayedCard
	for _, playedCard := range hand.PlayedCards {
		card := playedCard.Card.Card()
		if suit == activeSuit {
			// A trump card was played
			if (card.Suit == suit || card.Suit == Wild) && card.Value > winningCard.Card.Card().Value {
				winningCard = playedCard
			}
		} else {
			// A cold card will win as no trump was played
			if card.Suit == activeSuit {
				if card.ColdValue > winningCard.Card.Card().ColdValue {
					winningCard = playedCard
				}
			}
		}
	}
	if winningCard.Card == "" {
		return PlayedCard{}, errors.New("winning card not found")
	}
	return winningCard, nil
}

func findBestTrump(cards []PlayedCard, suit Suit) (PlayedCard, bool, error) {
	if len(cards) == 0 {
		return PlayedCard{}, false, errors.New("no cards played")
	}

	// 1. Verify at least one trump was played
	trumpPlayed := false
	for _, playedCard := range cards {
		if playedCard.Card.Card().Suit == suit || playedCard.Card.Card().Suit == Wild {
			trumpPlayed = true
			break
		}
	}

	if !trumpPlayed {
		return PlayedCard{}, false, nil
	}

	// 2. Get the best card
	var winningCard PlayedCard
	for _, c := range cards {
		card := c.Card.Card()
		if (card.Suit == suit || card.Suit == Wild) && (winningCard.Card == "" || card.Value > winningCard.Card.Card().Value) {
			winningCard = c
		}
	}
	return winningCard, true, nil
}

func findWinningCardsForRound(round Round) ([]PlayedCard, error) {
	if round.Suit == "" {
		return nil, errors.New("suit not set")
	}
	winningCards := make([]PlayedCard, 5)
	if len(round.CompletedHands) != 5 {
		return nil, errors.New("round not complete")
	}

	for i, hand := range round.CompletedHands {
		winningCard, err := findWinningCard(hand, round.Suit)
		if err != nil {
			return nil, err
		}
		winningCards[i] = winningCard
	}

	return winningCards, nil
}

func findPlayer(playerID string, players []Player) (Player, error) {
	for _, player := range players {
		if player.ID == playerID {
			return player, nil
		}
	}
	return Player{}, errors.New("player not found")
}

func getTeamID(playerID string, players []Player) (string, error) {
	for _, player := range players {
		if player.ID == playerID {
			return player.TeamID, nil
		}
	}
	return "", errors.New("player not found")
}

func checkForJink(winningCards []PlayedCard, players []Player, goerId string) (bool, error) {
	goer, err := findPlayer(goerId, players)
	if err != nil {
		return false, err
	}
	if goer.Call != Jink {
		return false, nil
	}

	if len(winningCards) != 5 {
		return false, errors.New("invalid number of winning cards")
	}
	// Jink is only valid if there are more than 2 players
	if len(players) < 3 {
		return false, nil
	}

	for _, winningCard := range winningCards {
		teamID, err := getTeamID(winningCard.PlayerID, players)
		if err != nil {
			return false, err
		}
		if teamID != goer.TeamID {
			return false, nil
		}
	}
	return true, nil
}

func findWinningTeam(players []Player, round Round) (string, error) {
	// 1. If only one team >= 110 -> they are the winner
	winningTeams := getTeamsOver110(players)
	if len(winningTeams) == 1 {
		for teamID := range winningTeams {
			return teamID, nil
		}
	}

	// 2. If more than one team >= 110 but one is the goer -> the goer is the winning team
	goerTeamID := ""
	for _, player := range players {
		if player.ID == round.GoerID {
			goerTeamID = player.TeamID
			break
		}
	}
	if winningTeams[goerTeamID] {
		return goerTeamID, nil
	}

	// 3. Else first team >= 110 is the winner
	return findFirstTeamToPass110(players, round)
}

func getTeamsOver110(players []Player) map[string]bool {
	teamsOver110 := make(map[string]bool)
	for _, player := range players {
		if player.Score >= 110 {
			teamsOver110[player.TeamID] = true
		}
	}
	return teamsOver110
}

func findFirstTeamToPass110(players []Player, round Round) (string, error) {
	winningCards, err := findWinningCardsForRound(round)
	if err != nil {
		return "", err
	}

	bestTrump, trumpPlayed, err := findBestTrump(winningCards, round.Suit)
	if err != nil {
		return "", err
	}

	// Go backwards through the hands and find the first team to pass 110
	for i := len(winningCards) - 1; i >= 0; i-- {
		card := winningCards[i]
		if err != nil {
			return "", err
		}
		teamID, err := getTeamID(card.PlayerID, players)
		if err != nil {
			return "", err
		}

		for j, p := range players {
			if p.TeamID == teamID {
				if trumpPlayed && bestTrump.Card == card.Card {
					players[j].Score -= 10
				} else {
					players[j].Score -= 5
				}
			}
		}
		winningTeams := getTeamsOver110(players)
		if len(winningTeams) == 1 {
			for t := range winningTeams {
				return t, nil
			}
		}
	}

	return "", errors.New("no winning team found")
}

func calculateScores(winningCards []PlayedCard, players []Player, suit Suit) (map[string]int, error) {
	// 1. Get the best card
	bestCard, trumpPlayed, err := findBestTrump(winningCards, suit)
	if err != nil {
		return nil, err
	}

	// Calculate the scores per player
	scoresPlayer := make(map[string]int)
	for _, winningCard := range winningCards {
		if err != nil {
			return nil, err
		}
		for _, p := range players {
			if p.ID == winningCard.PlayerID {
				if trumpPlayed && winningCard.Card == bestCard.Card {
					scoresPlayer[p.ID] += 10
				} else {
					scoresPlayer[p.ID] += 5
				}
			}
		}
	}

	// Aggregate the scores per team
	scoresTeam := make(map[string]int)
	for playerID, score := range scoresPlayer {
		teamID, err := getTeamID(playerID, players)
		if err != nil {
			return nil, err
		}
		scoresTeam[teamID] += score
	}
	return scoresTeam, nil
}
