package scripts

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateAdmin(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
}
