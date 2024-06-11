package storage

import "context"

type BookingStorer interface {
	BookingRoom(context.Context) error
}
