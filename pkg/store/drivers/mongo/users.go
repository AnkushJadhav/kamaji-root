package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/AnkushJadhav/kamaji-root/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllUsers retrieves all the users from MongoDB persistant storage
func (mdb *Driver) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	cur, err := mdb.dbs[dbPrimary].Collection(colUsers).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	result := make([]*models.User, 0)
	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetUserByID retrieves a user from the MongoDB persistant storage based on id
func (mdb *Driver) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	doc := mdb.dbs[dbPrimary].Collection(colUsers).FindOne(ctx, bson.D{{colUsersAtrID, id}})
	if doc.Err() != nil {
		if doc.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, doc.Err()
	}

	result := &models.User{}
	if err := doc.Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetUserByEmail finds a user based on the given email
func (mdb *Driver) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	doc := mdb.dbs[dbPrimary].Collection(colUsers).FindOne(ctx, bson.D{{colUsersAtrEmail, email}})
	if doc.Err() != nil {
		if doc.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, doc.Err()
	}

	result := &models.User{}
	if err := doc.Decode(result); err != nil {
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
	docs, err := mdb.dbs[dbPrimary].Collection(colUsers).DeleteMany(ctx, bson.M{colUsersAtrID: bson.M{"$in": ids}})
	if err != nil {
		return -1, err
	}
	return int(docs.DeletedCount), nil
}

// SetUserUsername deletes a user from the MongoDB persistant storage based on id
func (mdb *Driver) SetUserUsername(ctx context.Context, id, username string) error {
	if doc := mdb.dbs[dbPrimary].Collection(colUsers).FindOneAndUpdate(ctx, bson.D{}, bson.D{{"$set", bson.D{{colUsersAtrUsername, username}}}}); doc.Err() != nil {
		return doc.Err()
	}

	return nil
}

// SetUserPassword deletes a user from the MongoDB persistant storage based on id
func (mdb *Driver) SetUserPassword(ctx context.Context, id, password string) error {
	if doc := mdb.dbs[dbPrimary].Collection(colUsers).FindOneAndUpdate(ctx, bson.D{}, bson.D{{"$set", bson.D{{colUsersAtrPassword, password}}}}); doc.Err() != nil {
		return doc.Err()
	}

	return nil
}
