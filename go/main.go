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

func main() {

	fmt.Println("Hello world")
	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		fmt.Printf("%d of %d\n", i, N)
		wg.Add(1)
		go func() {
			fmt.Println("A slow running goroutine has started....")
			time.Sleep(Seconds * time.Second)
			wg.Done()
			fmt.Println("A slow running goroutine is done....")
		}()
	}

	wg.Wait()
	fmt.Println("Everything has shut down, goodbye")
}
