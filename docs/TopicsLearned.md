
[comment]: <> (This file is to log the topics leraned in `Go` programming on each example)

### Overview of My Learning 📘 ###

## Client and Server Integration in Go lang using [http package](https://pkg.go.dev/net/http). ##

<i>Topics Covered</i>: Implement REST services using [gorilla/mux package](https://pkg.go.dev/github.com/gorilla/mux#section-readme).

<ol>

- `gorilla/mux` package is used for implementation of REST API, HTTP request and response.
- Exposed sample REST APIs `[GET, POST, PUT, DELETE]`.
- Start `http` server with config parameters `ReadTimeout` and `WriteTimeout` 
- The REST APIs are implemented using different encoding/decoding approaches available in `Go` packages - `json.Marshal`, `json.UnMarshall`, `json.NewDecoder().Decode()`, `json.NewEncoder().Encode()`
- Use `Go` packages - `io`, `encoding/json`

</ol>

## Introduce and implement request / response interceptor [Middleware](https://pkg.go.dev/golang.org/x/pkgsite/internal/middleware) in `Go` application. ##

<i>Topics Covered</i>: Add middleware (interceptor ) in a `Go` application.

<ol>

- A middleware allows to intercept any request or response in `Go`. This concept is kind of
similar functionality of `Filters` in `Java` programming language. 
- An example to demonstrate the implementation of `Request`, `Response` Interceptors in REST application .
- Write your own custom `http.ResponseWriter` in current REST application. This response writer is
use to log response data or details before writing to the standard Response Writer. 
- Pass calls to multiple middlewares or we can say `Filter Chaining` concept in `Go` application

</ol>

## Integrate Redis database using [go-redis](https://github.com/redis/go-redis) package into `golang` application. ##

<i>Topics Covered</i>: Integrate `Redis Database` in a `Go` application.

<ol>

- Redis is in-memory data storage which is used as database. 
- Use package [go-redis](https://github.com/redis/go-redis) to integrate in our golang application.
- Implement REST services to Add, Get key(s) from locally running redis database.

</ol>

## Incorporate [CORS middleware](https://pkg.go.dev/github.com/rs/cors) into the Golang application's HTTP REST services. ##

<i>Topics Covered</i>: Set-Up `CORS Policy` in URL shortner application.

<ol>

- Add new Middleware to allow external applications to invoke the REST API from cross-origin request. 
- Configure `Cross Origin Resource Sharing (CORS)` to allow request from external domain or application.

</ol>

## Develop a Base62 encoding system to generate distinct short URLs for each unique integer input. ##

## Implement a Base10 decoding mechanism to extract the original integer value from a given short URL identifier. ##