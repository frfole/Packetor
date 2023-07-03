package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type ClientCommand struct {
	ActionID int32
}

func (p ClientCommand) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	action, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return ClientCommand{ActionID: action}, nil
}

func (p ClientCommand) IsValid() (reason error) {
	if p.ActionID < 0 || 1 < p.ActionID {
		return fmt.Errorf("action must be in <0; 1> was %d", p.ActionID)
	}
	return nil
}
