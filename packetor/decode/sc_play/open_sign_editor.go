package sc_play

import "Packetor/packetor/decode"

type OpenSignEditor struct {
	Location    decode.Position
	IsFrontText bool
}

func (p OpenSignEditor) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	isFront, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return OpenSignEditor{
		Location:    loc,
		IsFrontText: isFront,
	}, nil
}

func (p OpenSignEditor) IsValid() (reason error) {
	return nil
}
