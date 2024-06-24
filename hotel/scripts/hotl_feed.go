package scripts

import (
	"context"
	"errors"
	"os"

	"github.com/HsiaoCz/beast-clone/hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HotelFeed struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewHotelFeed() (*HotelFeed, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		return nil, err
	}

	return &HotelFeed{
		client: client,
		coll:   client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("HOTELCOLL")),
	}, nil
}

func (h *HotelFeed) CretaeHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	filter := bson.M{
		"localtion": hotel.Localtion,
		"name":      hotel.Name,
	}
	res := h.coll.FindOne(ctx, filter)
	if res.Err() != mongo.ErrNoDocuments {
		return nil, errors.New("record exists")
	}
	result, err := h.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = result.InsertedID.(primitive.ObjectID)
	return hotel, nil
}
