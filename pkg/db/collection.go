package db

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionI[T any] interface {
	FindOne(ctx context.Context, filter bson.M) (T, bool, error)
	Find(ctx context.Context, filter bson.M) ([]T, error)
	FindOneAndUpdate(ctx context.Context, filter bson.M, update bson.M) (T, error)
	FindOneAndReplace(ctx context.Context, filter bson.M, replacement T) (T, error)
	UpdateOne(ctx context.Context, t T, id string) error
	Upsert(ctx context.Context, t T, id string) error
	Aggregate(ctx context.Context, pipeline interface{}) (*mongo.Cursor, error)
}

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
		return nil, err
	}

	defer func() {
		if cerr := cur.Close(ctx); cerr != nil {
			log.Println("Failed to close cursor:", cerr)
		}
	}()

	for cur.Next(ctx) {
		var t T
		if err := cur.Decode(&t); err != nil {
			log.Println("Error decoding result:", err)
			// Decide here whether to continue or return an error
			return nil, err
		}
		ts = append(ts, t)
	}

	if err := cur.Err(); err != nil {
		return nil, err
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

func (c *Collection[T]) UpdateOne(ctx context.Context, t T, id string) error {
	_, err := c.Col.UpdateOne(ctx, bson.M{
		"_id": id,
	}, bson.M{"$set": t})

	return err
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

func (c *Collection[T]) Aggregate(ctx context.Context, pipeline interface{}) (*mongo.Cursor, error) {
	return c.Col.Aggregate(ctx, pipeline)
}
