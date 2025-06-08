package main

import (
	// "fmt"
	"log"
	"os"

	"log/slog"
)

func main() {
	log.Println("standard log")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	myslog.Info("hello again", "key", "val", "age", 25)

}
