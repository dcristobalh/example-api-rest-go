package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dcristobalh/api-rest-go/src"
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

	// Get all clothes
	router.HandleFunc("/clothes/", GetClothes).Methods("GET")

	// Create a clothes
	router.HandleFunc("/clothes/", CreateClothes).Methods("POST")

	// Delete a specific clothes by the clothesID
	router.HandleFunc("/clothes/{clothesid}", DeleteClothes).Methods("DELETE")

	// Delete all clothes
	router.HandleFunc("/clothes/", DeleteAllClothes).Methods("DELETE")

	// serve the app
	fmt.Println("Server is up & running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetClothes(w http.ResponseWriter, r *http.Request) {
	db := src.SetupDB()

	src.PrintMessage("Getting clothes...")

	// Get all clothes from clothes table that don't have clothesID = "1"
	rows, err := db.Query("SELECT * FROM clothes")

	// check errors
	src.CheckErr(err)

	// var response []JsonResponse
	var clothes []Clothes

	// Foreach clothes
	for rows.Next() {
		var id int
		var clothesID string
		var clothesName string

		err = rows.Scan(&id, &clothesID, &clothesName)

		// check errors
		src.CheckErr(err)

		clothes = append(clothes, Clothes{ClothesID: clothesID, ClothesName: clothesName})
	}

	var response = JsonResponse{Type: "success", Data: clothes}

	json.NewEncoder(w).Encode(response)
}

func CreateClothes(w http.ResponseWriter, r *http.Request) {
	clothesID := r.FormValue("clothesid")
	clothesName := r.FormValue("clothesname")

	var response = JsonResponse{}

	if clothesID == "" || clothesName == "" {
		response = JsonResponse{Type: "error", Message: "You are missing clothesID or clothesName parameter."}
	} else {
		db := src.SetupDB()

		src.PrintMessage("Inserting clothes into DB")

		fmt.Println("Inserting new clothes with ID: " + clothesID + " and name: " + clothesName)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO clothes(clothesID, clothesName) VALUES($1, $2) returning id;", clothesID, clothesName).Scan(&lastInsertID)

		// check errors
		src.CheckErr(err)

		response = JsonResponse{Type: "success", Message: "The clothes has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete a clothes

// response and request handlers
func DeleteClothes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	clothesID := params["clothesid"]

	var response = JsonResponse{}

	if clothesID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing clothesID parameter."}
	} else {
		db := src.SetupDB()

		src.PrintMessage("Deleting clothes from DB")

		_, err := db.Exec("DELETE FROM clothes where clothesID = $1", clothesID)

		// check errors
		src.CheckErr(err)

		response = JsonResponse{Type: "success", Message: "The clothes has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete all clothes

// response and request handlers
func DeleteAllClothes(w http.ResponseWriter, r *http.Request) {
	db := src.SetupDB()

	src.PrintMessage("Deleting all clothes...")

	_, err := db.Exec("DELETE FROM clothes")

	// check errors
	src.CheckErr(err)

	src.PrintMessage("All clothes have been deleted successfully!")

	var response = JsonResponse{Type: "success", Message: "All clothes have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}
