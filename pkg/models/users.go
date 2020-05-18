package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID       string    `bson:"id"`
	TS       time.Time `bson:"ts"`
	Email    string    `bson:"email"`
	Username string    `bson:"username,omitempty"`
	Password string    `bson:"password,omitempty"`
	RoleID   int       `bson:"roleid"`
}
