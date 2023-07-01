package sc_play

import "Packetor/packetor/decode"

type SetHeadRotation struct {
	EntityID int32
	HeadYaw  decode.Angle
}

func (p SetHeadRotation) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	yaw, err := reader.ReadAngle()
	if err != nil {
		return nil, err
	}
	return SetHeadRotation{
		EntityID: eid,
		HeadYaw:  yaw,
	}, nil
}

func (p SetHeadRotation) IsValid() (reason error) {
	return nil
}
