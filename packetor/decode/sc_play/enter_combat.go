package sc_play

import "Packetor/packetor/decode"

type EnterCombat struct {
}

func (p EnterCombat) Read(_ decode.PacketReader) (packet decode.Packet, err error) {
	return EnterCombat{}, nil
}

func (p EnterCombat) IsValid() (reason error) {
	return nil
}
