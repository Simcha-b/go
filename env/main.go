package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// os.Setenv("ENV_VAR", "Hello, World!")
	// value := os.Getenv("ENV_VAR")
	// fmt.Println("The value of ENV_VAR is:", value)
	for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Printf("%s : %s\n", pair[0], pair[1])
    }
}
