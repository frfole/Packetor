package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type SoundEffect struct {
	SoundID         int32
	SoundName       string
	HasFixedRange   bool
	Range           float32
	SoundCategory   int32
	EffectPositionX int32
	EffectPositionY int32
	EffectPositionZ int32
	Volume          float32
	Pitch           float32
	Seed            int64
}

func (p SoundEffect) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	sid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	sn := ""
	hasFixedRange := false
	fixedRange := float32(0)
	if sid == 0 {
		sn, err = reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
		hasFixedRange, err = reader.ReadBoolean()
		if err != nil {
			return nil, err
		}
		if hasFixedRange {
			fixedRange, err = reader.ReadFloat()
			if err != nil {
				return nil, err
			}
		}
	}
	cat, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	x, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	y, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	z, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	vol, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	pitch, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	seed, err := reader.ReadLong()
	if err != nil {
		return nil, err
	}
	return SoundEffect{
		SoundID:         sid,
		SoundName:       sn,
		HasFixedRange:   hasFixedRange,
		Range:           fixedRange,
		SoundCategory:   cat,
		EffectPositionX: x,
		EffectPositionY: y,
		EffectPositionZ: z,
		Volume:          vol,
		Pitch:           pitch,
		Seed:            seed,
	}, nil
}

func (p SoundEffect) IsValid() (reason error) {
	if p.SoundCategory < 0 || 9 < p.SoundCategory {
		return fmt.Errorf("category must be in <0; 9> was %d", p.SoundCategory)
	}
	return nil
}
