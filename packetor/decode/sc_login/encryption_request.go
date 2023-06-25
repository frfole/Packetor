package sc_login

import "Packetor/packetor/decode"

type EncryptionRequest struct {
	ServerID       string
	PublicKeyLen   int32
	PublicKey      []byte
	VerifyTokenLen int32
	VerifyToken    []byte
}

func (p EncryptionRequest) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	sid, err := reader.ReadString(20)
	if err != nil {
		return nil, err
	}
	pkLen, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	pk, err := reader.ReadBytesExact(int(pkLen))
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
	return EncryptionRequest{
		ServerID:       sid,
		PublicKeyLen:   pkLen,
		PublicKey:      pk,
		VerifyTokenLen: vtLen,
		VerifyToken:    vt,
	}, nil
}

func (p EncryptionRequest) IsValid() (reason error) {
	return nil
}
