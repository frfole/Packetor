package sc_play

import (
	"Packetor/packetor/decode"
	"github.com/gofrs/uuid/v5"
)

type SpawnEntity struct {
	EntityID   int32
	EntityUuid uuid.UUID
	Type       int32
	X          float64
	Y          float64
	Z          float64
	Pitch      decode.Angle
	Yaw        decode.Angle
	HeadYaw    decode.Angle
	Data       int32
	VelocityX  int16
	VelocityY  int16
	VelocityZ  int16
}

func (p SpawnEntity) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	uid, err := reader.ReadUuid()
	if err != nil {
		return nil, err
	}
	eType, err := reader.ReadVarInt()
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
	pitch, err := reader.ReadAngle()
	if err != nil {
		return nil, err
	}
	yaw, err := reader.ReadAngle()
	if err != nil {
		return nil, err
	}
	headYaw, err := reader.ReadAngle()
	if err != nil {
		return nil, err
	}
	data, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	velX, err := reader.ReadShort()
	if err != nil {
		return nil, err
	}
	velY, err := reader.ReadShort()
	if err != nil {
		return nil, err
	}
	velZ, err := reader.ReadShort()
	if err != nil {
		return nil, err
	}
	return SpawnEntity{
		EntityID:   eid,
		EntityUuid: uid,
		Type:       eType,
		X:          x,
		Y:          y,
		Z:          z,
		Pitch:      pitch,
		Yaw:        yaw,
		HeadYaw:    headYaw,
		Data:       data,
		VelocityX:  velX,
		VelocityY:  velY,
		VelocityZ:  velZ,
	}, nil
}

func (p SpawnEntity) IsValid() (reason error) {
	return nil
}
