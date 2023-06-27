package sc_play

import "Packetor/packetor/decode"

type BlockUpdate struct {
	Location decode.Position
	BlockID  int32
}

func (p BlockUpdate) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	bid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return BlockUpdate{
		Location: loc,
		BlockID:  bid,
	}, nil
}

func (p BlockUpdate) IsValid() (reason error) {
	return nil
}
