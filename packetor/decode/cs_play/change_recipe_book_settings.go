package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type ChangeRecipeBookSettings struct {
	BookID       int32
	BookOpen     bool
	FilterActive bool
}

func (p ChangeRecipeBookSettings) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	bid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	isOpen, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	filter, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return ChangeRecipeBookSettings{
		BookID:       bid,
		BookOpen:     isOpen,
		FilterActive: filter,
	}, nil
}

func (p ChangeRecipeBookSettings) IsValid() (reason error) {
	if p.BookID < 0 || 3 < p.BookID {
		return fmt.Errorf("book id must be in <0; 3> was %d", p.BookID)
	}
	return nil
}
