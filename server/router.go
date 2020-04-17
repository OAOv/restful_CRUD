package server

import (
	"github.com/OAOv/restful_CRUD/handler"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func InitRouter(router *gin.Engine) {
	g1 := router.Group("")
	{
		g1.POST("/user", CreateUser)
		g1.GET("/users", handler.GetUsers)
		g1.GET("/user/:id", GetUser)
		g1.PATCH("/user/:id", UpdateUser)
		g1.DELETE("/user/:id", DeleteUser)
	}
}

func FRouter(router *mux.Router) {
	router.HandleFunc("/", handler.TmplHandler)
}
