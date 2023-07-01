package sc_play

import "Packetor/packetor/decode"

type SetBorderCenter struct {
	X float64
	Z float64
}

func (p SetBorderCenter) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	x, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	z, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	return SetBorderCenter{
		X: x,
		Z: z,
	}, nil
}

func (p SetBorderCenter) IsValid() (reason error) {
	return nil
}
