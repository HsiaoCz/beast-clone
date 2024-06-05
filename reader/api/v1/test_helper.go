package v1

import (
	"context"
	"testing"

	"github.com/HsiaoCz/beast-clone/reader/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdbUri  = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.5"
	testDBName = "reader-reservation-test"
)

type testStore struct {
	client *mongo.Client
	store  *storage.Store
}

func setup(t *testing.T) *testStore {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdbUri))
	if err != nil {
		t.Fatal(err)
	}
	userStore := storage.NewMongoUserStore(client, client.Database(testDBName).Collection("users"))

	return &testStore{
		client: client,
		store: &storage.Store{
			User: userStore,
		},
	}
}

func (ts *testStore) tearDown(t *testing.T) {
	if err := ts.client.Database(testDBName).Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}
