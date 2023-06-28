package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"Packetor/packetor/nbt"
	"errors"
	"fmt"
)

type Login struct {
	EntityID            int32
	IsHardcore          bool
	GameMode            uint8
	PreviousGameMode    uint8
	Dimensions          []string
	RegistryCodec       nbt.Compound
	DimensionType       string
	DimensionName       string
	HashedSeed          int64
	MaxPlayers          int32
	ViewDistance        int32
	SimulationDistance  int32
	ReducedDebugInfo    bool
	EnableRespawnScreen bool
	IsDebug             bool
	IsFlat              bool
	HasDeathLocation    bool
	DeathDimension      string
	DeathLocation       decode.Position
	PortalCooldown      int32
}

func (p Login) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	hardcore, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	gm, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	prevGm, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	dimensions := make([]string, count)
	for i := int32(0); i < count; i++ {
		dimensions[i], err = reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
	}
	registry, err := reader.ReadNbt()
	if err != nil {
		return nil, err
	}
	dimType, err := reader.ReadString()
	if err != nil {
		return nil, err
	}
	dimName, err := reader.ReadString()
	if err != nil {
		return nil, err
	}
	seedHash, err := reader.ReadLong()
	if err != nil {
		return nil, err
	}
	maxPlayers, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	viewDist, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	simDist, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	redDebug, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	respawnScreen, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	isDebug, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	isFlat, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	hasDeath, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	dDimName := ""
	dLoc := decode.Position(0)
	if hasDeath {
		dDimName, err = reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
		dLoc, err = reader.ReadPosition()
		if err != nil {
			return nil, err
		}
	}
	portCooldown, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return Login{
		EntityID:            eid,
		IsHardcore:          hardcore,
		GameMode:            gm,
		PreviousGameMode:    prevGm,
		Dimensions:          dimensions,
		RegistryCodec:       registry,
		DimensionType:       dimType,
		DimensionName:       dimName,
		HashedSeed:          seedHash,
		MaxPlayers:          maxPlayers,
		ViewDistance:        viewDist,
		SimulationDistance:  simDist,
		ReducedDebugInfo:    redDebug,
		EnableRespawnScreen: respawnScreen,
		IsDebug:             isDebug,
		IsFlat:              isFlat,
		HasDeathLocation:    hasDeath,
		DeathDimension:      dDimName,
		DeathLocation:       dLoc,
		PortalCooldown:      portCooldown,
	}, nil
}

func (p Login) IsValid() (reason error) {
	if 3 < p.GameMode {
		return fmt.Errorf("game mode must be in <0 ;3> was %d", p.GameMode)
	} else if 3 < p.PreviousGameMode {
		return fmt.Errorf("previous game mode must be in <0; 3> was %d", p.PreviousGameMode)
	}
	// TODO: validate registry codec?
	return nil
}
