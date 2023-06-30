package sc_play

import "Packetor/packetor/decode"

type UpdateEntityPositionRotation struct {
	EntityID int32
	DeltaX   int16
	DeltaY   int16
	DeltaZ   int16
	Yaw      decode.Angle
	Pitch    decode.Angle
	OnGround bool
}

func (p UpdateEntityPositionRotation) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	dx, err := reader.ReadShort()
	if err != nil {
		return nil, err
	}
	dy, err := reader.ReadShort()
	if err != nil {
		return nil, err
	}
	dz, err := reader.ReadShort()
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
	return UpdateEntityPositionRotation{
		EntityID: eid,
		DeltaX:   dx,
		DeltaY:   dy,
		DeltaZ:   dz,
		Yaw:      yaw,
		Pitch:    pitch,
		OnGround: onGround,
	}, nil
}

func (p UpdateEntityPositionRotation) IsValid() (reason error) {
	return nil
}
