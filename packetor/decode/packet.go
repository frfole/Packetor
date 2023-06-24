package decode

type Packet interface {
	Read(reader PacketReader) (packet Packet, err error)
	IsValid() (reason error)
}
