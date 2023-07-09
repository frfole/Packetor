package cs_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type ClickContainer struct {
	WindowID    uint8
	StateID     int32
	Slot        int16
	Button      uint8
	Mode        int32
	Items       map[int16]decode.Slot
	CarriedItem decode.Slot
}

func (p ClickContainer) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	wid, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	sid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	slot, err := reader.ReadShort()
	if err != nil {
		return nil, err
	}
	button, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	mode, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	itemLen, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if itemLen < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", itemLen), error2.ErrDecodeTooSmall)
	}
	items := map[int16]decode.Slot{}
	for i := int32(0); i < itemLen; i++ {
		idx, err := reader.ReadShort()
		if err != nil {
			return nil, err
		}
		items[idx], err = reader.ReadSlot()
		if err != nil {
			return nil, err
		}
	}
	carried, err := reader.ReadSlot()
	if err != nil {
		return nil, err
	}
	return ClickContainer{
		WindowID:    wid,
		StateID:     sid,
		Slot:        slot,
		Button:      button,
		Mode:        mode,
		Items:       items,
		CarriedItem: carried,
	}, nil
}

func (p ClickContainer) IsValid() (reason error) {
	// TODO: validate
	return nil
}
