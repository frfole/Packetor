package cs_play

import "Packetor/packetor/decode"

type ClickContainerButton struct {
	WindowID int8
	ButtonID int8
}

func (p ClickContainerButton) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	wid, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	bid, err := reader.ReadSByte()
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
