package store

import (
	"time"
)

// Model is an interface that every data model in the application must implement
type Model interface {
	GetID() string
	SetID(string) Model
	GetSystemTimestamp() time.Time
	SetSystemTimestamp() Model
}

// Store is the interface that any storage driver in the application must implement
type Store interface {
	Connect() (Store, error)
	//Disconnect() error

	GetAll(string) (int, []Model, error)
	//GetMany(string, map[string]interface{}, interface{}) (int, []Model, error)
	//GetOne(string, map[string]interface{}, interface{}) (int, Model, error)
	CreateOne(string, Model) (Model, error)
	//CreateMany(string, []Model) (int, error)
	//UpdateOne(string, Model) (int, error)
	//UpdateMany(string, []Model) (int, error)
}
