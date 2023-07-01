package sc_play

import "Packetor/packetor/decode"

type SetDefaultSpawnPosition struct {
	Location decode.Position
	Angle    float32
}

func (p SetDefaultSpawnPosition) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	angle, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	return SetDefaultSpawnPosition{
		Location: loc,
		Angle:    angle,
	}, nil
}

func (p SetDefaultSpawnPosition) IsValid() (reason error) {
	return nil
}
