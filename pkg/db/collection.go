package db

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection[T any] struct {
	Col *mongo.Collection
}

func (c *Collection[T]) FindOne(ctx context.Context, filter bson.M) (T, bool, error) {
	var t T
	err := c.Col.FindOne(ctx, filter).Decode(&t)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return t, false, nil
		}
		return t, false, err
	}

	return t, true, nil
}

func (c *Collection[T]) Find(ctx context.Context, filter bson.M) ([]T, error) {
	var ts []T
	cur, err := c.Col.Find(ctx, filter)
	if err != nil {
		return []T{}, err
	}

	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cur, ctx)

	for cur.Next(ctx) {
		var t T
		err := cur.Decode(&t)
		if err != nil {
			log.Println("Error decoding result:", err)
			continue
		}
		ts = append(ts, t)
	}

	if err := cur.Err(); err != nil {
		return []T{}, err
	}

	return ts, nil
}

func (c *Collection[T]) FindOneAndUpdate(ctx context.Context, filter bson.M, update bson.M) (T, error) {
	var t T
	err := c.Col.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&t)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return t, nil
		}
		return t, err
	}

	return t, nil
}

func (c *Collection[T]) FindOneAndReplace(ctx context.Context, filter bson.M, replacement T) (T, error) {
	var t T
	err := c.Col.FindOneAndReplace(ctx, filter, replacement, options.FindOneAndReplace().SetReturnDocument(options.After)).Decode(&t)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return t, nil
		}
		return t, err
	}

	return t, nil
}

func (c *Collection[T]) Upsert(ctx context.Context, t T, id string) error {

	// Create filter
	filter := bson.M{
		"_id": id,
	}

	_, exists, err := c.FindOne(ctx, filter)

	if err != nil {
		return err
	}

	if exists {
		// Update the existing object
		_, err := c.Col.UpdateOne(ctx, filter, bson.M{"$set": t})

		if err != nil {
			return err
		}
	} else {
		// Insert a new object
		_, err := c.Col.InsertOne(ctx, t)

		if err != nil {
			return err
		}
	}

	return nil

}
