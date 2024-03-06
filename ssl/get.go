package ssl

import (
	"crypto/tls"
	"marcopolo/trace"
	"net"
	"time"
)

func VerifyAll(ip, host string) ([]string, bool, error) {
	var found bool
	if err := VerifyHost(ip, host); err == nil {
		found = true
	}
	commonNames, err := GetCommonNames(ip)
	if err != nil {
		return nil, false, trace.NewOrAdd(1, "ssl", "VerifyAll", err, "")
	}
	return commonNames, found, nil
}
func GetCommonNames(ip string) ([]string, error) {
	ipPort := ip + ":443"
	dialer := &net.Dialer{
		Timeout: 10 * time.Second,
	}
	conn, err := tls.DialWithDialer(dialer, "tcp", ipPort, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return nil, trace.NewOrAdd(1, "ssl", "GetCommonNames", err, "")
	}
	defer conn.Close()

	state := conn.ConnectionState()
	var commonNames []string
	for _, cert := range state.PeerCertificates {
		commonNames = append(commonNames, cert.Subject.CommonName)
	}
	return commonNames, nil
}

// yes separated functions, sometimes it's not enough with the IP and need to send the host to tls conection
func VerifyHost(ip, host string) error {
	ipPort := ip + ":443"
	dialer := &net.Dialer{
		Timeout: 3 * time.Second,
	}
	conn, err := tls.DialWithDialer(dialer, "tcp", ipPort, &tls.Config{
		ServerName: host,
	})
	if err != nil {
		return trace.NewOrAdd(1, "ssl", "VerifyHost", err, "")
	}
	defer conn.Close()
	if err := conn.VerifyHostname(host); err != nil {
		return trace.NewOrAdd(2, "ssl", "VerifyHost", err, "")
	}
	return nil
}
