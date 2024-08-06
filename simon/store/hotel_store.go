package store

import (
	"context"

	"github.com/HsiaoCz/beast-clone/simon/st"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStorer interface {
	CreateHotel(context.Context, *st.Hotel) (*st.Hotel, error)
	UpdateHotel(context.Context, bson.M, bson.M) (*st.Hotel, error)
	GetHotels(context.Context, bson.M) ([]*st.Hotel, error)
	GetHotelByID(context.Context, primitive.ObjectID) (*st.Hotel, error)
	DeleteHotel(context.Context, primitive.ObjectID) error
}

type HotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func HotelStoreInit(client *mongo.Client, coll *mongo.Collection) *HotelStore {
	return &HotelStore{
		client: client,
		coll:   coll,
	}
}
func (h *HotelStore) CreateHotel(context.Context, *st.Hotel) (*st.Hotel, error) {
	return nil, nil
}
func (h *HotelStore) UpdateHotel(context.Context, bson.M, bson.M) (*st.Hotel, error) {
	return nil, nil
}
func (h *HotelStore) GetHotels(context.Context, bson.M) ([]*st.Hotel, error) {
	return nil, nil
}
func (h *HotelStore) GetHotelByID(context.Context, primitive.ObjectID) (*st.Hotel, error) {
	return nil, nil
}
func (h *HotelStore) DeleteHotel(context.Context, primitive.ObjectID) error {
	return nil
}
