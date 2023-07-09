package tracker

import (
	"Packetor/packetor/decode"
	"Packetor/packetor/decode/sc_play"
	"Packetor/packetor/nbt"
	"fmt"
)

type Tracker struct {
	ServerInfo       ServerInfo
	WorldTracker     WorldTracker
	InventoryTracker InventoryTracker
	PacketTracker    PacketTracker
}

func NewTracker() Tracker {
	return Tracker{
		ServerInfo: ServerInfo{
			DimensionTypes: map[string]DimensionType{},
		},
		WorldTracker: WorldTracker{
			Dimension:     "",
			DimensionType: DimensionTypeDefault,
			Chunks:        map[uint64]Chunk{},
		},
		InventoryTracker: InventoryTracker{
			HasOpenInventory:   false,
			WindowID:           -1,
			IsSecondaryOpen:    false,
			IsSecondaryHorse:   false,
			SecondaryWindow:    nil,
			SecondarySlotCount: 0,
		},
		PacketTracker: newPacketTracker(),
	}
}

func (receiver *Tracker) OnLogin(packet sc_play.Login) error {
	codec := packet.RegistryCodec
	dimensionsTag := codec[""].(nbt.Compound)["minecraft:dimension_type"].(nbt.Compound)["value"].(nbt.List)
	receiver.ServerInfo.DimensionTypes = map[string]DimensionType{}
	for i := range dimensionsTag {
		dimensionType := dimensionsTag[i].(nbt.Compound)
		receiver.ServerInfo.DimensionTypes[string(dimensionType["name"].(nbt.String))] = DimensionType{
			MinY:   int32(dimensionType["element"].(nbt.Compound)["min_y"].(nbt.Int)),
			Height: int32(dimensionType["element"].(nbt.Compound)["height"].(nbt.Int)),
		}
	}

	err := receiver.ResetWorldTracker(packet.DimensionName, packet.DimensionType)
	if err != nil {
		return fmt.Errorf("failed to reset world tracker: %w", err)
	}
	return nil
}

func (receiver *Tracker) OnRespawn(basePacket decode.Packet, _ decode.PacketContext) error {
	packet, ok := basePacket.(sc_play.Respawn)
	if !ok {
		return nil
	}
	err := receiver.ResetWorldTracker(packet.DimensionName, packet.DimensionType)
	if err != nil {
		return fmt.Errorf("failed to reset world tracker: %w", err)
	}
	return nil
}
