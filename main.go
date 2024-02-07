// main : defines a standalone executable program, not a library.
// The main package is neccessary for a Go program that will compile
// into a executable
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/KumarVariable/go-for-url-shortner/server"
)

// entry point of executable program.
// each program must have exactly one main function and
// must be in main package. The main function takes no arguments
// and returns no values. When we run the executable program the
// the main function get executed first.
func main() {

	fmt.Println("initialize and set up application")

	server.HandleRequests()
	router := server.SetUpRoutes()

	server.StartHttpServer(router)

	stopServerListener()

}

// 1.Implements a listener for goroutine (concurrent thread to start HTTP server)
// 2.The listener is listening for stop signal on a dedicated signals channel
// (channels in Go are used for communication between goroutines).
// 3.A separate function to gracefully shutdown server
func stopServerListener() {

	// create a new channel with a capacity to hold one os.Signal
	// 1 indicates buffer size of signal value at a time
	signals := make(chan os.Signal, 1)

	// link signals channel to incoming operating system signals
	// 1. os.Interrupt: signal when Ctrl+C is pressed
	// 2. syscall.SIGTERM: signal sent from OS to request the program
	// to stop running
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// loop to read signals.This loop will automatically
	// block and wait for message to come into channel.
	// This loop will be executed for every signal received
	for signal := range signals {

		fmt.Printf("Received signal: %v, shutting down server\n", signal)

		// add logic to perform any code clean up

		break
	}

	fmt.Println("Server has been stopped")

}
