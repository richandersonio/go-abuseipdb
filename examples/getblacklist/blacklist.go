package main

import (
	"fmt"
	"os"

	"github.com/richandersonio/go-abuseipdb/abuseipdb"
)

func main() {
	apikey := os.Getenv("ABUSEIPDB_API_KEY")

	// apikey string, confidenceMinimum int, limit int
	var response []Entry
	var err error

	err, response = abuseipdb.Blacklist(apikey, 99, 100)
	if err != nil {
		fmt.Println("error checking ip: ", err)
		return
	}
}
