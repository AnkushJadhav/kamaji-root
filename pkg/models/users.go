package models

import (
	"time"
)

const (
	// RoleAdmin is the admin role for a user
	RoleAdmin = iota
	// RoleNode is the node role for a user
	RoleNode
)

// User represents a user in the system
type User struct {
	ID       string    `bson:"id"`
	TS       time.Time `bson:"ts"`
	Email    string    `bson:"email"`
	Username string    `bson:"username,omitempty"`
	Password string    `bson:"password,omitempty"`
	RoleID   int       `bson:"roleId"`
	Nodes    []Node    `bson:"nodes"`
}
