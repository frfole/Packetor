package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type ChangeDifficulty struct {
	Difficulty uint8
	Locked     bool
}

func (p ChangeDifficulty) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	difficulty, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	locked, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return ChangeDifficulty{
		Difficulty: difficulty,
		Locked:     locked,
	}, nil
}

func (p ChangeDifficulty) IsValid() (reason error) {
	if 3 < p.Difficulty {
		return fmt.Errorf("unknown difficulty %d", p.Difficulty)
	}
	return nil
}
