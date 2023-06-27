package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type SetContainerSlot struct {
	WindowID int8
	StateID  int32
	Slot     int16
	Data     decode.Slot
}

func (p SetContainerSlot) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	wid, err := reader.ReadSByte()
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
	data, err := reader.ReadSlot()
	if err != nil {
		return nil, err
	}
	return SetContainerSlot{
		WindowID: wid,
		StateID:  sid,
		Slot:     slot,
		Data:     data,
	}, nil
}

func (p SetContainerSlot) IsValid() (reason error) {
	if p.Slot < 0 {
		return fmt.Errorf("slot must be atleast 0 was %d", p.Slot)
	}
	return nil
}
