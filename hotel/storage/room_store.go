package storage

import "context"

type RoomStorer interface {
	CreateRoom(context.Context)error
}
