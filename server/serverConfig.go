package server

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const DEFAULT_PORT = 9999
const DEFAULT_HOST_NAME = "localhost"

// Represent HTTP server configuration
type HTTPServerConfig struct {
	Port         int
	Hostname     string
	Address      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// constructor to provide new instance of HTTPServerConfig struct
// with pre-determined values
func GetHttpServerConfig() *HTTPServerConfig {

	defaultPort := DEFAULT_PORT
	defaultHostname := DEFAULT_HOST_NAME

	return &HTTPServerConfig{
		Port:         defaultPort,
		Hostname:     defaultHostname,
		ReadTimeout:  2 * time.Second, // max allowed time to read request by server,
		WriteTimeout: 2 * time.Second, // max allowed time to write response by server
		Address:      defaultHostname + ":" + strconv.Itoa(defaultPort),
	}
}

// Function to start HTTP server and handle incoming
// request on incoming connections
func StartHttpServer(router *mux.Router) {

	// non-blocking goroutine(lightweight thread) to
	// start server in a separate thread
	go startServer(router)

}

// Start HTTP server
func startServer(router *mux.Router) {

	// ListenAndServe() is a blocking call. This means that the function
	// will not return until the server stops either due to an error or
	// shutdown signal to server
	server := &http.Server{
		Handler:      router,
		Addr:         GetHttpServerConfig().Address,
		ReadTimeout:  GetHttpServerConfig().ReadTimeout,
		WriteTimeout: GetHttpServerConfig().WriteTimeout,
	}

	error := server.ListenAndServe()
	if error != nil {
		log.Panic("Failed to start server on address: ", GetHttpServerConfig().Address, " , Error:", error)
	}

	log.Printf(" The HTTP server is operational on port %d and can be accessed at the address %s", GetHttpServerConfig().Port, GetHttpServerConfig().Address)
}
