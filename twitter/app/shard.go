package app

import (
	"encoding/json"
	"net/http"
)

type Handlerfunc func(w http.ResponseWriter, r *http.Request) error

func TransferHandler(hf Handlerfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := hf(w, r); err != nil {
			if e, ok := err.(ErrorResp); ok {
				WriteJson(w, e.Status, &e)
			}
			errResp := ErrorResp{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			WriteJson(w, errResp.Status, &errResp)
		}
	}
}

func WriteJson(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
