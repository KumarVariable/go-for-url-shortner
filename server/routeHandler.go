package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KumarVariable/go-for-url-shortner/util"
)

// Function to register the Handler/Endpoints
// for given request URLs (pattern)
func HandleRequests() {
	http.HandleFunc("/test", pingTest)

}

// Request Handler - Test Server Uptime
func pingTest(w http.ResponseWriter, r *http.Request) {

	serverUptime := util.GetServerUptime()
	uptimeString := util.FormatDuration(serverUptime)

	log.Println("response returned for ping test ", uptimeString)

	// write response to Response Writer.
	fmt.Fprintf(w, " server is running since : "+uptimeString)
}
