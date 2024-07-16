package scripts

import (
	"testing"

	"github.com/HsiaoCz/beast-clone/beast/db"
	"github.com/joho/godotenv"
)

func TestCreateComment(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
	if err := db.Init(); err != nil {
		t.Fatal(err)
	}
}
