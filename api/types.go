package main

type Category string

const (
	Car         Category = "car"
	Electronics Category = "electronics"
)

type Product struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
}


