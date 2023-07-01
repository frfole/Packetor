package sc_play

import "Packetor/packetor/decode"

type SetTitleText struct {
	Text string
}

func (p SetTitleText) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	text, err := reader.ReadChat()
	if err != nil {
		return nil, err
	}
	return SetTitleText{Text: text}, nil
}

func (p SetTitleText) IsValid() (reason error) {
	return nil
}
