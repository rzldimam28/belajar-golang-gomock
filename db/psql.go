package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "leesrcyng__"
	DB_NAME     = "gomock"
)

func New() *sql.DB {

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	  DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error connecting to db: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error ping db: %s", err.Error())
	}

	return db
}