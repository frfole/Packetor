package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type PlayerAbilities struct {
	Flags uint8
}

func (p PlayerAbilities) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	flags, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	return PlayerAbilities{Flags: flags}, nil
}

func (p PlayerAbilities) IsValid() (reason error) {
	if (p.Flags &^ 0b10) != 0 {
		return fmt.Errorf("unknown flags %b", p.Flags)
	}
	return nil
}
