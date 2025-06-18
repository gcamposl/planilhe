package models

// Password represents the password change request format
type Password struct {
	NewPassword    string `json:"newPassword"`
	ActualPassword string `json:"actual"`
}
