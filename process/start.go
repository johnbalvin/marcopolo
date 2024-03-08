package process

import (
	"marcopolo/trace"
	"net/url"
	"time"
)

func Start(websiteURL *url.URL, requestBody []byte, keywords [][]byte, ip string, tcpTimeout time.Duration) (bool, int, []byte, error) {
	redirectFound, keywordsFound, buffer, err := checkWithHTTP(websiteURL.Host, requestBody, keywords, ip, tcpTimeout)
	if err != nil {
		return false, 1, nil, trace.NewOrAdd(1, "process", "Start", err, "")
	}
	if keywordsFound {
		return true, 1, buffer, nil
	}
	if !redirectFound {
		return false, 1, buffer, nil
	}
	found, buffer, err := checkWithHTTPs(websiteURL, keywords, ip)
	if err != nil {
		return false, 2, nil, trace.NewOrAdd(2, "process", "Start", err, "")
	}
	return found, 2, buffer, nil
}
