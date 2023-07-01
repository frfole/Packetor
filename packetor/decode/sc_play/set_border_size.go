package sc_play

import "Packetor/packetor/decode"

type SetBorderSize struct {
	Diameter float64
}

func (p SetBorderSize) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	d, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	return SetBorderSize{Diameter: d}, nil
}

func (p SetBorderSize) IsValid() (reason error) {
	return nil
}
