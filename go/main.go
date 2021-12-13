package main

import (
	"fmt"
	"time"
)

const (
	N    = 50
	Mins = 1
)

func main() {
	fmt.Println("Hello world")
	for i := 0; i < N; i++ {
		fmt.Printf("%d of %d", i, N)
		go func() {
			fmt.Println("A slow running goroutine....")
			time.Sleep(Mins * time.Minute)
		}()
	}
}
