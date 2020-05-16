package users

import (
	"github.com/AnkushJadhav/kamaji-root/store"
)

// CreateUserInput is the input required for CreateUser
type CreateUserInput struct {
	Username string
	Email    string
	Password string
	Role     int
}

// CreateUserOutput is the output returned by CreateUser
type CreateUserOutput struct {
	ID       string
	Email    string
	Password string
	Role     int
}

// CreateUser creates a user in the store based on input
func CreateUser(s store.Store, input CreateUserInput) (CreateUserOutput, error) {
	user := store.UserModel{
		Email:    input.Email,
		Password: input.Password,
		RoleID:   input.Role,
		Username: input.Username,
	}

	mUser, err := s.CreateOne(store.UserCollection, user)
	if err != nil {
		return CreateUserOutput{}, err
	}
	tUser := mUser.(store.UserModel)

	return CreateUserOutput{
		Email:    tUser.Email,
		Password: tUser.Password,
		Role:     tUser.RoleID,
	}, nil
}
