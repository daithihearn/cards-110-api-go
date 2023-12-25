package user

import (
	"cards-110-api/pkg/db"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionI interface {
	db.Collection[User]
}

type Collection struct {
	Col *mongo.Collection
}

func (c Collection) FindOne(ctx context.Context, filter bson.M) (User, bool, error) {
	var u User
	err := c.Col.FindOne(ctx, filter).Decode(&u)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return User{}, false, nil
		}
		return User{}, false, err
	}

	return u, true, err
}

func (c Collection) Find(ctx context.Context, filter bson.M) ([]User, error) {
	//TODO implement me
	panic("implement me")
}

func (c Collection) InsertMany(ctx context.Context, documents []User) error {
	//TODO implement me
	panic("implement me")
}

func (c Collection) Aggregate(ctx context.Context, pipeline mongo.Pipeline) (*mongo.Cursor, error) {
	//TODO implement me
	panic("implement me")
}

func (c Collection) FindOneAndUpdate(ctx context.Context, filter bson.M, update bson.M) (User, error) {
	var u User
	err := c.Col.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&u)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return User{}, nil
		}
		return User{}, err
	}

	return u, nil
}
