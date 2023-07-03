package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type ResourcePack struct {
	Result int32
}

func (p ResourcePack) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	res, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return ResourcePack{Result: res}, nil
}

func (p ResourcePack) IsValid() (reason error) {
	if p.Result < 0 || 3 < p.Result {
		return fmt.Errorf("result must be in <0; 3> was %d", p.Result)
	}
	return nil
}
