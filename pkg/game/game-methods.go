package game

import (
	"fmt"
	"log"
	"time"
)

func (g *Game) Me(playerID string) (Player, error) {
	for _, p := range g.Players {
		if p.ID == playerID {
			return p, nil
		}
	}
	return Player{}, fmt.Errorf("player not found in game")
}

func (g *Game) GetState(playerID string) (State, error) {
	// Get player
	me, err := g.Me(playerID)
	if err != nil {
		return State{}, err
	}

	// 1. Get max call
	maxCall := Pass
	for _, p := range g.Players {
		if p.Call > maxCall {
			maxCall = p.Call
		}
	}

	// 2. Add dummy if applicable
	iamGoer := g.CurrentRound.GoerID == playerID
	if iamGoer && g.CurrentRound.Status == Called && g.Dummy != nil {
		me.Cards = append(me.Cards, g.Dummy...)
	}

	// 3. Return player's game state
	gameState := State{
		ID:           g.ID,
		Revision:     g.Revision,
		Me:           me,
		IamSpectator: false,
		IsMyGo:       g.CurrentRound.CurrentHand.CurrentPlayerID == me.ID,
		IamGoer:      iamGoer,
		IamDealer:    g.CurrentRound.DealerID == me.ID,
		IamAdmin:     g.AdminID == me.ID,
		Cards:        me.Cards,
		Status:       g.Status,
		Round:        g.CurrentRound,
		MaxCall:      maxCall,
		Players:      g.Players,
	}

	return gameState, nil
}

func (g *Game) EndRound() error {
	// Add the current round to the completed rounds
	g.Completed = append(g.Completed, g.CurrentRound)

	// Get next dealer
	nextDealer, err := nextPlayer(g.Players, g.CurrentRound.DealerID)
	if err != nil {
		return err
	}

	// Create next hand
	nextPlayer, err := nextPlayer(g.Players, nextDealer.ID)
	if err != nil {
		return err
	}
	nextHand := Hand{
		Timestamp:       time.Now(),
		CurrentPlayerID: nextPlayer.ID,
	}

	// Create a new round
	g.CurrentRound = Round{
		Timestamp:   time.Now(),
		Number:      g.CurrentRound.Number + 1,
		DealerID:    nextDealer.ID,
		Status:      Calling,
		CurrentHand: nextHand,
	}

	// Deal the cards
	deck, hands, err := DealCards(ShuffleCards(NewDeck()), len(g.Players))
	if err != nil {
		return err
	}
	var dummy []CardName
	for i, hand := range hands {
		if i >= len(g.Players) {
			dummy = hand
			break
		}
		g.Players[i].Cards = hand
	}
	g.Dummy = dummy
	g.Deck = deck

	// Increment revision
	g.Revision++

	return nil
}

