package users

import (
	"github.com/AnkushJadhav/kamaji-root/store"
)

// GetAllUsersOutput is the output returned by GetAllUsers
type GetAllUsersOutput struct {
	ID    string
	Email string
	Role  int
}

// GetAllUsers gets all the users
func GetAllUsers(s store.Store) ([]GetAllUsersOutput, error) {
	l, mUsers, err := s.GetAll(store.UserCollection)
	if err != nil {
		return []GetAllUsersOutput{}, err
	}

	tUsers := make([]GetAllUsersOutput, l)
	for _, u := range mUsers {
		um := u.(store.UserModel)
		tu := GetAllUsersOutput{
			ID:    um.GetID(),
			Email: um.Email,
			Role:  um.RoleID,
		}
		tUsers = append(tUsers, tu)
	}

	return tUsers, nil
}

// CreateUserInput is the input required for CreateUser
type CreateUserInput struct {
	Email string
	Role  int
}

// CreateUserOutput is the output returned by CreateUser
type CreateUserOutput struct {
	ID    string
	Email string
	Role  int
}

// CreateUser creates a user in the store based on input
func CreateUser(s store.Store, input CreateUserInput) (CreateUserOutput, error) {
	user := store.UserModel{
		Email:  input.Email,
		RoleID: input.Role,
	}

	mUser, err := s.CreateOne(store.UserCollection, user)
	if err != nil {
		return CreateUserOutput{}, err
	}
	tUser := mUser.(store.UserModel)

	return CreateUserOutput{
		ID:    tUser.GetID(),
		Email: tUser.Email,
		Role:  tUser.RoleID,
	}, nil
}
