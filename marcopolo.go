package main

import "marcopolo/process"

func (in *Input) Marco(ip string) (bool, int, []byte, error) {
	return process.Start(in.URL, in.request, in.keyworkds, ip, in.TCPTimeout)
}
