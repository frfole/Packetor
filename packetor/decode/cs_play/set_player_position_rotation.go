package cs_play

import "Packetor/packetor/decode"

type SetPlayerPositionRotation struct {
	X        float64
	Y        float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (p SetPlayerPositionRotation) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
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
	grounded, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return SetPlayerPositionRotation{
		X:        x,
		Y:        y,
		Z:        z,
		Yaw:      yaw,
		Pitch:    pitch,
		OnGround: grounded,
	}, nil
}

func (p SetPlayerPositionRotation) IsValid() (reason error) {
	return nil
}
