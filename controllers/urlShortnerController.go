package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/KumarVariable/go-for-url-shortner/models"
	"github.com/KumarVariable/go-for-url-shortner/util"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

// Route handler to get existing short url for given
// longUrl from redis database
func GetShortUrl(redisClient *redis.Client) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Content-Type", "application/json")

		errResp := ErrorResponse{}

		var requestData models.Payload
		ctx := request.Context()

		longUrl := request.URL.Query().Get("longUrl")
		log.Printf("request received to get short url %+v ", longUrl)

		if longUrl == "" {

			errResp.ErrCode = http.StatusBadRequest
			errResp.ErrMsg = "Missing mandatory parameter: longUrl"

			log.Printf("Error to get short url: %v ", errResp)
			SendErrResponse(writer, errResp)
			return

		} else {
			requestData.LongUrl = longUrl
			err := FindByLongUrl(redisClient, ctx, &requestData)
			if err == nil {
				writer.WriteHeader(http.StatusOK)
				util.ResponseEncoder(writer, &requestData)
			} else {

				errResp.ErrCode = http.StatusNotFound
				errResp.ErrMsg = " No data found for given long url, Kindly create short url"
				SendErrResponse(writer, errResp)
			}
		}

	}
}

// Route handler to create short url with given longUrl
func CreateShortUrl(redisClient *redis.Client) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		errorCode := http.StatusInternalServerError
		errResp := ErrorResponse{}

		writer.Header().Set("Content-Type", "application/json")

		ctx := request.Context()

		requestData, err := util.RequestDecoder(request)
		if err != nil {
			errorMsg := "Error to decode request for create short url"

			errResp.ErrCode = errorCode
			errResp.ErrMsg = errorMsg

			SendErrResponse(writer, errResp)
			return
		}

		log.Printf("request received to create short url %+v ", requestData)
		// close request body when the function is returned
		defer request.Body.Close()

		if requestData.LongUrl == "" {
			errResp.ErrCode = http.StatusBadRequest
			errResp.ErrMsg = "Missing mandatory request parameter: longUrl"

			log.Printf("Error invalid request %v ", errResp)
			SendErrResponse(writer, errResp)
			return
		}

		if !IsShortUrlExistsForLongUrl(redisClient, ctx, &requestData) {

			writer.WriteHeader(http.StatusCreated)

			//increment counter
			counter := getUniqueCounterValue(redisClient, ctx)
			if counter == 0 {
				errResp.ErrCode = http.StatusInternalServerError
				errResp.ErrMsg = "Internal server error"

				log.Printf("Error to get unique counter %v ", errResp)
				SendErrResponse(writer, errResp)
				return
			}

			shortUrlId, base62String := generateShortUrlId(counter)
			if shortUrlId != 0 {

				shortUrl := "http://" + request.Host + "/" + base62String
				requestData.ShortUrl = shortUrl
				requestData.KeyId = shortUrlId
				requestData.ShortUrlId = base62String

				err = SaveData(redisClient, ctx, requestData)
				if err == nil {
					util.ResponseEncoder(writer, &requestData)

				} else {
					errResp.ErrCode = http.StatusInternalServerError
					errResp.ErrMsg = "Internal server error"

					log.Printf("Error could not save data into redis %v ", err)
					SendErrResponse(writer, errResp)
					return
				}

			} else {
				errResp.ErrCode = http.StatusInternalServerError
				errResp.ErrMsg = "Could not create short Url, Try after sometime "

				log.Printf("Error could not create short url %v ", err)
				SendErrResponse(writer, errResp)
				return
			}

		} else {
			writer.WriteHeader(http.StatusCreated)
			FindByLongUrl(redisClient, ctx, &requestData)
			util.ResponseEncoder(writer, &requestData)
		}

	}
}

