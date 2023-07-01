package sc_play

import "Packetor/packetor/decode"

type SetExperience struct {
	ExperienceBar   float32
	TotalExperience int32
	Level           int32
}

func (p SetExperience) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	bar, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	total, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	level, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return SetExperience{
		ExperienceBar:   bar,
		TotalExperience: total,
		Level:           level,
	}, nil
}

func (p SetExperience) IsValid() (reason error) {
	return nil
}