func (g *Game) Call(playerID string, call Call) error {
	// Check the game is active
	if g.Status != Active {
		return fmt.Errorf("game not active")
	}

	// Check current round is calling
	if g.CurrentRound.Status != Calling {
		return fmt.Errorf("round not calling")
	}

	// Check the player is the current player
	if g.CurrentRound.CurrentHand.CurrentPlayerID != playerID {
		return fmt.Errorf("not current player")
	}

	// If they are in the bunker (score < -30) they can only pass
	state, err := g.GetState(playerID)
	if err != nil {
		return err
	}
	if state.Me.Score < -30 && call != Pass {
		return fmt.Errorf("player in bunker")
	}

	// Check the call is valid i.e. > all previous calls or a pass
	// The dealer can take a call of greater than 10
	if call != Pass {
		callForComparison := call
		if state.IamDealer {
			callForComparison++
		}
		for _, p := range g.Players {
			if p.Call >= callForComparison {
				return fmt.Errorf("invalid call")
			}
		}
	}

	// Validate 10 call
	if call == Ten {
		if len(g.Players) != 6 {
			return fmt.Errorf("can only call 10 in doubles")
		}
	}

	// Set the player's call
	for i, p := range g.Players {
		if p.ID == playerID {
			g.Players[i].Call = call
			break
		}
	}

	// Set next player/round status
	if call == Jink {
		log.Printf("Jink called by %s", playerID)
		if state.IamDealer {
			// If the dealer calls Jink, calling is complete
			g.CurrentRound.Status = Called
			g.CurrentRound.GoerID = playerID
			g.CurrentRound.CurrentHand.CurrentPlayerID = playerID
		} else {
			// If any other player calls Jink, jump to the dealer
			g.CurrentRound.CurrentHand.CurrentPlayerID = g.CurrentRound.DealerID
		}
	} else if state.IamDealer {
		// Get the highest calls
		var topCall Call
		for _, p := range g.Players {
			if p.Call > topCall {
				topCall = p.Call
			}
		}

		if topCall <= Ten {
			log.Printf("No one called. Starting new round...")
			err = g.EndRound()
			return err
		}

		// Get the players who made the top call (the dealer may have taken a call, this will result in more than one player)
		var topCallPlayers []Player
		for _, p := range g.Players {
			if p.Call == topCall {
				topCallPlayers = append(topCallPlayers, p)
			}
		}
		if len(topCallPlayers) == 0 || len(topCallPlayers) > 2 {
			return fmt.Errorf("invalid call state. There are %d top callers of %d", len(topCallPlayers), topCall)
		}
		var takenPlayer Player
		var caller Player
		if len(topCallPlayers) == 2 {
			for _, p := range topCallPlayers {
				if p.ID == g.CurrentRound.DealerID {
					caller = p
				} else {
					takenPlayer = p
				}
			}
		} else {
			caller = topCallPlayers[0]
		}

		if takenPlayer.ID != "" {
			log.Printf("Dealer seeing call by %s", takenPlayer.ID)
			g.CurrentRound.DealerSeeing = true
			g.CurrentRound.CurrentHand.CurrentPlayerID = takenPlayer.ID
		} else {
			log.Printf("Call successful. %s is goer", caller.ID)
			g.CurrentRound.Status = Called
			g.CurrentRound.GoerID = caller.ID
			g.CurrentRound.CurrentHand.CurrentPlayerID = caller.ID
		}

	} else if g.CurrentRound.DealerSeeing {
		log.Printf("%s was taken by the dealer.", playerID)
		if call == Pass {
			log.Printf("%s is letting the dealer go.", playerID)
			g.CurrentRound.Status = Called
			g.CurrentRound.GoerID = g.CurrentRound.DealerID
			g.CurrentRound.CurrentHand.CurrentPlayerID = g.CurrentRound.DealerID
		} else {
			log.Printf("%s has raised the call.", playerID)
			g.CurrentRound.CurrentHand.CurrentPlayerID = g.CurrentRound.DealerID
			g.CurrentRound.DealerSeeing = false
		}
	} else {
		log.Printf("Calling not complete. Next player...")
		nextPlayer, err := nextPlayer(g.Players, playerID)
		if err != nil {
			return err
		}
		g.CurrentRound.CurrentHand.CurrentPlayerID = nextPlayer.ID
	}

	// Increment revision
	g.Revision++

	return nil
}

// MinKeep returns the minimum number of cards that must be kept by a player
func (g *Game) MinKeep() (int, error) {
	switch len(g.Players) {
	case 2:
		return 0, nil
	case 3:
		return 0, nil
	case 4:
		return 0, nil
	case 5:
		return 1, nil
	case 6:
		return 2, nil
	}
	return 0, fmt.Errorf("invalid number of players")
}

func (g *Game) SelectSuit(playerID string, suit Suit, cards []CardName) error {
	// Verify the at the round is in the called state
	if g.CurrentRound.Status != Called {
		return fmt.Errorf("round must be in the called state to select a suit")
	}

	// Verify that the player is the goer
	if g.CurrentRound.GoerID != playerID {
		return fmt.Errorf("only the goer can select the suit")
	}

	// Verify the number of cards selected is valid (<=5 and >= minKeep)
	minKeep, err := g.MinKeep()
	if err != nil {
		return err
	}
	if len(cards) > 5 || len(cards) < minKeep {
		return fmt.Errorf("invalid number of cards selected")
	}

	// Verify the cards are valid (must be either in the player's hand or the dummy's hand and must be unique
	state, err := g.GetState(playerID)
	if err != nil {
		return err
	}
	if !containsAllUnique(state.Cards, cards) {
		return fmt.Errorf("invalid card selected")
	}

	// Update the round
	g.CurrentRound.Status = Buying
	g.CurrentRound.Suit = suit

	// Set my cards
	for i, p := range g.Players {
		if p.ID == playerID {
			g.Players[i].Cards = cards
			break
		}
	}

	// Remove the dummy player
	g.Dummy = nil

	// Set the next player
	np, err := nextPlayer(g.Players, playerID)
	if err != nil {
		return err
	}
	g.CurrentRound.CurrentHand.CurrentPlayerID = np.ID

	// Increment revision
	g.Revision++

	return nil
}

