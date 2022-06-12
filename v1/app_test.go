package abuseipdb

import (
	"fmt"
	"os"
	"testing"
)

func TestCheckIPNoIPAddressErrors(t *testing.T) {
	// generate an error by passing no key and IP address
	_, err := CheckIP("", "")
	if err == nil {
		t.Fatal("CheckIP should have returned an error")
	}
}

func TestKnownBadIP(t *testing.T) {
	// check a bad IP address with a known confidence score of 100
	apikey := os.Getenv("ABUSEIPDB_API_KEY")
	resp, err := CheckIP(apikey, "209.141.57.178")
	if err != nil {
		t.Fatalf("error checking ip: %s ", err)
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

func TestBlackList(t *testing.T) {
	// check a bad IP address with a known confidence score of 100
	apikey := os.Getenv("ABUSEIPDB_API_KEY")
	if apikey == "" {
		t.Fatal("ABUSEIPDB_API_KEY not set")
	}

	list, err := Blacklist(apikey, 99, 1000)
	if err != nil {
		t.Fatalf("error checking ip: %s", err)
	}

	if len(list) != 1000 {
		t.Fatalf("Did not return 1000 entries as expected")
	}

	return
}
