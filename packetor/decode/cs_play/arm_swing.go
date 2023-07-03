package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type ArmSwing struct {
	Hand int32
}

func (p ArmSwing) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	hand, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return ArmSwing{Hand: hand}, nil
}

func (p ArmSwing) IsValid() (reason error) {
	if p.Hand < 0 || 1 < p.Hand {
		return fmt.Errorf("hand must be in <0; 1> was %d", p.Hand)
	}
	return nil
}
