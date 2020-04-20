package server

import (
	"github.com/OAOv/restful_CRUD/handler"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func InitRouter(router *gin.Engine) {
	userHandler := &handler.UserAPI{}
	g1 := router.Group("")
	{
		g1.POST("/user", userHandler.CreateUser)
		g1.GET("/users", userHandler.GetUsers)
		g1.GET("/user/:id", userHandler.GetUser)
		g1.PATCH("/user/:id", userHandler.UpdateUser)
		g1.DELETE("/user/:id", userHandler.DeleteUser)
	}

	recordHandler := &handler.RecordAPI{}
	g2 := router.Group("")
	{
		g2.POST("/record", recordHandler.CreateRecord)
		g2.GET("/records", recordHandler.GetRecords)
		g2.GET("/record/:id", recordHandler.GetRecord)
		g2.GET("/records/user/:id", recordHandler.GetRecordByUser)
		g2.PATCH("/record/:id", recordHandler.UpdateRecord)
		g2.DELETE("/record/:id", recordHandler.DeleteRecord)
	}
}

func FRouter(router *mux.Router) {
	fh := &handler.FHandler{}

	router.HandleFunc("/", fh.TmplHandler)
}
