package sc_play

import (
	"Packetor/packetor/decode"
	"Packetor/packetor/nbt"
)

type TagQueryResponse struct {
	TransactionID int32
	NBT           nbt.Compound
}

func (p TagQueryResponse) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	tid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data, err := reader.ReadNbt()
	if err != nil {
		return nil, err
	}
	return TagQueryResponse{
		TransactionID: tid,
		NBT:           data,
	}, nil
}

func (p TagQueryResponse) IsValid() (reason error) {
	return nil
}
