package middleware

import (
	"bytes"
	"log"
	"net/http"
)

// Represents custom ResponseWriter
type logResponseWriterStruct struct {
	http.ResponseWriter              // embed to inherit the standard ResponseWriter
	statusCode          int          // capture and store status code for logging purpose
	httpMethod          string       // capture and store request HTTP method
	responseBody        bytes.Buffer // capture and store response body
}

// Middleware to perform any request logging, header manipulation
// or Response Write hijacking before passing control to the next handler.
func InterceptRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// perform your application logic before serving any request

		log.Printf("<---> processing request for url and method: %+v %+v ", r.URL, r.Method)

		// pass call to next middleware or the final handler
		next.ServeHTTP(w, r)

	})
}

// Middleware to perform any logging before HTTP response write to ResponseWriter
func InterceptResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// wrap the original ResponseWriter to inspect the response
		// after the handler has written response
		logResponse := logResponseWriter(w, r.Method)

		// pass call to next middleware or the final handler.
		// Writes to our custom ResponseWriter
		next.ServeHTTP(logResponse, r)

		// After the handler has written the response, log the details
		log.Printf("<---> processed response for URL: %s, Method: %s, Status Code: %d", r.URL.Path, r.Method, logResponse.statusCode)

	})
}

// Function to create and return new instance of logResponseWriterStruct
// parameter: http.ResponseWriter which the function wraps.
func logResponseWriter(w http.ResponseWriter, methodName string) *logResponseWriterStruct {
	return &logResponseWriterStruct{
		ResponseWriter: w,
		httpMethod:     methodName,
		statusCode:     http.StatusOK, // default is 200
	}
}

// Function is an implementation of Write to capture the reponse body and
// forward it to the underlying response writer
func (logResponse *logResponseWriterStruct) Write(b []byte) (int, error) {

	// capture the body
	logResponse.responseBody.Write(b)

	// Forward the response body to the original writer
	return logResponse.ResponseWriter.Write(b)

}

// Function is an implementation of WriteHeader to capture the HTTP status and
// custom header to the original ResponseWriter
func (logResponse *logResponseWriterStruct) WriteHeader(statusCode int) {

	// capture the status code
	logResponse.statusCode = statusCode

	// set the custom header to the original writer
	logResponse.ResponseWriter.Header().Set("X-Powered-By", "sample/application")

	// write the status code to the original writer
	logResponse.ResponseWriter.WriteHeader(statusCode)

}

// Function to integrate basic CORS support into HTTP services.
// This is to ensure that our application can handle requests
// from different origins securely, specifically tailored
// for web applications.
func SetUpCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		origin := request.Header.Get("origin")
		log.Println(" origin of the incoming HTTP request: ", origin)

		// allow for only known origin - a webapp application to create short url
		if origin == "http://127.0.0.1:3000" || origin == "http://localhost:3000" {
			log.Println(" inside if origin ")
			writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		// uncomment below if you want to allow for all origins
		// writer.Header().Set("Access-Control-Allow-Origin", "*")

		// specify the allowed HTTP methods
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		// indicate which headers can be included in the requests
		writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// handle pre-flight(OPTIONS methos) request
		// Helps to validate that the server accepts request from incoming Origin,
		// and the services are accessible from external code or domain.
		if request.Method == "OPTIONS" {
			log.Println("........  Pre flight options requests ......")
			writer.WriteHeader(http.StatusOK)
			return
		}

		// next handler to execute in the middleware chain.
		next.ServeHTTP(writer, request)

	})
}
