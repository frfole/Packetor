package sc_play

import "Packetor/packetor/decode"

type UpdateEntityPosition struct {
	EntityID int32
	DeltaX   int16
	DeltaY   int16
	DeltaZ   int16
	OnGround bool
}

func (p UpdateEntityPosition) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
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
	onGround, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return UpdateEntityPosition{
		EntityID: eid,
		DeltaX:   dx,
		DeltaY:   dy,
		DeltaZ:   dz,
		OnGround: onGround,
	}, nil
}

func (p UpdateEntityPosition) IsValid() (reason error) {
	return nil
}
