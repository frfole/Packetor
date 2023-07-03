package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type SelectTrade struct {
	SelectedSlot int32
}

func (p SelectTrade) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	slot, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return SelectTrade{SelectedSlot: slot}, nil
}

func (p SelectTrade) IsValid() (reason error) {
	if p.SelectedSlot < 0 {
		return fmt.Errorf("selected slot must be atleast 0 was %d", p.SelectedSlot)
	}
	return nil
}
