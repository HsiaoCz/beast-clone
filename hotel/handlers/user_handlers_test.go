package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/HsiaoCz/beast-clone/hotel/types"
	"github.com/gofiber/fiber/v2"
)

func TestCreateUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandlers(tdb.store)
	app.Post("/user", userHandler.HandleCreateUser)

	params := types.CreateUserParam{
		FirstName: "james",
		LastName:  "foo",
		Email:     "james@gmail.com",
		Password:  "james123asd",
		IsAdmin:   false,
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	userParam := types.CreateUserParam{}
	json.NewDecoder(resp.Body).Decode(&userParam)
	if userParam.FirstName != params.FirstName {
		t.Errorf("expected firstname %s but got %s", params.FirstName, userParam.FirstName)
	}
	if userParam.LastName != params.LastName {
		t.Errorf("expected firstname %s but got %s", params.LastName, userParam.LastName)
	}
	if userParam.Password != params.Password {
		t.Errorf("expected firstname %s but got %s", params.Password, userParam.Password)
	}
	if userParam.Email != params.Email {
		t.Errorf("expected firstname %s but got %s", params.FirstName, userParam.FirstName)
	}
	if userParam.IsAdmin != params.IsAdmin {
		t.Errorf("expected firstname %v but got %v", params.IsAdmin, userParam.IsAdmin)
	}
}

func TestGetUserByID(t *testing.T) {

	tdb := setup(t)
	defer tdb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandlers(tdb.store)
	app.Post("/user", userHandler.HandleCreateUser)
	app.Get("/user", userHandler.HandleGetUserByID)
	params := types.CreateUserParam{
		FirstName: "james",
		LastName:  "foo",
		Email:     "james@gmail.com",
		Password:  "james123asd",
		IsAdmin:   false,
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	user := types.User{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Fatal(err)
	}

	req1 := httptest.NewRequest("GET", "/user", nil)
	req1.URL.Query().Add("uid", user.ID.String())
	resp1, err := app.Test(req1)
	if err != nil {
		t.Fatal(err)
	}

	user1 := types.User{}
	if err := json.NewDecoder(resp1.Body).Decode(&user1); err != nil {
		t.Fatal(err)
	}
	if user.ID.String() != user1.ID.String() {
		t.Errorf("exception uid %s but get %s", user.ID.String(), user1.ID.String())
	}
}
