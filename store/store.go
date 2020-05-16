package store

import (
	"github.com/AnkushJadhav/kamaji-root/logger"
)

// Result represents a single document returned by a storage driver
type Result interface{}

// Store is the interface that any storage driver in the application must implement
type Store interface {
	Connect() error
	Disconnect() error

	GetAll(string) ([]Result, error)
	GetMany(string, interface{}) ([]Result, error)
	GetOne(string, interface{}) (Result, error)
	PutOne(string, interface{}) error
	PutMany(string, []interface{}) error
}

// Driver is the type of storage driver
type Driver uint8

const (
	// Mongo is the driver identitfier for MongoDB storage driver
	Mongo Driver = iota
)

// NewStorageDriver provides a new driver for storage type driver
func NewStorageDriver(driver Driver) (*Store, error) {
	switch driver {
	case Mongo:
		logger.Debugln("WIP!")
	}

	return nil, nil
}
