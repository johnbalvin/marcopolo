package main

import (
	"log"
	"marcopolo/asn"
	"marcopolo/ssl"
	"marcopolo/utils"
	"sync"
	"time"
)

func (input *Input) SearchIPsWithKeywords(stopOnAsnIPFound bool, ipsRanges asn.IPRangeWrapper, threadsNumber int, tcpTimeout, progressEach time.Duration, chanProgress chan Progress, chanIPsFound chan IPFoundMeta) []IPFoundMeta {
	input.setInputs()
	bufferSize := input.BufferSize
	if bufferSize == 0 {
		bufferSize = 1024
	}
	var finishedEverything bool
	var progress float64
	var countFirst, countSecond uint32
	var IPsFound []IPFoundMeta
	var mutex sync.Mutex
	var wg sync.WaitGroup
	asnIDIPFound := make(map[string]bool)
	chanIP := make(chan IPFoundMeta, threadsNumber)
	if chanProgress != nil {
		chanProgress <- Progress{
			CountFirst:  countFirst,
			CountSecond: countSecond,
			TotalIPs:    ipsRanges.Quantity,
		}
		seconds := int(progressEach.Seconds())
		if seconds == 0 {
			progressEach = time.Minute
		}
		go func() {
			for {
				time.Sleep(progressEach)
				mutex.Lock()
				chanProgress <- Progress{
					CountFirst:    countFirst,
					CountSecond:   countSecond,
					Progress:      progress,
					TotalIPs:      ipsRanges.Quantity,
					IpsFoundCount: len(IPsFound),
				}
				mutex.Unlock()
				if finishedEverything {
					break
				}
			}
		}()
	}
	currentIps := make([]uint32, len(ipsRanges.IPs))
	for range threadsNumber {
		go func() {
			for ipMetadata := range chanIP {
				polo, whichStep, buffer, err := input.Marco(ipMetadata.IP.IP, tcpTimeout)
				mutex.Lock()
				progress++
				if whichStep == 1 {
					countFirst++
				} else {
					countSecond++
				}
				mutex.Unlock()
				if err != nil {
					//log.Println("err: ", ipMetadata.IP.IP, err)
					wg.Done()
					continue
				}
				if polo {
					mutex.Lock()
					if stopOnAsnIPFound {
						asnIDIPFound[ipMetadata.AsnID] = true
						//cleaning the progress for the ASN found
						for asnIndex, ip := range ipsRanges.IPs {
							if ipMetadata.AsnID != ip.AsnID {
								continue
							}
							currentIP := currentIps[asnIndex]
							if currentIP > ip.End {
								continue
							}
							remainingIPsNumber := ip.End - currentIP
							progress += float64(remainingIPsNumber)
							currentIps[asnIndex] = ip.End + 1
						}
					}
					IPsFound = append(IPsFound, ipMetadata)
					mutex.Unlock()
					if chanIPsFound != nil {
						ipMetadata.Body = buffer
						chanIPsFound <- ipMetadata
					}
				}
				wg.Done()
			}
		}()
	}
	// I know it looks weird but with this I make sure the workers won't focus on one specific IP range from an row from the asn file but rather distrubute
	//the work among all of the  rows, this will look better if I make detail explanation in future
	//these will make faster searches, so lets say theare are 100 rows to check, and the Ip is row 80, if were to do it normally
	//one row per row it will take a lof of time, but if I instead distribute the work on all the rows, like this:
	// first I check row[0]: ipstart
	for asnIndex, ip := range ipsRanges.IPs {
		currentIps[asnIndex] = ip.Start
	}
	for {
		var atLeastOneIPSent bool
		for asnIndex, ip := range ipsRanges.IPs {
			if stopOnAsnIPFound {
				mutex.Lock()
				if asnIDIPFound[ip.AsnID] {
					mutex.Unlock()
					continue
				}
				mutex.Unlock()
			}
			currentIP := currentIps[asnIndex]
			if currentIP > ip.End {
				continue
			}
			wg.Add(1)
			ipString := utils.Uint32ToIP(currentIP)
			currentIP++
			currentIps[asnIndex] = currentIP
			chanIP <- IPFoundMeta{IP: IP{IP: ipString}, AsnIndex: asnIndex, AsnName: ip.AsnName, AsnID: ip.AsnID}
			atLeastOneIPSent = true
		}
		if !atLeastOneIPSent {
			break
		}
	}
	//
	wg.Wait()
	log.Println("finished all, ips found: ", len(IPsFound))
	finishedEverything = true
	close(chanIP)
	return IPsFound
}

