package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type ChangeDifficulty struct {
	Difficulty uint8
}

func (p ChangeDifficulty) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	difficulty, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	return ChangeDifficulty{Difficulty: difficulty}, nil
}

func (p ChangeDifficulty) IsValid() (reason error) {
	if 3 < p.Difficulty {
		return fmt.Errorf("difficulty must be in range <0; 3> was %d", p.Difficulty)
	}
	return nil
}
