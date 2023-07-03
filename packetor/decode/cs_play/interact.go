package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type Interact struct {
	EntityID int32
	Type     int32
	TargetX  float32
	TargetY  float32
	TargetZ  float32
	Hand     int32
	Sneaking bool
}

func (p Interact) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	iType, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	tx := float32(0)
	ty := float32(0)
	tz := float32(0)
	hand := int32(0)
	if iType == 2 {
		tx, err = reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		ty, err = reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		tz, err = reader.ReadFloat()
		if err != nil {
			return nil, err
		}
	}
	if iType == 0 || iType == 2 {
		hand, err = reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
	}
	sneaking, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return Interact{
		EntityID: eid,
		Type:     iType,
		TargetX:  tx,
		TargetY:  ty,
		TargetZ:  tz,
		Hand:     hand,
		Sneaking: sneaking,
	}, nil
}

func (p Interact) IsValid() (reason error) {
	if p.Type < 0 || 2 < p.Type {
		return fmt.Errorf("type must be in <0; 2> was %d", p.Type)
	}
	if p.Type == 0 || p.Type == 2 {
		if p.Hand < 0 || 1 < p.Hand {
			return fmt.Errorf("hand must be in <0; 1> was %d", p.Hand)
		}
	}
	return nil
}
