package abuseipdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
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
	resp, err := CheckIP(apikey, "144.91.79.5")
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

func TestBlackList(t *testing.T) {
	return
	// check a bad IP address with a known confidence score of 100
	apikey = os.Getenv("ABUSEIPDB_API_KEY")

	list, err := Blacklist(apikey, 99, 100000)
	if err != nil {
		t.Fatalf("error checking ip: %s", err)
	}

	fmt.Printf(strconv.Itoa(len(list)))

	f, err := os.Create("blacklist.txt")
	if err != nil {
		return
	}
	defer f.Close()

	b, _ := json.Marshal(list)
	f.Write(b)
}

func TestReadBlacklistFile(t *testing.T) {
	// read file
	data, err := ioutil.ReadFile("blacklist10000.txt")
	if err != nil {
		fmt.Print(err)
		return
	}

	var blacklistEntries []Entry
	err = json.Unmarshal(data, &blacklistEntries)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("entries in the blacklist slice -> " + strconv.Itoa(len(blacklistEntries)))

	for _, blacklistEntry := range blacklistEntries {
		//		fmt.Println(i, blacklistEntry.IPAddress)
		if blacklistEntry.IPAddress == "196.43.155.209" {
			fmt.Printf("IP found in blacklist")
		}
	}
	// change the seed based on the time
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 10; i++ {
		ip := blacklistEntries[rand.Intn(len(blacklistEntries))].IPAddress
		//		ip = "144.91.79.5"
		fmt.Println("Random entry:", ip)
		resp, err := CheckIP(apikey, ip)
		if err != nil {
			fmt.Println("error checking ip: ", err)
			return
		}
		if resp.AbuseConfidenceScore != 100 {
			t.Fatalf(("confidence level not 100 for blackisted ip: %s"), ip)
		}
	}
}
