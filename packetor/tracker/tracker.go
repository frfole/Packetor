package tracker

import (
	"Packetor/packetor/decode/sc_play"
	"Packetor/packetor/nbt"
)

type Tracker struct {
	ServerInfo   ServerInfo
	WorldTracker WorldTracker
}

func NewTracker() Tracker {
	return Tracker{
		ServerInfo: ServerInfo{
			DimensionTypes: map[string]DimensionType{},
		},
		WorldTracker: WorldTracker{
			Dimension:     "",
			DimensionType: nil,
		},
	}
}

func (receiver Tracker) OnLogin(packet sc_play.Login) error {
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
		return err
	}
	return nil
}
