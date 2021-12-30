package src

import (
	"database/sql"
	"fmt"
)

// Set variables to connection database
const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "mysecretpassword"
	DB_NAME     = "posgres"
)

// Connect to the database
func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	CheckErr(err)

	return DB
}
