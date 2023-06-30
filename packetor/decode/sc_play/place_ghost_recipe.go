package sc_play

import "Packetor/packetor/decode"

type PlaceGhostRecipe struct {
	WindowID int8
	Recipe   string
}

func (p PlaceGhostRecipe) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	wid, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	recipe, err := reader.ReadIdentifier()
	if err != nil {
		return nil, err
	}
	return PlaceGhostRecipe{
		WindowID: wid,
		Recipe:   recipe,
	}, nil
}

func (p PlaceGhostRecipe) IsValid() (reason error) {
	return nil
}
