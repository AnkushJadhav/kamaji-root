package users

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/AnkushJadhav/kamaji-root/pkg/utils"

	"github.com/AnkushJadhav/kamaji-root/pkg/models"
	"github.com/AnkushJadhav/kamaji-root/pkg/store"
)

// PasswordDoesNotMatchPolicy is the error returned when the user password does not match password policy
type PasswordDoesNotMatchPolicy struct{}

func (p *PasswordDoesNotMatchPolicy) Error() string {
	return fmt.Sprintf("password does not match policy")
}

// UserDoesNotExist is the error returned when the requested user does not exist in the system
type UserDoesNotExist struct{}

func (u *UserDoesNotExist) Error() string {
	return fmt.Sprintf("user does not exist")
}

// AuthCredentialMismatch is the error returned whe nthe credentaisl provided and stored do not match
type AuthCredentialMismatch struct{}

func (a *AuthCredentialMismatch) Error() string {
	return fmt.Sprintf("inavlaid password")
}

// GetAllUsers gets all users in the system
func GetAllUsers(ctx context.Context, store store.Driver) ([]*models.User, error) {
	return store.GetAllUsers(ctx)
}

// GetUserByID gets a user based on id
func GetUserByID(ctx context.Context, store store.Driver, id string) (*models.User, error) {
	return store.GetUserByID(ctx, id)
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

// RegisterUser updates a user with their registration information post successful sign in at node
func RegisterUser(ctx context.Context, store store.Driver, id, username, password string) error {
	if !isValid(password) {
		return &PasswordDoesNotMatchPolicy{}
	}

	hashedPwd, err := utils.GenerateBcryptHash([]byte(password))
	if err != nil {
		return err
	}
	toUpdate := models.User{
		Username: username,
		Password: string(hashedPwd),
	}
	_, err = store.UpdateUsersByIDs(ctx, []string{id}, toUpdate)
	if err != nil {
		return err
	}
	return nil
}

// AuthenticateUser verifies whether the the user is valid
func AuthenticateUser(ctx context.Context, store store.Driver, email, password string) error {
	user, err := store.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	if user == nil {
		return &UserDoesNotExist{}
	}

	if !utils.IsHashValid([]byte(password), []byte(user.Password)) {
		return &AuthCredentialMismatch{}
	}

	return nil
}

func isValid(password string) bool {
	policy := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\$%\^&\*])(?=.{8,})`)
	return policy.Match([]byte(password))
}
