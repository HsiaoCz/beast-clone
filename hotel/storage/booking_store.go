package storage

import (
	"context"

	"github.com/HsiaoCz/beast-clone/hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingStorer interface {
	GetBookings(context.Context,bson.M)([]*types.Booking,error)
	CreateBooking(context.Context,*types.Booking)(*types.Booking,error)
	GetBookingByID(context.Context,primitive.ObjectID)(*types.Booking,error)
	UpdateBooking(context.Context,primitive.ObjectID)(*types.Booking,error)
}

