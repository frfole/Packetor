package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type UpdateSectionBlocks struct {
	ChunkSectionPosition uint64
	Blocks               []int64
}

func (p UpdateSectionBlocks) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	pos, err := reader.ReadULong()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	blocks := make([]int64, count)
	for i := int32(0); i < count; i++ {
		blocks[i], err = reader.ReadVarLong()
		if err != nil {
			return nil, err
		}
	}
	return UpdateSectionBlocks{
		ChunkSectionPosition: pos,
		Blocks:               blocks,
	}, nil
}

func (p UpdateSectionBlocks) IsValid() (reason error) {
	return nil
}
