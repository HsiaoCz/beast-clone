package scripts

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BookingFeed struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewBookingFeed() (*BookingFeed, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		return nil, err
	}
	return &BookingFeed{client: client, coll: client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("BOOKINGCOLL"))}, nil
}
