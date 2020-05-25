package system

import (
	"context"

	"github.com/AnkushJadhav/kamaji-root/pkg/models"
	"github.com/AnkushJadhav/kamaji-root/pkg/store"
)

// GetBootupState gets the bootup state of the kamaji-root server
func GetBootupState(ctx context.Context, store store.Driver) (models.Bootupstate, error) {
	state, err := store.GetBootupState(ctx)
	if err != nil {
		return -1, err
	}

	return state, nil
}

// IsRootTokenValid checks whether token is the same as that entered in config file at bootup
func IsRootTokenValid(ctx context.Context, store store.Driver, token string) (bool, error) {
	configtoken, err := store.GetRootToken(ctx)
	if err != nil {
		return false, err
	}

	return (token == configtoken), nil
}
