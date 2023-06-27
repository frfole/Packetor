package sc_play

import "Packetor/packetor/decode"

type DamageEvent struct {
	EntityID    int32
	SrcTypeID   int32
	SrcCauseId  int32
	SrcDirectID int32
	HasSrcPos   bool
	SrcPosX     float64
	SrcPosY     float64
	SrcPosZ     float64
}

func (p DamageEvent) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	srcTypeId, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	srcCauseId, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	srcDirectId, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	hasPos, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	var posX, posY, posZ float64
	if hasPos {
		posX, err = reader.ReadDouble()
		if err != nil {
			return nil, err
		}
		posY, err = reader.ReadDouble()
		if err != nil {
			return nil, err
		}
		posZ, err = reader.ReadDouble()
		if err != nil {
			return nil, err
		}
	} else {
		posX = 0
		posY = 0
		posZ = 0
	}
	return DamageEvent{
		EntityID:    eid,
		SrcTypeID:   srcTypeId,
		SrcCauseId:  srcCauseId,
		SrcDirectID: srcDirectId,
		HasSrcPos:   hasPos,
		SrcPosX:     posX,
		SrcPosY:     posY,
		SrcPosZ:     posZ,
	}, nil
}

func (p DamageEvent) IsValid() (reason error) {
	return nil
}
