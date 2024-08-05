package storage

import (
	"context"

	"github.com/HsiaoCz/beast-clone/hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStorer interface {
	CreateHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	UpdateHotel(context.Context, bson.M, bson.M) (*types.Hotel, error)
	GetHotels(context.Context, bson.M) ([]*types.Hotel, error)
	GetHotelByID(context.Context, primitive.ObjectID) (*types.Hotel, error)
	DeleteHotel(context.Context, primitive.ObjectID) error
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client, coll *mongo.Collection) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   coll,
	}
}

func (h *MongoHotelStore) CreateHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	return nil, nil
}

func (h *MongoHotelStore) UpdateHotel(ctx context.Context, filter bson.M, update bson.M) (*types.Hotel, error) {
	return nil, nil
}

func (h *MongoHotelStore) GetHotels(ctx context.Context, filter bson.M) ([]*types.Hotel, error) {
	return nil, nil
}

func (h *MongoHotelStore) GetHotelByID(ctx context.Context, hotelID primitive.ObjectID) (*types.Hotel, error) {
	return nil, nil
}

func (h *MongoHotelStore) DeleteHotel(ctx context.Context, hotelID primitive.ObjectID) error {
	return nil
}
