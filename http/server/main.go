package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {

	http.HandleFunc("/hello", hello)

	fmt.Println("server running in port 8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
