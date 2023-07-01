package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type UpdateTeamsInfo struct {
	DisplayName       string
	FriendlyFlags     uint8
	NameTagVisibility string
	CollisionRule     string
	TeamColor         int32
	TeamPrefix        string
	TeamSuffix        string
}

type UpdateTeams struct {
	TeamName  string
	Operation uint8
	Info      UpdateTeamsInfo
	Members   []string
}

func (p UpdateTeams) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	teamName, err := reader.ReadString()
	if err != nil {
		return nil, err
	}
	operation, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	info := UpdateTeamsInfo{}
	var members []string
	if operation == 0 || operation == 2 {
		info.DisplayName, err = reader.ReadChat()
		if err != nil {
			return nil, err
		}
		info.FriendlyFlags, err = reader.ReadUByte()
		if err != nil {
			return nil, err
		}
		info.NameTagVisibility, err = reader.ReadString0(40)
		if err != nil {
			return nil, err
		}
		info.CollisionRule, err = reader.ReadString0(40)
		if err != nil {
			return nil, err
		}
		info.TeamColor, err = reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		info.TeamPrefix, err = reader.ReadChat()
		if err != nil {
			return nil, err
		}
		info.TeamSuffix, err = reader.ReadChat()
		if err != nil {
			return nil, err
		}
	}
	if operation == 0 || operation == 3 || operation == 4 {
		count, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if count < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
		}
		members = make([]string, count)
		for i := int32(0); i < count; i++ {
			members[i], err = reader.ReadString()
			if err != nil {
				return nil, err
			}
		}
	}
	return UpdateTeams{
		TeamName:  teamName,
		Operation: operation,
		Info:      info,
		Members:   members,
	}, nil
}

func (p UpdateTeams) IsValid() (reason error) {
	if 4 < p.Operation {
		return fmt.Errorf("operation must be in <0; 4> was %d", p.Operation)
	}
	if p.Operation == 0 || p.Operation == 2 {
		if (p.Info.FriendlyFlags &^ 0b11) != 0 {
			return fmt.Errorf("unknown friendly flags %b", p.Info.FriendlyFlags)
		}
		if p.Info.NameTagVisibility != "always" && p.Info.NameTagVisibility != "hideForOtherTeams" &&
			p.Info.NameTagVisibility != "hideForOwnTeam" && p.Info.NameTagVisibility != "never" {
			return fmt.Errorf("unknown name tag visibility")
		}
		if p.Info.NameTagVisibility != "always" && p.Info.NameTagVisibility != "pushOtherTeams" &&
			p.Info.NameTagVisibility != "pushOwnTeam" && p.Info.NameTagVisibility != "never" {
			return fmt.Errorf("unknown collision rule")
		}
		if p.Info.TeamColor < 0 || 21 < p.Info.TeamColor {
			return fmt.Errorf("team color must be in <0; 21> was %d", p.Info.TeamColor)
		}
	}
	// TODO: validate members name/uuid
	return nil
}
