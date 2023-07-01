package sc_play

import "Packetor/packetor/decode"

type SetActionBarText struct {
	Text string
}

func (p SetActionBarText) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	text, err := reader.ReadChat()
	if err != nil {
		return nil, err
	}
	return SetActionBarText{Text: text}, nil
}

func (p SetActionBarText) IsValid() (reason error) {
	return nil
}
