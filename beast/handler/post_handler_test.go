package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HsiaoCz/beast-clone/beast/data"
	"github.com/HsiaoCz/beast-clone/beast/db"
	"github.com/HsiaoCz/beast-clone/beast/types"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func TestCreatePost(t *testing.T) {
	if err := godotenv.Load("../,env"); err != nil {
		t.Fatal(err)
	}
	if err := db.Init(); err != nil {
		t.Fatal(err)
	}
	app := fiber.New()
	postHandler := NewPostHandler(data.NewPostStore(db.Get()))
	app.Post("/post", postHandler.HandleCreatePost)

	params := types.CreatePostParams{
		UserID:   "",
		Content:  "something",
		PostPath: "./posts/12232.txt",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	userParam := types.CreateUserParams{}
	json.NewDecoder(resp.Body).Decode(&userParam)

}
