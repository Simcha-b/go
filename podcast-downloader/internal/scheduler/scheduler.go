package scheduler

import (
	"time"
)

type Task func()

func Schedule(interval time.Duration, task Task) {
	go func() {
		for {
			task()
			time.Sleep(interval)
		}
	}()
}
