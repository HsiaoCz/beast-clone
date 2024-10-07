package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/HsiaoCz/beast-clone/underface/types"
	"github.com/joho/godotenv"
)

// man this is fucked
func TestCreateUser(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}
	router := http.NewServeMux()
	userHandlers := &UserHandlers{}
	router.HandleFunc("POST /user", TransferHandlerfunc(userHandlers.HandleCreateUser))
	port := os.Getenv("PORT")

	go func() {
		if err := http.ListenAndServe(port, router); err != nil {
			log.Fatal(err)
		}
	}()

	params := types.User{}
	b, _ := json.Marshal(params)
	client := &http.Client{}
	resp, err := client.Post("127.0.0.1:3001/user", "application/json", bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}

	// req and response
	user := types.User{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Fatal(err)
	}

}

func TestCreateUserWith(t *testing.T) {
	params := types.User{}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(b))
	w := httptest.NewRecorder()
	userHandler := &UserHandlers{}
	userHandler.HandleCreateUser(w, req)
	resp := w.Result()

	user := types.User{}

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", user)

}
