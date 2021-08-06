package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func doStuff(w http.ResponseWriter, req *http.Request) {
	fmt.Println("One new request...")
	wg.Add(1)
	go func() {
		duration := time.Duration(10) * time.Minute
		time.Sleep(duration)
		fmt.Println("One request sucessfully finished...")
		wg.Done()
	}()
	fmt.Fprintf(w, "Doing background stuff for 10 minutes...\n")
}
func healthCheck(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Ok\n")
}

func main() {
	fmt.Println("Hello world.")
	server := http.Server{
		Addr: fmt.Sprintf(":%d", 8090),
	}
	http.HandleFunc("/", healthCheck)
	http.HandleFunc("/do-stuff", doStuff)

	gracefulStop := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint
		fmt.Println("Shutting down gracefully...")
		wg.Wait()
		if err := server.Shutdown(context.Background()); err != nil {
			fmt.Println("HTTP Server Shutdown failed")
		}
		close(gracefulStop)
	}()

	fmt.Println("Listening on 8090...")
	err := server.ListenAndServe()
	fmt.Println(err)

	<-gracefulStop
	fmt.Println("Everything has shut down, goodbye")
}
