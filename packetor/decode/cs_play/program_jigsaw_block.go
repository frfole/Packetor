package cs_play

import "Packetor/packetor/decode"

type ProgramJigsawBlock struct {
	Location   decode.Position
	Name       string
	Target     string
	Pool       string
	FinalState string
	JointType  string
}

func (p ProgramJigsawBlock) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	name, err := reader.ReadIdentifier()
	if err != nil {
		return nil, err
	}
	target, err := reader.ReadIdentifier()
	if err != nil {
		return nil, err
	}
	pool, err := reader.ReadIdentifier()
	if err != nil {
		return nil, err
	}
	finalState, err := reader.ReadString()
	if err != nil {
		return nil, err
	}
	jointType, err := reader.ReadString()
	if err != nil {
		return nil, err
	}
	return ProgramJigsawBlock{
		Location:   loc,
		Name:       name,
		Target:     target,
		Pool:       pool,
		FinalState: finalState,
		JointType:  jointType,
	}, nil
}

func (p ProgramJigsawBlock) IsValid() (reason error) {
	return nil
}
