package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type JigsawGenerate struct {
	Location    decode.Position
	Levels      int32
	KeepJigsaws bool
}

func (p JigsawGenerate) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	levels, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	keep, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return JigsawGenerate{
		Location:    loc,
		Levels:      levels,
		KeepJigsaws: keep,
	}, nil
}

func (p JigsawGenerate) IsValid() (reason error) {
	if p.Levels < 0 {
		return fmt.Errorf("levels must be atleast 0 was %d", p.Levels)
	}
	return nil
}
