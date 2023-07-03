package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type ClientInformation struct {
	Locale              string
	ViewDistance        int8
	ChatMode            int32
	ChatColors          bool
	DisplayedSkinParts  uint8
	MainHand            int32
	EnableTextFiltering bool
	AllowServerListing  bool
}

func (p ClientInformation) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	lang, err := reader.ReadString0(16)
	if err != nil {
		return nil, err
	}
	viewDist, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	chatMode, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	colors, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	skinParts, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	mainHand, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	textFiltering, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	listing, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return ClientInformation{
		Locale:              lang,
		ViewDistance:        viewDist,
		ChatMode:            chatMode,
		ChatColors:          colors,
		DisplayedSkinParts:  skinParts,
		MainHand:            mainHand,
		EnableTextFiltering: textFiltering,
		AllowServerListing:  listing,
	}, nil
}

func (p ClientInformation) IsValid() (reason error) {
	if p.ViewDistance < 0 {
		return fmt.Errorf("view distance must be atleast 0 was %d", p.ViewDistance)
	}
	if p.ChatMode < 0 || 2 < p.ChatMode {
		return fmt.Errorf("chat mode must be in <0; 2> was %d", p.ChatMode)
	}
	if p.MainHand < 0 || 1 < p.MainHand {
		return fmt.Errorf("main hand must be in <0; 1> was %d", p.MainHand)
	}
	return nil
}
