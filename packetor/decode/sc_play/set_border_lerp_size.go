package sc_play

import "Packetor/packetor/decode"

type SetBorderLerpSize struct {
	OldDiameter float64
	NewDiameter float64
	Speed       int64
}

func (p SetBorderLerpSize) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	oldD, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	newD, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	speed, err := reader.ReadVarLong()
	if err != nil {
		return nil, err
	}
	return SetBorderLerpSize{
		OldDiameter: oldD,
		NewDiameter: newD,
		Speed:       speed,
	}, nil
}

func (p SetBorderLerpSize) IsValid() (reason error) {
	return nil
}
