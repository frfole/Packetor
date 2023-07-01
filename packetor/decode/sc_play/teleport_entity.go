package sc_play

import "Packetor/packetor/decode"

type TeleportEntity struct {
	EntityID int32
	X        float64
	Y        float64
	Z        float64
	Yaw      decode.Angle
	Pitch    decode.Angle
	OnGround bool
}

func (p TeleportEntity) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
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
	onGround, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return TeleportEntity{
		EntityID: eid,
		X:        x,
		Y:        y,
		Z:        z,
		Yaw:      yaw,
		Pitch:    pitch,
		OnGround: onGround,
	}, nil
}

func (p TeleportEntity) IsValid() (reason error) {
	return nil
}
