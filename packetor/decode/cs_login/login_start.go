package cs_login

import (
	"Packetor/packetor/decode"
	"fmt"
	"github.com/gofrs/uuid/v5"
)

type LoginStart struct {
	Username string
	HasUuid  bool
	Uuid     uuid.UUID
}

func (p LoginStart) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	username, err := reader.ReadString0(16)
	if err != nil {
		return nil, err
	}
	hasUuid, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	var pUuid uuid.UUID
	if hasUuid {
		pUuid, err = reader.ReadUuid()
		if err != nil {
			return nil, err
		}
	}
	return LoginStart{
		Username: username,
		HasUuid:  hasUuid,
		Uuid:     pUuid,
	}, nil
}

func (p LoginStart) IsValid() (reason error) {
	if len(p.Username) < 1 {
		return fmt.Errorf("username length must be in (0; 16> was %d", len(p.Username))
	}
	return nil
}
