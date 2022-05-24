# bitdotio-golang-example
Example of using golang with bit.io

bit.io works with any golang PostgreSQL library; in this example we use `pgx`. 

Setup:
`go get github.com/jackc/pgx/v4`

Build:
`go build`

Grab your bit.io API key:
* Log into bit.io and navigate to the database you would like to access
* Click the green "Connect"
* Click on your API key (the "Password" field in the connect menu) to copy it

Run:
`BITDOTIO_APIKEY=<your api key> ./bitdotio-golang-example` to save the key to your environment.

To run the program: `go run filename.go`

