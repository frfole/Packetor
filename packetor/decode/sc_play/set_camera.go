package sc_play

import "Packetor/packetor/decode"

type SetCamera struct {
	CameraID int32
}

func (p SetCamera) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	cid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return SetCamera{CameraID: cid}, nil
}

func (p SetCamera) IsValid() (reason error) {
	return nil
}
