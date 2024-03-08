package main

import (
	"marcopolo/process"
	"time"
)

func (in *Input) Marco(ip string, tcpTimeout time.Duration) (bool, int, []byte, error) {
	return process.Start(in.URL, in.request, in.keyworkds, ip, tcpTimeout)
}
