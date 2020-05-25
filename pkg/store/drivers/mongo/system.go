package mongo

import (
	"context"

	"github.com/AnkushJadhav/kamaji-root/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetBootupState gets the bootup state of the system
func (mdb *Driver) GetBootupState(ctx context.Context) (models.Bootupstate, error) {
	doc := mdb.dbs[dbPrimary].Collection(colSystem).FindOne(ctx, bson.D{})
	if doc.Err() != nil {
		return -1, doc.Err()
	}

	result := &models.SystemConfig{}
	if err := doc.Decode(result); err != nil {
		return -1, err
	}

	return result.BootupState, nil
}

// SetBootupState updates the bootup state of the system
func (mdb *Driver) SetBootupState(ctx context.Context, state models.Bootupstate) error {
	doc := mdb.dbs[dbPrimary].Collection(colSystem).FindOne(ctx, bson.D{})
	if doc.Err() != nil {
		return doc.Err()
	}

	result := &models.SystemConfig{}
	if err := doc.Decode(result); err != nil {
		return err
	}

	result.BootupState = state
	_, err := mdb.dbs[dbPrimary].Collection(colSystem).UpdateOne(ctx, bson.D{}, result)
	if err != nil {
		return err
	}
	return nil
}

// GetRootToken gets the bootup state of the system
func (mdb *Driver) GetRootToken(ctx context.Context) (string, error) {
	doc := mdb.dbs[dbPrimary].Collection(colSystem).FindOne(ctx, bson.D{})
	if doc.Err() != nil {
		return "", doc.Err()
	}

	result := &models.SystemConfig{}
	if err := doc.Decode(result); err != nil {
		return "", err
	}

	return result.RootToken, nil
}

// SetRootToken updates the bootup state of the system
func (mdb *Driver) SetRootToken(ctx context.Context, token string) error {
	doc := mdb.dbs[dbPrimary].Collection(colSystem).FindOne(ctx, bson.D{})
	if doc.Err() != nil {
		return doc.Err()
	}

	result := &models.SystemConfig{}
	if err := doc.Decode(result); err != nil {
		return err
	}

	result.RootToken = token
	_, err := mdb.dbs[dbPrimary].Collection(colSystem).UpdateOne(ctx, bson.D{}, result)
	if err != nil {
		return err
	}
	return nil
}
