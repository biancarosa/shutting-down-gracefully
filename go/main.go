package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	N       = 50
	Seconds = 100
)

type Status string

const (
	Started     Status = "STARTED"
	Finished    Status = "FINISHED"
	Interrupted Status = "INTERRUPTED"
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
		go func() {
			sigint := make(chan os.Signal, 1)
			signal.Notify(sigint, os.Interrupt)
			<-sigint
			// Is it possible to have a race condition here? Yes.
			wg.Done()
			task.Status = Interrupted
			fmt.Printf("Interrupted Task ID %d\n", task.ID)

		}()
		go func(t *Task) {
			fmt.Println("A slow running goroutine started....")
			time.Sleep(Seconds * time.Second)
			fmt.Println("A slow running goroutine finished....")
			task.Status = Finished
			wg.Done()
		}(&task)
	}
	wg.Wait()
	for _, t := range tasks {
		fmt.Printf("Task ID %d has status %s\n", t.ID, t.Status)
	}
	fmt.Println("Everything has shut down, goodbye")
}
