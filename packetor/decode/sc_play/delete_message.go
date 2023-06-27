package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type DeleteMessage struct {
	Signature []byte
}

func (p DeleteMessage) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	sig, err := reader.ReadBytesExact(int(count))
	if err != nil {
		return nil, err
	}
	return DeleteMessage{Signature: sig}, nil
}

func (p DeleteMessage) IsValid() (reason error) {
	return nil
}
