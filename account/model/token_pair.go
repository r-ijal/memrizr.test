package model

// TokenPair defines domain model and its json and db representations
type TokenPair struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}
