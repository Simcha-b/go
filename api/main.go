package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var Products = []Product{
	{Id: "1", Name: "Toyota Corolla", Category: Car},
	{Id: "2", Name: "Samsung Galaxy S21", Category: Electronics},
	{Id: "3", Name: "Apple MacBook Pro", Category: Electronics},
	{Id: "4", Name: "Honda Civic", Category: Car},
	{Id: "5", Name: "Sony WH-1000XM4", Category: Electronics},
}

func home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "ברוך הבא ל-API שלנו!")
}

func getAllProductsByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	category := params["category"]
	fmt.Println("הקטגוריה המבוקשת היא:", category)
	p := []Product{}
	for _, product := range Products {
		if strings.EqualFold(string(product.Category), category) {
			p = append(p, product)
		}
	}
	if len(p) == 0 {
		http.Error(w, "לא נמצאו מוצרים בקטגוריה זו", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func main() {
	r := mux.NewRouter()
	fmt.Println("Starting server on :8080")
	r.HandleFunc("/", home)
	r.HandleFunc("/products/{category}", getAllProductsByCategory).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
