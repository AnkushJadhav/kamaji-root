package store

import (
	"time"
)

// UserCollection is the collection to store users
const UserCollection = "users"

// UserModel represents a user
type UserModel struct {
	Email     string
	Username  string
	Password  string
	RoleID    int
}

// GetID returns the id of the user document
func (um UserModel) GetID() string {
	return ""
}

// GetTimestamp returns the id of the user document
func (um UserModel) GetTimestamp() time.Time {
	return time.Now()
}
