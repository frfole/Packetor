package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type CommandSuggestionsMatch struct {
	Match      string
	HasTooltip bool
	Tooltip    string
}

type CommandSuggestionsResponse struct {
	ID      int32
	Start   int32
	Length  int32
	Matches []CommandSuggestionsMatch
}

func (p CommandSuggestionsResponse) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	tid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	start, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	length, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	matches := make([]CommandSuggestionsMatch, count)
	for i := int32(0); i < count; i++ {
		match, err := reader.ReadString0(32767)
		if err != nil {
			return nil, err
		}
		hasTooltip, err := reader.ReadBoolean()
		if err != nil {
			return nil, err
		}
		tooltip := ""
		if hasTooltip {
			tooltip, err = reader.ReadChat()
			if err != nil {
				return nil, err
			}
		}
		matches[i] = CommandSuggestionsMatch{
			Match:      match,
			HasTooltip: hasTooltip,
			Tooltip:    tooltip,
		}
	}
	return CommandSuggestionsResponse{
		ID:      tid,
		Start:   start,
		Length:  length,
		Matches: matches,
	}, nil
}

func (p CommandSuggestionsResponse) IsValid() (reason error) {
	return nil
}
