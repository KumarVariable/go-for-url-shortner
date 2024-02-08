package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/KumarVariable/go-for-url-shortner/models"
)

// Route handler to get short url for long url
func GetShortUrl(writer http.ResponseWriter, request *http.Request) {

	var requestData models.UrlShortnerData

	// close request body when function is returned
	defer request.Body.Close()

	// read the request body (JSON) data directly from stream
	// efficient for large JSON payloads
	err := json.NewDecoder(request.Body).Decode(&requestData)
	if err != nil {
		errorCode := http.StatusInternalServerError
		log.Printf("Error decoding get short url request: %v ", err)
		SendServerErrResponse(writer, err.Error(), errorCode)
		return
	}

	log.Printf("request received to get short url %+v ", requestData)

	isValidReq := IsValidRequest(requestData.LongUrl, writer)
	if !isValidReq {
		errorCode := http.StatusBadRequest
		errorMsg := "Mandatory request parameter LongUrl missing"
		log.Printf("Error to get short url: %v ", errorMsg)
		SendBadRequestResponse(writer, errorMsg, errorCode)
		return
	}

	requestData.Key = "234"
	requestData.ShortUrl = "existing.googl.co"

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	// encode and send the response data
	err = json.NewEncoder(writer).Encode(requestData)
	if err != nil {
		errorCode := http.StatusInternalServerError
		log.Printf("Error encoding get short url response: %v", err)
		SendServerErrResponse(writer, err.Error(), errorCode)
		return
	}

}

// Route handler to create short url from incoming long url
func CreateShortUrl(writer http.ResponseWriter, request *http.Request) {

	var requestData models.UrlShortnerData

	// close request body when the function is returned
	defer request.Body.Close()

	// Decode the incoming json request data directly from stream
	// efficient for large JSON payloads
	err := json.NewDecoder(request.Body).Decode(&requestData)
	if err != nil {
		errorCode := http.StatusInternalServerError
		log.Printf("Error decoding request: %v ", err)
		SendServerErrResponse(writer, err.Error(), errorCode)
		return
	}

	log.Printf("request received to create short url %+v ", requestData)

	isValidReq := IsValidRequest(requestData.LongUrl, writer)
	if !isValidReq {
		errorCode := http.StatusBadRequest
		errorMsg := "Mandatory request parameter LongUrl missing"
		log.Printf("Error decoding create short url request: %v ", errorMsg)
		SendBadRequestResponse(writer, errorMsg, errorCode)
		return
	}

	requestData.Key = "234"
	requestData.ShortUrl = "googl.co"

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	// Encode and send the response data
	err = json.NewEncoder(writer).Encode(requestData)
	if err != nil {
		errorCode := http.StatusInternalServerError
		log.Printf("Error encoding create short url response: %v", err)
		SendServerErrResponse(writer, err.Error(), errorCode)
		return
	}

}

// Route handler to update existing short url for long url
func UpdateShortUrl(writer http.ResponseWriter, request *http.Request) {
	var requestData models.UrlShortnerData

	// close request body after function is returned
	defer request.Body.Close()

	// Decode the incoming json request data directly from stream
	// efficient for handling of large JSON payload
	err := json.NewDecoder(request.Body).Decode(&requestData)
	if err != nil {
		errorCode := http.StatusInternalServerError
		log.Printf("Error decoding update short url request: %v ", err)
		SendServerErrResponse(writer, err.Error(), errorCode)
		return
	}

	log.Printf("request received to update short url %+v ", requestData)

	isValidReq := IsValidRequest(requestData.LongUrl, writer)
	if !isValidReq {
		errorCode := http.StatusBadRequest
		errorMsg := "Mandatory request parameter LongUrl missing"
		log.Printf("Error decoding update short url request: %v ", errorMsg)
		SendBadRequestResponse(writer, errorMsg, errorCode)
		return
	}

	requestData.Key = "2341"
	requestData.ShortUrl = "updated.googl.co"

	writer.WriteHeader(http.StatusCreated)

	// write data to the connection as part of HTTP reply
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		errorCode := http.StatusInternalServerError
		log.Printf("Error to marshal update short url response: %v", err)
		SendServerErrResponse(writer, err.Error(), errorCode)
		return
	}

	// sends JSON data as the response body.
	_, err = writer.Write(jsonData)
	if err != nil {
		errorCode := http.StatusInternalServerError
		log.Printf("Error to write update short url response: %v", err)
		SendServerErrResponse(writer, err.Error(), errorCode)
		return
	}

}

// Route handler to delete existing short url for long url
// Implementation is done using `json.Unmarshal` and `json.Marshal`
// for learning purpose only
func DeleteShortUrl(writer http.ResponseWriter, request *http.Request) {

	// read the entire request body into a byte slice.
	// read entire body into memory, not much efficient
	// for large request body
	byteData, err := io.ReadAll(request.Body)
	if err != nil {
		errorCode := http.StatusInternalServerError
		log.Printf("Error to read request body for delete short url: %v", err)
		SendServerErrResponse(writer, err.Error(), errorCode)
		return
	}

	var reqData models.ShortUrlRequest

	// unmarshall - decode JSON data
	err = json.Unmarshal(byteData, &reqData)
	if err != nil {
		errorCode := http.StatusInternalServerError
		errMsg := "Error parsing JSON request for delete short url"
		log.Printf("Error parsing JSON request for delete short url: %v", err)
		SendServerErrResponse(writer, errMsg, errorCode)
		return
	}

	response := models.ShortUrlRequest{
		ShortUrl: "",
		Status:   "DELETED",
		Message:  "Short URL deleted successfully",
	}

	// convert response struct into JSON byte slice
	// before writing to the response writer
	responseBody, err := json.Marshal(response)
	if err != nil {
		errorCode := http.StatusInternalServerError
		errMsg := "Error parsing JSON response for delete short url"
		log.Printf("Error parsing JSON response for delete short url: %v", err)
		SendServerErrResponse(writer, errMsg, errorCode)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseBody)

}
