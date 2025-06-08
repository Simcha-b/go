package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, _ := http.Get("https://gobyexample.com/")
	// if err != nil {
	// 	panic(err)
	// }

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(r)
	// 	}
	// 	res.Body.Close()
	// }()

	fmt.Println(res.Status)
}
