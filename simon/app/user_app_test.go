package app

import (
	"net/http"
	"testing"
)

func TestCreateUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	router := http.NewServeMux()

	userApp := UserAppInit(tdb.store)
	router.HandleFunc("POST /user", TransferHandlerfunc(userApp.HandleCreateUser))

}
