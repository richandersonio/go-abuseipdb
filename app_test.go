package abuseipdb

import (
	"fmt"
	"os"
	"testing"
)

var apikey string

func TestCheckIPNoIPAddressErrors(t *testing.T) {
	// general an error by passing no IP address
	apikey = os.Getenv("ABUSEIPDB_API_KEY")
	_, err := CheckIP("", "")
	if err == nil {
		t.Fatal("CheckIP should have returned an error")
	}
}
func TestCheckIPValid(t *testing.T) {
	// check a bad IP address with a known confidence score of 100
	apikey = os.Getenv("ABUSEIPDB_API_KEY")
	resp, err := CheckIP(apikey, "5.101.0.209")
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
