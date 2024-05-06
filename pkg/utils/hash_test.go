package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomSaltSizeErr(t *testing.T) {
	os.Setenv("SALT_SIZE", "R")
	assert.Panics(t, func() { generateRandomSalt() })
}

func TestGenerateRandomSalt(t *testing.T) {

}

func TestHashPassword(t *testing.T) {
	password := "password"
	salt := []byte("salt")

	expectedHash := generateExpectedHash(password, salt)

	hash := HashPassword(password, salt)

	assert.Equal(t, expectedHash, hash, "Hashed password does not match expected hash")
}

func TestDoPasswordsMatch(t *testing.T) {
	password := "password"
	salt := []byte("salt")
	hash := HashPassword(password, salt)

	// Testing correct password
	assert.True(t, DoPasswordsMatch(hash, password, salt), "Passwords should match")

	// Testing incorrect password
	assert.False(t, DoPasswordsMatch(hash, "incorrectPassword", salt), "Passwords should not match")
}

func generateExpectedHash(password string, salt []byte) string {
	passwordBytes := []byte(password)
	sha512Hasher := sha512.New()
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	hashedPasswordBytes := sha512Hasher.Sum(nil)
	hashedPasswordHex := hex.EncodeToString(hashedPasswordBytes)
	return hashedPasswordHex
}
