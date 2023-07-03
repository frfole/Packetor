package cs_play

import "Packetor/packetor/decode"

type PaddleBoat struct {
	LeftPaddleTurning  bool
	RightPaddleTurning bool
}

func (p PaddleBoat) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	left, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	right, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return PaddleBoat{
		LeftPaddleTurning:  left,
		RightPaddleTurning: right,
	}, nil
}

func (p PaddleBoat) IsValid() (reason error) {
	return nil
}
