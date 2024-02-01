package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KumarVariable/go-for-url-shortner/controllers"
	"github.com/KumarVariable/go-for-url-shortner/util"
	"github.com/gorilla/mux"
)

// Function to register the Handler/Endpoints
// for given request URLs (pattern)
func HandleRequests() {
	http.HandleFunc("/test", pingTest)

}

// Function to register route(s), handlers using gorilla/mux.
// mux - stands for HTTP request multiplexer.
func SetUpRoutes() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/test", pingTest).Methods("GET")

	router.HandleFunc("/get-short-url", controllers.GetShortUrl).Methods("GET")
	router.HandleFunc("/create-short-url", controllers.CreateShortUrl).Methods("POST")
	router.HandleFunc("/update-short-url", controllers.UpdateShortUrl).Methods("PUT")

	router.HandleFunc("/delete-short-url", controllers.DeleteShortUrl).Methods("DELETE")

	return router

}

// Route handler to get Server Uptime
func pingTest(w http.ResponseWriter, r *http.Request) {

	log.Println("request received for ping test ")

	serverUptime := util.GetServerUptime()
	uptimeString := util.FormatDuration(serverUptime)

	log.Println("response returned for ping test ", uptimeString)

	// write response to Response Writer.
	fmt.Fprintf(w, " server is running since : "+uptimeString)
}
