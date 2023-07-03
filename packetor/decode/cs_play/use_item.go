package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type UseItem struct {
	Hand     int32
	Sequence int32
}

func (p UseItem) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	hand, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	sequence, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return UseItem{
		Hand:     hand,
		Sequence: sequence,
	}, nil
}

func (p UseItem) IsValid() (reason error) {
	if p.Hand < 0 || 1 < p.Hand {
		return fmt.Errorf("hand must be in <0; 1> was %d", p.Hand)
	}
	return nil
}
