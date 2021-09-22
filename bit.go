package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {
	apiKey, ok := os.LookupEnv("BITDOTIO_APIKEY")
	if !ok {
		panic("BITDOTIO_APIKEY not set in environment")
	}
	userName := "adam"
	repoName := "sensors"
	connectUrl := fmt.Sprintf("postgres://%s:%s@db.bit.io:5432/%s", userName, apiKey, repoName)
	conn, err := pgx.Connect(context.Background(), connectUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	fullRepoName := fmt.Sprintf(`"%s/%s"`, userName, repoName)
	// Let's see when we have bad AQI in San Francisco. 
	sqlQuery := fmt.Sprintf(`SELECT datetime, pm_10 FROM %s.measurements WHERE pm_10 >= 20.0 AND pm_10 <= 1000 ORDER BY datetime;`, fullRepoName)
	rows, err := conn.Query(context.Background(), sqlQuery)

	if err != nil {
		fmt.Println("query error")
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var datetime time.Time
		var pm_10 float32

		err = rows.Scan(&datetime, &pm_10)

		if err != nil {
			panic(err)
		}
		fmt.Println(datetime, pm_10)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
