package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type UpdateObjectives struct {
	ObjectiveName  string
	Mode           int8
	ObjectiveValue string
	Type           int32
}

func (p UpdateObjectives) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	name, err := reader.ReadString()
	mode, err := reader.ReadSByte()
	value := ""
	oType := int32(0)
	switch mode {
	case 0, 2:
		value, err = reader.ReadChat()
		if err != nil {
			return nil, err
		}
		oType, err = reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
	}
	return UpdateObjectives{
		ObjectiveName:  name,
		Mode:           mode,
		ObjectiveValue: value,
		Type:           oType,
	}, nil
}

func (p UpdateObjectives) IsValid() (reason error) {
	if p.Mode < 0 || 2 < p.Mode {
		return fmt.Errorf("mode must be in <0; 2> was %d", p.Mode)
	}
	if p.Mode == 0 || p.Mode == 2 {
		if p.Type < 0 || 1 < p.Type {
			return fmt.Errorf("type must be in <0; 1> was %d", p.Type)
		}
	}
	return nil
}
