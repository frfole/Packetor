package sc_play

import "Packetor/packetor/decode"

type SetBorderWarningDelay struct {
	WarningTime int32
}

func (p SetBorderWarningDelay) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	t, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return SetBorderWarningDelay{WarningTime: t}, nil
}

func (p SetBorderWarningDelay) IsValid() (reason error) {
	return nil
}
