package deck

import (
	"cards-110-api/pkg/db"
	"context"
	"errors"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson"
)

type ServiceI interface {
	Shuffle(ctx context.Context, gameId string) error
	NextCard(ctx context.Context, gameId string) (Card, error)
	Save(ctx context.Context, deck Deck) error
	Get(ctx context.Context, gameId string) (Deck, bool, error)
}
type Service struct {
	Col db.CollectionI[Deck]
}

// Shuffle the deck for a game.
func (ds *Service) Shuffle(ctx context.Context, gameId string) error {
	// Get the deck
	deck, found, err := ds.Get(ctx, gameId)

	if err != nil {
		return err
	}
	if !found {
		return errors.New("deck not found")
	}

	shuffledCards := shuffleCards(deck.Cards)

	var deckCards []Card
	for _, card := range shuffledCards {
		if card != EmptyCard {
			deckCards = append(deckCards, card)
		}
	}

	deck.Cards = deckCards

	return ds.Save(ctx, deck)
}

// NextCard Get the next card.
func (ds *Service) NextCard(ctx context.Context, gameId string) (Card, error) {
	deck, found, err := ds.Get(ctx, gameId)
	if err != nil {
		return Card{}, err
	}
	if !found {
		return Card{}, errors.New("deck not found")
	}
	if len(deck.Cards) == 0 {
		return Card{}, errors.New("no cards left")
	}

	card := deck.Cards[len(deck.Cards)-1] // Simulate popping from the stack
	deck.Cards = deck.Cards[:len(deck.Cards)-1]

	if err := ds.Save(ctx, deck); err != nil {
		return Card{}, err
	}
	return card, nil
}

func (ds *Service) Save(ctx context.Context, deck Deck) error {
	return ds.Col.Upsert(ctx, deck, deck.ID)
}

func (ds *Service) Get(ctx context.Context, gameId string) (Deck, bool, error) {
	return ds.Col.FindOne(ctx, bson.M{"_id": gameId})
}

func shuffleCards(cards []Card) []Card {
	shuffled := make([]Card, len(cards))
	perm := rand.Perm(len(cards))

	for i, v := range perm {
		shuffled[v] = cards[i]
	}

	return shuffled
}
