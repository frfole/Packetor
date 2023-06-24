package sc_login

import (
	"Packetor/packetor/decode"
	"fmt"
	"github.com/gofrs/uuid/v5"
)

type LoginSuccessProperty struct {
	Name      string
	Value     string
	IsSigned  bool
	Signature string
}

type LoginSuccess struct {
	Uuid       uuid.UUID
	Username   string
	Properties []LoginSuccessProperty
}

func (p LoginSuccess) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	uid, err := reader.ReadUuid()
	if err != nil {
		return nil, err
	}
	username, err := reader.ReadString(16)
	if err != nil {
		return nil, err
	}
	return LoginSuccess{
		Uuid:       uid,
		Username:   username,
		Properties: []LoginSuccessProperty{},
	}, nil
}

func (p LoginSuccess) IsValid() (reason error) {
	if len(p.Username) < 1 {
		return fmt.Errorf("username length must be in (0; 16> was %d", len(p.Username))
	}
	return nil
}
