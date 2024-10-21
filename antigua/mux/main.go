package main

import "net/http"

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /user/{id}", HandleGetUserByID)
	http.ListenAndServe(":3001", router)
}

func HandleGetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_ = id
}
