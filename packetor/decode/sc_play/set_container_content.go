package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type SetContainerContent struct {
	WindowID    uint8
	StateID     int32
	Slots       []decode.Slot
	CarriedItem decode.Slot
}

func (p SetContainerContent) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	wid, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	sid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	slots := make([]decode.Slot, count)
	for i := int32(0); i < count; i++ {
		slots[i], err = reader.ReadSlot()
		if err != nil {
			return nil, err
		}
	}
	carried, err := reader.ReadSlot()
	if err != nil {
		return nil, err
	}
	return SetContainerContent{
		WindowID:    wid,
		StateID:     sid,
		Slots:       slots,
		CarriedItem: carried,
	}, nil
}

func (p SetContainerContent) IsValid() (reason error) {
	// TODO: validate?
	return nil
}
