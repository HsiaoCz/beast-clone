package app

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateHotel(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
}
