# Welcome

This package is a small wrapper for some of the APIs provided by AbuseIPDB API (https://www.abuseipdb.com).

It can be used to if an IP address is known to be engaging in hacking attempts or other malicious behavior.

I use it on my site (https://richanderson.io) where I check every inbound request.  If a bad IP is detected I send back a fake page.

## Example code

### seting up the api key

The example code and tests depend on an API key. Once you've registered for an account at https://www.abuseipdb.com, you have to setup an environment variable the code uses:

```bash
export ABUSEIPDB_API_KEY=YOUR_API_KEY
```

### Simple example

You can run the example code as follows:

```bash
cd examples
go get github.com/richandersonio/go-abuseipdb/abuseipdb
go run example.go 
```

This example calls the API and outputs the basic information about the iP address 144.91.79.5 to the console:

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

```bash
export ABUSEIPDB_API_KEY=YOUR_API_KEY
go test -v
```

## TODO:

Create basic go docs - see https://github.com/googleapis/google-cloud-go/blob/master/datastore/doc.go
