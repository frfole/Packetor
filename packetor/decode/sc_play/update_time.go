package sc_play

import "Packetor/packetor/decode"

type UpdateTime struct {
	WorldAge  int64
	TimeOfDay int64
}

func (p UpdateTime) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	age, err := reader.ReadLong()
	if err != nil {
		return nil, err
	}
	tod, err := reader.ReadLong()
	if err != nil {
		return nil, err
	}
	return UpdateTime{
		WorldAge:  age,
		TimeOfDay: tod,
	}, nil
}

func (p UpdateTime) IsValid() (reason error) {
	return nil
}
