package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

// Define main function where we will start our server and handle routes
func main() {

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all movies
	router.HandleFunc("/clothes/", GetClothes).Methods("GET")

	// Create a movie
	router.HandleFunc("/clothes/", CreateClothes).Methods("POST")

	// Delete a specific movie by the movieID
	router.HandleFunc("/clothes/{clothesid}", DeleteClothes).Methods("DELETE")

	// Delete all movies
	router.HandleFunc("/clothes/", DeleteClothes).Methods("DELETE")

	// serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
