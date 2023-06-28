package sc_play

import "Packetor/packetor/decode"

type HurtAnimation struct {
	EntityID int32
	Yaw      float32
}

func (p HurtAnimation) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	yaw, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	return HurtAnimation{
		EntityID: eid,
		Yaw:      yaw,
	}, nil
}

func (p HurtAnimation) IsValid() (reason error) {
	return nil
}
