package sc_play

import (
	"Packetor/packetor/decode"
	"github.com/gofrs/uuid/v5"
)

type SpawnPlayer struct {
	EntityID   int32
	PlayerUuid uuid.UUID
	X          float64
	Y          float64
	Z          float64
	Yaw        decode.Angle
	Pitch      decode.Angle
}

func (p SpawnPlayer) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	uid, err := reader.ReadUuid()
	if err != nil {
		return nil, err
	}
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
	yaw, err := reader.ReadAngle()
	if err != nil {
		return nil, err
	}
	pitch, err := reader.ReadAngle()
	if err != nil {
		return nil, err
	}
	return SpawnPlayer{
		EntityID:   eid,
		PlayerUuid: uid,
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        yaw,
		Pitch:      pitch,
	}, nil
}

func (p SpawnPlayer) IsValid() (reason error) {
	return nil
}
