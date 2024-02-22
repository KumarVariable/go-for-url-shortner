package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/KumarVariable/go-for-url-shortner/models"
)

// constant that define `0-9` for first 10 values
// `a-z` and `A-Z` for the next 26 characters respectively
const base62Characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// To convert a decimal/integer value to base62
// 1.Divide the number by 62
// 2. Get the quotient and remainder
// 3. With the remainder find the corresponding charachter into
// base62Characters constant.
// Remember: base62 system starts its index at zero, the last character "Z"
// is at position sixty-one, and "1" is at position one.
func ConvertToBase62String(number int64) string {

	base62String := ""

	if number == 0 {
		return "0"
	}

	for number > 0 {

		// remainder operation
		remainder := number % 62
		base62String = string(base62Characters[remainder]) + base62String

		// division operation
		number = number / 62

	}
	return base62String
}

// Function to convert base62 to base10
// 1. Identify the position of each base62 character
// 2. Calculate the powers of 62 for each character based on its position.
// 3. Multiply each character by its corresponding power of 62
// 4. Add all the results to obtain decimal representation
func ConvertToBase10Decimal(base62String string) int64 {

	// final number
	var num int64 = 0

	// represents length of base-62 numeral system used for encoding i.e 62
	base := len(base62Characters)

	for index, char := range base62String {

		// the power of base 62 that will be applied to
		// current character's nummeric value
		power := len(base62String) - index - 1

		// find numeric value of current character, multiple with base raise to power
		// and add the summation to `num`
		num += int64(strings.Index(base62Characters, string(char))) * int64(pow(base, power))

	}

	return num

}

// Helper function to calculate the power of a number
func pow(base, exponent int) int {

	// set initial value to 1 because any number raise to
	// the power of 0 is 1
	result := 1

	// start loop as long as power is greater than baseLength
	for exponent > 0 {

		// keep on multiplying
		result = result * base

		// decrement the value of power by 1, to terminate the for loop
		exponent--

	}

	return result
}

// Function to decode the incoming json request data
// directly from stream.Also, efficient in case of large JSON payloads
func RequestDecoder(request *http.Request) (models.Payload, error) {

	var requestData models.Payload

	err := json.NewDecoder(request.Body).Decode(&requestData)
	if err != nil {
		log.Printf("Error decoding request: %v ", err)
		return requestData, err
	}
	return requestData, nil
}

// Function to encode the response data
func ResponseEncoder(writer io.Writer, requestData *models.Payload) error {

	err := json.NewEncoder(writer).Encode(requestData)
	if err != nil {
		log.Printf("Error encoding response: %v ", err)

		return err
	}
	return nil
}

// Helper function to validate the incoming request for
// url shortner data
func HasValidRequestParams(requestData *models.Payload) CustomErrorCodes {

	if requestData.LongUrl == "" && requestData.ShortUrl == "" {
		return MISSING_REQUEST_PARAMS
	} else if requestData.ShortUrl != "" {
		return SHORT_URL_PARAM_FOUND
	} else if requestData.LongUrl != "" {
		return LONG_URL_PARAM_FOUND
	}
	return VALID_REQUEST_PARAMS
}
