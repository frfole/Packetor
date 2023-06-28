package sc_play

import "Packetor/packetor/decode"

type UnloadChunk struct {
	ChunkX int32
	ChunkZ int32
}

func (p UnloadChunk) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	cx, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	cz, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	return UnloadChunk{
		ChunkX: cx,
		ChunkZ: cz,
	}, nil
}

func (p UnloadChunk) IsValid() (reason error) {
	return nil
}
