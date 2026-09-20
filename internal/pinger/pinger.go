package pinger

import "pinguin/internal/pinger/icmp"

type Pinger struct {
	Host string
}

func (p *Pinger) Ping(packet *icmp.Packet) {

}
