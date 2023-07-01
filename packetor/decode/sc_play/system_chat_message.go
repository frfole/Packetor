package sc_play

import "Packetor/packetor/decode"

type SystemChatMessage struct {
	Content string
	Overlay bool
}

func (p SystemChatMessage) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	text, err := reader.ReadChat()
	if err != nil {
		return nil, err
	}
	overlay, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return SystemChatMessage{
		Content: text,
		Overlay: overlay,
	}, nil
}

func (p SystemChatMessage) IsValid() (reason error) {
	return nil
}
