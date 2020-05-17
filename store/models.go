package store

import (
	"time"
)

// UserCollection is the collection to store users
const UserCollection = "users"

// UserModel represents a user
type UserModel struct {
	ID       string
	TS       time.Time
	Email    string
	Username string
	Password string
	RoleID   int
}

// GetID returns the id of the user document
func (um UserModel) GetID() string {
	return um.ID
}

// SetID sets the id of the user document
func (um UserModel) SetID(id string) Model {
	um.ID = id
	return um
}

// GetSystemTimestamp returns the timestamp of the user document
func (um UserModel) GetSystemTimestamp() time.Time {
	return um.TS
}

// SetSystemTimestamp sets the timestap of the user document
func (um UserModel) SetSystemTimestamp() Model {
	um.TS = time.Now()
	return um
}