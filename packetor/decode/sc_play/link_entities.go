package sc_play

import "Packetor/packetor/decode"

type LinkEntities struct {
	AttachedEntityID int32
	HoldingEntityID  int32
}

func (p LinkEntities) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	attached, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	holding, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	return LinkEntities{
		AttachedEntityID: attached,
		HoldingEntityID:  holding,
	}, nil
}

func (p LinkEntities) IsValid() (reason error) {
	return nil
}
