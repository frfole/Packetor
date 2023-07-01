package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type UpdateTags struct {
	Tags map[string]map[string][]int32
}

func (p UpdateTags) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	count1, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count1 < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count1), error2.ErrDecodeTooSmall)
	}
	tags := map[string]map[string][]int32{}
	for i := int32(0); i < count1; i++ {
		tagType, err := reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
		count2, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if count2 < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count2), error2.ErrDecodeTooSmall)
		}
		tags[tagType] = map[string][]int32{}
		for j := int32(0); j < count2; j++ {
			tagName, err := reader.ReadIdentifier()
			if err != nil {
				return nil, err
			}
			count3, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			} else if count3 < 0 {
				return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count3), error2.ErrDecodeTooSmall)
			}
			tags[tagType][tagName] = make([]int32, count3)
			for k := int32(0); k < count3; k++ {
				bid, err := reader.ReadVarInt()
				if err != nil {
					return nil, err
				}
				tags[tagType][tagName][k] = bid
			}
		}
	}
	return UpdateTags{Tags: tags}, nil
}

func (p UpdateTags) IsValid() (reason error) {
	return nil
}
