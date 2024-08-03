package storage

import (
	"context"

	"github.com/HsiaoCz/beast-clone/hotel/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelStorer interface {
	CreateHotel(context.Context,*types.Hotel) (*types.Hotel,error)
	UpdateHotel(context.Context)error
	GetHotels(context.Context)([]*types.Hotel,error)
	GetHotelByID(context.Context,primitive.ObjectID)(*types.Hotel,error)
	DeleteHotel(context.Context,primitive.ObjectID)error
}
