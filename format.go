package main

import (
	"marcopolo/asn"
	"net/url"
	"time"
)

type Result struct {
	Domain    string
	AsnsFound []AsnFound
}
type AsnFound struct {
	ID   string
	Name string
	IPs  []IP
}
type IP struct {
	IP               string
	HashSSLVerified  *bool    `json:",omitempty"`
	CommonSSLCNNames []string `json:",omitempty"`
}
type Input struct {
	URL        *url.URL
	Keyworkds  []string
	keyworkds  [][]byte
	Asn        asn.Asn
	BufferSize int
	TCPTimeout time.Duration
	request    []byte
}

type Progress struct {
	CountFirst    uint32
	CountSecond   uint32
	Progress      float64
	TotalIPs      uint32
	IpsFoundCount int
}

type IPFoundMeta struct {
	IP       IP
	AsnName  string
	AsnID    string
	AsnIndex int
	Body     []byte
}
