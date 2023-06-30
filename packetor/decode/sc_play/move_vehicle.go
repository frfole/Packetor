package sc_play

import "Packetor/packetor/decode"

type MoveVehicle struct {
	X     float64
	Y     float64
	Z     float64
	Yaw   float32
	Pitch float32
}

func (p MoveVehicle) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
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
	pitch, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	yaw, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	return MoveVehicle{
		X:     x,
		Y:     y,
		Z:     z,
		Yaw:   yaw,
		Pitch: pitch,
	}, nil
}

func (p MoveVehicle) IsValid() (reason error) {
	return nil
}
