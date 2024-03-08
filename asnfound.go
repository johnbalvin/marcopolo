package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func (input Input) SearchByKeywords(stopOnASNFound bool, threadsNumber int, tcpTimeout, sslTimeout time.Duration, asnPath, outputFolder string) {
	for _, priority := range input.Asn.PrioritiesNames {
		fmt.Printf("Priority name: *%s*\n", priority)
	}
	checkProgressEach := time.Minute
	chanProgress := make(chan Progress)
	chanIPsFound := make(chan IPFoundMeta)
	//printing progress
	log.Println("threads number size: ", threadsNumber, input.URL.Host)
	go func() {
		for progressData := range chanProgress {
			log.Printf("Progress 3: %.2f %% %.0f/%d amount step 1-2: *%.2f %% - %.2f%%*  ips found: %d\n", (progressData.Progress/float64(progressData.TotalIPs))*100, progressData.Progress, progressData.TotalIPs, (float64(progressData.CountFirst)/float64(progressData.Progress))*100, (float64(progressData.CountSecond)/progressData.Progress)*100, progressData.IpsFoundCount)
		}
	}()
	//---
	//saving results as soon as it finds one
	if err := os.MkdirAll(fmt.Sprintf("%s/body", outputFolder), 0700); err != nil {
		log.Fatalf("failed creating a folder, please check the permissions of code: %s", err)
	}
	file, err := os.OpenFile(fmt.Sprintf("%s/ips.csv", outputFolder), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	file.WriteString("asn index file, asn id, asn name, ip\n")
	go func() {
		for ipMeta := range chanIPsFound {
			file.WriteString(strconv.Itoa(ipMeta.AsnIndex) + "," + ipMeta.AsnID + "," + ipMeta.AsnName + "," + ipMeta.IP.IP + "\n")
			fileName := strings.ReplaceAll(ipMeta.IP.IP, ".", "_")
			os.WriteFile(fmt.Sprintf("%s/body/%s.html", outputFolder, fileName), ipMeta.Body, 0644)
		}
	}()
	//
	ipsFound, err := input.GetIPsFromPriorities(stopOnASNFound, threadsNumber, tcpTimeout, checkProgressEach, chanProgress, chanIPsFound, asnPath)
	if err != nil {
		log.Fatalf("failed running the code: %s", err)
	}
	time.Sleep(time.Second * 3)
	file.Close()
	finalResult := Result{
		Domain: input.URL.Host,
	}
	//grouping the results into it's own ASN number
	for _, currentIPrange := range ipsFound {
		var found bool
		for k, asnFound := range finalResult.AsnsFound {
			if asnFound.ID == currentIPrange.AsnID {
				found = true
				finalResult.AsnsFound[k].IPs = append(finalResult.AsnsFound[k].IPs, currentIPrange.IP)
				break
			}
		}
		if !found {
			finalResult.AsnsFound = append(finalResult.AsnsFound, AsnFound{ID: currentIPrange.AsnID, Name: currentIPrange.AsnName, IPs: []IP{currentIPrange.IP}})
		}
	}
	//
	f, _ := os.Create(fmt.Sprintf("%s/result.json", outputFolder))
	json.NewEncoder(f).Encode(finalResult)
	//if you want to verify the ssl
	finalResult.SetSSL(input.URL.Host, sslTimeout)
	//---------------------
	f2, _ := os.Create(fmt.Sprintf("%s/result_ssl_verification.json", outputFolder))
	json.NewEncoder(f2).Encode(finalResult)
}

// formulario 8832
// eian 32-0763100
