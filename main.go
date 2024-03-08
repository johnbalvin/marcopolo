package main

import (
	"time"
)

// you need to play with the variables theads and tcp timouet number, change the values to the ones that adjust to your network enviroment
// increase the threads if you think you will network enviroment will be able to handle more threads
// decrease the tcp timeout if you think you will network enviroment is fast enough
// threadsNumber and tcpTimeout are values that depend on your enviroment, so make you sure make the right setup
// USE CABLE AND NOT WIFI, DIRECT CABLE WILL BE BETTER FOR THIS PROJECT
func main() {
	input := SecureState //this is the default host, fill the variables as you need it
	threadsNumber := 80
	tcpTimeout := time.Second
	sslTimeout := time.Second * 5 //it's ok this one to be bigger than the TCP timeout, at the end it will search an small portion of IPs so no need to worry
	asnPath := "./asn.csv"
	outputFolder := "./results"
	stopOnASNFound := true // it will stop once an IP is found on an ASN number, still will search on others ASN
	input.SearchByKeywords(stopOnASNFound, threadsNumber, tcpTimeout, sslTimeout, asnPath, outputFolder)
	//Uncoment in case you need an IP with valid SSL certificate
	//	stopOnSSlFound := true // it will stop once a valid SSL for that domain is found
	//	input.SearchBySSLCertificatesOnly(stopOnSSlFound, threadsNumber, sslTimeout, asnPath, outputFolder)
}
