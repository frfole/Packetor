package sc_play

import "Packetor/packetor/decode"

type AcknowledgeBlockChange struct {
	SequenceID int32
}

func (p AcknowledgeBlockChange) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	sid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return AcknowledgeBlockChange{SequenceID: sid}, nil
}

func (p AcknowledgeBlockChange) IsValid() (reason error) {
	return nil
}
