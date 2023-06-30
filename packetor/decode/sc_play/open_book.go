package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type OpenBook struct {
	Hand int32
}

func (p OpenBook) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	hand, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return OpenBook{Hand: hand}, nil
}

func (p OpenBook) IsValid() (reason error) {
	if p.Hand < 0 || 1 < p.Hand {
		return fmt.Errorf("hand must be in range <0; 1> was %d", p.Hand)
	}
	return nil
}
