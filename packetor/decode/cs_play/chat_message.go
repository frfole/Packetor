package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type ChatMessage struct {
	Message      string
	Timestamp    int64
	Salt         int64
	HasSignature bool
	Signature    []byte
	MessageCount int32
	Acknowledged decode.BitSet
}

func (p ChatMessage) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	msg, err := reader.ReadString0(256)
	if err != nil {
		return nil, err
	}
	ts, err := reader.ReadLong()
	if err != nil {
		return nil, err
	}
	salt, err := reader.ReadLong()
	if err != nil {
		return nil, err
	}
	hasSig, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	var sig []byte
	if hasSig {
		sig, err = reader.ReadBytesExact(256)
		if err != nil {
			return nil, err
		}
	}
	msgCount, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	ack, err := reader.ReadBitSet0(20)
	if err != nil {
		return nil, err
	}
	return ChatMessage{
		Message:      msg,
		Timestamp:    ts,
		Salt:         salt,
		HasSignature: hasSig,
		Signature:    sig,
		MessageCount: msgCount,
		Acknowledged: ack,
	}, nil
}

func (p ChatMessage) IsValid() (reason error) {
	if p.MessageCount < 0 {
		return fmt.Errorf("cannot acknowledge less than 0 messages")
	}
	return nil
}
