package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"path"

	"github.com/KumarVariable/go-for-url-shortner/models"
	"github.com/KumarVariable/go-for-url-shortner/util"
	"github.com/redis/go-redis/v9"
)

// Helper function to increase the value of the "counter" key in Redis
// by +1 each time it's called.This counter value is then used to
// generate a base62 string.
// - `redisClient` is the connection to the Redis database.
// - `ctx` is the context for this operation, allowing for things
// like timeouts or cancellations.
func IncrementCounter(redisClient *redis.Client, ctx context.Context) int64 {

	result, err := redisClient.Incr(ctx, util.REDIS_KEY_TO_GET_UNIQUE_ID).Result()
	if err != nil {
		log.Println(" Error caught to increment counter into redis database")
		return 0
	}

	if result < int64(util.INITIAL_COUNTER_VALUE) {
		answer, _ := redisClient.Set(ctx, util.REDIS_KEY_TO_GET_UNIQUE_ID, util.INITIAL_COUNTER_VALUE, 0).Result()
		log.Println("Answer:::  ", answer)
		if answer == "OK" {
			return int64(util.INITIAL_COUNTER_VALUE)
		}
	}

	return result

}

// Function to check whether the mapping for long url and short url
// already exists in redis database.
func IsShortUrlExistsForLongUrl(redisClient *redis.Client, ctx context.Context, payload *models.Payload) bool {

	indexKey := fmt.Sprintf(util.REDIS_KEY_TO_STORE_LONG_URL_ID+":%s", payload.LongUrl)
	redisKey, err := redisClient.Get(ctx, indexKey).Result()
	if err != nil {
		log.Println("long url mapping key not found", err.Error())
		return false
	}
	log.Println(" short url already exists for long url ", redisKey)
	return true

}

// Function to find data related to short url by longUrl
func FindByLongUrl(redisClient *redis.Client, ctx context.Context, payload *models.Payload) error {

	longUrlKey := fmt.Sprintf(util.REDIS_KEY_TO_STORE_LONG_URL_ID+":%s", payload.LongUrl)

	redisKey, err := redisClient.Get(ctx, longUrlKey).Result()
	if err != nil {
		log.Println("long url mapping key not found", err.Error())
		return err
	}

	result, err := redisClient.Get(ctx, redisKey).Result()
	if err != nil {
		log.Printf("data not found for key %v", redisKey)
		return err
	}

	// Convert the result to a byte slice before unmarshalling
	err = json.Unmarshal([]byte(result), payload)
	if err != nil {
		log.Println("error unmarshalling result into payload", err.Error())
		return err
	}

	return nil

}

// Function to find data related to short url by short url id
func FindByShortUrl(redisClient *redis.Client, ctx context.Context, payload *models.Payload) error {

	shortUrl := payload.ShortUrl
	shortUrlId := ""
	if shortUrl != "" {
		shortUrlId = path.Base(shortUrl)
	} else {
		shortUrlId = payload.ShortUrlId
	}

	log.Println("short url id ", shortUrlId)

	shortUrlKey := fmt.Sprintf(util.REDIS_KEY_TO_STORE_SHORT_URL_ID+":%s", shortUrlId)

	redisKey, err := redisClient.Get(ctx, shortUrlKey).Result()
	if err != nil {
		log.Println("short url mapping key not found", err.Error())
		return err
	}

	result, err := redisClient.Get(ctx, redisKey).Result()
	if err != nil {
		log.Printf("data not found for key %v", redisKey)
		return err
	}

	// Convert the result to a byte slice before unmarshalling
	err = json.Unmarshal([]byte(result), payload)
	if err != nil {
		log.Println("error unmarshalling result into payload", err.Error())
		return err
	}

	return nil

}

// Function to save data to redis database for short url generator application
func SaveData(redisClient *redis.Client, ctx context.Context, payload models.Payload) error {

	// serialize the struct `Payload` to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error serializing Payload at saveData() ", err.Error())
		return err
	}

	// our redis key will be base10
	redisKey := fmt.Sprintf(util.REDIS_KEY_TO_STORE_SHORT_URL+":%d", payload.KeyId)

	// maintain secondary index to track short url id
	shortUrlKey := fmt.Sprintf(util.REDIS_KEY_TO_STORE_SHORT_URL_ID+":%s", payload.ShortUrlId)

	// maintain secondary index to track long url
	longUrlKey := fmt.Sprintf(util.REDIS_KEY_TO_STORE_LONG_URL_ID+":%s", payload.LongUrl)

	// store entire `payloadJSON` under a single key `redisKey` into
	// redis database. Zero expiration means the key has no expiration time
	err = redisClient.Set(ctx, redisKey, payloadJSON, 0).Err()
	if err != nil {
		log.Println("failed to store short url details:", err.Error())
		return err
	} else {

		err = redisClient.Set(ctx, shortUrlKey, redisKey, 0).Err()
		if err != nil {
			log.Println("failed to store short url mapping:", err.Error())
			return err
		}
		err = redisClient.Set(ctx, longUrlKey, redisKey, 0).Err()
		if err != nil {
			log.Println("failed to store long url mapping:", err.Error())
			return err
		}
		log.Println("data saved into redis")
	}

	return nil

}

// Function to delete data related to short url
func DeleteData(redisClient *redis.Client, ctx context.Context, payload *models.Payload) error {

	shortUrl := payload.ShortUrl
	shortUrlId := path.Base(shortUrl)

	longUrlKey := fmt.Sprintf(util.REDIS_KEY_TO_STORE_LONG_URL_ID+":%s", payload.LongUrl)
	shortUrlKey := fmt.Sprintf(util.REDIS_KEY_TO_STORE_SHORT_URL_ID+":%s", shortUrlId)
	redisKey := fmt.Sprintf(util.REDIS_KEY_TO_STORE_SHORT_URL+":%d", payload.KeyId)

	result, err := redisClient.Del(ctx, longUrlKey, shortUrlKey, redisKey).Result()
	if err != nil {
		log.Println("failed to delete record for long url:", err.Error())
		return err
	}

	log.Printf(" Total count of keys removed %v", result)

	return nil

}
