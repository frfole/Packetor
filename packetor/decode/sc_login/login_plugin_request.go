package sc_login

import "Packetor/packetor/decode"

type LoginPluginRequest struct {
	MessageID int32
	Channel   string
	Data      []byte
}

func (p LoginPluginRequest) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	mid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	cid, err := reader.ReadIdentifier()
	if err != nil {
		return nil, err
	}
	data, err := reader.ReadBytes(1048576)
	if err != nil {
		return nil, err
	}
	return LoginPluginRequest{
		MessageID: mid,
		Channel:   cid,
		Data:      data,
	}, nil
}

func (p LoginPluginRequest) IsValid() (reason error) {
	return nil
}
