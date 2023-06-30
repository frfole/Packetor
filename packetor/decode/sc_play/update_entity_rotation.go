package sc_play

import "Packetor/packetor/decode"

type UpdateEntityRotation struct {
	EntityID int32
	Yaw      decode.Angle
	Pitch    decode.Angle
	OnGround bool
}

func (p UpdateEntityRotation) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	yaw, err := reader.ReadAngle()
	if err != nil {
		return nil, err
	}
	pitch, err := reader.ReadAngle()
	if err != nil {
		return nil, err
	}
	onGround, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return UpdateEntityRotation{
		EntityID: eid,
		Yaw:      yaw,
		Pitch:    pitch,
		OnGround: onGround,
	}, nil
}

func (p UpdateEntityRotation) IsValid() (reason error) {
	return nil
}
