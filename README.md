Example go code for using the AbuseIPDB API to check if an IP address has been reported for malicious behavior
website:  https://www.abuseipdb.com

# Running tests

The tests depend on an API key.  You have to set this as an environment variable before running the tests:

export ABUSEIPDB_API_KEY=YOUR_API_KEY
go test

# TODO:

Create basic go docs - see https://github.com/googleapis/google-cloud-go/blob/master/datastore/doc.go
