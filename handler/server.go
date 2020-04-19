package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RunHTTPServer() {
	router := mux.NewRouter()
	Router(router)

	go func() {
		err := http.ListenAndServe(":8000", router)
		if err != nil {
			panic(err.Error())
		}
	}()

	fRouter := mux.NewRouter()
	FRouter(fRouter)
	http.ListenAndServe(":8010", fRouter)
}
