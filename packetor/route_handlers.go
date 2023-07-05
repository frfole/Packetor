package packetor

import (
	"Packetor/packetor/decode"
	"Packetor/packetor/decode/cs_handshake"
	"Packetor/packetor/decode/sc_login"
	"github.com/sirupsen/logrus"
	"reflect"
)

func (r *Route) handleHandshakeC(packet decode.Packet) (err error) {
	r.state = byte(packet.(cs_handshake.Handshake).NextState)
	return nil
}

func (r *Route) handleLoginSuccess(_ decode.Packet) (err error) {
	r.state = 3
	return nil
}
func (r *Route) handleSetCompression(packet decode.Packet) (err error) {
	r.compress = packet.(sc_login.SetCompression).Compression()
	return nil
}

func (r *Route) printPacket(packet decode.Packet) (err error) {
	logrus.Info(reflect.TypeOf(packet), packet)
	return nil
}
