package sc_play

import "Packetor/packetor/decode"

type Ping struct {
	ID int32
}

func (p Ping) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	id, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	return Ping{ID: id}, nil
}

func (p Ping) IsValid() (reason error) {
	return nil
}
