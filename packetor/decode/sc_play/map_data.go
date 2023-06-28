package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type MapDataIcon struct {
	Type           int32
	X              int8
	Z              int8
	Direction      int8
	HasDisplayName bool
	DisplayName    string
}

type MapData struct {
	MapID   int32
	Scale   int8
	Locked  bool
	Icons   []MapDataIcon
	Columns uint8
	Rows    uint8
	X       int8
	Z       int8
	Data    []byte
}

func (p MapData) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	mid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	scale, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	locked, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	hasIcons, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	var icons []MapDataIcon
	if hasIcons {
		count, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if count < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
		}
		icons = make([]MapDataIcon, count)
		for i := int32(0); i < count; i++ {
			iType, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			ix, err := reader.ReadSByte()
			if err != nil {
				return nil, err
			}
			iz, err := reader.ReadSByte()
			if err != nil {
				return nil, err
			}
			iDir, err := reader.ReadSByte()
			if err != nil {
				return nil, err
			}
			hasDN, err := reader.ReadBoolean()
			if err != nil {
				return nil, err
			}
			iDN := ""
			if hasDN {
				iDN, err = reader.ReadChat()
				if err != nil {
					return nil, err
				}
			}
			icons[i] = MapDataIcon{
				Type:           iType,
				X:              ix,
				Z:              iz,
				Direction:      iDir,
				HasDisplayName: hasDN,
				DisplayName:    iDN,
			}
		}
	}
	col, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	rows := uint8(0)
	x := int8(0)
	z := int8(0)
	var data []byte
	if col != 0 {
		rows, err = reader.ReadUByte()
		if err != nil {
			return nil, err
		}
		x, err = reader.ReadSByte()
		if err != nil {
			return nil, err
		}
		z, err = reader.ReadSByte()
		if err != nil {
			return nil, err
		}
		count, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if count < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
		}
		data, err = reader.ReadBytesExact(int(count))
		if err != nil {
			return nil, err
		}
	}
	return MapData{
		MapID:   mid,
		Scale:   scale,
		Locked:  locked,
		Icons:   icons,
		Columns: col,
		Rows:    rows,
		X:       x,
		Z:       z,
		Data:    data,
	}, nil
}

func (p MapData) IsValid() (reason error) {
	if p.Icons != nil {
		for i, icon := range p.Icons {
			if icon.Type < 0 || 26 < icon.Type {
				return fmt.Errorf("icon type must be in range <0; 26> was %d at index %d", icon.Type, i)
			}
		}
	}
	return nil
}
