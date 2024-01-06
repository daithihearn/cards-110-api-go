package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type ServiceI interface {
	GetUser(ctx context.Context, id string) (User, bool, error)
	UpdateUser(ctx context.Context, id string, req UpdateProfileRequest) (User, error)
}
type Service struct {
	Col CollectionI
}

func (s *Service) GetUser(ctx context.Context, id string) (User, bool, error) {

	filter := bson.M{
		"_id": id,
	}

	u, h, err := s.Col.FindOne(ctx, filter)

	if err != nil {
		return User{}, h, err
	}

	return u, h, nil
}

func (s *Service) UpdateUser(ctx context.Context, id string, req UpdateProfileRequest) (User, error) {
	filter := bson.M{
		"_id": id,
	}

	// Get the current user profile
	u, h, err := s.Col.FindOne(ctx, filter)

	// If the picture is locked and the force flag isn't set, don't update the picture
	if h && u.PictureLocked && !req.ForceUpdate {
		req.Picture = u.Picture
	}

	update := bson.M{
		"$set": bson.M{
			"name":       req.Name,
			"picture":    req.Picture,
			"lastAccess": time.Now(),
		},
	}

	u, err = s.Col.FindOneAndUpdate(ctx, filter, update)

	if err != nil {
		return User{}, err
	}

	return u, nil
}
