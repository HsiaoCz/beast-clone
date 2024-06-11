package storage

import "context"

type HotelStorer interface {
	CreateHotel(context.Context) error
}
