package store

import (
	"time"
)

// Model is an interface that every data model in the application must implement
type Model interface {
	GetID() string
	GetTimestamp() time.Time
}

// Store is the interface that any storage driver in the application must implement
type Store interface {
	Connect() (Store, error)
	//Disconnect() error

	//GetAll(string) (int64, []Model, error)
	//GetMany(string, map[string]interface{}, interface{}) (int64, []Model, error)
	//GetOne(string, map[string]interface{}, interface{}) (int64, Model, error)
	CreateOne(string, Model) (Model, error)
	//CreateMany(string, []Model) (int64, error)
	//UpdateOne(string, Model) (int64, error)
	//UpdateMany(string, []Model) (int64, error)
}
