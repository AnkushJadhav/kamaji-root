package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"context"

	"github.com/AnkushJadhav/kamaji-root/pkg/utils"

	"github.com/AnkushJadhav/kamaji-root/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo provides the mongo store driver implementation
type Mongo struct {
	rootClient *mongo.Client
	dbClient   *mongo.Database
}

// NewMongoDriver initialises a mongo client driver
func NewMongoDriver(dst string) (Mongo, error) {
	co := &options.ClientOptions{}
	co = co.ApplyURI(dst)

	c, err := mongo.NewClient(co)
	if err != nil {
		return Mongo{}, err
	}
	return Mongo{
		rootClient: c,
	}, nil
}

// Connect connects to dst
func (mdb Mongo) Connect() (store.Store, error) {
	err := mdb.rootClient.Connect(context.Background())
	if err != nil {
		return mdb, err
	}

	mdb.dbClient = mdb.rootClient.Database("kamaji-root")
	return mdb, nil
}

// GetAll gets all the documents from a collection
func (mdb Mongo) GetAll(col string) (int, []store.Model, error) {
	var m []store.Model
	c, err := mdb.dbClient.Collection(col).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return -1, m, err
	}
	c.All(context.Background(), m)

	return len(m), m, nil
}

// CreateOne creates a document in collection col
func (mdb Mongo) CreateOne(col string, m store.Model) (store.Model, error) {
	m = m.SetID(utils.GenerateUUID())
	m = m.SetSystemTimestamp()
	mdb.dbClient.Collection(col).InsertOne(context.Background(), m)
	return m, nil
}
