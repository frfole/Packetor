package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type FeatureFlags struct {
	Features []string
}

func (p FeatureFlags) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	features := make([]string, count)
	for i := int32(0); i < count; i++ {
		features[i], err = reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
	}
	return FeatureFlags{Features: features}, nil
}

func (p FeatureFlags) IsValid() (reason error) {
	return nil
}
