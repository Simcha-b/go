package main

import "fmt"

func main()  {
	fmt.Println("i ❤️ Pointers!!")

	// var a *int
	b := 2
	var p = &b

	fmt.Println("value of a is:", p)
	fmt.Println("value of a is:", *p)
}