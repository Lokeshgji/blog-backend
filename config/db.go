package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB() {

	dsn := "user=postgres password=postgres dbname=blog sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
