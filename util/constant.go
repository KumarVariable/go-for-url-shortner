package util

import "time"

// global variable to hold the server start time
var SERVER_STARTED_AT time.Time

// intial unique counter value in redis database
// This unique id will be used to create unique short url
const INITIAL_COUNTER_VALUE int = 100000000000

// CustomErrorCodes holds unique identifiers for
// different types of errors that can occur.
type CustomErrorCodes int

const (
	VALID_REQUEST_PARAMS   CustomErrorCodes = iota + 1000 // Represents a valid request with no errors.
	MISSING_REQUEST_PARAMS                                // Represents a required parameter is missing from the request.
	LONG_URL_PARAM_FOUND                                  // Represents that the request contains the parameter for a long URL.
	SHORT_URL_PARAM_FOUND                                 // Represents that the request contains the parameter for a long URL.

	// Declare more error codes as per need
)

// errorMessages: To map our CustomErrorCodes to human-readable error messages.
var errorMessages = map[CustomErrorCodes]string{
	VALID_REQUEST_PARAMS:   "Valid Request",
	MISSING_REQUEST_PARAMS: "Missing required parameter longUrl or shortUrl",
	LONG_URL_PARAM_FOUND:   "Missing request parameter",
	SHORT_URL_PARAM_FOUND:  "Missing request parameter",

	// Declare more error codes as per need
}

// Function to return a human-readable error message for a given error code.
func GetCustomErrorMsgs(errCode CustomErrorCodes) string {

	errMsg := errorMessages[errCode]

	// If error code is not recognized
	if errMsg == "" {
		return "UNKNOWN_ERROR"
	}
	return errMsg // Return the matching error message.

}
