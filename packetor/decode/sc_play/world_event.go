package sc_play

import (
	"Packetor/packetor/decode"
)

type WorldEvent struct {
	Event                 int32
	Location              decode.Position
	Data                  int32
	DisableRelativeVolume bool
}

func (p WorldEvent) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	event, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	loc, err := reader.ReadPosition()
	if err != nil {
		return nil, err
	}
	data, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	disRelVol, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	return WorldEvent{
		Event:                 event,
		Location:              loc,
		Data:                  data,
		DisableRelativeVolume: disRelVol,
	}, nil
}

func (p WorldEvent) IsValid() (reason error) {
	// TODO: validate
	return nil
}
