package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type PlayerInput struct {
	Sideways float32
	Forward  float32
	Flags    uint8
}

func (p PlayerInput) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	sideways, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	forward, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	flags, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	return PlayerInput{
		Sideways: sideways,
		Forward:  forward,
		Flags:    flags,
	}, nil
}

func (p PlayerInput) IsValid() (reason error) {
	if (p.Flags &^ 0b11) != 0 {
		return fmt.Errorf("unknown flags %b", p.Flags)
	}
	return nil
}
