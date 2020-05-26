package store

import (
	"context"

	"github.com/AnkushJadhav/kamaji-root/pkg/models"
)

// Driver is an interface that every persistant storage adapter in the application must implement
type Driver interface {
	Connect() error
	Disconnect() error

	InitSystemConfig(context.Context, string) error
	GetBootupState(context.Context) (models.Bootupstate, error)
	SetBootupState(context.Context, models.Bootupstate) error
	GetRootToken(context.Context) (string, error)
	SetRootToken(context.Context, string) error
	GetJWTToken(context.Context) (string, error)
	SetJWTToken(context.Context, string) error

	GetAllUsers(context.Context) ([]*models.User, error)
	GetUserByEmail(context.Context, string) (*models.User, error)
	GetUserByID(context.Context, string) (*models.User, error)
	CreateUser(context.Context, *models.User) error
	DeleteUserByIDs(context.Context, []string) (int, error)
	UpdateUsersByIDs(context.Context, []string, models.User) (int, error)

	AddNodeToUser(context.Context, string, models.Node) (int, error)
}
