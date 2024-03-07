package server

import (
	"github.com/KumarVariable/go-for-url-shortner/controllers"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

// Function to register route(s), handlers using gorilla/mux.
// mux - stands for HTTP request multiplexer.
func SetUpRoutes(redisClient *redis.Client) *mux.Router {

	router := mux.NewRouter()

	ConfigMiddleware(router)
	ConfigureCustomHandlers(router)

	// create a sub-router for short url based operations
	urlsRouter := router.PathPrefix("/urls").Subrouter()

	// create a sub-router for redis key based operations
	keysRouter := router.PathPrefix("/key").Subrouter()

	// test ping local running redis database
	router.HandleFunc("/test", controllers.PingTest).Methods("GET")

	// short-url related operations
	urlsRouter.HandleFunc("/get-short-url", controllers.GetShortUrl(redisClient)).Methods("GET")
	urlsRouter.HandleFunc("/create-short-url", controllers.CreateShortUrl(redisClient)).Methods("POST")
	urlsRouter.HandleFunc("/update-short-url", controllers.UpdateShortUrl(redisClient)).Methods("POST")
	urlsRouter.HandleFunc("/delete-short-url", controllers.DeleteShortUrl(redisClient)).Methods("GET")

	urlsRouter.HandleFunc("/custom-short-url", controllers.CreateCustomShortUrl(redisClient)).Methods("POST")

	// test redis database operations
	keysRouter.HandleFunc("/get-key", controllers.GetKeyFromRedis).Methods("GET")
	keysRouter.HandleFunc("/get-all-keys", controllers.GetAllKeysFromRedis).Methods("GET")
	keysRouter.HandleFunc("/add-key", controllers.StoreKeyValue).Methods("POST")

	// redirect route - use pattern to match short url id containing only
	// alphanumeric characters to avoid clashes with other routers.
	// `+` in the regular expression ensures that the shortUrlID contains
	// at least one character, to prevent the route from matching empty strings.
	router.HandleFunc("/{shortUrlID:[A-Za-z0-9]+}", controllers.RedirectToOriginalUrl(redisClient)).Methods("GET")

	return router

}
