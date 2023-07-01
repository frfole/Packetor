package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type SetPassengers struct {
	EntityID   int32
	Passengers []int32
}

func (p SetPassengers) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	passengers := make([]int32, count)
	for i := int32(0); i < count; i++ {
		passengers[i], err = reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
	}
	return SetPassengers{
		EntityID:   eid,
		Passengers: passengers,
	}, nil
}

func (p SetPassengers) IsValid() (reason error) {
	for i, passenger := range p.Passengers {
		if passenger == p.EntityID {
			return fmt.Errorf("enity %d cannot ride itself at index %d", passenger, i)
		}
	}
	return nil
}
