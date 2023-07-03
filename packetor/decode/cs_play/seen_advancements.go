package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type SeenAdvancements struct {
	Action int32
	TabID  string
}

func (p SeenAdvancements) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	action, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	tid := ""
	if action == 0 {
		tid, err = reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
	}
	return SeenAdvancements{
		Action: action,
		TabID:  tid,
	}, nil
}

func (p SeenAdvancements) IsValid() (reason error) {
	if p.Action < 0 || 1 < p.Action {
		return fmt.Errorf("action must be in <0; 1> was %d", p.Action)
	}
	return nil
}
