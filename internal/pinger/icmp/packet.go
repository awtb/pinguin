package icmp

import "fmt"

type Packet struct {
	Type           uint8
	Code           uint8
	Checksum       uint16
	Identifier     uint16
	SequenceNumber uint16
}

func BuildPacket(packetType uint8, code uint8, identifier uint16, seqNum uint16) *Packet {
	packet := &Packet{
		Type:           packetType,
		Code:           code,
		Checksum:       0,
		Identifier:     identifier,
		SequenceNumber: seqNum,
	}

	packet.Checksum = calculateChecksum(packet)

	return packet
}

func calculateChecksum(packet *Packet) uint16 {
	bytes := packet.Marshal()

	pt1 := uint32(bytes[0])<<8 | uint32(bytes[1])
	pt2 := uint32(bytes[2])<<8 | uint32(bytes[3])
	pt3 := uint32(bytes[4])<<8 | uint32(bytes[5])
	pt4 := uint32(bytes[6])<<8 | uint32(bytes[7])

	var sum uint32 = (pt1 + pt2 + pt3 + pt4)

	sum = (sum & 0xFFFF) + (sum >> 16)
	return ^uint16(sum)
}

func (packet *Packet) Marshal() []byte {
	data := make([]byte, 8)

	data[0] = packet.Type
	data[1] = packet.Code

	// Full checksum is bigger than 1 byte
	// Protocol specifies that the checksum is stored in 2 bytes
	// The high byte is stored first
	data[2] = byte(packet.Checksum >> 8) // Use bitwise right shift to get it
	// Then the low byte is stored
	data[3] = byte(packet.Checksum) // Just cast to byte to get the low byte (it trims automatically)

	// Exactly the same shit with the identifier
	data[4] = byte(packet.Identifier >> 8)
	data[5] = byte(packet.Identifier)

	// And twith the sequence number
	data[6] = byte(packet.SequenceNumber >> 8)
	data[7] = byte(packet.SequenceNumber)

	return data
}

func Unmarshal(data []byte) (Packet, error) {
	if len(data) < 8 {
		return Packet{}, fmt.Errorf("icmp: packet too short: got %d bytes, need at least 8", len(data))
	}

	return Packet{
		Type:           data[0],
		Code:           data[1],
		Checksum:       uint16(data[2])<<8 | uint16(data[3]),
		Identifier:     uint16(data[4])<<8 | uint16(data[5]),
		SequenceNumber: uint16(data[6])<<8 | uint16(data[7]),
	}, nil
}
