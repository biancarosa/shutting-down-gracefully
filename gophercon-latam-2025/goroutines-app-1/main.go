package main

import (
	"fmt"
	"time"
)

const (
	N       = 10
	Seconds = 10
)

func main() {
	fmt.Println("Hello world")

	for i := 0; i < N; i++ {
		fmt.Printf("%d of %d\n", i, N)
		go func() {
			fmt.Println("A slow running goroutine started....")
			time.Sleep(Seconds * time.Second)
			fmt.Println("A slow running goroutine finished....")
		}()
	}
	fmt.Println("Everything has shut down, goodbye")
}