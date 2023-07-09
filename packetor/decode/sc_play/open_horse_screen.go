package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type OpenHorseScreen struct {
	WindowID  uint8
	SlotCount int32
	EntityID  int32
}

func (p OpenHorseScreen) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	wid, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	eid, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	return OpenHorseScreen{
		WindowID:  wid,
		SlotCount: count,
		EntityID:  eid,
	}, nil
}

func (p OpenHorseScreen) IsValid() (reason error) {
	if p.SlotCount < 0 {
		return fmt.Errorf("slot count must be atleast 0 was %v", p.SlotCount)
	}
	return nil
}
