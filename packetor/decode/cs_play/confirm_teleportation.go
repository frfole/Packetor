package cs_play

import "Packetor/packetor/decode"

type ConfirmTeleportation struct {
	TeleportID int32
}

func (p ConfirmTeleportation) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	tid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return ConfirmTeleportation{TeleportID: tid}, nil
}

func (p ConfirmTeleportation) IsValid() (reason error) {
	return nil
}
