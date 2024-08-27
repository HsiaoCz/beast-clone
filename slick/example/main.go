package main

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/slick"
)

func main() {
	slick := slick.New()
	http.ListenAndServe(":30001", slick)
}
