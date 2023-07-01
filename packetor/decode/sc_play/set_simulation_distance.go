package sc_play

import "Packetor/packetor/decode"

type SetSimulationDistance struct {
	SimulationDistance int32
}

func (p SetSimulationDistance) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	dist, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return SetSimulationDistance{SimulationDistance: dist}, nil
}

func (p SetSimulationDistance) IsValid() (reason error) {
	return nil
}
