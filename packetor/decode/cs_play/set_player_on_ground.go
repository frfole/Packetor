package cs_play

import "Packetor/packetor/decode"

type SetPlayerOnGround struct {
	OnGround bool
}

func (p SetPlayerOnGround) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	grounded, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return SetPlayerOnGround{OnGround: grounded}, nil
}

func (p SetPlayerOnGround) IsValid() (reason error) {
	return nil
}
