package process

import "net/http"

var headers = http.Header{
	"Accept":          []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
	"Accept-Language": []string{"en"},
	"User-Agent":      []string{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"},
	"Sec-Fetch-Dest":  []string{"document"},
	"Sec-Fetch-Mode":  []string{"navigate"},
	"Sec-Fetch-Site":  []string{"none"},
	"Sec-Ch-Ua":       []string{`"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`},
	"Connection":      []string{"close"},
}

const bufferSize = 1024
