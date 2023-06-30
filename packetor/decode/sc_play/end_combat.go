package sc_play

import "Packetor/packetor/decode"

type EndCombat struct {
	Duration int32
}

func (p EndCombat) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	dur, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return EndCombat{Duration: dur}, nil
}

func (p EndCombat) IsValid() (reason error) {
	return nil
}
