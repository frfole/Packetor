package sc_play

import "Packetor/packetor/decode"

type SetCooldown struct {
	ItemID        int32
	CooldownTicks int32
}

func (p SetCooldown) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	iid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	t, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return SetCooldown{
		ItemID:        iid,
		CooldownTicks: t,
	}, nil
}

func (p SetCooldown) IsValid() (reason error) {
	return nil
}
