package cs_play

import "Packetor/packetor/decode"

type Pong struct {
	ID int32
}

func (p Pong) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	id, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	return Pong{ID: id}, nil
}

func (p Pong) IsValid() (reason error) {
	return nil
}
