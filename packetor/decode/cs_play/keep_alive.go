package cs_play

import "Packetor/packetor/decode"

type KeepAlive struct {
	KeepAliveID int64
}

func (p KeepAlive) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	id, err := reader.ReadLong()
	if err != nil {
		return nil, err
	}
	return KeepAlive{KeepAliveID: id}, nil
}

func (p KeepAlive) IsValid() (reason error) {
	return nil
}
