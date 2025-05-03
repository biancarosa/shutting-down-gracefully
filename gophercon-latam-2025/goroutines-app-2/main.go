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
		fmt.Printf("%d of %d: Scheduling go routine\n", i, N)
		go func(i int) {
			fmt.Printf("%d of %d: A slow running goroutine started....\n", i, N)
			time.Sleep(Seconds * time.Second)
			fmt.Printf("%d of %d: A slow running goroutine finished....\n", i, N)
		}(i)
	}

	time.Sleep(Seconds * time.Second * 2)

	fmt.Println("Everything has shut down, goodbye")
}