package sc_play

import "Packetor/packetor/decode"

type SetBorderWarningDistance struct {
	WarningBlocks int32
}

func (p SetBorderWarningDistance) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	dist, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return SetBorderWarningDistance{WarningBlocks: dist}, nil
}

func (p SetBorderWarningDistance) IsValid() (reason error) {
	return nil
}
