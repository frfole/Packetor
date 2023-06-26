package sc_play

import "Packetor/packetor/decode"

type SpawnExperienceOrb struct {
	EntityID int32
	X        float64
	Y        float64
	Z        float64
	Count    int16
}

func (p SpawnExperienceOrb) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	x, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	y, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	z, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadShort()
	if err != nil {
		return nil, err
	}
	return SpawnExperienceOrb{
		EntityID: eid,
		X:        x,
		Y:        y,
		Z:        z,
		Count:    count,
	}, nil
}

func (p SpawnExperienceOrb) IsValid() (reason error) {
	return nil
}
