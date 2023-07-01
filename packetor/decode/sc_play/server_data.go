package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type ServerData struct {
	MOTD               string
	HasIcon            bool
	Icon               []byte
	EnforcesSecureChat bool
}

func (p ServerData) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	motd, err := reader.ReadChat()
	if err != nil {
		return nil, err
	}
	hasIcon, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	var icon []byte
	if hasIcon {
		count, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if count < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
		}
		icon, err = reader.ReadBytesExact(int(count))
	}
	esc, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return ServerData{
		MOTD:               motd,
		HasIcon:            hasIcon,
		Icon:               icon,
		EnforcesSecureChat: esc,
	}, nil
}

func (p ServerData) IsValid() (reason error) {
	// TODO: validate icon
	return nil
}
