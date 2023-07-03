package cs_play

import (
	"Packetor/packetor/decode"
	"fmt"
)

type ProgramStructureBlock struct {
	Location  decode.Position
	Action    int32
	Mode      int32
	Name      string
	OffsetX   int8
	OffsetY   int8
	OffsetZ   int8
	SizeX     int8
	SizeY     int8
	SizeZ     int8
	Mirror    int32
	Rotation  int32
	Metadata  string
	Integrity float32
	Seed      int64
	Flags     uint8
}

func (p ProgramStructureBlock) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	action, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	mode, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	name, err := reader.ReadString()
	if err != nil {
		return nil, err
	}
	ox, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	oy, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	oz, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	sx, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	sy, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	sz, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	mirror, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	rotation, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	metadata, err := reader.ReadString0(128)
	if err != nil {
		return nil, err
	}
	integrity, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	seed, err := reader.ReadVarLong()
	if err != nil {
		return nil, err
	}
	flags, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	return ProgramStructureBlock{
		Location:  loc,
		Action:    action,
		Mode:      mode,
		Name:      name,
		OffsetX:   ox,
		OffsetY:   oy,
		OffsetZ:   oz,
		SizeX:     sx,
		SizeY:     sy,
		SizeZ:     sz,
		Mirror:    mirror,
		Rotation:  rotation,
		Metadata:  metadata,
		Integrity: integrity,
		Seed:      seed,
		Flags:     flags,
	}, nil
}

func (p ProgramStructureBlock) IsValid() (reason error) {
	if p.Action < 0 || 3 < p.Action {
		return fmt.Errorf("action must be in <0; 3> was %d", p.Action)
	}
	if p.Mode < 0 || 3 < p.Mode {
		return fmt.Errorf("mode must be in <0; 3> was %d", p.Mode)
	}
	if p.SizeX < 0 {
		return fmt.Errorf("size x must be atleast 0 was %d", p.SizeX)
	}
	if p.SizeY < 0 {
		return fmt.Errorf("size y must be atleast 0 was %d", p.SizeY)
	}
	if p.SizeZ < 0 {
		return fmt.Errorf("size z must be atleast 0 was %d", p.SizeZ)
	}
	if p.Mirror < 0 || 2 < p.Mirror {
		return fmt.Errorf("mirrot must be in <0; 2> was %d", p.Mirror)
	}
	if p.Rotation < 0 || 3 < p.Rotation {
		return fmt.Errorf("rotation must be in <0; 3> was %d", p.Rotation)
	}
	if !(0 <= p.Integrity && p.Integrity <= 1) {
		return fmt.Errorf("integrity must be in <0; 1> was %v", p.Integrity)
	}
	if (p.Flags &^ 0b111) != 0 {
		return fmt.Errorf("unknown flags %b", p.Flags)
	}
	return nil
}
