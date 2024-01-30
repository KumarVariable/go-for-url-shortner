package server

import (
	"log"
	"net/http"
	"time"

	"github.com/KumarVariable/go-for-url-shortner/util"
)

// Function to start HTTP server and handle incoming
// request on incoming connections
func StartHttpServer() string {

	hostName := "localhost"
	address := hostName + ":" + "9999"

	util.SERVER_STARTED_AT = time.Now()

	// non-blocking goroutine(lightweight thread) to start server
	// in a separate thread
	go startServer(address)

	return address
}

func startServer(address string) {

	// http.ListenAndServe() is a blocking call. This means that the function
	// will not return until the server stops either due to an error or
	// shutdown signal to server
	error := http.ListenAndServe(address, nil)
	if error != nil {
		log.Panic("Failed to start server on address: ", address, " , Error:", error)
	}
}
