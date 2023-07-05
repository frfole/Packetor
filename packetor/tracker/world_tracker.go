package tracker

import (
	"Packetor/packetor/decode"
	"Packetor/packetor/decode/sc_play"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
)

type (
	ChunkSection struct {
		Blocks *decode.PaletteContainer
	}
	Chunk struct {
		ChunkX   int32
		ChunkZ   int32
		Sections []*ChunkSection
	}
	WorldTracker struct {
		Dimension     string
		DimensionType DimensionType
		Chunks        map[uint64]Chunk
	}
)

func (receiver *Tracker) ResetWorldTracker(dimName string, dimTypeName string) error {
	dimensionType, ok := receiver.ServerInfo.DimensionTypes[dimTypeName]
	if !ok {
		return fmt.Errorf("unknown dimenstion type %v of dimension %v", dimTypeName, dimName)
	}
	receiver.WorldTracker = WorldTracker{
		Dimension:     dimName,
		DimensionType: dimensionType,
		Chunks:        map[uint64]Chunk{},
	}
	logrus.Info(receiver.WorldTracker)
	return nil
}

func (receiver *WorldTracker) UpdateChunk(packet decode.Packet) error {
	if chunkPacket, ok := packet.(sc_play.ChunkData); ok {
		chunkIndex := ((uint64(chunkPacket.ChunkX) & 0x3FFFFF) << 42) | ((uint64(chunkPacket.ChunkZ) & 0x3FFFFF) << 20)
		sections := make([]*ChunkSection, len(chunkPacket.Data))
		for i := 0; i < len(chunkPacket.Data); i++ {
			sections[i] = &ChunkSection{Blocks: &chunkPacket.Data[i].Blocks}
		}
		receiver.Chunks[chunkIndex] = Chunk{
			ChunkX:   chunkPacket.ChunkX,
			ChunkZ:   chunkPacket.ChunkZ,
			Sections: sections,
		}
	} else if chunkPacket, ok := packet.(sc_play.UnloadChunk); ok {
		chunkIndex := ((uint64(chunkPacket.ChunkX) & 0x3FFFFF) << 42) | ((uint64(chunkPacket.ChunkZ) & 0x3FFFFF) << 20)
		delete(receiver.Chunks, chunkIndex)
	} else if chunkPacket, ok := packet.(sc_play.BlockUpdate); ok {
		err := receiver.SetBlockState(chunkPacket.Location.X(), chunkPacket.Location.Y(), chunkPacket.Location.Z(), chunkPacket.BlockID)
		if err != nil {
			return errors.Join(fmt.Errorf("failed to set block from block update packet"), err)
		}
	} else if explosion, ok := packet.(sc_play.Explosion); ok {
		if len(explosion.Changes) == 0 {
			return nil
		}
		x := int32(math.Floor(explosion.X))
		y := int32(math.Floor(explosion.Y))
		z := int32(math.Floor(explosion.Z))
		for _, change := range explosion.Changes {
			err := receiver.SetBlockState(x+int32(change.X), y+int32(change.Y), z+int32(change.Z), 0) // TODO: dont assume state id of air is always 0
			if err != nil {
				return errors.Join(fmt.Errorf("failed to set block after explosion"), err)
			}
		}
	} else if sectionPacket, ok := packet.(sc_play.UpdateSectionBlocks); ok {
		chunkIndex := sectionPacket.ChunkSectionPosition &^ 0xfffff
		sectionIdx := int32(sectionPacket.ChunkSectionPosition<<44>>44) - (receiver.DimensionType.MinY / 16)
		chunk, ok := receiver.Chunks[chunkIndex]
		if !ok {
			return fmt.Errorf("chunk[%v %v] not loaded", chunk.ChunkX, chunk.ChunkZ)
		}
		if sectionIdx < 0 || len(chunk.Sections) <= int(sectionIdx) {
			return fmt.Errorf("section %v does not exists", sectionIdx)
		}
		section := chunk.Sections[sectionIdx]
		for _, block := range sectionPacket.Blocks {
			section.Blocks.SetState(int((block>>8)&0xf), int(block&0xf), int((block>>4)&0xf), int32(block>>12))
		}
	}
	return nil
}

func (receiver *WorldTracker) GetBlockState(x int32, y int32, z int32) int32 {
	chunkX := x >> 4
	sectionIdx := (y - receiver.DimensionType.MinY) >> 4
	chunkZ := z >> 4
	chunkIndex := ((uint64(chunkX) & 0x3FFFFF) << 42) | ((uint64(chunkZ) & 0x3FFFFF) << 20)
	chunk, ok := receiver.Chunks[chunkIndex]
	if !ok {
		return -1
	}
	if sectionIdx < 0 || len(chunk.Sections) <= int(sectionIdx) {
		return -1
	}
	section := chunk.Sections[sectionIdx]
	return section.Blocks.GetState(int(x&0xf), int(y&0xf), int(z&0xf))
}

func (receiver *WorldTracker) SetBlockState(x int32, y int32, z int32, stateId int32) error {
	chunkX := x >> 4
	sectionIdx := (y - receiver.DimensionType.MinY) >> 4
	chunkZ := z >> 4
	chunkIndex := ((uint64(chunkX) & 0x3FFFFF) << 42) | ((uint64(chunkZ) & 0x3FFFFF) << 20)
	chunk, ok := receiver.Chunks[chunkIndex]
	if !ok {
		return fmt.Errorf("chunk[%v %v] not loaded", chunkX, chunkZ)
	}
	if sectionIdx < 0 || len(chunk.Sections) <= int(sectionIdx) {
		return fmt.Errorf("section %v does not exists", sectionIdx)
	}
	section := chunk.Sections[sectionIdx]
	section.Blocks.SetState(int(x&0xf), int(y&0xf), int(z&0xf), stateId)
	return nil
}
