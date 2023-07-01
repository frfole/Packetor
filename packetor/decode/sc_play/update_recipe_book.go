package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type UpdateRecipeBook struct {
	Action          int32
	CraftingOpen    bool
	CraftingFilter  bool
	SmeltingOpen    bool
	SmeltingFilter  bool
	FurnaceOpen     bool
	FurnaceFilter   bool
	SmokerOpen      bool
	SmokerFilter    bool
	RecipeChangeIDs []string
	RecipeInitIDs   []string
}

func (p UpdateRecipeBook) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	action, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	cOpen, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	cFilter, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	smeltOpen, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	smeltFilter, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	fOpen, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	fFilter, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	smokeOpen, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	smokeFilter, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	changes := make([]string, count)
	for i := int32(0); i < count; i++ {
		changes[i], err = reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
	}
	var inits []string
	if action == 0 {
		count, err = reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if count < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
		}
		inits = make([]string, count)
		for i := int32(0); i < count; i++ {
			inits[i], err = reader.ReadIdentifier()
			if err != nil {
				return nil, err
			}
		}
	}
	return UpdateRecipeBook{
		Action:          action,
		CraftingOpen:    cOpen,
		CraftingFilter:  cFilter,
		SmeltingOpen:    smeltOpen,
		SmeltingFilter:  smeltFilter,
		FurnaceOpen:     fOpen,
		FurnaceFilter:   fFilter,
		SmokerOpen:      smokeOpen,
		SmokerFilter:    smokeFilter,
		RecipeChangeIDs: changes,
		RecipeInitIDs:   inits,
	}, nil
}

func (p UpdateRecipeBook) IsValid() (reason error) {
	if p.Action < 0 || 2 < p.Action {
		return fmt.Errorf("action must be in <0; 2> was %d", p.Action)
	}
	return nil
}
