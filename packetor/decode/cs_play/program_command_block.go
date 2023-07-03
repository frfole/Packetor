package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type ProgramCommandBlock struct {
	Location decode.Position
	Command  string
	Mode     int32
	Flags    uint8
}

func (p ProgramCommandBlock) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	cmd, err := reader.ReadString()
	if err != nil {
		return nil, err
	}
	mode, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	flags, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	return ProgramCommandBlock{
		Location: loc,
		Command:  cmd,
		Mode:     mode,
		Flags:    flags,
	}, nil
}

func (p ProgramCommandBlock) IsValid() (reason error) {
	if p.Mode < 0 || 2 < p.Mode {
		return fmt.Errorf("mode must be in <0; 2> was %d", p.Mode)
	}
	if (p.Flags &^ 0b111) != 0 {
		return fmt.Errorf("unknown flags %b", p.Flags)
	}
	return nil
}
