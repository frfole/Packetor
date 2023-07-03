package cs_play

import "Packetor/packetor/decode"

type SetCreativeModeSlot struct {
	Slot int16
	Item decode.Slot
}

func (p SetCreativeModeSlot) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	slot, err := reader.ReadShort()
	if err != nil {
		return nil, err
	}
	item, err := reader.ReadSlot()
	if err != nil {
		return nil, err
	}
	return SetCreativeModeSlot{
		Slot: slot,
		Item: item,
	}, nil
}

func (p SetCreativeModeSlot) IsValid() (reason error) {
	return nil
}
