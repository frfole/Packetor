package sc_play

import "Packetor/packetor/decode"

type EntityEvent struct {
	EntityID     int32
	EntityStatus int8
}

func (p EntityEvent) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	status, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	return EntityEvent{
		EntityID:     eid,
		EntityStatus: status,
	}, nil
}

func (p EntityEvent) IsValid() (reason error) {
	return nil
}
