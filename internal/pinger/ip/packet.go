package ip

type Packet struct {
	Version         uint8
	IHL             uint8
	DSCP            uint8
	TotalLength     uint16
	ECN             uint8
	Flags           uint8
	FragmentOffset  uint16
	Identification  uint16
	TTL             uint8
	Protocol        uint8
	Checksum        uint16
	SourceAddr      [4]byte
	DestinationAddr [4]byte
	Payload         []byte
}

func BuildPacket() {
	//TODO: Implement me.
}

func (packet *Packet) Marhsal() {
	//TODO: Implement me.
}

func Unmarshal(data []byte) {
	//TODO: Implement me.
}
