package main

import (
	"net/netip"
	"pinguin/internal/pinger"
)

func main() {
	println("pinger v.0.0.0")

	var pinger = pinger.Pinger{}
	pinger.InitializeConnection()

	ip, err := netip.ParseAddr("8.8.8.8")

	if err != nil {
		println("pizdec")
		return
	}

	reply, err := pinger.Ping(ip)

	if err != nil {
		println(err.Error())
	}

	println(reply.Packet.Type)
	println(reply.From.String())
	println(reply.Packet.Identifier)
	println(reply.Packet.Code)
	println(reply.Packet.Checksum)

}
