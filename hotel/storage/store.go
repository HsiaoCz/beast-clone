package storage

type Store struct {
	User UserStorer
	Book BookingStorer
	Hotel HotelStorer
}
