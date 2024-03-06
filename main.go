package main

func main() {
	input := SecureState
	threadsNumber := 10
	asnPath := "./asn.csv"
	input.stopOnEachAsnFound(threadsNumber, asnPath)
	//input.stopOnIPCertificateFound(threadsNumber,asnPath)
}
