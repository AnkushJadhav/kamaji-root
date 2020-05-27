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
	ID       string    `bson:"id" json:"id"`
	TS       time.Time `bson:"ts" json:"ts"`
	Email    string    `bson:"email" json:"email"`
	Username string    `bson:"username,omitempty" json:"username"`
	Password string    `bson:"password,omitempty" json:"-"`
	RoleID   int       `bson:"roleId" json:"role"`
	Nodes    []Node    `bson:"nodes" json:"nodes"`
}
