package pinger

import (
	"net"
	"pinguin/internal/pinger/icmp"
)

type Pinger struct {
	Host string
}

func (p *Pinger) Ping() (Reply, error) {
	conn, err := net.ListenPacket("ip4:icmp", "0.0.0.0")

	if err != nil {
		return Reply{}, err
	}

	var packet = icmp.BuildPacket(8, 0, 30, 0)
	var marshaledPacket = icmp.Marshal(packet)

	var neeededAddr = &net.IPAddr{
		IP: net.ParseIP("8.8.8.8"), // TODO: Remove the const
	}

	var _, conError = conn.WriteTo(marshaledPacket, neeededAddr)

	if conError != nil {
		println("Connection error")
		println(conError.Error())
	}

	receivedPayload := make([]byte, 1500)

	var nReceivedBytes, addr, fetchingErr = conn.ReadFrom(receivedPayload)

	if fetchingErr != nil {
		println("fetching error")
		println(fetchingErr)
		return Reply{}, fetchingErr
	}

	if nReceivedBytes == 0 {
		println("Zero bytes received.")
		return Reply{}, nil
	}

	println("Received packets num: ", nReceivedBytes)
	println("Packets from addrs: ", addr.String())

	return Reply{}, nil

}
