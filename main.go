package main

import (
	"database/sql"
	"fmt"

	"github.com/dcristobalh/api-rest-go/src"
)

// Set variables to connection database
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "mysecretpassword"
	DB_NAME     = "posgres"
)

// Create structs to define and get data from database
type Clothes struct {
	ClothesID   string `json:"clothesid"`
	ClothesName string `json:"clothesname"`
}

type JsonResponse struct {
	Type    string    `json:"type"`
	Data    []Clothes `json:"data"`
	Message string    `json:"message"`
}

// Connect to the database
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	src.CheckErr(err)

	return DB
}
