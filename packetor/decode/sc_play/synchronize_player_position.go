package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type SynchronizePlayerPosition struct {
	X          float64
	Y          float64
	Z          float64
	Yaw        float32
	Pitch      float32
	Flags      int8
	TeleportID int32
}

func (p SynchronizePlayerPosition) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	x, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	y, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	z, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	yaw, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	pitch, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	flags, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	tid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return SynchronizePlayerPosition{
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        yaw,
		Pitch:      pitch,
		Flags:      flags,
		TeleportID: tid,
	}, nil
}

func (p SynchronizePlayerPosition) IsValid() (reason error) {
	if (p.Flags &^ 0b11111) != 0 {
		return fmt.Errorf("unknown flags %b", p.Flags)
	}
	return nil
}
