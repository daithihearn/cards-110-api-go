package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection[T any] interface {
	FindOne(ctx context.Context, filter bson.M) (T, bool, error)
	Find(ctx context.Context, filter bson.M) ([]T, error)
	InsertMany(ctx context.Context, documents []T) error
	FindOneAndUpdate(ctx context.Context, filter bson.M, update bson.M) (T, error)
	Aggregate(ctx context.Context, pipeline mongo.Pipeline) (*mongo.Cursor, error)
}
