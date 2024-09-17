package store

import (
	"context"

	"github.com/HsiaoCz/beast-clone/hotel/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomStorer interface {
	CreateRoom(context.Context,*types.Room)error
	GetRooms(context.Context)([]*types.Room,error)
	GetRoomByID(context.Context,primitive.ObjectID)(*types.Room,error)
	DeleteRoomByID(context.Context,primitive.ObjectID)error
}
