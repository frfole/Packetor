package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type UseItemOn struct {
	Hand            int32
	Location        decode.Position
	Face            int32
	CursorPositionX float32
	CursorPositionY float32
	CursorPositionZ float32
	InsideBlock     bool
	Sequence        int32
}

func (p UseItemOn) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	hand, err := reader.ReadVarInt()
	loc, err := reader.ReadPosition()
	face, err := reader.ReadVarInt()
	x, err := reader.ReadFloat()
	y, err := reader.ReadFloat()
	z, err := reader.ReadFloat()
	inside, err := reader.ReadBoolean()
	sequence, err := reader.ReadVarInt()
	return UseItemOn{
		Hand:            hand,
		Location:        loc,
		Face:            face,
		CursorPositionX: x,
		CursorPositionY: y,
		CursorPositionZ: z,
		InsideBlock:     inside,
		Sequence:        sequence,
	}, nil
}

func (p UseItemOn) IsValid() (reason error) {
	if p.Hand < 0 || 1 < p.Hand {
		return fmt.Errorf("hand must be in <0; 1> was %d", p.Hand)
	}
	if p.Face < 0 || 5 < p.Face {
		return fmt.Errorf("face must be in <0; 5> was %d", p.Face)
	}
	return nil
}
