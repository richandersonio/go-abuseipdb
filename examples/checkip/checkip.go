package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/richandersonio/go-abuseipdb/abuseipdb"
)

func check(ip string) {
	apikey := os.Getenv("ABUSEIPDB_API_KEY")
	resp, err := abuseipdb.CheckIP(apikey, ip)
	if err != nil {
		fmt.Println("error checking ip: ", err)
		return
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

func main() {
	fmt.Println("Checking IP Address...")
	check("209.141.57.178")

	for {
		fmt.Println()
		fmt.Print("Enter another IP address to check (just hit enter to quit):")
		reader := bufio.NewReader(os.Stdin)
		ans, _ := reader.ReadString('\n')
		ans = strings.TrimRight(ans, "\n")
		if ans == "" {
			return
		}
		fmt.Println("Checking IP Address...")
		check(ans)
	}
}