// Route handler to update short url Id for existing long url
func UpdateShortUrl(redisClient *redis.Client) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()
		errResp := ErrorResponse{}

		requestData, err := util.RequestDecoder(request)
		if err != nil {
			errResp.ErrCode = http.StatusInternalServerError
			errResp.ErrMsg = "Error to decode delete short url request"

			log.Printf("Error decoding update short url request: %v ", err)

			SendErrResponse(writer, errResp)
			return
		}

		// close request body when function is returned
		defer request.Body.Close()
		log.Printf("request received to update short url %+v ", requestData)

		writer.Header().Set("Content-Type", "application/json")

		customErrCode := util.HasValidRequestParams(&requestData)
		if customErrCode == util.MISSING_REQUEST_PARAMS {
			errResp.ErrCode = http.StatusBadRequest
			errResp.ErrMsg = util.GetCustomErrorMsgs(customErrCode)

			log.Printf("Error to update short url: %v ", errResp)
			SendErrResponse(writer, errResp)
			return

		} else if customErrCode == util.SHORT_URL_PARAM_FOUND {

			errResp.ErrCode = http.StatusBadRequest
			errResp.ErrMsg += util.GetCustomErrorMsgs(customErrCode) + " longUrl"

			log.Printf("Error to update short url: %v ", errResp)
			SendErrResponse(writer, errResp)
			return

		} else {

			// Find existing data
			if IsShortUrlExistsForLongUrl(redisClient, ctx, &requestData) {

				writer.WriteHeader(http.StatusOK)
				FindByLongUrl(redisClient, ctx, &requestData)

				// Delete existing data
				DeleteData(redisClient, ctx, &requestData)

				// Create new short url
				counter := getUniqueCounterValue(redisClient, ctx)
				if counter == 0 {
					errResp.ErrCode = http.StatusInternalServerError
					errResp.ErrMsg = "Internal server error"

					log.Printf("Error to get unique counter %v ", errResp)
					SendErrResponse(writer, errResp)
					return
				}

				shortUrlId, base62String := generateShortUrlId(counter)
				if shortUrlId != 0 {

					shortUrl := "http://" + request.Host + "/" + base62String
					requestData.ShortUrl = shortUrl
					requestData.KeyId = shortUrlId
					requestData.ShortUrlId = base62String

					err = SaveData(redisClient, ctx, requestData)
					if err == nil {
						util.ResponseEncoder(writer, &requestData)

					} else {
						errResp.ErrCode = http.StatusInternalServerError
						errResp.ErrMsg = "Internal server error"

						log.Printf("Error could not save data into redis %v ", err)
						SendErrResponse(writer, errResp)
						return
					}

				} else {
					errResp.ErrCode = http.StatusInternalServerError
					errResp.ErrMsg = "Could not update short Url, Try after sometime "

					log.Printf("Error could not update short url %v ", err)
					SendErrResponse(writer, errResp)
					return
				}

			} else {
				errResp.ErrCode = http.StatusNotFound
				errResp.ErrMsg = "No record found to update.Please create short url "
				log.Printf("Error could not update short url %v ", err)
				SendErrResponse(writer, errResp)
				return
			}
		}
	}
}

// Route handler to delete existing short url for long url
func DeleteShortUrl(redisClient *redis.Client) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()
		errResp := ErrorResponse{}
		requestData := models.Payload{}

		longUrl := request.URL.Query().Get("longUrl")
		log.Printf("request received to delete short url %+v ", longUrl)

		if longUrl == "" {
			errResp.ErrCode = http.StatusBadRequest
			errResp.ErrMsg = "Missing required parameter longUrl"
			log.Printf("Error to delete short url: %v ", errResp)
			SendErrResponse(writer, errResp)
			return
		} else {

			requestData.LongUrl = longUrl
			err := FindByLongUrl(redisClient, ctx, &requestData)
			if err == nil {

				err = DeleteData(redisClient, ctx, &requestData)
				if err == nil {
					writer.WriteHeader(http.StatusNoContent)
					return

				} else {
					errResp.ErrCode = http.StatusInternalServerError
					errResp.ErrMsg = " Error to delete record "
					SendErrResponse(writer, errResp)
					return
				}

			} else {
				errResp.ErrCode = http.StatusNotFound
				errResp.ErrMsg = " No data found for given long url, Kindly create short url"
				SendErrResponse(writer, errResp)
				return
			}
		}

	}
}

