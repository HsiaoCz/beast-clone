package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HsiaoCz/beast-clone/reader/models"
	"github.com/gin-gonic/gin"
)

func TestCreateUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	app := gin.Default()
	userHandler := NewUserHandler(tdb.store)
	app.POST("/user", userHandler.HandleCreateUser)

	params := models.UserCreateParams{
		Username: "james.bound",
		Email:    "james@gmail.com",
		Password: "james123asd",
		IsAdmin:  false,
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	userParam := models.UserCreateParams{}
	json.NewDecoder(resp.Body).Decode(&userParam)
	if userParam.Username != params.Username {
		t.Errorf("expected firstname %s but got %s", params.Password, userParam.Password)
	}
	if userParam.Password != params.Password {
		t.Errorf("expected firstname %s but got %s", params.Password, userParam.Password)
	}
	if userParam.Email != params.Email {
		t.Errorf("expected firstname %s but got %s", params.Email, userParam.Email)
	}
	if userParam.IsAdmin != params.IsAdmin {
		t.Errorf("expected firstname %v but got %v", params.IsAdmin, userParam.IsAdmin)
	}
}
