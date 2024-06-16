package middlewares

import "net/http"

type Middlewarefunc func(http.Handler) http.Handler

func SetMiddlewares(handler http.Handler, middlewares ...Middlewarefunc) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
