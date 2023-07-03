package cs_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type EditBook struct {
	Slot     int32
	Pages    []string
	HasTitle bool
	Title    string
}

func (p EditBook) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	slot, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	pagesLen, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if pagesLen < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", pagesLen), error2.ErrDecodeTooSmall)
	}
	pages := make([]string, pagesLen)
	for i := int32(0); i < pagesLen; i++ {
		pages[i], err = reader.ReadString0(8192)
		if err != nil {
			return nil, err
		}
	}
	hasTitle, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	title := ""
	if hasTitle {
		title, err = reader.ReadString0(128)
		if err != nil {
			return nil, err
		}
	}
	return EditBook{
		Slot:     slot,
		Pages:    pages,
		HasTitle: hasTitle,
		Title:    title,
	}, nil
}

func (p EditBook) IsValid() (reason error) {
	return nil
}
