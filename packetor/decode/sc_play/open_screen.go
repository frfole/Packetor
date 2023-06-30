package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type OpenScreen struct {
	WindowID    int32
	WindowType  int32
	WindowTitle string
}

func (p OpenScreen) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	wid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	wt, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	title, err := reader.ReadChat()
	if err != nil {
		return nil, err
	}
	return OpenScreen{
		WindowID:    wid,
		WindowType:  wt,
		WindowTitle: title,
	}, nil
}

func (p OpenScreen) IsValid() (reason error) {
	if p.WindowType < 0 || 23 < p.WindowType {
		return fmt.Errorf("window type must be in <0; 23> was %d", p.WindowType)
	}
	return nil
}
