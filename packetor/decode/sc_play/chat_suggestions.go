package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type ChatSuggestions struct {
	Action  int32
	Entries []string
}

func (p ChatSuggestions) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	action, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	entries := make([]string, count)
	for i := int32(0); i < count; i++ {
		entries[i], err = reader.ReadString()
		if err != nil {
			return nil, err
		}
	}
	return ChatSuggestions{
		Action:  action,
		Entries: entries,
	}, nil
}

func (p ChatSuggestions) IsValid() (reason error) {
	if p.Action < 0 || 2 < p.Action {
		return fmt.Errorf("action must be in range <0; 3> was %d", p.Action)
	}
	return nil
}
