package underwood

import "net/http"

type Engine struct {
}

type Config struct {
	ErrHandler func(w http.ResponseWriter, r *http.Request) error
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
}
