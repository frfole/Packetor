package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type EntityAnimation struct {
	EntityID  int32
	Animation uint8
}

func (p EntityAnimation) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	anim, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	return EntityAnimation{
		EntityID:  eid,
		Animation: anim,
	}, nil
}

func (p EntityAnimation) IsValid() (reason error) {
	if p.Animation > 5 {
		return fmt.Errorf("unknown animation %d", p.Animation)
	}
	return nil
}
