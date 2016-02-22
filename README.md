# goan - GoAnalytics
[![GoDoc](https://godoc.org/github.com/kevineaton/goan/lib?status.svg)](https://godoc.org/github.com/kevineaton/goan/lib)

A very basic analytics microservice that uses MongoDB or MySQL (soon) to store basic events that can be queried.

## Why?
I was bored and wanted to learn GoLang. I needed a project to work on so I created this.

## What does it do?
It uses a Gin-based RESTful API to take in event entries and store them in a data store. You can then pull
those entries out for analysis. Eventually, there will be filtering, searching, and a frontend. For now, it
is pretty bare bones and nowhere near ready for real use.

## What does it use?
Currently it uses:

- Gin for the API
- mgo For the Mongo driver
- Testify for unit testing

## Environment variables
You should set the following environment variables to setup the service:

- `GOAN_AUTHTOKEN`: The API auth token (passed as auth in query string or form posts) that allows access to the system. If not specified, 
a random token will be generated on each startup
- `GOAN_API_PORT`: The API port goan will listen on. If not specified, the default is 44889
- `GOAN_DBTYPE`: Currently only "mongo", although "mysql" support is coming soon. If not specified, the default is "mongo"
- `GOAN_DBURL`: The database connection string; if set, all other DB* variables will be ignored
- `GOAN_DBNAME`: The database name to connect to. If not specified, the default is goan
- `GOAN_DBHOST`: The database host to connect to. If not specified, the default is localhost
- `GOAN_DBPORT`: The database port to connect to. If not specified, the default is 27017
- `GOAN_DBUSER`: The database user to connect to. If not specified, the default is the database default
- `GOAN_DBPASSWORD`: The database password to connect to. If not specified, the default is the database default

## Running
To run, type:

`GIN_MODE=release GOAN_AUTHTOKEN=agoodencryptedtoken go run main.go`

## Testing
To test with a local mongodb installed (environment variables can be specified otherwise):

`cd ./lib`

`GIN_MODE=release go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html`
