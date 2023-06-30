package sc_play

import "Packetor/packetor/decode"

type CombatDeath struct {
	PlayerID int32
	Message  string
}

func (p CombatDeath) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	pid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	msg, err := reader.ReadChat()
	if err != nil {
		return nil, err
	}
	return CombatDeath{
		PlayerID: pid,
		Message:  msg,
	}, nil
}

func (p CombatDeath) IsValid() (reason error) {
	return nil
}
