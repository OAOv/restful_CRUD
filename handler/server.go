package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RunHTTPServer() {
	routerSQL := mux.NewRouter()
	Router(routerSQL)

	go func() {
		err := http.ListenAndServe(":8000", routerSQL)
		if err != nil {
			panic(err.Error())
		}
	}()

	fRouter := mux.NewRouter()
	FRouter(fRouter)
	http.ListenAndServe(":8010", fRouter)
}
