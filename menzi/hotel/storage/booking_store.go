package storage

import (
	"context"

	"github.com/HsiaoCz/beast-clone/hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingStorer interface {
	GetBookings(context.Context, bson.M) ([]*types.Booking, error)
	CreateBooking(context.Context, *types.Booking) (*types.Booking, error)
	GetBookingByID(context.Context, primitive.ObjectID) (*types.Booking, error)
	UpdateBooking(context.Context, primitive.ObjectID) (*types.Booking, error)
}

type MongoBookingStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoBookingStore(client *mongo.Client, coll *mongo.Collection) *MongoBookingStore {
	return &MongoBookingStore{
		client: client,
		coll:   coll,
	}
}

func (b *MongoBookingStore) CreateBooking(ctx context.Context, booking *types.Booking) (*types.Booking, error) {
	return nil, nil
}
func (b *MongoBookingStore) GetBookings(ctx context.Context, filter bson.M) ([]*types.Booking, error) {
	return nil, nil
}
func (b *MongoBookingStore) GetBookingByID(ctx context.Context, bookingID primitive.ObjectID) (*types.Booking, error) {
	return nil, nil
}
func (b *MongoBookingStore) UpdateBooking(ctx context.Context, userID primitive.ObjectID) (*types.Booking, error) {
	return nil, nil
}
