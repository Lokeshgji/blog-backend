package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB() {

	dsn := "postgresql://postgres.yugaprmiwjxrpiwekotz:%40zerologic51295%23@aws-1-ap-northeast-1.pooler.supabase.com:5432/postgres"

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
