package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type Respawn struct {
	DimensionType      string
	DimensionName      string
	SeedHash           int64
	GameMode           uint8
	PreviousGameMode   int8
	IsDebug            bool
	IsFlat             bool
	Flags              int8
	HasDeathLocation   bool
	DeathDimensionName string
	DeathLocation      decode.Position
	PortalCooldown     int32
}

func (p Respawn) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	dimType, err := reader.ReadIdentifier()
	if err != nil {
		return nil, err
	}
	dimName, err := reader.ReadIdentifier()
	if err != nil {
		return nil, err
	}
	hash, err := reader.ReadLong()
	if err != nil {
		return nil, err
	}
	gm, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	prevGM, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	debug, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	flat, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	flags, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	hasDeath, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	deathDim := ""
	var deathLoc decode.Position
	if hasDeath {
		deathDim, err = reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
		deathLoc, err = reader.ReadPosition()
		if err != nil {
			return nil, err
		}
	}
	portalCooldown, err := reader.ReadVarInt()
	return Respawn{
		DimensionType:      dimType,
		DimensionName:      dimName,
		SeedHash:           hash,
		GameMode:           gm,
		PreviousGameMode:   prevGM,
		IsDebug:            debug,
		IsFlat:             flat,
		Flags:              flags,
		HasDeathLocation:   hasDeath,
		DeathDimensionName: deathDim,
		DeathLocation:      deathLoc,
		PortalCooldown:     portalCooldown,
	}, nil
}

func (p Respawn) IsValid() (reason error) {
	if 3 < p.GameMode {
		return fmt.Errorf("gamemode must be in <0; 3> was %d", p.GameMode)
	}
	if p.PreviousGameMode < -1 || 3 < p.PreviousGameMode {
		return fmt.Errorf("previous gamemode must be in <-1; 3> was %d", p.PreviousGameMode)
	}
	if (p.Flags &^ 0b11) != 0 {
		return fmt.Errorf("unknown flags %b", p.Flags)
	}
	return nil
}
