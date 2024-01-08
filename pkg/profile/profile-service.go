package profile

import (
	"cards-110-api/pkg/db"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type ServiceI interface {
	Get(ctx context.Context, id string) (Profile, bool, error)
	Save(ctx context.Context, p Profile) error
}
type Service struct {
	Col db.Collection[Profile]
}

func (s *Service) Get(ctx context.Context, id string) (Profile, bool, error) {
	return s.Col.FindOne(ctx, bson.M{"_id": id})
}

func (s *Service) Save(ctx context.Context, p Profile) error {
	return s.Col.Upsert(ctx, p, p.ID)
}
