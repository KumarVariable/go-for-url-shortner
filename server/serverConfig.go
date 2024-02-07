package server

import (
	"log"
	"net/http"
	"time"

	"github.com/KumarVariable/go-for-url-shortner/util"
	"github.com/gorilla/mux"
)

// Function to start HTTP server and handle incoming
// request on incoming connections
func StartHttpServer(router *mux.Router) string {

	hostName := "localhost"
	address := hostName + ":" + "9999"

	util.SERVER_STARTED_AT = time.Now()

	// non-blocking goroutine(lightweight thread) to
	// start server in a separate thread

	// go startServer(address, nil)
	go startServer(address, router)

	return address
}

func startServer(address string, router *mux.Router) {

	// ListenAndServe() is a blocking call. This means that the function
	// will not return until the server stops either due to an error or
	// shutdown signal to server
	server := &http.Server{
		Handler:      router,
		Addr:         address,
		ReadTimeout:  2 * time.Second, // max allowed time to read request by server
		WriteTimeout: 2 * time.Second, // max allowed time to write response by server
	}

	error := server.ListenAndServe()
	if error != nil {
		log.Panic("Failed to start server on address: ", address, " , Error:", error)
	}
}
