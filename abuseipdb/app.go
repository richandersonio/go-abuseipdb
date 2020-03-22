package abuseipdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/* Structs for this response format:
   	jStr2 := `
       {
           "data":{
   			"totalReports": ["1111"],
   			"abuseConfidenceScore": ["2222"],
   			"ipAddress": "333"
   		}
       }
   	`
*/

// AbuseipdbEntry struct details the data known about an IP address
type Entry struct {
	TotalReports         int       `json:"totalReports"`
	AbuseConfidenceScore int       `json:"abuseConfidenceScore"`
	IPAddress            string    `json:"ipAddress"`
	IsPublic             bool      `json:"isPublic"`
	IPVersion            int       `json:"ipVersion"`
	IsWhiteListed        bool      `json:"isWhitelisted"`
	CountryCode          string    `json:"countryCode"`
	UsageType            string    `json:"usageType"`
	ISP                  string    `json:"isp"`
	Domain               string    `json:"domain"`
	NumberDistinctUsers  int       `json:"numDistinctUsers"`
	LastReportedAt       time.Time `json:"lastReportedAt"`
}

type abuseipdbReponse struct {
	Data Entry `json:"data"`
}

func crunchResponse(jStr string) (response Entry) {
	var cont abuseipdbReponse
	json.Unmarshal([]byte(jStr), &cont)
	//fmt.Printf("%+v\n", cont)
	return cont.Data
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

	//fmt.Println(string(page))
	var ipEntry Entry
	ipEntry = crunchResponse(string(page))
	return ipEntry, nil
}
