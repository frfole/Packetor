package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"Packetor/packetor/nbt"
	"errors"
	"fmt"
	"io"
)

type ChunkDataBlockEntity struct {
	XZ   uint8
	Y    int16
	Type int32
	Data nbt.Compound
}

type ChunkDataSkyLight []byte
type ChunkDataBlockLight []byte
type ChunkDataSection struct {
	BlockCount int16
	Blocks     decode.PaletteContainer
	Biomes     decode.PaletteContainer
}

type ChunkData struct {
	ChunkX              int32
	ChunkZ              int32
	Heightmaps          nbt.Compound
	Data                []ChunkDataSection
	BlockEntities       []ChunkDataBlockEntity
	SkyLightMask        decode.BitSet
	BlockLightMash      decode.BitSet
	EmptySkyLightMask   decode.BitSet
	EmptyBlockLightMash decode.BitSet
	SkyLightArrays      []ChunkDataSkyLight
	BlockLightArrays    []ChunkDataBlockLight
}

func (p ChunkData) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	cx, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	cz, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	heightmaps, err := reader.ReadNbt()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	sectionsData, err := reader.ReadBytesExact(int(count))
	var sections []ChunkDataSection
	{
		height := reader.Context.WorldHeight
		sectionReader := decode.NewPacketReaderBytes(sectionsData)
		for i := int32(0); i < height; i++ {
			blockCount, err := sectionReader.ReadShort()
			if err != nil && errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				return nil, err
			}
			blocks, err := sectionReader.ReadPaletteContainer(16, 8)
			if err != nil && errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				return nil, err
			}
			biomes, err := sectionReader.ReadPaletteContainer(4, 3)
			if err != nil && errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				return nil, err
			}
			sections = append(sections, ChunkDataSection{
				BlockCount: blockCount,
				Blocks:     blocks,
				Biomes:     biomes,
			})
		}
	}
	if err != nil {
		return nil, err
	}
	count, err = reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	bEntities := make([]ChunkDataBlockEntity, count)
	for i := int32(0); i < count; i++ {
		xz, err := reader.ReadUByte()
		if err != nil {
			return nil, err
		}
		y, err := reader.ReadShort()
		if err != nil {
			return nil, err
		}
		eType, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		eData, err := reader.ReadNbt()
		if err != nil {
			return nil, err
		}
		bEntities[i] = ChunkDataBlockEntity{
			XZ:   xz,
			Y:    y,
			Type: eType,
			Data: eData,
		}
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
	count, err = reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	slArr := make([]ChunkDataSkyLight, count)
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
	count, err = reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	blArr := make([]ChunkDataBlockLight, count)
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
	return ChunkData{
		ChunkX:              cx,
		ChunkZ:              cz,
		Heightmaps:          heightmaps,
		Data:                sections,
		BlockEntities:       bEntities,
		SkyLightMask:        slMask,
		BlockLightMash:      blMask,
		EmptySkyLightMask:   eslMask,
		EmptyBlockLightMash: eblMask,
		SkyLightArrays:      slArr,
		BlockLightArrays:    blArr,
	}, nil
}

func (p ChunkData) IsValid() (reason error) {
	return nil
}
