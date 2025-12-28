package abuseipdb

import (
	"encoding/json"
	"errors"
	"io"
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

type abuseipdbResponse struct {
	Data Entry `json:"data"`
}

type abuseipdbBlacklistResponse struct {
	Data []Entry `json:"data"`
}

func crunchResponse(jStr string) Entry {
	var cont abuseipdbResponse
	json.Unmarshal([]byte(jStr), &cont)
	return cont.Data
}

// Blacklist retrieves a list of black listed IP addresses
func Blacklist(apikey string, confidenceMinimum int, limit int) ([]Entry, error) {
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
		return nil, err
	}
	defer resp.Body.Close()

	page, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cont abuseipdbBlacklistResponse
	json.Unmarshal(page, &cont)

	return cont.Data, nil
}

// CheckIP checks an IP address against the AbuseIPDB database
func CheckIP(apikey string, ipAddress string) (Entry, error) {
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
	query.Add("ipAddress", ipAddress)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return emptyEntry, err
	}
	defer resp.Body.Close()

	page, err := io.ReadAll(resp.Body)
	if err != nil {
		return emptyEntry, err
	}

	ipEntry := crunchResponse(string(page))
	return ipEntry, nil
}
