package sc_play

import "Packetor/packetor/decode"

type RemoveEntityEffect struct {
	EntityID int32
	EffectID int32
}

func (p RemoveEntityEffect) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	effect, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return RemoveEntityEffect{
		EntityID: eid,
		EffectID: effect,
	}, nil
}

func (p RemoveEntityEffect) IsValid() (reason error) {
	return nil
}
