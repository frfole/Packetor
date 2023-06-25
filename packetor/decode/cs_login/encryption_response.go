package cs_login

import "Packetor/packetor/decode"

type EncryptionResponse struct {
	SharedSecretLen int32
	SharedSecret    []byte
	VerifyTokenLen  int32
	VerifyToken     []byte
}

func (p EncryptionResponse) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	sLen, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	s, err := reader.ReadBytesExact(int(sLen))
	if err != nil {
		return nil, err
	}
	vtLen, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	vt, err := reader.ReadBytesExact(int(vtLen))
	if err != nil {
		return nil, err
	}
	return EncryptionResponse{
		SharedSecretLen: sLen,
		SharedSecret:    s,
		VerifyTokenLen:  vtLen,
		VerifyToken:     vt,
	}, nil
}

func (p EncryptionResponse) IsValid() (reason error) {
	return nil
}
