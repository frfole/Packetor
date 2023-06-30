package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
)

type PlayerChatMessagePrevious struct {
	MessageID int32
	Signature []byte
}

type PlayerChatMessage struct {
	Sender               uuid.UUID
	Index                int32
	Signature            []byte
	Message              string
	Timestamp            int64
	Salt                 int64
	PreviousMessages     []PlayerChatMessagePrevious
	HasUnsignedContent   bool
	UnsignedContent      string
	FilterType           int32
	FilterTypeBits       decode.BitSet
	ChatType             int32
	NetworkName          string
	HasNetworkTargetName bool
	NetworkTargetNAme    string
}

func (p PlayerChatMessage) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	sender, err := reader.ReadUuid()
	if err != nil {
		return nil, err
	}
	idx, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	hasSig, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	var sig []byte
	if hasSig {
		sig, err = reader.ReadBytesExact(256)
		if err != nil {
			return nil, err
		}
	}
	msg, err := reader.ReadString0(256)
	if err != nil {
		return nil, err
	}
	ts, err := reader.ReadLong()
	if err != nil {
		return nil, err
	}
	salt, err := reader.ReadLong()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	prevMsgs := make([]PlayerChatMessagePrevious, count)
	for i := int32(0); i < count; i++ {
		mid, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		var sigPrev []byte
		if mid == 0 {
			sigPrev, err = reader.ReadBytesExact(256)
			if err != nil {
				return nil, err
			}
		}
		prevMsgs[i] = PlayerChatMessagePrevious{
			MessageID: mid,
			Signature: sigPrev,
		}
	}
	hasUnsigned, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	unsignedContent := ""
	if hasUnsigned {
		unsignedContent, err = reader.ReadChat()
		if err != nil {
			return nil, err
		}
	}
	filterType, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	var filterBits decode.BitSet
	if filterType == 2 {
		filterBits, err = reader.ReadBitSet()
		if err != nil {
			return nil, err
		}
	}
	chatType, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	netName, err := reader.ReadChat()
	if err != nil {
		return nil, err
	}
	hasTarget, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	target := ""
	if hasTarget {
		target, err = reader.ReadChat()
		if err != nil {
			return nil, err
		}
	}
	return PlayerChatMessage{
		Sender:               sender,
		Index:                idx,
		Signature:            sig,
		Message:              msg,
		Timestamp:            ts,
		Salt:                 salt,
		PreviousMessages:     prevMsgs,
		HasUnsignedContent:   hasUnsigned,
		UnsignedContent:      unsignedContent,
		FilterType:           filterType,
		FilterTypeBits:       filterBits,
		ChatType:             chatType,
		NetworkName:          netName,
		HasNetworkTargetName: hasTarget,
		NetworkTargetNAme:    target,
	}, nil
}

func (p PlayerChatMessage) IsValid() (reason error) {
	if p.FilterType < 0 || 2 < p.FilterType {
		return fmt.Errorf("filter type must be in <0; 2> was %d", p.FilterType)
	}
	return nil
}
