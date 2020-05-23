package mongo

import (
	"context"

	"github.com/AnkushJadhav/kamaji-root/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

// AddNodeToUser updates a user with data identified by id
func (mdb *Driver) AddNodeToUser(ctx context.Context, userID string, node models.Node) (int, error) {
	docs, err := mdb.dbs[dbPrimary].Collection(colUsers).UpdateOne(ctx, bson.D{{atrID, userID}}, bson.M{"$push": bson.D{{atrNodes, node}}})
	if err != nil {
		return -1, err
	}

	return int(docs.MatchedCount), nil
}
