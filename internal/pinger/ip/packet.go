package ip

import "fmt"

type IPProtocolVer uint8
type ProtocolType uint8

const (
	IPV4 IPProtocolVer = 4
	IPV6 IPProtocolVer = 6
)

const (
	ProtocolICMP ProtocolType = 1
	ProtocolTCP  ProtocolType = 6
	ProtocolUDP  ProtocolType = 17
)

type Packet struct {
	Version         IPProtocolVer
	IHL             uint8
	DSCP            uint8
	ECN             uint8
	TotalLength     uint16
	Identification  uint16
	Flags           uint8
	FragmentOffset  uint16
	TTL             uint8
	Protocol        ProtocolType
	Checksum        uint16
	SourceAddr      []byte
	DestinationAddr []byte
	Payload         []byte
}

func BuildPacket(
	ttl uint8,
	protocol ProtocolType,
	sourceAddr []byte,
	destinationAddr []byte,
	version IPProtocolVer,
	payload []byte,
) (*Packet, error) {
	if version != IPV4 {
		return nil, fmt.Errorf("Only IPV4 is currently supported.")
	}

	packet := Packet{
		TTL:             ttl,
		Version:         version,
		Protocol:        protocol,
		Payload:         payload,
		SourceAddr:      sourceAddr,
		DestinationAddr: destinationAddr,
	}

	// Notes
	// 1. How to calculate checksum? Learn algo from IPV4 spec
	// 2. Maybe a more high-level type for both source & destinations?
	// 3. TotalLength should be private? We calculate it by ourself, by the way.
	// 4. Should we provide machenisms to set other params, such as DSCP / ECN / TTL / Identification?
	// 5. Realated to 4, probably we can do exactly the same with DSCP / ECN / TTL (provide a list of consts as enums)

	return &packet, nil
}

func (packet *Packet) Marshal() []byte {
	//TODO: Implement me.
	return nil
}

func Unmarshal(data []byte) *Packet {
	//TODO: Implement me.
	return nil
}
