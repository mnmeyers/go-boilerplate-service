package repositories

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"
)

const TEST_DB = "testy"

func TestThatGetClientWorksAsExpected(t *testing.T) {
	expect := assert.New(t)
	client, err := mongoDB.getClient()
	expect.Nil(err)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	expect.Nil(err)
}

func TestInsertIndexesIsSuccessful(t *testing.T) {
	expect := assert.New(t)
	client, err := mongoDB.getClient()
	expect.Nil(err)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	defer client.Disconnect(ctx)
	err = insertIndexes(TEST_DB)
	expect.Nil(err)
	collection := client.Database(TEST_DB).Collection(COLLECTION_NAME)
	cursor, err := collection.Indexes().List(ctx)
	defer collection.Drop(ctx)

	expect.Nil(err)
	var indexes []index
	err = cursor.All(ctx, &indexes)
	expect.Nil(err)
	// _id is always added by default
	expect.Equal(3, len(indexes))
	expectedCustomerIdIndex := index{
		Key: map[string]int{
			"customerId": 1,
		},
		Name: "customerId_1",
	}

	expect.Equal(expectedCustomerIdIndex, indexes[1])

	expectedPaymentIdIndex := index{
		Key: map[string]int{
			"paymentId": 1,
		},
		Name: "paymentId_1",
	}

	expect.Equal(expectedPaymentIdIndex, indexes[2])
}