package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

const (
	N    = 50
	Mins = 1
)

func main() {

	fmt.Println("Hello world")

	goroutines := make(chan bool, N)
	for i := 0; i < N; i++ {
		fmt.Printf("%d of %d\n", i, N)
		go func() {
			fmt.Println("A slow running goroutine....")
			time.Sleep(Mins * time.Minute)
			goroutines <- true
		}()
	}

	gracefulStop := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		for i := 0; i < N; i++ {
			<-goroutines
		}
		close(gracefulStop)
	}()

	<-gracefulStop
	fmt.Println("Everything has shut down, goodbye")
}
