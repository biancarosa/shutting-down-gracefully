package main

import (
	"fmt"
	"time"

	"sync"
)


const (
	N       = 10
	Seconds = 10
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Hello world")

	for i := 0; i < N; i++ {
		fmt.Printf("%d of %d: Scheduling go routine\n", i, N)
		wg.Add(1)
		go func(i int) {
			fmt.Printf("%d of %d: A slow running goroutine started....\n", i, N)
			time.Sleep(Seconds * time.Second)
			fmt.Printf("%d of %d: A slow running goroutine finished....\n", i, N)
			wg.Done()
		}(i)
	}
	
	wg.Wait()

	fmt.Println("Everything has shut down, goodbye")
}