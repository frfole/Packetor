package cs_play

import "Packetor/packetor/decode"

type QueryBlockEntityTag struct {
	TransactionID int32
	Location      decode.Position
}

func (p QueryBlockEntityTag) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	tid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	return QueryBlockEntityTag{
		TransactionID: tid,
		Location:      loc,
	}, nil
}

func (p QueryBlockEntityTag) IsValid() (reason error) {
	return nil
}
