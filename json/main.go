package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string
	Password string `json:"pass"`
	Age      int `json: age`
	Tags     []string
}

func main() {

	user1 := user{
		"John", 
		"12345",
		30,
		nil,
	}
	EncodeJson(user1)
}

// EncodeJson encodes a user struct to JSON format
func EncodeJson(u user) {
	jsonData, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Printf("Encoded JSON: %s", jsonData)
}


