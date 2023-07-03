package cs_play

import "Packetor/packetor/decode"

type ProgramCommandBlockMinecart struct {
	EntityID    int32
	Command     string
	TrackOutput bool
}

func (p ProgramCommandBlockMinecart) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	eid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	cmd, err := reader.ReadString()
	if err != nil {
		return nil, err
	}
	track, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return ProgramCommandBlockMinecart{
		EntityID:    eid,
		Command:     cmd,
		TrackOutput: track,
	}, nil
}

func (p ProgramCommandBlockMinecart) IsValid() (reason error) {
	return nil
}
