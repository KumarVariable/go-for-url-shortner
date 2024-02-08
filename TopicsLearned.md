
[comment]: <> (This file is to log the topics leraned in `Go` programming on each example)

### Examples Summary

- A sample hands-on learning to expose REST services using `Go` programming language.

<i>Topics Covered:</i>

<ol>

- `http` package in `Go` language.
- Concurrency in `Golang` using lightweight thread `goroutine`.
- Implement and expose REST endpoint using `Golang` application.
- Use `Go` packages - `time`, `fmt`, `log`

</ol>

- Implement REST services using `gorilla/mux` package

<i>Topics Covered:</i>

<ol>

- `gorilla/mux` package is used for implementation of REST API, HTTP request and response.
- Exposed sample REST APIs `[GET, POST, PUT, DELETE]`.
- Start `http` server with config parameters `ReadTimeout` and `WriteTimeout` 
- The REST APIs are implemented using different encoding/decoding approaches available in `Go` packages - `json.Marshal`, `json.UnMarshall`, `json.NewDecoder().Decode()`, `json.NewEncoder().Encode()`
- Use `Go` packages - `io`, `encoding/json`

</ol>

- Add middleware (interceptor ) in a `Go` application.

<i>Topics Covered:</i>

<ol>

- A middleware allows to intercept any request or response in `Go`. This concept is kind of
similar functionality of `Filters` in `Java` programming language. 
- An example to demonstrate the implementation of `Request`, `Response` Interceptors in REST application .
- Write your own custom `http.ResponseWriter` in current REST application. This response writer is
use to log response data or details before writing to the standard Response Writer. 
- Pass calls to multiple middlewares or we can say `Filter Chaining` concept in `Go` application

</ol>

- Integrate `Redis Database` in a `Go` application.

<i>Topics Covered:</i>

<ol>

- Redis is in-memory data storage which is used as database. 
- Use package [go-redis](https://github.com/redis/go-redis) to integrate in our golang application.
- Implement REST services to Add, Get key(s) from locally running redis database.

</ol>