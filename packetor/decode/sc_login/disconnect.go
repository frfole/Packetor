package sc_login

import "Packetor/packetor/decode"

type Disconnect struct {
	Reason string
}

func (p Disconnect) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	reason, err := reader.ReadChat()
	if err != nil {
		return nil, err
	}
	return Disconnect{Reason: reason}, nil
}

func (p Disconnect) IsValid() (reason error) {
	// TODO: maybe validate by parsing
	return nil
}
