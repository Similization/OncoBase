package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"os"
	"strconv"
)

// Salt is the salt value used for password hashing.
var Salt = []byte(os.Getenv("SALT"))

// generateRandomSalt generates a random salt of the specified size.
func generateRandomSalt() []byte {
	saltSize, err := strconv.Atoi(os.Getenv("SALT_SIZE"))
	if err != nil {
		panic(err.Error())
	}
	salt := make([]byte, saltSize)
	_, err = rand.Read(salt[:])
	if err != nil {
		panic(err.Error())
	}
	return salt
}

// HashPassword generates a hash for the given password using SHA-512 algorithm and the provided salt.
func HashPassword(password string, salt []byte) string {
	passwordBytes := []byte(password)
	sha512Hasher := sha512.New()
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	hashedPasswordBytes := sha512Hasher.Sum(nil)
	hashedPasswordHex := hex.EncodeToString(hashedPasswordBytes)
	return hashedPasswordHex
}

// DoPasswordsMatch checks if the provided hashed password matches the current password using the given salt.
func DoPasswordsMatch(hashedPassword, currPassword string, salt []byte) bool {
	currPasswordHash := HashPassword(currPassword, salt)
	return hashedPassword == currPasswordHash
}
