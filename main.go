package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doStuff(w http.ResponseWriter, req *http.Request) {
	wg.Add(1)
	go func() {
		duration := time.Duration(10) * time.Minute
		time.Sleep(duration)
		wg.Done()
	}()
	fmt.Fprintf(w, "Doing background stuff for 10 minutes...\n")
}
func main() {
	fmt.Println("Hello world.")
	server := http.Server{
		Addr: fmt.Sprintf(":%d", 8090),
	}
	http.HandleFunc("/", doStuff)

	gracefulStop := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		fmt.Println("Shutting down gracefully...")
		wg.Wait()
		if err := server.Shutdown(context.Background()); err != nil {
			fmt.Println("HTTP Server Shutdown failed")
		}
		close(gracefulStop)
	}()

	fmt.Println("Listening on 8090...")
	server.ListenAndServe()
}
