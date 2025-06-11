package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func main() {

	ctx := context.Background()
	res, err := http.Get("http://localhost:8090/hello")
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
		res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Status)
	fmt.Println(string(body))
}
