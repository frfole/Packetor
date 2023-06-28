package sc_play

import "Packetor/packetor/decode"

type InitializeWorldBorder struct {
	X                      float64
	Z                      float64
	OldDiameter            float64
	NewDiameter            float64
	Speed                  int64
	PortalTeleportBoundary int32
	WarningBlocks          int32
	WarningTime            int32
}

func (p InitializeWorldBorder) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	x, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	z, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	oldD, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	newD, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	speed, err := reader.ReadVarLong()
	if err != nil {
		return nil, err
	}
	portal, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	warnBlocks, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	warnTime, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return InitializeWorldBorder{
		X:                      x,
		Z:                      z,
		OldDiameter:            oldD,
		NewDiameter:            newD,
		Speed:                  speed,
		PortalTeleportBoundary: portal,
		WarningBlocks:          warnBlocks,
		WarningTime:            warnTime,
	}, nil
}

func (p InitializeWorldBorder) IsValid() (reason error) {
	return nil
}
