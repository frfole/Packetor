package sc_play

import "Packetor/packetor/decode"

type SetHealth struct {
	Health         float32
	Food           int32
	FoodSaturation float32
}

func (p SetHealth) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	health, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	food, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	saturation, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	return SetHealth{
		Health:         health,
		Food:           food,
		FoodSaturation: saturation,
	}, nil
}

func (p SetHealth) IsValid() (reason error) {
	return nil
}
