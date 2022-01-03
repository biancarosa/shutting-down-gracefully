package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	N       = 50
	Seconds = 10
)

type Status string

const (
	Started  Status = "STARTED"
	Finished Status = "FINISHED"
)

type Task struct {
	ID     int
	Status Status
}

func main() {

	fmt.Println("Hello world")

	var wg sync.WaitGroup
	tasks := make([]*Task, N)
	for i := 0; i < N; i++ {
		fmt.Printf("%d of %d\n", i, N)
		task := Task{
			ID:     i,
			Status: Started,
		}
		wg.Add(1)
		tasks[i] = &task
		go func(t *Task) {
			fmt.Println("A slow running goroutine started....")
			time.Sleep(Seconds * time.Second)
			task.Status = Finished
			wg.Done()
		}(&task)
	}

	wg.Wait()

	for i, t := range tasks {
		fmt.Printf("Index %d of Task ID %d has status %s\n", i, t.ID, t.Status)
	}
	fmt.Println("Everything has shut down, goodbye")
}
