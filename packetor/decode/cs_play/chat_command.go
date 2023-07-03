package cs_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type ChatCommand struct {
	Command      string
	Timestamp    int64
	Salt         int64
	Arguments    map[string][]byte
	MessageCount int32
	Acknowledged decode.BitSet
}

func (p ChatCommand) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	cmd, err := reader.ReadString0(256)
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
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	arguments := map[string][]byte{}
	for i := int32(0); i < count; i++ {
		argName, err := reader.ReadString()
		if err != nil {
			return nil, err
		}
		arguments[argName], err = reader.ReadBytesExact(256)
		if err != nil {
			return nil, err
		}
	}
	msgCount, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	acknowledged, err := reader.ReadBitSet0(20)
	if err != nil {
		return nil, err
	}
	return ChatCommand{
		Command:      cmd,
		Timestamp:    ts,
		Salt:         salt,
		Arguments:    arguments,
		MessageCount: msgCount,
		Acknowledged: acknowledged,
	}, nil
}

func (p ChatCommand) IsValid() (reason error) {
	if p.MessageCount < 0 {
		return fmt.Errorf("cannot acknowledge less than 0 messages")
	}
	return nil
}
