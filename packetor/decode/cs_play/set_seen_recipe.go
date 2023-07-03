package cs_play

import "Packetor/packetor/decode"

type SetSeenRecipe struct {
	RecipeID string
}

func (p SetSeenRecipe) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	rid, err := reader.ReadIdentifier()
	if err != nil {
		return nil, err
	}
	return SetSeenRecipe{RecipeID: rid}, nil
}

func (p SetSeenRecipe) IsValid() (reason error) {
	return nil
}
