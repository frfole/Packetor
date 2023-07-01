package sc_play

import "Packetor/packetor/decode"

type PickupItem struct {
	CollectedID int32
	CollectorID int32
	ItemCount   int32
}

func (p PickupItem) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	collected, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	collector, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	amount, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return PickupItem{
		CollectedID: collected,
		CollectorID: collector,
		ItemCount:   amount,
	}, nil
}

func (p PickupItem) IsValid() (reason error) {
	return nil
}
