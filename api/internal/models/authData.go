package models

// AuthData represents the data returned by the login endpoint
type AuthData struct {
	ID    uint64 `json:"id"`
	Token string `json:"token"`
}
