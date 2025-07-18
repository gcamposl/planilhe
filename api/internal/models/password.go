package models

// Password represents the password change request format
type Password struct {
	NewPassword    string `json:"new"`
	ActualPassword string `json:"actual"`
}
