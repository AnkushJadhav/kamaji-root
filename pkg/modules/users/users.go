package users

import (
	"context"
	"time"

	"github.com/AnkushJadhav/kamaji-root/pkg/utils"

	"github.com/AnkushJadhav/kamaji-root/pkg/models"
	"github.com/AnkushJadhav/kamaji-root/pkg/store"
)

// GetAllUsers gets all users in the system
func GetAllUsers(ctx context.Context, store store.Driver) ([]models.User, error) {
	return store.GetAllUsers(ctx)
}

// CreateUser creates a new user in the system
func CreateUser(ctx context.Context, store store.Driver, email string, roleID int) (models.User, error) {
	user := models.User{
		Email:  email,
		RoleID: roleID,
	}
	user.ID = utils.GenerateUUID()
	user.TS = time.Now()

	if err := store.CreateUser(ctx, &user); err != nil {
		return models.User{}, err
	}

	return user, nil
}
