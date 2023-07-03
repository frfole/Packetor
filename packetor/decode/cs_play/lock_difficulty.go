package cs_play

import "Packetor/packetor/decode"

type LockDifficulty struct {
	Locked bool
}

func (p LockDifficulty) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	locked, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return LockDifficulty{Locked: locked}, nil
}

func (p LockDifficulty) IsValid() (reason error) {
	return nil
}
