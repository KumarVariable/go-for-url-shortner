package controllers

import (
	"net/http"
)

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
