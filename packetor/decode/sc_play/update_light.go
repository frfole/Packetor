package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type UpdateLightSkyLight []byte
type UpdateLightBlockLight []byte

type UpdateLight struct {
	ChunkX              int32
	ChunkZ              int32
	SkyLightMask        decode.BitSet
	BlockLightMask      decode.BitSet
	EmptySkyLightMask   decode.BitSet
	EmptyBlockLightMask decode.BitSet
	SkyLightArrays      []UpdateLightSkyLight
	BlockLightArrays    []UpdateLightBlockLight
}

func (p UpdateLight) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	cx, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	cz, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	slMask, err := reader.ReadBitSet()
	if err != nil {
		return nil, err
	}
	blMask, err := reader.ReadBitSet()
	if err != nil {
		return nil, err
	}
	eslMask, err := reader.ReadBitSet()
	if err != nil {
		return nil, err
	}
	eblMask, err := reader.ReadBitSet()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	slArr := make([]UpdateLightSkyLight, count)
	for i := int32(0); i < count; i++ {
		length, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if length < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", length), error2.ErrDecodeTooSmall)
		}
		slArr[i], err = reader.ReadBytesExact(int(length))
		if err != nil {
			return nil, err
		}
	}
	blArr := make([]UpdateLightBlockLight, count)
	for i := int32(0); i < count; i++ {
		length, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if length < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", length), error2.ErrDecodeTooSmall)
		}
		blArr[i], err = reader.ReadBytesExact(int(length))
		if err != nil {
			return nil, err
		}
	}
	return UpdateLight{
		ChunkX:              cx,
		ChunkZ:              cz,
		SkyLightMask:        slMask,
		BlockLightMask:      blMask,
		EmptySkyLightMask:   eslMask,
		EmptyBlockLightMask: eblMask,
		SkyLightArrays:      slArr,
		BlockLightArrays:    blArr,
	}, nil
}

func (p UpdateLight) IsValid() (reason error) {
	return nil
}
