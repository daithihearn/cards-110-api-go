package settings

import (
	"cards-110-api/pkg/db"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type ServiceI interface {
	Get(ctx context.Context, userId string) (Settings, bool, error)
	Save(ctx context.Context, settings Settings) error
}

type Service struct {
	Col db.Collection[Settings]
}

// Get the settings for a user.
func (s *Service) Get(ctx context.Context, userId string) (Settings, bool, error) {
	return s.Col.FindOne(ctx, bson.M{"_id": userId})
}

// Save the settings for a user.
func (s *Service) Save(ctx context.Context, settings Settings) error {
	return s.Col.Upsert(ctx, settings, settings.ID)
}
