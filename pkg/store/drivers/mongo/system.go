package mongo

import (
	"context"

	"github.com/AnkushJadhav/kamaji-root/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

// InitSystemConfig initialises the persistant system config with zero values
func (mdb *Driver) InitSystemConfig(ctx context.Context, id string) error {
	conf := models.SystemConfig{
		ID: id,
	}
	_, err := mdb.dbs[dbPrimary].Collection(colSystem).InsertOne(ctx, conf)
	if err != nil {
		return err
	}

	return nil
}

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
	doc := mdb.dbs[dbPrimary].Collection(colSystem).FindOneAndUpdate(ctx, bson.D{}, bson.D{{"$set", bson.D{{colSystemAtrBootupstate, state}}}})
	if doc.Err() != nil {
		return doc.Err()
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
	doc := mdb.dbs[dbPrimary].Collection(colSystem).FindOneAndUpdate(ctx, bson.D{}, bson.D{{"$set", bson.D{{colSystemAtrRoottoken, token}}}})
	if doc.Err() != nil {
		return doc.Err()
	}
	return nil
}

// GetJWTToken gets the bootup state of the system
func (mdb *Driver) GetJWTToken(ctx context.Context) (string, error) {
	doc := mdb.dbs[dbPrimary].Collection(colSystem).FindOne(ctx, bson.D{})
	if doc.Err() != nil {
		return "", doc.Err()
	}

	result := &models.SystemConfig{}
	if err := doc.Decode(result); err != nil {
		return "", err
	}

	return result.JWTSecret, nil
}

// SetJWTToken updates the bootup state of the system
func (mdb *Driver) SetJWTToken(ctx context.Context, token string) error {
	doc := mdb.dbs[dbPrimary].Collection(colSystem).FindOneAndUpdate(ctx, bson.D{}, bson.D{{"$set", bson.D{{colSystemAtrJWTToken, token}}}})
	if doc.Err() != nil {
		return doc.Err()
	}
	return nil
}
