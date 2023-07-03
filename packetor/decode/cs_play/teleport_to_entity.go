package cs_play

import (
	"Packetor/packetor/decode"
	"github.com/gofrs/uuid/v5"
)

type TeleportToEntity struct {
	Target uuid.UUID
}

func (p TeleportToEntity) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	target, err := reader.ReadUuid()
	if err != nil {
		return nil, err
	}
	return TeleportToEntity{Target: target}, nil
}

func (p TeleportToEntity) IsValid() (reason error) {
	return nil
}
