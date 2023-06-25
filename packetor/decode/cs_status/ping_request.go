package cs_status

import "Packetor/packetor/decode"

type PingRequest struct {
	Payload int64
}

func (p PingRequest) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	payload, err := reader.ReadLong()
	if err != nil {
		return nil, err
	}
	return PingRequest{
		Payload: payload,
	}, nil
}

func (p PingRequest) IsValid() (reason error) {
	return nil
}