func (input *Input) SearchIPsWithSSLMatches(stopOnSSlFound bool, ipsRanges asn.IPRangeWrapper, threadsNumber int, sslTimeout, progressEach time.Duration, chanProgress chan Progress, chanIPsFound chan IPFoundMeta) []IPFoundMeta {
	var finishedEverything bool
	var progress float64
	var countFirst, countSecond uint32
	var IPsFound []IPFoundMeta
	var mutex sync.Mutex
	var wg sync.WaitGroup
	chanIP := make(chan IPFoundMeta, threadsNumber)
	if chanProgress != nil {
		chanProgress <- Progress{
			CountFirst:  countFirst,
			CountSecond: countSecond,
			TotalIPs:    ipsRanges.Quantity,
		}
		seconds := int(progressEach.Seconds())
		if seconds == 0 {
			progressEach = time.Minute
		}
		go func() {
			for {
				time.Sleep(progressEach)
				mutex.Lock()
				chanProgress <- Progress{
					CountFirst:    countFirst,
					CountSecond:   countSecond,
					Progress:      progress,
					TotalIPs:      ipsRanges.Quantity,
					IpsFoundCount: len(IPsFound),
				}
				mutex.Unlock()
				if finishedEverything {
					break
				}
			}
		}()
	}
	for range threadsNumber {
		go func() {
			for ipMetadata := range chanIP {
				if finishedEverything {
					wg.Done()
					continue
				}
				err := ssl.VerifyHost(ipMetadata.IP.IP, input.URL.Host, sslTimeout)
				mutex.Lock()
				progress++
				mutex.Unlock()
				if err != nil {
					//log.Println("err: ", ip.IP, err)
					wg.Done()
					continue
				}
				mutex.Lock()
				if stopOnSSlFound {
					finishedEverything = true
				}
				IPsFound = append(IPsFound, ipMetadata)
				mutex.Unlock()
				if chanIPsFound != nil {
					chanIPsFound <- ipMetadata
				}
				wg.Done()
			}
		}()
	}
	// I know it looks weird but with this I make sure the workers won't focus on one specific IP range from an row from the asn file but rather distrubute
	//the work among all of the  rows, this will look better if I make detail explanation in future
	//these will make faster searches, so lets say theare are 100 rows to check, and the Ip is row 80, if were to do it normally
	//one row per row it will take a lof of time, but if I instead distribute the work on all the rows, like this:
	// first I check row[0]: ipstart
	currentIps := make([]uint32, len(ipsRanges.IPs))
	for asnIndex, ip := range ipsRanges.IPs {
		currentIps[asnIndex] = ip.Start
	}
	for {
		var atLeastOneIPSent bool
		if finishedEverything {
			break
		}
		for ipRangeIndex, ip := range ipsRanges.IPs {
			if finishedEverything {
				break
			}
			currentIP := currentIps[ipRangeIndex]
			if currentIP > ip.End {
				continue
			}
			wg.Add(1)
			ipString := utils.Uint32ToIP(currentIP)
			chanIP <- IPFoundMeta{IP: IP{IP: ipString}, AsnIndex: ipRangeIndex, AsnName: ip.AsnName, AsnID: ip.AsnID}
			atLeastOneIPSent = true
			currentIP++
			currentIps[ipRangeIndex] = currentIP
		}
		if !atLeastOneIPSent {
			break
		}
	}
	//
	wg.Wait()
	log.Println("finished all, ips found: ", len(IPsFound))
	finishedEverything = true
	close(chanIP)
	return IPsFound
}

/*
	for asnIndex, ip := range ipsRanges.IPs {
		wg.Add(1)
		go func(asnIndex int, ip IpRange) {
			for ipUint := ip.Start; ipUint <= ip.End; ipUint++ {
				wg.Add(1)
				ipString := utils.Uint32ToIP(ipUint)
				chanIP <- IPFound{Index: asnIndex, IP: ipString, AsnName: ip.AsnName, AsnID: ip.AsnID}
			}
			wg.Done()
		}(asnIndex, ip)
	}
*/
