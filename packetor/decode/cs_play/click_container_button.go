package cs_play

import "Packetor/packetor/decode"

type ClickContainerButton struct {
	WindowID uint8
	ButtonID uint8
}

func (p ClickContainerButton) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	wid, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	bid, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	return ClickContainerButton{
		WindowID: wid,
		ButtonID: bid,
	}, nil
}

func (p ClickContainerButton) IsValid() (reason error) {
	return nil
}
