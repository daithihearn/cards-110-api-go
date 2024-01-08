package profile

import "time"

type Profile struct {
	ID            string    `bson:"_id,omitempty" json:"-"`
	Name          string    `bson:"name" json:"name"`
	Picture       string    `bson:"picture" json:"picture"`
	PictureLocked bool      `bson:"pictureLocked" json:"pictureLocked"`
	LastAccess    time.Time `bson:"lastAccess" json:"lastAccess"`
}

type UpdateProfileRequest struct {
	Name        string `json:"name" binding:"required"`
	Picture     string `json:"picture,omitempty"`
	ForceUpdate bool   `json:"forceUpdate"`
}
