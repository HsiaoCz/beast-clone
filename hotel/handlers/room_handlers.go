package handlers

import "github.com/HsiaoCz/beast-clone/hotel/storage"

type RoomHandlers struct {
	store *storage.Store
}

func NewRoomHandlers(store *storage.Store) *RoomHandlers {
	return &RoomHandlers{
		store: store,
	}
}
