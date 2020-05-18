package mongo

import (
	"context"

	"github.com/AnkushJadhav/kamaji-root/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllUsers retrieves all the users from MongoDB persistant storage
func (mdb *Driver) GetAllUsers(ctx context.Context) ([]models.User, error) {
	cur, err := mdb.dbs[dbPrimary].Collection(colUsers).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	result := make([]models.User, 0)
	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateUser creates a user in the MongoDB persistan storage
func (mdb *Driver) CreateUser(ctx context.Context, user *models.User) error {
	_, err := mdb.dbs[dbPrimary].Collection(colUsers).InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserByIDs deletes a user from the MongoDB persistant storage based on id
func (mdb *Driver) DeleteUserByIDs(ctx context.Context, ids []string) (int, error) {
	deleted, err := mdb.dbs[dbPrimary].Collection(colUsers).DeleteMany(ctx, bson.M{atrID: bson.M{"$in": ids}})
	if err != nil {
		return -1, err
	}
	return int(deleted.DeletedCount), nil
}