// Route handler to create custom short url for long url
func CreateCustomShortUrl(redisClient *redis.Client) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		errorCode := http.StatusInternalServerError
		errResp := ErrorResponse{}

		writer.Header().Set("Content-Type", "application/json")

		ctx := request.Context()

		requestData, err := util.RequestDecoder(request)
		if err != nil {
			errorMsg := "Error to decode request for create custom short url"

			errResp.ErrCode = errorCode
			errResp.ErrMsg = errorMsg

			SendErrResponse(writer, errResp)
			return
		}

		log.Printf("request received to create custom short url %+v ", requestData)
		// close request body when the function is returned
		defer request.Body.Close()

		if requestData.LongUrl == "" {
			errResp.ErrCode = http.StatusBadRequest
			errResp.ErrMsg = "Missing mandatory request parameter: longUrl"

			log.Printf("Error invalid request %v ", errResp)
			SendErrResponse(writer, errResp)
			return
		}

		if !IsShortUrlExistsForLongUrl(redisClient, ctx, &requestData) {

			writer.WriteHeader(http.StatusCreated)

			//increment counter
			counter := getUniqueCounterValue(redisClient, ctx)
			if counter == 0 {
				errResp.ErrCode = http.StatusInternalServerError
				errResp.ErrMsg = "Internal server error"

				log.Printf("Error to get unique counter %v ", errResp)
				SendErrResponse(writer, errResp)
				return
			}

			shortUrlId, _ := generateShortUrlId(counter)
			base62String := requestData.ShortUrl
			if shortUrlId != 0 {

				shortUrl := "http://" + request.Host + "/" + base62String
				requestData.ShortUrl = shortUrl
				requestData.KeyId = shortUrlId
				requestData.ShortUrlId = base62String

				err = SaveData(redisClient, ctx, requestData)
				if err == nil {
					util.ResponseEncoder(writer, &requestData)

				} else {
					errResp.ErrCode = http.StatusInternalServerError
					errResp.ErrMsg = "Internal server error"

					log.Printf("Error could not save data into redis %v ", err)
					SendErrResponse(writer, errResp)
					return
				}

			} else {
				errResp.ErrCode = http.StatusInternalServerError
				errResp.ErrMsg = "Could not create custom short Url, Try after sometime "

				log.Printf("Error could not create custom short url %v ", err)
				SendErrResponse(writer, errResp)
				return
			}

		} else {
			writer.WriteHeader(http.StatusCreated)
			FindByLongUrl(redisClient, ctx, &requestData)
			util.ResponseEncoder(writer, &requestData)
		}

	}
}

// Route Handler to redirect short url to the known longUrl
func RedirectToOriginalUrl(redisClient *redis.Client) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()
		payload := models.Payload{}
		routeVariable := mux.Vars(request)
		// log.Println("route variables ::  ", routeVariable)

		shortUrlId := routeVariable["shortUrlID"]
		log.Println("route to be redirected to ::  ", shortUrlId)

		// close the request after function execution
		defer request.Body.Close()

		errResp := ErrorResponse{}

		if shortUrlId != "" {
			payload.ShortUrlId = shortUrlId
			err := FindByShortUrl(redisClient, ctx, &payload)
			if err != nil {
				errResp.ErrCode = http.StatusNotFound
				errResp.ErrMsg = "Redirect url does not exists "

				log.Printf("No long url mapped to short url id %v ", errResp)
				SendErrResponse(writer, errResp)
				return
			}
			http.Redirect(writer, request, payload.LongUrl, http.StatusMovedPermanently)
		}
	}
}

// Helper function to get unique counter id from redis database
func getUniqueCounterValue(redisClient *redis.Client, ctx context.Context) int64 {
	//increment counter
	counter := IncrementCounter(redisClient, ctx)
	return counter

}

// Helper function to get unique base62 string
func generateShortUrlId(counter int64) (int64, string) {

	shortUrlId := int64(0)

	// generate random string with base62 algorithm
	base62String := util.ConvertToBase62String(counter)
	if base62String != "0" {
		shortUrlId = util.ConvertToBase10Decimal(base62String)
	}
	return shortUrlId, base62String

}
