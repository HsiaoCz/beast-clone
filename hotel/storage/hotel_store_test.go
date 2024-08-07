package storage

import "testing"

func TestCreateHotel(t *testing.T) {
	if err := Init(); err != nil {
		t.Fatal(err)
	}
	hotelColl := client.Database(test_db).Collection("hotels")
	hotelStore := NewMongoHotelStore(client, hotelColl)
	_ = hotelStore
}
