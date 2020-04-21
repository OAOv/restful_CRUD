package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func RunHTTPServer() {
	router := gin.Default()
	InitRouter(router)
	go router.Run(":8000")

	fRouter := mux.NewRouter()
	FRouter(fRouter)
	http.ListenAndServe(":8010", fRouter)
}
