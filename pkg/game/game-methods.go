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

	// 1. Find dummy
	var dummy Player
	if g.CurrentRound.GoerID == playerID {
		for _, p := range g.Players {
			if p.ID == "dummy" {
				dummy = p
				break
			}
		}
	}

	// 2. Get max call
	maxCall := Pass
	for _, p := range g.Players {
		if p.Call > maxCall {
			maxCall = p.Call
		}
	}

	// 3. Add dummy if applicable
	iamGoer := g.CurrentRound.GoerID == playerID
	if iamGoer && g.CurrentRound.Status == Called && dummy.ID != "" {
		me.Cards = append(me.Cards, dummy.Cards...)
	}

	// 4. Return player's game state
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
		Players:      make([]Player, 0),
	}

	for _, p := range g.Players {
		if p.ID != "dummy" {
			gameState.Players = append(gameState.Players, p)
		}
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
	deck := ShuffleCards(NewDeck())
	g.Deck, g.Players = DealCards(deck, g.Players)

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
