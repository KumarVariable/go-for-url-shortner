# go-for-url-shortner

## Example ##
1) Client and Server Integration in Go lang using [http package](https://pkg.go.dev/net/http).
2) Implement REST services `[GET POST PUT DELETE]` in `Go` application using [gorilla/mux package](https://pkg.go.dev/github.com/gorilla/mux#section-readme).

| Endpoint Url                          |      HTTP Method     |  HTTP Response       |
|---------------------------------------|:--------------------:|---------------------:|
| localhost:9999/test                   |  GET                 | JSON Dummy Data      |
| localhost:9999/get-short-url          |  GET                 | JSON Dummy Data      |
| localhost:9999/create-short-url       |  POST                | JSON Dummy Data      |
| localhost:9999/update-short-url       |  PUT                 | JSON Dummy Data      |
| localhost:9999/delete-short-url       |  DELETE              | JSON Dummy Data      |

3) Introduce and implement request / response interceptor [Middleware](https://pkg.go.dev/golang.org/x/pkgsite/internal/middleware) in `Go` application.

4) Integrate Redis database using [go-redis](https://github.com/redis/go-redis) package into `golang` application.

| Endpoint Url                |    HTTP Method  |  HTTP Response      |
|-----------------------------|:---------------:|--------------------:|
| localhost:9999/get-key      |  GET            | JSON Dummy Data     |
| localhost:9999/get-all-keys |  GET            | JSON Dummy Data     |
| localhost:9999/add-key      |  POST           | JSON Dummy Data     |

Sample cURL calls:

Get Key

```text
curl --location --request GET 'http://localhost:9999/get-key' \
--header 'Content-Type: application/json' \
--data '{
    "longUrl" : "dbEntry2"
}'
```

Get All Keys

```text
curl --location 'http://localhost:9999/get-all-keys'
```

Add Key

```text
curl --location 'http://localhost:9999/add-key' \
--header 'Content-Type: application/json' \
--data '{
    "longUrl" : "dbEntry2"
}'
```
## Prerequisite ##

1. [Go](https://go.dev/doc/install)
2. [Redis](https://redis.io/docs/install/install-redis/)

<strong>NOTE: </strong>
This example is done based on legacy authentication method for redis database.<br/>
Edit the setting `requirepass` available in `redis.conf` file.<br/>
Keep Password property inside `server/redisConfig.go` file as empty, if the local running redis set up does not require any authentication.


## How to Run the sample `Go` Project ##

* Clone the project.
* Redis is installed and the service is up and running.
* Run command from terminal `go build .` to build the project. 
* Run command from terminal `go run .` to run the application or use your IDE to start `main.go` in `debug or non-debug` mode. 
* Open Browser, Run url `http://localhost:9999/test` to test the server is up and running.
* Change/modify the desired port number if you wish. Refer function `StartHttpServer()` at [Server Configuration](server/serverConfig.go).
* Use command `Ctrl+C` or `Stop` from IDE to shutdown server.


## Basic Go Commands ##
1. Command to check `Go` is installed on your machine.

```text
go version
```
2. Initialize/Create a `Go` project or module.Enable dependency tracking for your project.

```text
go mod init <my-project-name>
```
<strong>Note:</strong> Command will create `go.mod` file in project directory. The `go.mod` file provides information of:
- `name` of your project
- current `Go` version in use
- details of `libraries (project's dependencies)`

<u>Example to create module/project from Command line client:</u>

```text
mkdir go-for-url-shortner
cd go-for-url-shortner
go mod init github.com/KumarVariable/go-for-url-shortner
touch main.go
```

3. How to build a `Go` application ?

```text
go build main.go
```
- Command will compile packages and dependencies in a specific file (ex: main.go)
- Command generates an executable file in the current directory, (on Unix, it's typically named after the directory; on Windows, it will have an .exe suffix).

```text
go build .
```
- Here `.` (dot) represents a current directory.
- Command to compile the package that is in the current directory, along with any dependent files in that package.
- Command generates an executable file if the package is named as `main`, otherwise it compiles the package and produces a package archive.

4. How to run a `Go` application ?

```text
go run main.go
```
- Command to tell `Go` compiler to run and run `main.go` (a specific file). This command is helpful to quickly test a single file.

```text
go run .
```
- Here `.` (dot) represents a current directory.
- Command tells `Go` to compile and run the entire package in the current directory, not just a single file.
- Command also compiles multiple `.go` files which are part of package.

5. How to run executable created through command at point #3 ?

- Locate the executable and run `./go-for-url-shortner` from the terminal

```text
./go-for-url-shortner
```

6. How to tidy/clean up unnecessary dependencies from your `Go` project ?

```text
go mod tidy
```

