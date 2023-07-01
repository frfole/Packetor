package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type LookAt struct {
	Reference    int32
	TargetX      float64
	TargetY      float64
	TargetZ      float64
	IsEntity     bool
	EntityID     int32
	EntityTarget int32
}

func (p LookAt) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	ref, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	tx, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	ty, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	tz, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	isEntity, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	eid := int32(0)
	eTarget := int32(0)
	if isEntity {
		eid, err = reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		eTarget, err = reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
	}
	return LookAt{
		Reference:    ref,
		TargetX:      tx,
		TargetY:      ty,
		TargetZ:      tz,
		IsEntity:     isEntity,
		EntityID:     eid,
		EntityTarget: eTarget,
	}, nil
}

func (p LookAt) IsValid() (reason error) {
	if p.Reference < 0 || 1 < p.Reference {
		return fmt.Errorf("reference must be in <0; 1> was %d", p.Reference)
	}
	if p.EntityTarget < 0 || 1 < p.EntityTarget {
		return fmt.Errorf("target reference must be in <0; 1> was %d", p.EntityTarget)
	}
	return nil
}
