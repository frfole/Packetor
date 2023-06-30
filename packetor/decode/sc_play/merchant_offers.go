package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type MerchantOffersTrade struct {
	InputItem1      decode.Slot
	OutputItem      decode.Slot
	InputItem2      decode.Slot
	TradeDisabled   bool
	Uses            int32
	MaxUses         int32
	XP              int32
	SpecialPrice    int32
	PriceMultiplier float32
	Demand          int32
}

type MerchantOffers struct {
	WindowsID         int32
	Trades            []MerchantOffersTrade
	VillagerLevel     int32
	Experience        int32
	IsRegularVillager bool
	CanRestock        bool
}

func (p MerchantOffers) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	wid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	trades := make([]MerchantOffersTrade, count)
	for i := int32(0); i < count; i++ {
		in1, err := reader.ReadSlot()
		if err != nil {
			return nil, err
		}
		out1, err := reader.ReadSlot()
		if err != nil {
			return nil, err
		}
		in2, err := reader.ReadSlot()
		if err != nil {
			return nil, err
		}
		disabled, err := reader.ReadBoolean()
		if err != nil {
			return nil, err
		}
		uses, err := reader.ReadInt()
		if err != nil {
			return nil, err
		}
		usesMax, err := reader.ReadInt()
		if err != nil {
			return nil, err
		}
		xp, err := reader.ReadInt()
		if err != nil {
			return nil, err
		}
		specialP, err := reader.ReadInt()
		if err != nil {
			return nil, err
		}
		priceMult, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		demand, err := reader.ReadInt()
		if err != nil {
			return nil, err
		}
		trades[i] = MerchantOffersTrade{
			InputItem1:      in1,
			OutputItem:      out1,
			InputItem2:      in2,
			TradeDisabled:   disabled,
			Uses:            uses,
			MaxUses:         usesMax,
			XP:              xp,
			SpecialPrice:    specialP,
			PriceMultiplier: priceMult,
			Demand:          demand,
		}
	}
	level, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	xp, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	regular, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	canRestock, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return MerchantOffers{
		WindowsID:         wid,
		Trades:            trades,
		VillagerLevel:     level,
		Experience:        xp,
		IsRegularVillager: regular,
		CanRestock:        canRestock,
	}, nil
}

func (p MerchantOffers) IsValid() (reason error) {
	return nil
}
