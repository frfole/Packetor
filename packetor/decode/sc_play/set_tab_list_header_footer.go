package sc_play

import "Packetor/packetor/decode"

type SetTabListHeaderFooter struct {
	Header string
	Footer string
}

func (p SetTabListHeaderFooter) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	header, err := reader.ReadChat()
	if err != nil {
		return nil, err
	}
	footer, err := reader.ReadChat()
	if err != nil {
		return nil, err
	}
	return SetTabListHeaderFooter{
		Header: header,
		Footer: footer,
	}, nil
}

func (p SetTabListHeaderFooter) IsValid() (reason error) {
	return nil
}
