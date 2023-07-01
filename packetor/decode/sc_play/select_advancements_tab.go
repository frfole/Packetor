package sc_play

import "Packetor/packetor/decode"

type SelectAdvancementsTab struct {
	HasID bool
	ID    string
}

func (p SelectAdvancementsTab) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	hasId, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	id := ""
	if hasId {
		id, err = reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
	}
	return SelectAdvancementsTab{
		HasID: hasId,
		ID:    id,
	}, nil
}

func (p SelectAdvancementsTab) IsValid() (reason error) {
	return nil
}
