package cs_play

import "Packetor/packetor/decode"

type PlaceRecipe struct {
	WindowID int8
	Recipe   string
	MakeAll  bool
}

func (p PlaceRecipe) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	wid, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	recipe, err := reader.ReadIdentifier()
	if err != nil {
		return nil, err
	}
	makeAll, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return PlaceRecipe{
		WindowID: wid,
		Recipe:   recipe,
		MakeAll:  makeAll,
	}, nil
}

func (p PlaceRecipe) IsValid() (reason error) {
	return nil
}
