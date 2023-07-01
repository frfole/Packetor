package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type StopSound struct {
	Flags    uint8
	Category int32
	Sound    string
}

func (p StopSound) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	flags, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	cat := int32(0)
	if (flags & 0x01) == 0x01 {
		cat, err = reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
	}
	sound := ""
	if (flags & 0x02) == 0x02 {
		sound, err = reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
	}
	return StopSound{
		Flags:    flags,
		Category: cat,
		Sound:    sound,
	}, nil
}

func (p StopSound) IsValid() (reason error) {
	if (p.Flags &^ 0b11) != 0 {
		return fmt.Errorf("unknown flags %b", p.Flags)
	}
	if (p.Flags & 0x01) == 0x01 {
		if p.Category < 0 || 9 < p.Category {
			return fmt.Errorf("category must be in <0; 9> was %d", p.Category)
		}
	}
	return nil
}
