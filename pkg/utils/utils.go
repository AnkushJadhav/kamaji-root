package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GenerateUUID generates a RFC compliant unique ID
func GenerateUUID() string {
	return uuid.New().String()
}

// GenerateBcryptHash generates the bcrypt hash of input with a cost of 10
func GenerateBcryptHash(input []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(input, bcrypt.DefaultCost)
}

// IsHashValid checks whether hash is calid for plain
func IsHashValid(hash, plain []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hash, plain); err != nil {
		return false
	}

	return true
}
