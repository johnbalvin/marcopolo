package utils

import (
	"fmt"
	"log"
	"net"
)

func CountIPsInRange(startIP, endIP string) uint32 {
	startUint := IpToUint32(startIP)
	endUint := IpToUint32(endIP)
	if startUint > endUint {
		log.Fatalln("countIPsInRange 13 ", startIP, endIP)
	}
	return endUint - startUint + 1
}

func IpToUint32(ipStr string) uint32 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		log.Fatalln("ipToUint32 1: ", ipStr)
		return 0
	}
	ip = ip.To4()
	if ip == nil {
		log.Fatalln("ipToUint32 2: ", ipStr)
		return 0
	}
	var ipUint uint32
	for i := 0; i < 4; i++ {
		ipUint |= uint32(ip[i]) << ((3 - uint(i)) * 8)
	}
	return ipUint
}
func Uint32ToIP(ipUint uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ipUint>>24), byte(ipUint>>16)&0xFF, byte(ipUint>>8)&0xFF, byte(ipUint)&0xFF)
}
