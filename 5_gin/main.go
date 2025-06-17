package main

import (
	"database/sql"
	"golang101/app/api"
	"golang101/app/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver      = "mysql"
	dbSource      = "root:password@tcp(127.0.0.1:3306)/golang_development?parseTime=true"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := model.NewStore(db)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
