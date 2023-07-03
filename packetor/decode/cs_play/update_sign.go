package cs_play

import "Packetor/packetor/decode"

type UpdateSign struct {
	Location    decode.Position
	IsFrontText bool
	Line1       string
	Line2       string
	Line3       string
	Line4       string
}

func (p UpdateSign) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	isFront, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	line1, err := reader.ReadString0(384)
	if err != nil {
		return nil, err
	}
	line2, err := reader.ReadString0(384)
	if err != nil {
		return nil, err
	}
	line3, err := reader.ReadString0(384)
	if err != nil {
		return nil, err
	}
	line4, err := reader.ReadString0(384)
	if err != nil {
		return nil, err
	}
	return UpdateSign{
		Location:    loc,
		IsFrontText: isFront,
		Line1:       line1,
		Line2:       line2,
		Line3:       line3,
		Line4:       line4,
	}, nil
}

func (p UpdateSign) IsValid() (reason error) {
	return nil
}
