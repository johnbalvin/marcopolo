package utils

import (
	"log"
	"net/url"
)

func ParseURL(urlToUse string) *url.URL {
	url, err := url.Parse(urlToUse)
	if err != nil {
		log.Fatalln("wrong url, plese check the url", urlToUse)
	}
	return url
}
