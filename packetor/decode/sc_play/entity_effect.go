package sc_play

import (
	"Packetor/packetor/decode"
	"Packetor/packetor/nbt"
	"fmt"
)

type EntityEffect struct {
	EntityID      int32
	EffectID      int32
	Amplifier     uint8
	Duration      int32
	Flags         uint8
	HasFactorData bool
	FactorCodec   nbt.Compound
}

func (p EntityEffect) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	effectId, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	amplifier, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	dur, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	flags, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	hasData, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	var factor nbt.Compound
	if hasData {
		factor, err = reader.ReadNbt()
		if err != nil {
			return nil, err
		}
	}
	return EntityEffect{
		EntityID:      eid,
		EffectID:      effectId,
		Amplifier:     amplifier,
		Duration:      dur,
		Flags:         flags,
		HasFactorData: hasData,
		FactorCodec:   factor,
	}, nil
}

func (p EntityEffect) IsValid() (reason error) {
	// TODO: validate effect id and factor codec
	if (p.Flags &^ 0b111) != 0 {
		return fmt.Errorf("unknown flags %b", p.Flags)
	}
	return nil
}
