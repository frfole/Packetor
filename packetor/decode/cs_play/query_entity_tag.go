package cs_play

import "Packetor/packetor/decode"

type QueryEntityTag struct {
	TransactionID int32
	EntityID      int32
}

func (p QueryEntityTag) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	tid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return QueryEntityTag{
		TransactionID: tid,
		EntityID:      eid,
	}, nil
}

func (p QueryEntityTag) IsValid() (reason error) {
	return nil
}
