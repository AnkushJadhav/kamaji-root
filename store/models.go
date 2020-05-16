package store

import (
	"time"
)

// UserCollection is the collection to store users
const UserCollection = "users"

// UserModel represents a user
type UserModel struct {
	id        string
	Email     string
	Username  string
	Password  string
	RoleID    int
	createdAt time.Time
}

// GetID returns the id of the user document
func (um UserModel) GetID() string {
	return um.id
}

// GetTimestamp returns the id of the user document
func (um UserModel) GetTimestamp() time.Time {
	return um.createdAt
}
