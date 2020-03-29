package abuseipdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Entry struct details the data known about an IP address
type Entry struct {
	TotalReports         int       `json:"totalReports,omitempty"`
	AbuseConfidenceScore int       `json:"abuseConfidenceScore,omitempty"`
	IPAddress            string    `json:"ipAddress,omitempty"`
	IsPublic             bool      `json:"isPublic,omitempty"`
	IPVersion            int       `json:"ipVersion,omitempty"`
	IsWhiteListed        bool      `json:"isWhitelisted,omitempty"`
	CountryCode          string    `json:"countryCode,omitempty"`
	UsageType            string    `json:"usageType,omitempty"`
	ISP                  string    `json:"isp,omitempty"`
	Domain               string    `json:"domain,omitempty"`
	NumberDistinctUsers  int       `json:"numDistinctUsers,omitempty"`
	LastReportedAt       time.Time `json:"lastReportedAt,omitempty"`
}

type abuseipdbReponse struct {
	Data Entry `json:"data"`
}

type abuseipdbBlacklistReponse struct {
	Data []Entry `json:"data"`
}

func crunchResponse(jStr string) (response Entry) {
	var cont abuseipdbReponse
	json.Unmarshal([]byte(jStr), &cont)
	return cont.Data
}

// Blacklist retrieves a list of black listed IP addresess
func Blacklist(apikey string, confidenceMinimum int, limit int) (response []Entry, returnError error) {

	if confidenceMinimum < 25 {
		return nil, errors.New("confidenceMinimum must be in the range 25-100")
	}

	if confidenceMinimum > 100 {
		return nil, errors.New("confidenceMinimum must be in the range 25-100")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.abuseipdb.com/api/v2/blacklist", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Key", apikey)
	req.Header.Add("Accept", "application/json")

	query := req.URL.Query()
	query.Add("confidenceMinimum", strconv.Itoa(confidenceMinimum))
	query.Add("limit", strconv.Itoa(limit))

	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error calling api: ", err)
		return nil, err
	}

	page, e := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if e != nil {
		log.Fatal(e)
		return nil, err
	}

	var cont abuseipdbBlacklistReponse
	json.Unmarshal([]byte(page), &cont)
	fmt.Println(len(cont.Data))
	return cont.Data, nil
}

// CheckIP checks an IP address against the AbuseIPDB database
func CheckIP(apikey string, ipAddress string) (response Entry, returnError error) {
	var emptyEntry Entry
	if ipAddress == "" {
		return emptyEntry, errors.New("no ipAddress specified")
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.abuseipdb.com/api/v2/check", nil)
	if err != nil {
		return emptyEntry, err
	}

	req.Header.Add("Key", apikey)
	req.Header.Add("Accept", "application/json")

	query := req.URL.Query()
	query.Add("maxAgeInDays", "90")
	//query.Add("verbose", "")
	query.Add("ipAddress", ipAddress)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error calling api: ", err)
		return emptyEntry, err
	}

	page, e := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if e != nil {
		log.Fatal(e)
		return emptyEntry, err
	}

	var ipEntry Entry
	ipEntry = crunchResponse(string(page))
	return ipEntry, nil
}
