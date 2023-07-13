package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func Connect() {
	connectionString := "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"
	d, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	// Test the connection
	err = d.Ping()
	if err != nil {
		panic(err)
	}

	db = d
	fmt.Println("Connected to PostgreSQL database!")
}

func GetDB() *sql.DB {
	return db
}
