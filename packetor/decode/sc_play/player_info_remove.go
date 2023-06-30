package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
)

type PlayerInfoRemove struct {
	Players []uuid.UUID
}

func (p PlayerInfoRemove) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	players := make([]uuid.UUID, count)
	for i := int32(0); i < count; i++ {
		players[i], err = reader.ReadUuid()
		if err != nil {
			return nil, err
		}
	}
	return PlayerInfoRemove{Players: players}, nil
}

func (p PlayerInfoRemove) IsValid() (reason error) {
	return nil
}
