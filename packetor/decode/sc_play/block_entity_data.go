package sc_play

import (
	"Packetor/packetor/decode"
	"Packetor/packetor/nbt"
)

type BlockEntityData struct {
	Location decode.Position
	Type     int32
	Data     nbt.Compound
}

func (p BlockEntityData) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	bType, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	compound, err := reader.ReadNbt()
	if err != nil {
		return nil, err
	}
	return BlockEntityData{
		Location: loc,
		Type:     bType,
		Data:     compound,
	}, nil
}

func (p BlockEntityData) IsValid() (reason error) {
	return nil
}
