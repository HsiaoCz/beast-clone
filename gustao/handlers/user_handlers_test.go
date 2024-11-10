package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HsiaoCz/beast-clone/gustao/data"
	"github.com/HsiaoCz/beast-clone/gustao/db"
	"github.com/HsiaoCz/beast-clone/gustao/types"
	"github.com/joho/godotenv"
)

func TestCreateUser(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
	if err := db.Init(); err != nil {
		t.Fatal(err)
	}

	userHandler := UserHandlersInit(data.UserDataInit(db.Get()))

	app := http.NewServeMux()

	app.HandleFunc("POST /api/v1/user", TransferHandlerfunc(userHandler.HandleCreateUser))

	parmas := types.CreateUserReq{
		Username: "shawcz",
		Email:    "shawcz@gmail.com",
		Password: "shawcz123",
		Age:      26,
		Birthday: "1998/04/01",
		Gender:   "男",
	}

	b, _ := json.Marshal(parmas)
	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	userParams := types.CreateUserReq{}

	json.NewDecoder(resp.Body).Decode(&userParams)

	if parmas.Username != userParams.Username {
		t.Errorf("want %s but got %s", parmas.Username, userParams.Username)
	}
}
