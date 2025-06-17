package main

import (
	"database/sql"
	"golang101/app/api"
	"golang101/app/model"
	"golang101/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	Env, err := config.LoadEnv("./config")
	if err != nil {
		log.Fatal("cannot load configuration", err)
	}

	db, err := sql.Open(Env.DBDriver, Env.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := model.NewStore(db)
	server := api.NewServer(store)
	err = server.Start(Env.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
