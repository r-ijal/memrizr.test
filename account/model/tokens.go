package model

import "github.com/google/uuid"

// TokenPair defines domain model and its json and db representations
// type TokenPair struct {
// 	IDToken      string `json:"idToken"`
// 	RefreshToken string `json:"refreshToken"`
// }

type TokenPair struct {
	IDToken      IDToken
	RefreshToken RefreshToken
}

type IDToken struct {
	SS string `json:"idToken"`
}

// stores token properties that are accessed in multiple app layers
type RefreshToken struct {
	SS  string    `json:"refreshToken"`
	ID  uuid.UUID `json:"-"`
	UID uuid.UUID `json:"-"`
}
