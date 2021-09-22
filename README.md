# bitdotio-golang-example
Example of using golang with bit.io

bit.io works with any golang PostgreSQL library; in this example we use `pgx`. 

Setup:
`go get github.com/jackc/pgx/v4`

Build:
`go build`

Grab your bit.io API key:
* Log into bit.io
* Click the green "Connect" button in the upper left
* Click on your API key to copy it

Run:
`BITDOTIO_APIKEY=<your api key> ./bitdotio-golang-example`

