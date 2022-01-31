package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TODO: Use env variable
func NewPostgreSQLDatabase() *sqlx.DB {
	dsn := "host=localhost user=shirotama password=Rizkiatama15DB dbname=rasalibrary sslmode=disable port=5432 TimeZone=Asia/Jakarta"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
