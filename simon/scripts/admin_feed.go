package scripts

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AdminFeed struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func AdminFeedInit() (*AdminFeed, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		return nil, err
	}
	return &AdminFeed{
		client: client,
		coll:   client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("USERCOLL")),
	}, nil
}
