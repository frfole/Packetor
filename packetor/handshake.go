package packetor

import (
	"Packetor/packetor/decode"
	"Packetor/packetor/decode/cs_handshake"
)

func (r *Route) handleHandshakeC(packet decode.Packet) (err error) {
	r.state = byte(packet.(cs_handshake.Handshake).NextState)
	println("next state", r.state)
	return nil
}
