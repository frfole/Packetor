package sc_play

import "Packetor/packetor/decode"

type ClearTitles struct {
	Reset bool
}

func (p ClearTitles) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	rst, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return ClearTitles{Reset: rst}, nil
}

func (p ClearTitles) IsValid() (reason error) {
	return nil
}
