package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type PlayerAction struct {
	Status   int32
	Location decode.Position
	Face     uint8
	Sequence int32
}

func (p PlayerAction) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	status, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	face, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	sequence, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return PlayerAction{
		Status:   status,
		Location: loc,
		Face:     face,
		Sequence: sequence,
	}, nil
}

func (p PlayerAction) IsValid() (reason error) {
	if p.Status < 0 || 6 < p.Status {
		return fmt.Errorf("status must be in <0; 6> was %d", p.Status)
	}
	if 5 < p.Face {
		return fmt.Errorf("face must be in <0; 5> was %d", p.Face)
	}
	return nil
}
