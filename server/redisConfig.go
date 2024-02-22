package server

import (
	"context"
	"log"
	"strconv"

	"github.com/KumarVariable/go-for-url-shortner/models"
	"github.com/KumarVariable/go-for-url-shortner/util"
	"github.com/redis/go-redis/v9"
)

// Represents Redis Configuration
type RedisConfig struct {
	Address  string
	Password string
	Database int
	PoolSize int
}

// constructor to get fresh instance of RedisConfig struct
// with pre-determined values
func GetRedisConfig() *RedisConfig {

	return &RedisConfig{
		Address:  util.REDIS_CONNECTION_URL, // default redis connection url.
		Password: util.REDIS_DB_PASSKEY,     // leave empty if no authorization is set
		Database: util.DEFAULT_REDIS_DATABASE,
		PoolSize: util.DEFAULT_POOL_SIZE,
	}

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

// function to set up initial counter(integer value) required for unique ID
// generation for short url generator application
func SetUpCounter(redisClient *redis.Client, ctx context.Context) {

	result := redisClient.Get(ctx, util.REDIS_KEY_TO_GET_UNIQUE_ID)

	currentCounter := result.Val()
	if currentCounter == "" {
		redisClient.Set(ctx, util.REDIS_KEY_TO_GET_UNIQUE_ID, util.INITIAL_COUNTER_VALUE, 0)
		log.Printf(" set counter: %d into redis database ", util.INITIAL_COUNTER_VALUE)

	} else {

		intVal, err := strconv.Atoi(currentCounter)
		if err != nil {
			log.Printf("Error caught in SetUpCounter: %s\n", err.Error())
		}

		log.Printf("current counter: %d available into redis database ", intVal)
	}

}
