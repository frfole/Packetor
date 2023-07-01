package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type EntitySoundEffect struct {
	SoundID       int32
	SoundName     string
	HasFixedRange bool
	Range         float32
	SoundCategory int32
	EntityID      int32
	Volume        float32
	Pitch         float32
	Seed          int64
}

func (p EntitySoundEffect) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
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
	eid, err := reader.ReadVarInt()
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
	return EntitySoundEffect{
		SoundID:       sid,
		SoundName:     sn,
		HasFixedRange: hasFixedRange,
		Range:         fixedRange,
		SoundCategory: cat,
		EntityID:      eid,
		Volume:        vol,
		Pitch:         pitch,
		Seed:          seed,
	}, nil
}

func (p EntitySoundEffect) IsValid() (reason error) {
	if p.SoundCategory < 0 || 9 < p.SoundCategory {
		return fmt.Errorf("category must be in <0; 9> was %d", p.SoundCategory)
	}
	return nil
}
