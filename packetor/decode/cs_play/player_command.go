package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type PlayerCommand struct {
	EntityID  int32
	ActionID  int32
	JumpBoost int32
}

func (p PlayerCommand) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	aid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	jumpBoost, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return PlayerCommand{
		EntityID:  eid,
		ActionID:  aid,
		JumpBoost: jumpBoost,
	}, nil
}

func (p PlayerCommand) IsValid() (reason error) {
	if p.ActionID < 0 || 8 < p.ActionID {
		return fmt.Errorf("action id must be in <0; 8> was %d", p.ActionID)
	}
	return nil
}
