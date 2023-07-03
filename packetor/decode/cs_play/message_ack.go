package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type MessageAck struct {
	MessageCount int32
}

func (p MessageAck) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return MessageAck{MessageCount: count}, nil
}

func (p MessageAck) IsValid() (reason error) {
	if p.MessageCount < 0 {
		return fmt.Errorf("cannot acknowledge less than 0 messages")
	}
	return nil
}
