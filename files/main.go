package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("start working with files!!")
	
	context:= "to be or not to be, this is the question"

	file, err := os.Create("./myFile.txt")

	checkErr(err)

	len, err := io.WriteString(file, context)

	checkErr(err)

	fmt.Println("length is: ", len)

	redFile("myFile.txt")

	defer file.Close()
}



func redFile(fileName string)  {
	con, err := os.ReadFile(fileName)
	checkErr(err)
	fmt.Println(string(con))
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}