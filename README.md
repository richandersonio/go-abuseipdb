# Welcome

This package is a small wrapper for some of the APIs provided by AbuseIPDB API (https://www.abuseipdb.com).

It can be used to determine if an IP address is known to be engaging in hacking attempts or other malicious behavior.

I use it on my site (https://richanderson.io) where I check every inbound request.


## Prerequisites

The package and examples require an API key from https://www.abuseipdb.com.  Go to their site to register and create one.

One you have your API key setup the environment variable "ABUSEIPDB_API_KEY":

```bash
export ABUSEIPDB_API_KEY=YOUR_API_KEY
```

## Examples

###  Check IP Example

This example calls the API and outputs basic information about the iP addresses entered in the console:

```bash
IPAddress: 209.141.57.178
TotalReports: 29
NumberDistinctUsers: 25
AbuseConfidenceScore: 100
IsPublic: true
IPVersion: 4
IsWhiteListed: false
CountryCode: DE
UsageType: Data Center/Web Hosting/Transit
ISP: Contabo GmbH
Domain: contabo.de
LastReportedAt: 2020-03-22 00:30:08 +0000 +0000
```

You can run the example locally on your machine:

```bash
cd examples/gocheck
go run .
```

####  Code example

'''go
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
'''

# History

## V2 created - 6/12/22

Changes:

Updated to 1.18
Fix bug in one of the examples
Module now uses github.com/richandersonio/go-abuseipdb/v2
