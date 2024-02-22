// main : defines a standalone executable program, not a library.
// The main package is neccessary for a Go program that will compile
// into a executable
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/KumarVariable/go-for-url-shortner/controllers"
	"github.com/KumarVariable/go-for-url-shortner/server"
)

// main() serves as the entry point of the executable program. In Go,
// each executable program requires exactly one main function and
// must be iwithin main package. The main() function takes no arguments
// nor does it returns any values. The main() function is the first
// function to get executed first, responsible for initializing and
// setting up the application environment.
func main() {

	fmt.Println("initialize and set up application")
	ctx := context.Background()

	redisClient := server.SetUpRedis()
	controllers.PingRedis(redisClient, ctx)

	server.SetUpCounter(redisClient, ctx)

	router := server.SetUpRoutes(redisClient)
	server.StartHttpServer(router)
	stopServerListener()

}

// Function to initialize a listener for OS signals to gracefully
// shutdown HTTP server. The listener is listening for stop signal
// (e.g., os.Interrupt, syscall.SIGTERM) on a dedicated signals channel
// (channels in Go are used for communication between goroutines).
// Upon receiving such a signal, it performs necessary cleanup
// before stopping the server. This mechanism ensures a graceful shutdown
// process, releasing resources and saving state as required.
func stopServerListener() {

	// Creates a channel for os.Signal with a buffer size of one.
	// `1` indicates buffer size of signal value at a time
	// The buffer size of one ensures at least one signal can be held if
	// the program is temporarily unable to read from the channel.
	signals := make(chan os.Signal, 1)

	// link signals channel to incoming operating system signals
	// 1. os.Interrupt: signal when Ctrl+C is pressed
	// 2. syscall.SIGTERM: signal sent from OS to request the program
	// to stop running
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// loop to read signals.This loop will automatically
	// block and wait for message to come into channel.
	// For each signal received, the loop processes the signal
	// appropriate action before breaking out of the loop to
	// shut down the server. This ensures the program remains
	// responsive to termination requests.
	for signal := range signals {

		fmt.Printf("Received signal: %v, shutting down server\n", signal)

		// add your logic or condition to perform any code clean up

		break
	}

	fmt.Println("Server has been stopped")

}
