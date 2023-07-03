package cs_play

import "Packetor/packetor/decode"

type CommandSuggestionsRequest struct {
	TransactionID int32
	Text          string
}

func (p CommandSuggestionsRequest) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	tid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	text, err := reader.ReadString0(32500)
	if err != nil {
		return nil, err
	}
	return CommandSuggestionsRequest{
		TransactionID: tid,
		Text:          text,
	}, nil
}

func (p CommandSuggestionsRequest) IsValid() (reason error) {
	return nil
}
