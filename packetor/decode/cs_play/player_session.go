package cs_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
)

type PlayerSession struct {
	SessionID    uuid.UUID
	PubKeyExpiry int64
	PubKey       []byte
	KeySig       []byte
}

func (p PlayerSession) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	sid, err := reader.ReadUuid()
	if err != nil {
		return nil, err
	}
	expiry, err := reader.ReadLong()
	if err != nil {
		return nil, err
	}
	pkLen, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if pkLen < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", pkLen), error2.ErrDecodeTooSmall)
	}
	pk, err := reader.ReadBytesExact(int(pkLen))
	if err != nil {
		return nil, err
	}
	sigLen, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if sigLen < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", sigLen), error2.ErrDecodeTooSmall)
	}
	sig, err := reader.ReadBytesExact(int(sigLen))
	if err != nil {
		return nil, err
	}
	return PlayerSession{
		SessionID:    sid,
		PubKeyExpiry: expiry,
		PubKey:       pk,
		KeySig:       sig,
	}, nil
}

func (p PlayerSession) IsValid() (reason error) {
	return nil
}
