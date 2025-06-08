package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(3 * time.Second)

	fmt.Println("start")
	<-timer.C
	fmt.Println("end")
}
