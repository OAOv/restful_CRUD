package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RunHTTPServer() {
	routerSQL := mux.NewRouter()
	Router(routerSQL)

	err := http.ListenAndServe(":8000", routerSQL)
	if err != nil {
		panic(err.Error())
	}
}
