package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type ExplosionChange struct {
	X int8
	Y int8
	Z int8
}

type Explosion struct {
	X             float64
	Y             float64
	Z             float64
	Strength      float32
	Changes       []ExplosionChange
	PlayerMotionX float32
	PlayerMotionY float32
	PlayerMotionZ float32
}

func (p Explosion) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	x, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	y, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	z, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	strength, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	changes := make([]ExplosionChange, count)
	for i := int32(0); i < count; i++ {
		cx, err := reader.ReadSByte()
		if err != nil {
			return nil, err
		}
		cy, err := reader.ReadSByte()
		if err != nil {
			return nil, err
		}
		cz, err := reader.ReadSByte()
		if err != nil {
			return nil, err
		}
		changes[i] = ExplosionChange{
			X: cx,
			Y: cy,
			Z: cz,
		}
	}
	pVelX, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	pVelY, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	pVelZ, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	return Explosion{
		X:             x,
		Y:             y,
		Z:             z,
		Strength:      strength,
		Changes:       changes,
		PlayerMotionX: pVelX,
		PlayerMotionY: pVelY,
		PlayerMotionZ: pVelZ,
	}, nil
}

func (p Explosion) IsValid() (reason error) {
	return nil
}
