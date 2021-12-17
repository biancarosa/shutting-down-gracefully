package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	N    = 50
	Mins = 1
)

func main() {

	fmt.Println("Hello world")

	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		fmt.Printf("%d of %d\n", i, N)
		wg.Add(1)
		go func() {
			fmt.Println("A slow running goroutine....")
			time.Sleep(Mins * time.Minute)
			wg.Done()
		}()
	}

	gracefulStop := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		wg.Wait()
		close(gracefulStop)
	}()

	<-gracefulStop
	fmt.Println("Everything has shut down, goodbye")
}
