package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type RemoveEntities struct {
	EntityIDs []int32
}

func (p RemoveEntities) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	eids := make([]int32, count)
	for i := int32(0); i < count; i++ {
		eids[i], err = reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
	}
	return RemoveEntities{EntityIDs: eids}, nil
}

func (p RemoveEntities) IsValid() (reason error) {
	return nil
}
