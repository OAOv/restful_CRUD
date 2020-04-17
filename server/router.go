package server

import (
	"github.com/OAOv/restful_CRUD/handler"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func InitRouter(router *gin.Engine) {
	gh := &handler.UserAPI{}
	g1 := router.Group("")
	{
		g1.POST("/user", gh.CreateUser)
		g1.GET("/users", gh.GetUsers)
		g1.GET("/user/:id", gh.GetUser)
		g1.PATCH("/user/:id", gh.UpdateUser)
		g1.DELETE("/user/:id", gh.DeleteUser)
	}
}

func FRouter(router *mux.Router) {
	fh := &handler.FHandler{}

	router.HandleFunc("/", fh.TmplHandler)
}
