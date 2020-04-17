package main

import (
	"github.com/OAOv/restful_CRUD/handler"
)

func main() {
	mysqlDB, err := repository.OpenDB()
	if err != nil {
		panic(err)
	}
	defer mysqlDB.Close()

	handler.RunHTTPServer()
}
