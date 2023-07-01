package sc_play

import "Packetor/packetor/decode"

type SetEntityVelocity struct {
	EntityID  int32
	VelocityX int16
	VelocityY int16
	VelocityZ int16
}

func (p SetEntityVelocity) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	vx, err := reader.ReadShort()
	if err != nil {
		return nil, err
	}
	vy, err := reader.ReadShort()
	if err != nil {
		return nil, err
	}
	vz, err := reader.ReadShort()
	if err != nil {
		return nil, err
	}
	return SetEntityVelocity{
		EntityID:  eid,
		VelocityX: vx,
		VelocityY: vy,
		VelocityZ: vz,
	}, nil
}

func (p SetEntityVelocity) IsValid() (reason error) {
	return nil
}
