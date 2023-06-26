package sc_play

import "Packetor/packetor/decode"

type SetBlockDestroyStage struct {
	EntityID     int32
	Location     decode.Position
	DestroyStage uint8
}

func (p SetBlockDestroyStage) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	stage, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	return SetBlockDestroyStage{
		EntityID:     eid,
		Location:     loc,
		DestroyStage: stage,
	}, nil
}

func (p SetBlockDestroyStage) IsValid() (reason error) {
	return nil
}
