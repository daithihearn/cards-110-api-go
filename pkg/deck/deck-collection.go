package deck

import (
	"cards-110-api/pkg/db"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type CollectionI interface {
	db.Collection[Deck]
}

type Collection struct {
	Col *mongo.Collection
}

func (c Collection) FindOne(ctx context.Context, filter bson.M) (Deck, bool, error) {
	var d Deck
	err := c.Col.FindOne(ctx, filter).Decode(&d)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Deck{}, false, nil
		}
		return Deck{}, false, err
	}

	return d, true, err
}

func (c Collection) Find(ctx context.Context, filter bson.M) ([]Deck, error) {
	var decks []Deck
	cur, err := c.Col.Find(ctx, filter)
	if err != nil {
		return []Deck{}, err
	}

	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cur, ctx)

	for cur.Next(ctx) {
		var deck Deck
		err := cur.Decode(&deck)
		if err != nil {
			log.Println("Error decoding deck:", err)
			continue
		}
		decks = append(decks, deck)
	}

	if err := cur.Err(); err != nil {
		return []Deck{}, err
	}

	return decks, nil
}

func (c Collection) InsertMany(ctx context.Context, documents []Deck) error {
	//TODO implement me
	panic("implement me")
}

func (c Collection) Aggregate(ctx context.Context, pipeline mongo.Pipeline) (*mongo.Cursor, error) {
	//TODO implement me
	panic("implement me")
}

func (c Collection) FindOneAndUpdate(ctx context.Context, filter bson.M, update bson.M) (Deck, error) {
	var d Deck
	err := c.Col.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&d)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Deck{}, nil
		}
		return Deck{}, err
	}

	return d, nil
}
