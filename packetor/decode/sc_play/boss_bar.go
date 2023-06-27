package sc_play

import (
	"Packetor/packetor/decode"
	"fmt"
	"github.com/gofrs/uuid/v5"
)

const (
	BossBarActAdd int32 = iota
	BossBarActRemove
	BossBarActUpdateHealth
	BossBarActUpdateTitle
	BossBarActUpdateStyle
	BossBarActUpdateFlags
)

type BossBarAction interface {
	Type() int32
}

type BossBarActionAdd struct {
	Title    string
	Health   float32
	Color    int32
	Division int32
	Flags    uint8
}

type BossBarActionRemove struct {
}

type BossBarActionUpdateHealth struct {
	Health float32
}

type BossBarActionUpdateTitle struct {
	Title string
}

type BossBarActionUpdateStyle struct {
	Color    int32
	Division int32
}

type BossBarActionUpdateFlags struct {
	Flags uint8
}

type BossBar struct {
	Uuid   uuid.UUID
	Action BossBarAction
}

func (a BossBarActionAdd) Type() int32 {
	return BossBarActAdd
}
func (a BossBarActionRemove) Type() int32 {
	return BossBarActRemove
}
func (a BossBarActionUpdateHealth) Type() int32 {
	return BossBarActUpdateHealth
}
func (a BossBarActionUpdateTitle) Type() int32 {
	return BossBarActUpdateTitle
}
func (a BossBarActionUpdateStyle) Type() int32 {
	return BossBarActUpdateStyle
}
func (a BossBarActionUpdateFlags) Type() int32 {
	return BossBarActUpdateFlags
}

func (p BossBar) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	uid, err := reader.ReadUuid()
	if err != nil {
		return nil, err
	}
	actionType, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	var action BossBarAction
	switch actionType {
	case BossBarActAdd:
		title, err := reader.ReadChat()
		if err != nil {
			return nil, err
		}
		health, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		color, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		division, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		flags, err := reader.ReadUByte()
		if err != nil {
			return nil, err
		}
		action = BossBarActionAdd{
			Title:    title,
			Health:   health,
			Color:    color,
			Division: division,
			Flags:    flags,
		}
	case BossBarActRemove:
		action = BossBarActionRemove{}
	case BossBarActUpdateHealth:
		health, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		action = BossBarActionUpdateHealth{Health: health}
	case BossBarActUpdateTitle:
		title, err := reader.ReadChat()
		if err != nil {
			return nil, err
		}
		action = BossBarActionUpdateTitle{Title: title}
	case BossBarActUpdateStyle:
		color, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		division, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		action = BossBarActionUpdateStyle{
			Color:    color,
			Division: division,
		}
	case BossBarActUpdateFlags:
		flags, err := reader.ReadUByte()
		if err != nil {
			return nil, err
		}
		action = BossBarActionUpdateFlags{Flags: flags}
	}
	return BossBar{
		Uuid:   uid,
		Action: action,
	}, nil
}

func (p BossBar) IsValid() (reason error) {
	if p.Action == nil {
		return fmt.Errorf("unknown boss bar action")
	}
	switch p.Action.Type() {
	case BossBarActAdd:
		action := p.Action.(BossBarActionAdd)
		if action.Health < 0 || 1 < action.Health {
			return fmt.Errorf("boss bar health must be in range <0; 1> was %v", action.Health)
		} else if action.Color < 0 || 6 < action.Color {
			return fmt.Errorf("unknown boss bar color index %d", action.Color)
		} else if action.Division < 0 || 4 < action.Division {
			return fmt.Errorf("unknown boss bar division index %d", action.Division)
		} else if (action.Flags &^ 0b111) != 0 {
			return fmt.Errorf("unknown boss bar flags %b", action.Flags)
		}
	case BossBarActUpdateHealth:
		action := p.Action.(BossBarActionUpdateHealth)
		if action.Health < 0 || 1 < action.Health {
			return fmt.Errorf("boss bar health must be in range <0; 1> was %v", action.Health)
		}
	case BossBarActUpdateStyle:
		action := p.Action.(BossBarActionUpdateStyle)
		if action.Color < 0 || 6 < action.Color {
			return fmt.Errorf("unknown boss bar color index %d", action.Color)
		} else if action.Division < 0 || 4 < action.Division {
			return fmt.Errorf("unknown boss bar division index %d", action.Division)
		}
	case BossBarActUpdateFlags:
		action := p.Action.(BossBarActionUpdateFlags)
		if (action.Flags &^ 0b111) != 0 {
			return fmt.Errorf("unknown boss bar flags %b", action.Flags)
		}
	}
	return nil
}
