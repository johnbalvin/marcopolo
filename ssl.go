package main

import (
	"log"
	"marcopolo/ssl"
	"sync"
	"time"
)

func (finalResult *Result) SetSSL(theads int, host string, sslTimeout time.Duration) {
	var wg sync.WaitGroup
	var progress int
	var size int
	var mutex sync.Mutex
	chanIndexes := make(chan [2]int)
	for range theads {
		go func() {
			for indexes := range chanIndexes {
				i := indexes[0]
				j := indexes[1]
				ipMeta := finalResult.AsnsFound[i].IPs[j]
				commonNames, sslHashHost, _ := ssl.VerifyAll(ipMeta.IP, host, sslTimeout)
				finalResult.AsnsFound[i].IPs[j].CommonSSLCNNames = commonNames
				finalResult.AsnsFound[i].IPs[j].HashSSLVerified = &sslHashHost
				mutex.Lock()
				progress++
				log.Printf("Progress ssl verification: %d/%d\n", progress, size)
				mutex.Unlock()
				wg.Done()
			}
		}()
	}
	for i, asnFound := range finalResult.AsnsFound {
		wg.Add(len(asnFound.IPs))
		size += len(asnFound.IPs)
		for j := range asnFound.IPs {
			chanIndexes <- [2]int{i, j}
		}
	}
	wg.Wait()
}
