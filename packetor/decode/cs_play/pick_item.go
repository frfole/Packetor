package cs_play

import "Packetor/packetor/decode"

type PickItem struct {
	Slot int32
}

func (p PickItem) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	slot, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return PickItem{Slot: slot}, nil
}

func (p PickItem) IsValid() (reason error) {
	return nil
}
