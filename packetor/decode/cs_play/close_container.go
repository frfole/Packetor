package cs_play

import "Packetor/packetor/decode"

type CloseContainer struct {
	WindowID int8
}

func (p CloseContainer) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	wid, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	return CloseContainer{WindowID: wid}, nil
}

func (p CloseContainer) IsValid() (reason error) {
	return nil
}
