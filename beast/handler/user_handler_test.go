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

func TestCreateUser(t *testing.T) {
	if err := godotenv.Load("../,env"); err != nil {
		t.Fatal(err)
	}
	if err := db.Init(); err != nil {
		t.Fatal(err)
	}
	app := fiber.New()
	userHandler := NewUserHandler(data.NewUserStore(db.Get()))
	app.Post("/user", userHandler.HandleCreateUser)

	params := types.CreateUserParams{
		Username:         "shawcz",
		Email:            "shawcz@gmail.com",
		Password:         "zsaa123242",
		Synopsis:         "something as great",
		Avatar:           "./picture/1233.jpg",
		Background_Image: "./bgi/1244.jpg",
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
