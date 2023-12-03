package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "1234"
	Name     = "bank_api"
)

func CreateConnection() *sql.DB {

	connStr := "user=postgres password=1234 dbname=bank_api sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err != nil {
		panic(err)
	}
	return db
}
