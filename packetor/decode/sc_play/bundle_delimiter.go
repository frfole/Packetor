package sc_play

import "Packetor/packetor/decode"

type BundleDelimiter struct {
}

func (p BundleDelimiter) Read(_ decode.PacketReader) (packet decode.Packet, err error) {
	return BundleDelimiter{}, nil
}

func (p BundleDelimiter) IsValid() (reason error) {
	return nil
}
