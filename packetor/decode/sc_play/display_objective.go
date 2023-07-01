package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type DisplayObjective struct {
	Position  int8
	ScoreName string
}

func (p DisplayObjective) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	pos, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	name, err := reader.ReadString()
	if err != nil {
		return nil, err
	}
	return DisplayObjective{
		Position:  pos,
		ScoreName: name,
	}, nil
}

func (p DisplayObjective) IsValid() (reason error) {
	if p.Position < 0 || 18 < p.Position {
		return fmt.Errorf("position must be in <0; 18> was %d", p.Position)
	}
	return nil
}
