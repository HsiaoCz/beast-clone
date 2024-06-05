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
