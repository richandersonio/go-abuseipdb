package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/richandersonio/go-abuseipdb/v2"
)

func main() {
	apikey := os.Getenv("ABUSEIPDB_API_KEY")

	// apikey string, confidenceMinimum int, limit int
	var response []abuseipdb.Entry
	var err error

	response, err = abuseipdb.Blacklist(apikey, 99, 100)
	if err != nil {
		fmt.Println("error checking ip: ", err)
		return
	}

	fmt.Println("entries downloaded in Blacklist " + strconv.Itoa((len(response))))
}
