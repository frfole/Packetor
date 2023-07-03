package sc_play

import (
	"Packetor/packetor/decode"
)

type SetEntityMetadataEntry struct {
	Index uint8
	Type  int32
	Value any
}

type SetEntityMetadata struct {
	EntityID int32
	Entries  []SetEntityMetadataEntry
}

func (p SetEntityMetadata) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	var entries []SetEntityMetadataEntry
	idx := uint8(0)
	for idx != 255 {
		idx, err = reader.ReadUByte()
		if err != nil {
			return nil, err
		}
		if idx == 255 {
			break
		}
		metaType, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		meta := SetEntityMetadataEntry{
			Index: idx,
			Type:  metaType,
		}
		switch meta.Type {
		case 0:
			meta.Value, err = reader.ReadSByte()
			if err != nil {
				return nil, err
			}
		case 1:
			meta.Value, err = reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
		case 2:
			meta.Value, err = reader.ReadVarLong()
			if err != nil {
				return nil, err
			}
		case 3:
			meta.Value, err = reader.ReadFloat()
			if err != nil {
				return nil, err
			}
		case 4:
			meta.Value, err = reader.ReadString()
			if err != nil {
				return nil, err
			}
		case 5:
			meta.Value, err = reader.ReadChat()
			if err != nil {
				return nil, err
			}
		case 6:
			hasValue, err := reader.ReadBoolean()
			if err != nil {
				return nil, err
			}
			if hasValue {
				meta.Value, err = reader.ReadSByte()
				if err != nil {
					return nil, err
				}
			}
		case 7:
			meta.Value, err = reader.ReadSlot()
			if err != nil {
				return nil, err
			}
		case 8:
			meta.Value, err = reader.ReadBoolean()
			if err != nil {
				return nil, err
			}
		case 9:
			rx, err := reader.ReadFloat()
			if err != nil {
				return nil, err
			}
			ry, err := reader.ReadFloat()
			if err != nil {
				return nil, err
			}
			rz, err := reader.ReadFloat()
			if err != nil {
				return nil, err
			}
			meta.Value = []float32{rx, ry, rz}
		case 10:
			meta.Value, err = reader.ReadPosition()
			if err != nil {
				return nil, err
			}
		case 11:
			has, err := reader.ReadBoolean()
			if err != nil {
				return nil, err
			}
			if has {
				meta.Value, err = reader.ReadPosition()
				if err != nil {
					return nil, err
				}
			}
		case 12:
			meta.Value, err = reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
		case 13:
			has, err := reader.ReadBoolean()
			if err != nil {
				return nil, err
			}
			if has {
				meta.Value, err = reader.ReadUuid()
				if err != nil {
					return nil, err
				}
			}
		case 14:
			meta.Value, err = reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
		case 15:
			meta.Value, err = reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			if meta.Value == 0 {
				meta.Value = nil
			}
		case 16:
			meta.Value, err = reader.ReadNbt()
			if err != nil {
				return nil, err
			}
		case 17:
			pid, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			particle, err := readParticle(pid, reader)
			if err != nil {
				return nil, err
			}
			meta.Value = struct {
				ParticleID int32
				Data       ParticleData
			}{ParticleID: pid, Data: particle}
		case 18:
			vType, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			vProf, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			vLevel, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			meta.Value = []int32{vType, vProf, vLevel}
		case 19:
			oeid, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			if eid == 0 {
				meta.Value = nil
			} else {
				meta.Value = oeid + 1
			}
		case 20, 21, 22:
			meta.Value, err = reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
		case 23:
			has, err := reader.ReadBoolean()
			if err != nil {
				return nil, err
			}
			if has {
				name, err := reader.ReadIdentifier()
				if err != nil {
					return nil, err
				}
				loc, err := reader.ReadPosition()
				if err != nil {
					return nil, err
				}
				meta.Value = struct {
					DimensionName string
					Location      decode.Position
				}{DimensionName: name, Location: loc}
			}
		case 24, 25:
			meta.Value, err = reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
		case 26:
			x, err := reader.ReadFloat()
			if err != nil {
				return nil, err
			}
			y, err := reader.ReadFloat()
			if err != nil {
				return nil, err
			}
			z, err := reader.ReadFloat()
			if err != nil {
				return nil, err
			}
			meta.Value = []float32{x, y, z}
		case 27:
			x, err := reader.ReadFloat()
			if err != nil {
				return nil, err
			}
			y, err := reader.ReadFloat()
			if err != nil {
				return nil, err
			}
			z, err := reader.ReadFloat()
			if err != nil {
				return nil, err
			}
			w, err := reader.ReadFloat()
			if err != nil {
				return nil, err
			}
			meta.Value = []float32{x, y, z, w}
		}
		entries = append(entries, meta)
	}
	return SetEntityMetadata{
		EntityID: eid,
		Entries:  entries,
	}, nil
}

func (p SetEntityMetadata) IsValid() (reason error) {
	// TODO: validate
	return nil
}
