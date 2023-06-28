package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

const (
	GameEventNoRespawn uint8 = iota
	GameEventBeginRain
	GameEventEndRain
	GameEventSetGameMode
	GameEventWinGame
	GameEventDemo
	GameEventArrowHit
	GameEventRainLevel
	GameEventThunderLevel
	GameEventPufferfishSting
	GameEventElderGuardian
	GameEventEnableRespawnScreen
)

type GameEvent struct {
	Event uint8
	Value float32
}

func (p GameEvent) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	event, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	value, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	return GameEvent{
		Event: event,
		Value: value,
	}, nil
}

func (p GameEvent) IsValid() (reason error) {
	if 11 < p.Event {
		return fmt.Errorf("unknown game event %d", p.Event)
	}
	// TODO: validate game event value
	return nil
}
