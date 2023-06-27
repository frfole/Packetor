package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type ChunkBiomesEntry struct {
	ChunkX int32
	ChunkZ int32
	Data   []byte
}

type ChunkBiomes struct {
	Entries []ChunkBiomesEntry
}

func (p ChunkBiomes) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	amount, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if amount < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", amount), error2.ErrDecodeTooSmall)
	}
	entries := make([]ChunkBiomesEntry, amount)
	for i := int32(0); i < amount; i++ {
		cx, err := reader.ReadInt()
		if err != nil {
			return nil, err
		}
		cz, err := reader.ReadInt()
		if err != nil {
			return nil, err
		}
		length, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if length < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", length), error2.ErrDecodeTooSmall)
		}
		data, err := reader.ReadBytesExact(int(length))
		if err != nil {
			return nil, err
		}
		entries[i] = ChunkBiomesEntry{
			ChunkX: cx,
			ChunkZ: cz,
			Data:   data,
		}
	}
	return ChunkBiomes{Entries: entries}, nil
}

func (p ChunkBiomes) IsValid() (reason error) {
	// TODO: validate
	return nil
}
