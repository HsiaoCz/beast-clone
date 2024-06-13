package underwood

import "net/http"

type Handlerfunc func(w http.ResponseWriter, r *http.Request) error

