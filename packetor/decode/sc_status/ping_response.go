package sc_status

import "Packetor/packetor/decode"

type PingResponse struct {
	Payload int64
}

func (p PingResponse) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	payload, err := reader.ReadLong()
	if err != nil {
		return nil, err
	}
	return PingResponse{
		Payload: payload,
	}, nil
}

func (p PingResponse) IsValid() (reason error) {
	return nil
}
