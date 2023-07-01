package sc_play

import "Packetor/packetor/decode"

type SetCenterChunk struct {
	ChunkX int32
	ChunkZ int32
}

func (p SetCenterChunk) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	cx, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	cz, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return SetCenterChunk{
		ChunkX: cx,
		ChunkZ: cz,
	}, nil
}

func (p SetCenterChunk) IsValid() (reason error) {
	return nil
}
