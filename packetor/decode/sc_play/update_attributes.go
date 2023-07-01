package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
)

type UpdateAttributesModifier struct {
	Uuid      uuid.UUID
	Amount    float64
	Operation uint8
}

type UpdateAttributesEntry struct {
	Value     float64
	Modifiers []UpdateAttributesModifier
}

type UpdateAttributes struct {
	EntityID int32
	Entries  map[string]UpdateAttributesEntry
}

func (p UpdateAttributes) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	count1, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count1 < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count1), error2.ErrDecodeTooSmall)
	}
	entries := map[string]UpdateAttributesEntry{}
	for i := int32(0); i < count1; i++ {
		key, err := reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
		value, err := reader.ReadDouble()
		if err != nil {
			return nil, err
		}
		count2, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if count2 < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count2), error2.ErrDecodeTooSmall)
		}
		modifiers := make([]UpdateAttributesModifier, count2)
		for j := int32(0); j < count2; j++ {
			uid, err := reader.ReadUuid()
			if err != nil {
				return nil, err
			}
			amount, err := reader.ReadDouble()
			if err != nil {
				return nil, err
			}
			op, err := reader.ReadUByte()
			if err != nil {
				return nil, err
			}
			modifiers[j] = UpdateAttributesModifier{
				Uuid:      uid,
				Amount:    amount,
				Operation: op,
			}
		}
		entries[key] = UpdateAttributesEntry{
			Value:     value,
			Modifiers: modifiers,
		}
	}
	return UpdateAttributes{
		EntityID: eid,
		Entries:  entries,
	}, nil
}

func (p UpdateAttributes) IsValid() (reason error) {
	for _, entry := range p.Entries {
		for i, modifier := range entry.Modifiers {
			if 2 < modifier.Operation {
				return fmt.Errorf("modifier operation must be in <0; 2> was %d at index %d", modifier.Operation, i)
			}
		}
	}
	return nil
}
