package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type UpdateAdvancementsTaskDisplay struct {
	Title             string
	Description       string
	Icon              decode.Slot
	FrameType         int32
	Flags             int32
	BackgroundTexture string
	X                 float32
	Y                 float32
}

type UpdateAdvancementsTask struct {
	HasParent     bool
	ParentID      string
	HasDisplay    bool
	Display       UpdateAdvancementsTaskDisplay
	Criteria      []string
	Requirements  [][]string
	SendTelemetry bool
}

type UpdateAdvancementsProgress struct {
	HasAchieved  bool
	AchievedDate int64
}

type UpdateAdvancements struct {
	ClearCurrent       bool
	AdvancementMapping map[string]UpdateAdvancementsTask
	ToRemove           []string
	ProgressMapping    map[string]map[string]UpdateAdvancementsProgress
}

func (p UpdateAdvancements) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	clearCurrent, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	tasks := map[string]UpdateAdvancementsTask{}
	for i := int32(0); i < count; i++ {
		key, err := reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
		hasParent, err := reader.ReadBoolean()
		if err != nil {
			return nil, err
		}
		parentId := ""
		if hasParent {
			parentId, err = reader.ReadIdentifier()
			if err != nil {
				return nil, err
			}
		}
		hasDisplay, err := reader.ReadBoolean()
		if err != nil {
			return nil, err
		}
		display := UpdateAdvancementsTaskDisplay{}
		if hasDisplay {
			display.Title, err = reader.ReadChat()
			if err != nil {
				return nil, err
			}
			display.Description, err = reader.ReadChat()
			if err != nil {
				return nil, err
			}
			display.Icon, err = reader.ReadSlot()
			if err != nil {
				return nil, err
			}
			display.FrameType, err = reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			display.Flags, err = reader.ReadInt()
			if err != nil {
				return nil, err
			}
			display.BackgroundTexture = ""
			if (display.Flags & 0x01) == 0x01 {
				display.BackgroundTexture, err = reader.ReadIdentifier()
				if err != nil {
					return nil, err
				}
			}
			display.X, err = reader.ReadFloat()
			if err != nil {
				return nil, err
			}
			display.Y, err = reader.ReadFloat()
			if err != nil {
				return nil, err
			}
		}
		length, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if length < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", length), error2.ErrDecodeTooSmall)
		}
		criteria := make([]string, length)
		for j := int32(0); j < length; j++ {
			criteria[j], err = reader.ReadIdentifier()
			if err != nil {
				return nil, err
			}
		}
		length, err = reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if length < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", length), error2.ErrDecodeTooSmall)
		}
		req := make([][]string, length)
		for j := int32(0); j < length; j++ {
			length2, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			} else if length2 < 0 {
				return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", length2), error2.ErrDecodeTooSmall)
			}
			req[j] = make([]string, length2)
			for k := int32(0); k < length2; k++ {
				req[j][k], err = reader.ReadString()
				if err != nil {
					return nil, err
				}
			}
		}
		telemetry, err := reader.ReadBoolean()
		tasks[key] = UpdateAdvancementsTask{
			HasParent:     hasParent,
			ParentID:      parentId,
			HasDisplay:    hasDisplay,
			Display:       display,
			Criteria:      criteria,
			Requirements:  req,
			SendTelemetry: telemetry,
		}
	}
	count, err = reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	toRemove := make([]string, count)
	for i := int32(0); i < count; i++ {
		toRemove[i], err = reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
	}
	count, err = reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	progress := map[string]map[string]UpdateAdvancementsProgress{}
	for i := int32(0); i < count; i++ {
		key1, err := reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
		progress[key1] = map[string]UpdateAdvancementsProgress{}
		count2, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if count2 < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count2), error2.ErrDecodeTooSmall)
		}
		for j := int32(0); j < count2; j++ {
			key2, err := reader.ReadIdentifier()
			if err != nil {
				return nil, err
			}
			obtained, err := reader.ReadBoolean()
			if err != nil {
				return nil, err
			}
			obtainedDate := int64(0)
			if obtained {
				obtainedDate, err = reader.ReadLong()
				if err != nil {
					return nil, err
				}
			}
			progress[key1][key2] = UpdateAdvancementsProgress{
				HasAchieved:  obtained,
				AchievedDate: obtainedDate,
			}
		}
	}
	return UpdateAdvancements{
		ClearCurrent:       clearCurrent,
		AdvancementMapping: tasks,
		ToRemove:           toRemove,
		ProgressMapping:    progress,
	}, nil
}

func (p UpdateAdvancements) IsValid() (reason error) {
	// TODO: validate task display
	return nil
}
