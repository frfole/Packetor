package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type UpdateScore struct {
	EntityName    string
	Action        int32
	ObjectiveName string
	Value         int32
}

func (p UpdateScore) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eName, err := reader.ReadString()
	if err != nil {
		return nil, err
	}
	action, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	oName, err := reader.ReadString()
	if err != nil {
		return nil, err
	}
	value := int32(0)
	if action == 0 {
		value, err = reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
	}
	return UpdateScore{
		EntityName:    eName,
		Action:        action,
		ObjectiveName: oName,
		Value:         value,
	}, nil
}

func (p UpdateScore) IsValid() (reason error) {
	if p.Action < 0 || 1 < p.Action {
		return fmt.Errorf("action must be in <0; 1> was %d", p.Action)
	}
	return nil
}
