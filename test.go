package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func Test() {
	input := Kroger
	threadsNumber := 80
	tcpTimeout := time.Second
	sslTimeout := time.Second * 5
	asnPath := "../asn.csv"
	outputFolder := "./results"
	stopOnASNFound := true
	input.SearchByKeywords(stopOnASNFound, threadsNumber, tcpTimeout, sslTimeout, asnPath, outputFolder)
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

type resultBefore struct {
	Domain string
	IPBulk []IPBulk
}
type IPBulk struct {
	AsnID   string
	AsnName string
	IPs     []string
}

func test2() {
	fName := "./results/starngage/starngage.json"
	processSSLName(fName)
}
func processSSLName(fileName string) {
	var finalResult resultBefore
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("err: ", err)
	}
	json.NewDecoder(f).Decode(&finalResult)
	final2 := Result{
		Domain: finalResult.Domain,
	}
	for _, bulk := range finalResult.IPBulk {
		asnFound := AsnFound{
			ID:   bulk.AsnID,
			Name: bulk.AsnName,
		}
		for _, ipDataBefore := range bulk.IPs {
			ipData := IP{
				IP: ipDataBefore,
			}
			asnFound.IPs = append(asnFound.IPs, ipData)
		}
		final2.AsnsFound = append(final2.AsnsFound, asnFound)
	}
	final2.SetSSL(final2.Domain, time.Second*3)
	f2, _ := os.Create(fileName)
	json.NewEncoder(f2).Encode(final2)
}
