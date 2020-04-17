package main

import (
	"github.com/OAOv/restful_CRUD/repository"
	"github.com/OAOv/restful_CRUD/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysqlDB, err := repository.OpenDB()
	if err != nil {
		panic(err)
	}
	defer mysqlDB.Close()

	server.RunHTTPServer()
}
