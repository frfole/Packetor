package sc_play

import "Packetor/packetor/decode"

type CloseContainer struct {
	WindowID uint8
}

func (p CloseContainer) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	wid, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	return CloseContainer{WindowID: wid}, nil
}

func (p CloseContainer) IsValid() (reason error) {
	return nil
}
