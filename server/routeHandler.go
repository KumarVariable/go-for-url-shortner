package server

import (
	"github.com/KumarVariable/go-for-url-shortner/controllers"
	"github.com/KumarVariable/go-for-url-shortner/models"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

// Function to register route(s), handlers using gorilla/mux.
// mux - stands for HTTP request multiplexer.
func SetUpRoutes() *mux.Router {

	router := mux.NewRouter()

	ConfigMiddleware(router)
	ConfigureCustomHandlers(router)

	router.HandleFunc("/test", controllers.PingTest).Methods("GET")
	router.HandleFunc("/get-short-url", controllers.GetShortUrl).Methods("GET")
	router.HandleFunc("/create-short-url", controllers.CreateShortUrl).Methods("POST")
	router.HandleFunc("/update-short-url", controllers.UpdateShortUrl).Methods("PUT")
	router.HandleFunc("/delete-short-url", controllers.DeleteShortUrl).Methods("DELETE")

	router.HandleFunc("/get-key", controllers.GetKeyFromRedis).Methods("GET")
	router.HandleFunc("/get-all-keys", controllers.GetAllKeysFromRedis).Methods("GET")
	router.HandleFunc("/add-key", controllers.StoreKeyValue).Methods("POST")

	return router

}

// Function to configure locally running redis server
func SetUpRedis() *redis.Client {

	redisOptions := redis.Options{
		Addr:     GetRedisConfig().Address,
		Password: GetRedisConfig().Password,
		DB:       GetRedisConfig().Database,
		PoolSize: GetRedisConfig().PoolSize,
	}

	redisClient := redis.NewClient(&redisOptions)
	models.RedisClient = redisClient

	return redisClient

}
