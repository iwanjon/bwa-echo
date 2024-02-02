package user

import (
	"database/sql"
	"time"
)

type User struct {
	// _id primitive.ObjectID
	ID string
	// ID             string `bson:"_id,omitempty"`
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	Token          sql.NullString
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
