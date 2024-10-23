package model

import (
	"github.com/google/uuid"
)

// User defines domain model and its json and db representations
type User struct {
	UID      uuid.UUID `db:"uid" json:"uid"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"Password" json:"-"`
	Name     string    `db:"name" json:"name"`
	ImageURL string    `db:"imageURL" json:"imageURL"`
	Website  string    `db:"website" json:"website"`
}
