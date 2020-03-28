package main

import (
	"fmt"
	"os"

	"github.com/richandersonio/go-abuseipdb/abuseipdb"
)

func main() {
	apikey := os.Getenv("ABUSEIPDB_API_KEY")
	err := abuseipdb.Blacklist(apikey, 99)
	if err != nil {
		fmt.Println("error checking ip: ", err)
		return
	}
}
