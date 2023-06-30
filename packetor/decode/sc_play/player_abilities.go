package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type PlayerAbilities struct {
	Flags       int8
	FlyingSpeed float32
	FOVModifier float32
}

func (p PlayerAbilities) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	flags, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	fSpeed, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	fovMod, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	return PlayerAbilities{
		Flags:       flags,
		FlyingSpeed: fSpeed,
		FOVModifier: fovMod,
	}, nil
}

func (p PlayerAbilities) IsValid() (reason error) {
	if (p.Flags &^ 0b1111) != 0 {
		return fmt.Errorf("unknown flags %b", p.Flags)
	}
	return nil
}
