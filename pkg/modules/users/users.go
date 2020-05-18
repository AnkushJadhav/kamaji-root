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

// DeleteUser deletes a user based on id
func DeleteUser(ctx context.Context, store store.Driver, id string) error {
	ids := make([]string, 1)
	ids[0] = id

	_, err := store.DeleteUserByIDs(ctx, ids)
	if err != nil {
		return err
	}

	return nil
}
