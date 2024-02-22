// This source file contains functions to test redis database
// operations.

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/KumarVariable/go-for-url-shortner/models"
	"github.com/KumarVariable/go-for-url-shortner/util"
	"github.com/redis/go-redis/v9"
)

// Route handler to get Server Uptime
func PingTest(w http.ResponseWriter, r *http.Request) {

	log.Println("request received for ping test ")

	serverUptime := util.GetServerUptime()
	uptimeString := util.FormatDuration(serverUptime)

	// write response to Response Writer.
	fmt.Fprintf(w, " server is running since : "+uptimeString)

	log.Println("response returned for ping test ", uptimeString)
}

// Function to verify the connection with the locally running Redis instance.
func PingRedis(redisClient *redis.Client, ctx context.Context) {

	con, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Panic("Could not connect to redis server ", err.Error())
	}
	log.Println("connection to redis server: ", con)

}

// Router to add/store a key into the Redis database.
func StoreKeyValue(w http.ResponseWriter, r *http.Request) {

	log.Println("request received to add key-value in redis")

	payload := models.Payload{}

	ctx := context.Background()
	redisClient := models.RedisClient

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set key,values, expiration time.
	// expiration time as 0 is to set the key with no expiration time
	err = redisClient.Set(ctx, payload.LongUrl, payload.LongUrl, 0).Err()
	if err != nil {
		log.Println("could not add entry into redis database ", err.Error())
	} else {
		log.Println("entry created into redis database ")
		payload.Message = "entry created :" + strconv.Itoa(http.StatusCreated)
	}

	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)

	// encode and send the response data
	err = json.NewEncoder(w).Encode(&payload)
	if err != nil {
		errorCode := http.StatusInternalServerError
		log.Printf("Error encoding store key into redis database: %v", err)
		sendErrorResponse(w, err.Error(), errorCode)
		return
	}

}

// Router to retrieve all keys (multiple) from redis database
func GetAllKeysFromRedis(w http.ResponseWriter, r *http.Request) {

	var responseBody []byte
	var err error

	ctx := context.Background()
	redisClient := models.RedisClient

	// close the request after execution of function
	defer r.Body.Close()

	// Use the Keys method with the pattern "*" to get all keys
	keys, err := redisClient.Keys(ctx, "*").Result()
	if err != nil {
		log.Println("error caught to get all keys stored in redis ", err.Error())
	}

	var data []models.Payload

	if len(keys) > 0 {

		for i := 0; i < len(keys); i++ {

			redisData := models.Payload{}
			redisData.KeyName = keys[i]

			data = append(data, redisData)

		}

		responseBody, err = json.Marshal(data)
		if err != nil {
			errorCode := http.StatusInternalServerError
			log.Println("error caught at get all redis keys, could not marshall keys to json ", err.Error())
			SendServerErrResponse(w, err.Error(), errorCode)
		}

	} else {
		log.Println("no data found in redis server")

		errResponse := ErrorResponse{
			ErrCode: 404,
			ErrMsg:  "no data found in redis server",
		}

		responseBody, err = json.Marshal(errResponse)
		if err != nil {
			errorCode := http.StatusInternalServerError
			log.Println("error caught at get all redis keys, could not marshall keys to json ", err.Error())
			SendServerErrResponse(w, err.Error(), errorCode)
		}

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)

}

// Router to retrieve key for longUrl from redis database
func GetKeyFromRedis(w http.ResponseWriter, r *http.Request) {

	var reqData models.Payload

	ctx := context.Background()
	redisClient := models.RedisClient

	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		errorCode := http.StatusInternalServerError
		sendErrorResponse(w, err.Error(), errorCode)
		return
	}

	result, err := redisClient.Get(ctx, reqData.LongUrl).Result()
	if err != nil {
		errorCode := http.StatusNotFound
		errMsg := "Could not find key in the system"
		sendErrorResponse(w, errMsg, errorCode)
		return
	}

	log.Println("response :::  ", result)

	reqData.ShortUrl = result

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)

	// encode and send the response data
	err = json.NewEncoder(w).Encode(reqData)
	if err != nil {
		errorCode := http.StatusInternalServerError
		log.Printf("Error encoding get key from redis database: %v", err)
		sendErrorResponse(w, err.Error(), errorCode)
		return
	}

}
