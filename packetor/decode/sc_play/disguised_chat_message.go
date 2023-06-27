package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type DisguisedChatMessage struct {
	Message       string
	ChatType      int32
	ChatTypeName  string
	HasTargetName bool
	TargetName    string
}

func (p DisguisedChatMessage) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	msg, err := reader.ReadChat()
	if err != nil {
		return nil, err
	}
	chatType, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	chatTypeName, err := reader.ReadChat()
	if err != nil {
		return nil, err
	}
	hasTarget, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	targetName := ""
	if hasTarget {
		targetName, err = reader.ReadChat()
		if err != nil {
			return nil, err
		}
	}
	return DisguisedChatMessage{
		Message:       msg,
		ChatType:      chatType,
		ChatTypeName:  chatTypeName,
		HasTargetName: hasTarget,
		TargetName:    targetName,
	}, nil
}

func (p DisguisedChatMessage) IsValid() (reason error) {
	if p.ChatType < 0 || 6 < p.ChatType {
		return fmt.Errorf("unknown chat type excepted <0; 6> was %d", p.ChatType)
	}
	return nil
}
