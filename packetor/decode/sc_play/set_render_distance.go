package sc_play

import "Packetor/packetor/decode"

type SetRenderDistance struct {
	ViewDistance int32
}

func (p SetRenderDistance) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	dist, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return SetRenderDistance{ViewDistance: dist}, nil
}

func (p SetRenderDistance) IsValid() (reason error) {
	return nil
}
