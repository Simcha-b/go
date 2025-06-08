package main

import (
	"fmt"
	"time"
)

type Task struct {
	id int
}

func (t *Task) Process() {
	fmt.Printf("processing task number %d",t.id)
	time.Sleep(time.Second * 2)
}

type workerPool struct{
	Tasks[]
}