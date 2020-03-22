# Welcome

This package is a small wrapper for some of the APIs provided by AbuseIPDB API (https://www.abuseipdb.com)
It can be used for checking bad actors / bad IPs

## Example code

### seting up the api key

The example code and tests depend on an API key. Once you've registered and got this from the abuseipdb.com website  you have an environment variable the code uses:

```bash
export ABUSEIPDB_API_KEY=YOUR_API_KEY
```

### Simple example

You can the example code as follows:

```bash
cd examples
go get github.com/richandersonio/go-abuseipdb/abuseipdb
go run example.go 
```

This example provides console output for a known dodgy IP address:

```bash
IPAddress: 144.91.79.5
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

## Running tests

The tests depends on an API key.  You have to set this as an environment variable before running the tests:

```bash
export ABUSEIPDB_API_KEY=YOUR_API_KEY
go test -v
```

## TODO:

Create basic go docs - see https://github.com/googleapis/google-cloud-go/blob/master/datastore/doc.go
