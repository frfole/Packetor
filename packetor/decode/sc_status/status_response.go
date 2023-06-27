package sc_status

import "Packetor/packetor/decode"

type StatusResponse struct {
	JsonResponse string
}

func (p StatusResponse) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	jsonRes, err := reader.ReadString0(32767)
	if err != nil {
		return nil, err
	}
	return StatusResponse{
		JsonResponse: jsonRes,
	}, nil
}

func (p StatusResponse) IsValid() (reason error) {
	return nil
}
