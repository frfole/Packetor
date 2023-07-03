package cs_play

import "Packetor/packetor/decode"

type SetPlayerRotation struct {
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (p SetPlayerRotation) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
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
	return SetPlayerRotation{
		Yaw:      yaw,
		Pitch:    pitch,
		OnGround: grounded,
	}, nil
}

func (p SetPlayerRotation) IsValid() (reason error) {
	return nil
}
