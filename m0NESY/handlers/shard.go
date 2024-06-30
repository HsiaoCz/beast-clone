package handlers

import "net/http"

type Handlerfunc func(w http.ResponseWriter, r *http.Request) error

type Map map[string]any

func TransferHandlerfunc(h Handlerfunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}