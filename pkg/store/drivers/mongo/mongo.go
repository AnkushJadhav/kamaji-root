package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// Driver is the adapter handling communication with the MongoDB persistant store
type Driver struct {
	client *mongo.Client
	dbs    map[string]*mongo.Database
}

// NewDriver initializes a driver for connecting to the MongoDB persistant store
func NewDriver(conn string) (*Driver, error) {
	if ok := (conn != ""); !ok {
		return nil, fmt.Errorf("invalid MongoDB connection string : %s", conn)
	}

	conf := initConfig(conn)
	client, err := initClient(conf)
	if err != nil {
		return nil, err
	}
	dbHandlers := initDatabaseHandlers(client)

	drvr := &Driver{
		client: client,
		dbs:    dbHandlers,
	}

	return drvr, nil
}

// Connect attempts to connect the driver mdb to the MongoDB persistant store
func (mdb *Driver) Connect() error {
	ctx := context.TODO()
	return mdb.client.Connect(ctx)
}

// Disconnect attempts to disconnect the driver mdb from the MongoDB persistant store
func (mdb *Driver) Disconnect() error {
	ctx := context.TODO()
	return mdb.client.Disconnect(ctx)
}

func initConfig(conn string) *options.ClientOptions {
	opts := &options.ClientOptions{}
	return opts.ApplyURI(conn)
}

func initClient(conf *options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func initDatabaseHandlers(client *mongo.Client) map[string]*mongo.Database {
	dbHandlers := make(map[string]*mongo.Database)

	dbHandlers[dbPrimary] = client.Database(dbPrimary)

	return dbHandlers
}
