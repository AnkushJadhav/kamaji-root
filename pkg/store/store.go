package store

import (
	"context"

	"github.com/AnkushJadhav/kamaji-root/pkg/models"
)

// Driver is an interface that every persistant storage adapter in the application must implement
type Driver interface {
	Connect() error
	Disconnect() error

	GetAllUsers(context.Context) ([]models.User, error)
	GetUserByID(context.Context, string) (models.User, error)
	CreateUser(context.Context, *models.User) error
	DeleteUserByIDs(context.Context, []string) (int, error)
}
