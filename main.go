package main

import (
	"database/sql"
	"fmt"

	chkerrors "github.com/dcristobalh/api-rest-go/src"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "mysecretpassword"
	DB_NAME     = "posgres"
)

// Connect to the database
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	chkerrors.CheckErr(err)

	return DB
}
