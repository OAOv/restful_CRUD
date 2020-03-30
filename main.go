package main

import (
	"github.com/OAOv/restful_CRUD/db"
	"github.com/OAOv/restful_CRUD/handler"
)

func main() {
	mysqlDB := db.OpenDB()
	handler.RunHTTPServer()

	defer mysqlDB.Close()
}
