package sc_play

import "Packetor/packetor/decode"

type SetTitleAnimationTimes struct {
	FadeIn  int32
	Stay    int32
	FadeOut int32
}

func (p SetTitleAnimationTimes) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	fin, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	stay, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	fout, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	return SetTitleAnimationTimes{
		FadeIn:  fin,
		Stay:    stay,
		FadeOut: fout,
	}, nil
}

func (p SetTitleAnimationTimes) IsValid() (reason error) {
	return nil
}
