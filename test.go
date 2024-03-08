package main

import (
	"log"
	"os"
	"time"
)

func Test() {
	input := Kroger
	threadsKeywords := 100
	threadsSSLVerification := 150
	tcpTimeout := time.Second
	sslTimeout := time.Second * 5
	asnPath := "../asn.csv"
	outputFolder := "./results"
	stopOnASNFound := true
	input.SearchByKeywords(stopOnASNFound, threadsKeywords, threadsSSLVerification, tcpTimeout, sslTimeout, asnPath, outputFolder)
}
func test01() {
	input := Mouser
	ip := "12.5.163.52"
	input.setInputs()
	tcpTimeout := time.Second
	found, number, buffer, err := input.Marco(ip, tcpTimeout)
	os.WriteFile("./rawdata2.html", buffer, 0644)
	if err != nil {
		log.Println("err: ", err)
		return
	}
	os.WriteFile("./rawdata.html", buffer, 0644)
	log.Printf("found: %+v, number: %d\n", found, number)
}
