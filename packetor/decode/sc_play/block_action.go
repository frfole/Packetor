package sc_play

import "Packetor/packetor/decode"

type BlockAction struct {
	Location    decode.Position
	ActionID    uint8
	ActionParam uint8
	BlockType   int32
}

func (p BlockAction) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	aid, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	ap, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	bType, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return BlockAction{
		Location:    loc,
		ActionID:    aid,
		ActionParam: ap,
		BlockType:   bType,
	}, nil
}

func (p BlockAction) IsValid() (reason error) {
	return nil
}