func (g *Game) Buy(id string, cards []CardName) error {
	// Verify the at the round is in the buying state
	if g.CurrentRound.Status != Buying {
		return fmt.Errorf("round must be in the buying state to buy cards")
	}

	// Verify that is the player's go
	if g.CurrentRound.CurrentHand.CurrentPlayerID != id {
		return fmt.Errorf("only the current player can buy cards")
	}

	// Verify the number of cards selected is valid (<=5 and >= minKeep)
	minKeep, err := g.MinKeep()
	if err != nil {
		return err
	}
	if len(cards) > 5 || len(cards) < minKeep {
		return fmt.Errorf("invalid number of cards selected")
	}

	// Verify the cards are valid (must be either in the player's hand or the dummy's hand and must be unique
	state, err := g.GetState(id)
	if err != nil {
		return err
	}
	if !containsAllUnique(state.Cards, cards) {
		return fmt.Errorf("invalid card selected")
	}

	// Get cards from the deck so the player has 5 cards
	deck, cards, err := BuyCards(g.Deck, cards)
	if err != nil {
		return err
	}

	g.Deck = deck

	// Set my cards
	for i, p := range g.Players {
		if p.ID == id {
			g.Players[i].Cards = cards
			break
		}
	}

	// If the current player is the dealer update the round status to playing
	if g.CurrentRound.DealerID == id {
		g.CurrentRound.Status = Playing

		// Set the next player when the dealer buys
		np, err := nextPlayer(g.Players, g.CurrentRound.GoerID)
		if err != nil {
			return err
		}
		g.CurrentRound.CurrentHand.CurrentPlayerID = np.ID
	} else {
		// Set the next player
		np, err := nextPlayer(g.Players, id)
		if err != nil {
			return err
		}
		g.CurrentRound.CurrentHand.CurrentPlayerID = np.ID
	}

	// Increment revision
	g.Revision++

	return nil
}

func (g *Game) Play(id string, card CardName) error {
	// Verify the at the round is in the playing state
	if g.CurrentRound.Status != Playing {
		return fmt.Errorf("round must be in the playing state to play a card")
	}

	// Verify that is the player's go
	if g.CurrentRound.CurrentHand.CurrentPlayerID != id {
		return fmt.Errorf("only the current player can play a card")
	}

	// Verify the card is valid
	state, err := g.GetState(id)
	if err != nil {
		return err
	}
	if !contains(state.Cards, card) {
		return fmt.Errorf("invalid card selected")
	}

	// Check that they are following suit
	if g.CurrentRound.CurrentHand.LeadOut == "" {
		// I must be leading out
		g.CurrentRound.CurrentHand.LeadOut = card
	} else {
		if !isFollowing(card, state.Cards, g.CurrentRound.CurrentHand, g.CurrentRound.Suit) {
			return fmt.Errorf("must follow suit")
		}
	}

	// Remove the card from the player's hand
	var cards []CardName
	for _, c := range state.Cards {
		if c != card {
			cards = append(cards, c)
		}
	}
	for i, p := range g.Players {
		if p.ID == id {
			g.Players[i].Cards = cards
			break
		}
	}

	// Add the card to the played cards
	pc := PlayedCard{
		PlayerID: id,
		Card:     card,
	}
	g.CurrentRound.CurrentHand.PlayedCards = append(g.CurrentRound.CurrentHand.PlayedCards, pc)

	// Check if the hand is complete
	if len(g.CurrentRound.CurrentHand.PlayedCards) < len(g.Players) {
		// Set the next player
		np, err := nextPlayer(g.Players, id)
		if err != nil {
			return err
		}
		g.CurrentRound.CurrentHand.CurrentPlayerID = np.ID
	} else {

		g.completeHand()

		// Check if the round is complete
		if len(g.CurrentRound.CompletedHands) == 5 {
			err := g.applyScores()
			if err != nil {
				return err
			}

			// Check if the game is complete
			if g.isGameOver() {
				err := g.completeGame()
				if err != nil {
					return err
				}
			} else {
				err := g.completeRound()
				if err != nil {
					return err
				}
			}
		}
	}

	// Increment revision
	g.Revision++

	return nil

}

