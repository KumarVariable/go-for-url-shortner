package server

import (
	"github.com/KumarVariable/go-for-url-shortner/controllers"
	"github.com/KumarVariable/go-for-url-shortner/middleware"
	"github.com/gorilla/mux"
)

// Function to configure middleware to intercept request
func ConfigMiddleware(router *mux.Router) {

	router.Use(middleware.InterceptRequest)

	// use middleware to intercept response
	router.Use(middleware.InterceptResponse)
}

// Function to configure custom handlers for Router
func ConfigureCustomHandlers(router *mux.Router) {

	router.MethodNotAllowedHandler = controllers.MethodNotAllowedHandler()
	router.NotFoundHandler = controllers.RouteNotFoundHandler()
}
