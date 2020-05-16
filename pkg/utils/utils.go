package utils

import (
	"github.com/google/uuid"
)

// GenerateUUID generates a RFC compliant unique ID
func GenerateUUID() string {
	return uuid.New().String()
}