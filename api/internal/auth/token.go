package auth

import (
	"api/internal/config"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken create a new token with user permissions
func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}

// Validate token
func ValidateToken(r *http.Request) error {
	return nil
}
