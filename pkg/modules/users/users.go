package users

import (
	"context"

	"github.com/AnkushJadhav/kamaji-root/pkg/models"
	"github.com/AnkushJadhav/kamaji-root/pkg/store"
)

// GetAllUsers gets all users in the system
func GetAllUsers(ctx context.Context, store store.Driver) ([]models.User, error) {
	return store.GetAllUsers(ctx)
}
