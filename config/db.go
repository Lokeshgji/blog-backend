package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB() {

	dsn := "postgresql://postgres:%40zerologic51295%23@db.yugaprmiwjxrpiwekotz.supabase.co:5432/postgres?sslmode=require"

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
