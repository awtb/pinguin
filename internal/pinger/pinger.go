package pinger

import (
	"fmt"
	"net"
	"net/netip"
	"pinguin/internal/pinger/icmp"
)

type Pinger struct {
	Host       string
	connection net.PacketConn
}

func (p *Pinger) InitializeConnection() error {
	conn, err := net.ListenPacket("ip4:icmp", "0.0.0.0")

	if err != nil {
		return err
	}

	p.connection = conn
	return nil
}

func buildPingPacket() *icmp.Packet {
	p := icmp.BuildPacket(8, 0, 30, 0)
	return p
}

func (p *Pinger) sendBytes(destination net.Addr, data []byte) error {
	_, err := p.connection.WriteTo(data, destination)

	if err != nil {
		return fmt.Errorf("Some problem with sending bytes to the needed destination")
	}

	return nil
}

func (p *Pinger) receiveBytes() (net.Addr, []byte, error) {
	buffer := make([]byte, 1500)
	receivedBytesCount, addr, err := p.connection.ReadFrom(buffer)

	if err != nil {
		println("fetching error")
		println(err)
		return nil, nil, err
	}

	if receivedBytesCount == 0 {
		return nil, nil, fmt.Errorf("Zero bytes received")
	}

	return addr, buffer, nil

}

func (p *Pinger) receivePacket() (*icmp.Packet, error) {
	_, bytes, err := p.receiveBytes()
	packet, err := icmp.Unmarshal(bytes)

	if err != nil {
		return nil, fmt.Errorf("Problem at deserialization")
	}

	return &packet, nil
}

func (p *Pinger) sendPacket(destination netip.Addr, packet *icmp.Packet) error {
	bytes := packet.Marshal()

	addr := &net.IPAddr{
		IP: net.IP(destination.AsSlice()),
	}

	err := p.sendBytes(addr, bytes)

	if err != nil {
		return fmt.Errorf("Please, call Pinger.InitializeConnection(), before ping.")
	}

	return nil
}

func (p *Pinger) Ping(addr netip.Addr) (Reply, error) {
	if p.connection == nil {
		return Reply{}, fmt.Errorf("Please, call Pinger.InitializeConnection(), before ping.")
	}

	pingPacket := buildPingPacket()

	err := p.sendPacket(addr, pingPacket)

	if err != nil {
		return Reply{}, err
	}

	receivedPacket, err := p.receivePacket()

	if err != nil {
		return Reply{}, nil
	}

	return Reply{
		Packet: receivedPacket,
		From:   net.ParseIP(addr.String()),
		RTT:    0,
		TTL:    0,
	}, nil

}
