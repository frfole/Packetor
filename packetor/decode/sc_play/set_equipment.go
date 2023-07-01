package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type SetEquipmentEntry struct {
	Slot int8
	Item decode.Slot
}

type SetEquipment struct {
	EntityID  int32
	Equipment []SetEquipmentEntry
}

func (p SetEquipment) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	var entries []SetEquipmentEntry
	slot := uint8(0b1000_0000)
	for (slot &^ 0x7f) != 0 {
		slot, err = reader.ReadUByte()
		if err != nil {
			return nil, err
		}
		item, err := reader.ReadSlot()
		if err != nil {
			return nil, err
		}
		entries = append(entries, SetEquipmentEntry{
			Slot: int8(slot & 0x7f),
			Item: item,
		})
	}
	return SetEquipment{
		EntityID:  eid,
		Equipment: entries,
	}, nil
}

func (p SetEquipment) IsValid() (reason error) {
	for i, entry := range p.Equipment {
		if entry.Slot < 0 || 5 < entry.Slot {
			return fmt.Errorf("slot must be in <0; 5> was %d at index %d", entry.Slot, i)
		}
	}
	return nil
}
