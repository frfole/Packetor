package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type SetHeldItem struct {
	Slot int8
}

func (p SetHeldItem) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	slot, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	return SetHeldItem{Slot: slot}, nil
}

func (p SetHeldItem) IsValid() (reason error) {
	if p.Slot < 0 || 8 < p.Slot {
		return fmt.Errorf("slot must be in <0; 8> was %d", p.Slot)
	}
	return nil
}
