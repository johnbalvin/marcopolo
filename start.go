package main

import (
	"log"
	"time"
)

func (input *Input) GetIPsFromPriorities(stopOnAsnIPFound bool, threadsNumber int, tcpTimeout, progressEach time.Duration, chanProgress chan Progress, chanIPsFound chan IPFoundMeta, asnPath string) ([]IPFoundMeta, error) {
	ipCollection, err := input.Asn.GetIPs(asnPath)
	if err != nil {
		return nil, err
	}
	log.Printf("total quantity: %d\n", ipCollection.Priorities.Quantity)
	ipsFound := input.SearchIPsWithKeywords(stopOnAsnIPFound, ipCollection.Priorities, threadsNumber, tcpTimeout, progressEach, chanProgress, chanIPsFound)
	return ipsFound, nil
}

func (input *Input) GetIPsWithSSLFromPriorities(stopOnSSlFound bool, threadsNumber int, sslTimeout, progressEach time.Duration, chanProgress chan Progress, chanIPsFound chan IPFoundMeta, asnPath string) ([]IPFoundMeta, error) {
	ipCollection, err := input.Asn.GetIPs(asnPath)
	if err != nil {
		return nil, err
	}
	log.Printf("total quantity: %d\n", ipCollection.Priorities.Quantity)
	ipsFound := input.SearchIPsWithSSLMatches(stopOnSSlFound, ipCollection.Priorities, threadsNumber, sslTimeout, progressEach, chanProgress, chanIPsFound)
	return ipsFound, nil
}
