package pinger

import (
	"net"
	"pinguin/internal/pinger/icmp"
	"time"
)

type Reply struct {
	Packet icmp.Packet
	From   net.IP
	RTT    time.Duration
	TTL    int
}