func (g *Game) completeHand() {
	// 1. Find the winner
	winningCard, err := determineWinner(g.CurrentRound.CurrentHand, g.CurrentRound.Suit)
	if err != nil {
		log.Printf("Error determining winner: %s", err.Error())
		return
	}

	// 2. Add the hand to the completed hands
	g.CurrentRound.CompletedHands = append(g.CurrentRound.CompletedHands, g.CurrentRound.CurrentHand)

	// 3. Set the next hand
	g.CurrentRound.CurrentHand = Hand{
		Timestamp:       time.Now(),
		CurrentPlayerID: winningCard.PlayerID,
		PlayedCards:     make([]PlayedCard, 0),
	}
}

func (g *Game) applyScores() error {
	// 1. Find winning card for each hand
	winningCards, err := findWinningCardsForRound(g.CurrentRound)
	if err != nil {
		return err
	}

	// 2. Check if there was a jink
	jinkHappened, err := checkForJink(winningCards, g.Players, g.CurrentRound.GoerID)
	if err != nil {
		return err
	}
	if jinkHappened {
		teamId, err := getTeamID(winningCards[0].PlayerID, g.Players)
		if err != nil {
			return err
		}
		// The successful team gets 60 points
		for i, p := range g.Players {
			if p.ID == teamId {
				g.Players[i].Score += 60
			}
		}
		return nil
	}

	// 3. Calculate scores
	scores, err := calculateScores(winningCards, g.Players, g.CurrentRound.Suit)

	// 4. If the goer didn't make their contract, set score to minus the contract
	goer, err := findPlayer(g.CurrentRound.GoerID, g.Players)
	if err != nil {
		return err
	}
	call := int(goer.Call)
	if scores[goer.TeamID] < call {
		scores[goer.TeamID] = -call
	}

	// 5. Apply scores
	for teamId, score := range scores {
		for i, p := range g.Players {
			if p.TeamID == teamId {
				g.Players[i].Score += score
			}
		}
	}

	return nil
}

func (g *Game) isGameOver() bool {
	for _, p := range g.Players {
		if p.Score >= 110 {
			return true
		}
	}
	return false
}

func (g *Game) completeRound() error {
	// 1. Add round to completed
	g.Completed = append(g.Completed, g.CurrentRound)

	// 2. Create next round
	nextDealer, err := nextPlayer(g.Players, g.CurrentRound.DealerID)
	if err != nil {
		return err
	}
	nextPlayer, err := nextPlayer(g.Players, nextDealer.ID)
	if err != nil {
		return err
	}
	nextHand := Hand{
		Timestamp:       time.Now(),
		CurrentPlayerID: nextPlayer.ID,
	}
	g.CurrentRound = Round{
		Timestamp:   time.Now(),
		Number:      g.CurrentRound.Number + 1,
		DealerID:    nextDealer.ID,
		Status:      Calling,
		CurrentHand: nextHand,
	}

	// 3. Clear cards
	for i := range g.Players {
		g.Players[i].Cards = make([]CardName, 0)
	}

	// 4. Deal cards
	deck, hands, err := DealCards(ShuffleCards(NewDeck()), len(g.Players))
	if err != nil {
		return err
	}
	var dummy []CardName
	for i, hand := range hands {
		if i >= len(g.Players) {
			dummy = hand
			break
		}
		g.Players[i].Cards = hand
	}
	g.Dummy = dummy
	g.Deck = deck

	// 5. Increment revision
	g.Revision++

	return nil
}

func (g *Game) completeGame() error {
	winningTeam, err := findWinningTeam(g.Players, g.CurrentRound)
	if err != nil {
		return err
	}

	for i, p := range g.Players {
		if p.TeamID == winningTeam {
			g.Players[i].Winner = true
		}
	}
	g.Status = Completed

	return nil
}
