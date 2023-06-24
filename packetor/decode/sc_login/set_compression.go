package sc_login

import (
	"Packetor/packetor/decode"
	"fmt"
)

type SetCompression int32

func (p SetCompression) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	value, err := reader.ReadVarInt()
	return SetCompression(value), err
}

func (p SetCompression) IsValid() (reason error) {
	if p < 0 {
		return fmt.Errorf("compression threshold must be in <0; inf) was %d", p)
	}
	return nil
}

func (p SetCompression) Compression() bool {
	return p > 0
}
