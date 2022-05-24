package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	apiKey, ok := os.LookupEnv("BITDOTIO_APIKEY") // "Password" from connect menu
	if !ok {
		panic("BITDOTIO_APIKEY not set in environment")
	}
	userName := "<bit.io username>"
	dbName := "dliden/2020_Census_Reapportionment"
	connectUrl := fmt.Sprintf("postgres://%s:%s@db.bit.io:5432/%s", userName, apiKey, dbName)
	conn, err := pgx.Connect(context.Background(), connectUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Let's see how the census population of Nevada has changed
	sqlQuery := fmt.Sprintf(`SELECT "Year", "Resident Population" FROM "dliden/2020_Census_Reapportionment"."Historical Apportionment" WHERE "Name" = 'Nevada';`)
	rows, err := conn.Query(context.Background(), sqlQuery)

	if err != nil {
		fmt.Println("query error")
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
        var Year int
        var Population int

		err = rows.Scan(&Year, &Population)

		if err != nil {
			panic(err)
		}
		fmt.Println(Year, Population)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
