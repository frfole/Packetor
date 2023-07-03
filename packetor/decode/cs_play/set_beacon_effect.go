package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type SetBeaconEffect struct {
	HasPrimaryEffect   bool
	PrimaryEffect      int32
	HasSecondaryEffect bool
	SecondaryEffect    int32
}

func (p SetBeaconEffect) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	has1, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	effect1 := int32(0)
	if has1 {
		effect1, err = reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
	}
	has2, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	effect2 := int32(0)
	if has2 {
		effect2, err = reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
	}
	return SetBeaconEffect{
		HasPrimaryEffect:   has1,
		PrimaryEffect:      effect1,
		HasSecondaryEffect: has2,
		SecondaryEffect:    effect2,
	}, nil
}

func (p SetBeaconEffect) IsValid() (reason error) {
	if p.HasPrimaryEffect {
		if p.PrimaryEffect < 0 {
			return fmt.Errorf("primary effect must be atleast 0 was %d", p.PrimaryEffect)
		}
	}
	if p.HasSecondaryEffect {
		if p.SecondaryEffect < 0 {
			return fmt.Errorf("secondary effect must be atleast 0 was %d", p.SecondaryEffect)
		}
	}
	return nil
}
