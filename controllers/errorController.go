package controllers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	ErrCode int
	ErrMsg  string
}

// Function to send error response to the client
func sendErrorResponse(writer http.ResponseWriter, errorMsg string, errorCode int) {
	http.Error(writer, errorMsg, errorCode)
}

func IsValidRequest(data string, writer http.ResponseWriter) bool {
	return data != ""
}

func SendBadRequestResponse(writer http.ResponseWriter, errorMsg string, errorCode int) {
	sendErrorResponse(writer, errorMsg, errorCode)
}

func SendServerErrResponse(writer http.ResponseWriter, errorMsg string, errorCode int) {
	sendErrorResponse(writer, errorMsg, errorCode)
}

// Function to handle when request method is not
// matched with any of known routes
func MethodNotAllowedHandler() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		errResponse := ErrorResponse{
			ErrCode: http.StatusMethodNotAllowed,
			ErrMsg:  "Method not allowed",
		}

		w.Header().Set("Content-Type", "application/json")

		// encode and send the response data
		json.NewEncoder(w).Encode(errResponse)

	})

}

// Function to handle when no route URL is matched
func RouteNotFoundHandler() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		errResponse := ErrorResponse{
			ErrCode: http.StatusNotFound,
			ErrMsg:  "Service not found",
		}

		w.Header().Set("Content-Type", "application/json")

		// encode and send the response data
		json.NewEncoder(w).Encode(errResponse)

	})

}
