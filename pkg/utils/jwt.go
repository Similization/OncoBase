package utils

import (
	"errors"
	"fmt"
	"med/pkg/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// jwtKey is the secret key used to sign JWT tokens.
var jwtKey = []byte(os.Getenv("JWT_KEY"))

// tokenTTL specifies the time-to-live (TTL) duration for JWT tokens.
const tokenTTL = 30 * time.Minute

// tokenClaims represents the custom claims to be included in JWT tokens.
type tokenClaims struct {
	UserId               int    `json:"id"`   // User ID associated with the token
	UserRole             string `json:"role"` // User role associated with the token
	jwt.RegisteredClaims        // Standard JWT claims
}

// GenerateJWT generates a JWT token for the provided user.
func GenerateJWT(user model.User) (string, error) {
	// Calculate token expiration time
	expirationTime := time.Now().Add(tokenTTL)

	// Create custom claims
	claims := &tokenClaims{
		UserId:   user.Id,
		UserRole: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create JWT token with custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token and return the resulting string
	return token.SignedString(jwtKey)
}

// ParseToken parses and validates the provided JWT access token.
func ParseToken(accessToken string) (*tokenClaims, error) {
	// Parse and validate the JWT token
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Validate the token signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		// Return the secret key for token validation
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Extract custom claims from the token
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}
