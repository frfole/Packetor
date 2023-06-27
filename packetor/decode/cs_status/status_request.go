package cs_status

import "Packetor/packetor/decode"

type StatusRequest struct {
}

func (p StatusRequest) Read(_ decode.PacketReader) (packet decode.Packet, err error) {
	return StatusRequest{}, nil
}

func (p StatusRequest) IsValid() (reason error) {
	return nil
}
