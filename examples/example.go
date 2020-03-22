package main

import (
	"fmt"
	"os"

	"github.com/richandersonio/go-abuseipdb/abuseipdb"
)

func main() {

	apikey := os.Getenv("ABUSEIPDB_API_KEY")
	resp, err := abuseipdb.CheckIP(apikey, "144.91.79.5")
	if err == nil {

		fmt.Println("error checking ip: ", err)
	}

	fmt.Println("IPAddress:", resp.IPAddress)
	fmt.Println("TotalReports:", resp.TotalReports)
	fmt.Println("NumberDistinctUsers:", resp.NumberDistinctUsers)
	fmt.Println("AbuseConfidenceScore:", resp.AbuseConfidenceScore)
	fmt.Println("IsPublic:", resp.IsPublic)
	fmt.Println("IPVersion:", resp.IPVersion)
	fmt.Println("IsWhiteListed:", resp.IsWhiteListed)
	fmt.Println("CountryCode:", resp.CountryCode)
	fmt.Println("UsageType:", resp.UsageType)
	fmt.Println("ISP:", resp.ISP)
	fmt.Println("Domain:", resp.Domain)
	fmt.Println("LastReportedAt:", resp.LastReportedAt)
}
