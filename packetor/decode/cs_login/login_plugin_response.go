package cs_login

import "Packetor/packetor/decode"

type LoginPluginResponse struct {
	MessageID  int32
	Successful bool
	Data       []byte
}

func (p LoginPluginResponse) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	mid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	successful, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	var data []byte
	if successful {
		data, err = reader.ReadBytes(1048576)
		if err != nil {
			return nil, err
		}
	}
	return LoginPluginResponse{
		MessageID:  mid,
		Successful: successful,
		Data:       data,
	}, nil
}

func (p LoginPluginResponse) IsValid() (reason error) {
	return nil
}
