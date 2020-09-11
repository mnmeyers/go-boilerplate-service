package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go-boilerplate/configurations"
	"sync"
)

const COLLECTION_NAME = "myCollection"
const DB_NAME = "myDB"

type Database interface {
	Repository
}

type MongoDB struct{}

type index struct {
	Key  map[string]int
	Name string
}

// Compile-time assertion that MongoDB implements Database
var _ Database = (*MongoDB)(nil)

// Thread-safe singleton of Database
var mongoDB *MongoDB
var once sync.Once

func GetDatabase() Database {
	once.Do(func() {
		mongoDB = &MongoDB{}
		err := insertIndexes(DB_NAME)
		if err != nil {
			fmt.Printf("Failed to add indexes to db: %v", err)
		}
	})

	return mongoDB
}

func (db *MongoDB) getClient() (*mongo.Client, error) {
	ctx := context.Background()
	uri := configurations.GetConfiguration().MongoURI
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func insertIndexes(dbName string) error {
	client, err := mongoDB.getClient()

	if err != nil {
		return err
	}
	// 2nd value is cancel function which I'm not using
	ctx := context.Background()
	defer client.Disconnect(ctx)
	collection := client.Database(dbName).Collection(COLLECTION_NAME)
	cursor, err := collection.Indexes().List(ctx)
	if err != nil {
		return err
	}

	var indexes []index
	err = cursor.All(ctx, &indexes)

	// _id is always added by default. Avoiding attempting to write
	// indexes that already exist
	if len(indexes) == 1 {
		_, err = collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
			{
				Keys: bson.D{{"someId", 1}},
			},
			{
				Keys: bson.D{{"someOtherField", 1}},
				Options: options.Index().SetUnique(true),
			},
		})

		if err != nil {
			return err
		}
	}

	return nil
}