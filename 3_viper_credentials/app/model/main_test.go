package model

import (
	"database/sql"
	"fmt"
	"golang101/config"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	Env, err := config.LoadEnv("../../config")
	if err != nil {
		log.Fatal("cannot load configuration, ", err)
	}

	db, err := sql.Open(Env.DBDriver, Env.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	truncateDatabase(db)
	testQueries = New(db)

	os.Exit(m.Run())
}

func truncateDatabase(db *sql.DB) {
	res, _ := db.Query("SHOW TABLES")

	var tableName string
	var tables []string

	for res.Next() {
		res.Scan(&tableName)

		if tableName != "schema_migrations" {
			tables = append(tables, tableName)
		}
	}

	truncateTables(db, tables)
}

func truncateTables(db *sql.DB, tables []string) {
	_, _ = db.Exec("SET FOREIGN_KEY_CHECKS=0;")

	for _, v := range tables {
		_, _ = db.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", v))
	}

	_, _ = db.Exec("SET FOREIGN_KEY_CHECKS=1;")
}
