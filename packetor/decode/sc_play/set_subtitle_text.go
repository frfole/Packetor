package sc_play

import "Packetor/packetor/decode"

type SetSubtitleText struct {
	Text string
}

func (p SetSubtitleText) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	text, err := reader.ReadChat()
	if err != nil {
		return nil, err
	}
	return SetSubtitleText{Text: text}, nil
}

func (p SetSubtitleText) IsValid() (reason error) {
	return nil
}
