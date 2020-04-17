package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func InitRouter(router *gin.Engine) {
	h := &Handler{}

	g1 := router.Group("")
	{
		g1.POST("/user", h.CreateUser)
		g1.GET("/users", h.GetUsers)
		g1.GET("/user/:id", h.GetUsers)
		g1.PATCH("/user/:id", h.UpdateUser)
		g1.DELETE("/user/:id", h.DeleteUser)
	}
}

func FRouter(router *mux.Router) {
	fh := &FHandler{}

	router.HandleFunc("/", fh.TmplHandler)
}
