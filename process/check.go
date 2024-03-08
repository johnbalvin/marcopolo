package process

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"marcopolo/trace"
	"net"
	"net/http"
	"net/url"
	"time"
)

// I'm not using the http package because for some IP the code will hang forever even thouth I had place a timeout on the client and on the dialer
func checkWithHTTP(host string, requestBody []byte, keywords [][]byte, ip string, tcpTimeout time.Duration) (bool, bool, []byte, error) {
	ipPort := fmt.Sprintf("%s:80", ip)
	conn, err := net.DialTimeout("tcp", ipPort, tcpTimeout)
	if err != nil {
		return false, false, nil, trace.NewOrAdd(1, "process", "checkWithHTTP", err, "")
	}
	defer conn.Close()
	deadLine := time.Now().Add(time.Second * 5)
	if err := conn.SetWriteDeadline(deadLine); err != nil {
		return false, false, nil, trace.NewOrAdd(2, "process", "checkWithHTTP", err, "")
	}
	if _, err := conn.Write(requestBody); err != nil {
		return false, false, nil, trace.NewOrAdd(3, "process", "checkWithHTTP", err, "")
	}
	if err := conn.SetReadDeadline(deadLine); err != nil {
		return false, false, nil, trace.NewOrAdd(4, "process", "checkWithHTTP", err, "")
	}
	buffer := make([]byte, bufferSize)
	if _, err := conn.Read(buffer); err != nil && err != io.EOF {
		return false, false, nil, trace.NewOrAdd(5, "process", "checkWithHTTP", err, "")
	}
	foundAllKeywords := true
	for _, keyword := range keywords {
		if !bytes.Contains(buffer, keyword) {
			foundAllKeywords = false
			break
		}
	}
	if foundAllKeywords {
		return false, true, buffer, nil
	}
	if bytes.Contains(buffer, []byte(host)) {
		return true, false, buffer, nil
	}
	return false, false, buffer, nil
}

func checkWithHTTPs(websiteURL *url.URL, keywords [][]byte, ip string) (bool, []byte, error) {
	ipPort := fmt.Sprintf("%s:443", ip)
	transport := &http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 5 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", websiteURL.String(), nil)
	if err != nil {
		return false, nil, trace.NewOrAdd(1, "process", "checkWithHTTPs", err, "")
	}
	req.Host = websiteURL.Host
	req.URL.Host = ipPort
	req.Header = headers
	resp, err := client.Do(req)
	if err != nil {
		return false, nil, trace.NewOrAdd(2, "process", "checkWithHTTPs", err, "")
	}
	defer resp.Body.Close()
	buffer := make([]byte, bufferSize)
	_, err = io.ReadFull(resp.Body, buffer)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return false, nil, trace.NewOrAdd(3, "process", "checkWithHTTPs", err, "")
	}
	for _, keyword := range keywords {
		if !bytes.Contains(buffer, keyword) {
			return false, buffer, nil
		}
	}
	return true, buffer, nil
}
