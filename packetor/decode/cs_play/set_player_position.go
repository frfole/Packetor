package cs_play

import "Packetor/packetor/decode"

type SetPlayerPosition struct {
	X        float64
	Y        float64
	Z        float64
	OnGround bool
}

func (p SetPlayerPosition) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
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
	grounded, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return SetPlayerPosition{
		X:        x,
		Y:        y,
		Z:        z,
		OnGround: grounded,
	}, nil
}

func (p SetPlayerPosition) IsValid() (reason error) {
	return nil
}
