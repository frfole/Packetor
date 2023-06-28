package sc_play

import "Packetor/packetor/decode"

type (
	ParticleData interface {
	}
	ParticleBlock       int32
	ParticleBlockMarker int32
	ParticleDust        struct {
		Red   float32
		Green float32
		Blue  float32
		Scale float32
	}
	ParticleDustTransition struct {
		FromRed   float32
		FromGreen float32
		FromBlue  float32
		Scale     float32
		ToRed     float32
		ToGreen   float32
		ToBlue    float32
	}
	ParticleFallingDust int32
	ParticleSculkCharge float32
	ParticleItem        decode.Slot
	ParticleVibration   struct {
		PositionSourceType string
		BlockPosition      decode.Position
		EntityID           int32
		EntityEyeHeight    float32
		Ticks              int32
	}
	ParticleShriek int32
)

type Particle struct {
	ParticleID    int32
	LongDistance  bool
	X             float64
	Y             float64
	Z             float64
	OffsetX       float32
	OffsetY       float32
	OffsetZ       float32
	MaxSpeed      float32
	ParticleCount int32
	Data          ParticleData
}

func (p Particle) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	pid, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	longDist, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	x, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	y, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	z, err := reader.ReadDouble()
	if err != nil {
		return nil, err
	}
	ox, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	oy, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	oz, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	speed, err := reader.ReadFloat()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadInt()
	if err != nil {
		return nil, err
	}
	var data ParticleData
	switch pid {
	case 2:
		value, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		data = ParticleBlock(value)
	case 3:
		value, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		data = ParticleBlockMarker(value)
	case 14:
		red, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		green, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		blue, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		scale, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		data = ParticleDust{
			Red:   red,
			Green: green,
			Blue:  blue,
			Scale: scale,
		}
	case 15:
		red, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		green, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		blue, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		scale, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		red2, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		green2, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		blue2, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		data = ParticleDustTransition{
			FromRed:   red,
			FromGreen: green,
			FromBlue:  blue,
			Scale:     scale,
			ToRed:     red2,
			ToGreen:   green2,
			ToBlue:    blue2,
		}
	case 25:
		value, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		data = ParticleFallingDust(value)
	case 33:
		value, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		data = ParticleSculkCharge(value)
	case 42:
		value, err := reader.ReadSlot()
		if err != nil {
			return nil, err
		}
		data = ParticleItem(value)
	case 43:
		posSrcType, err := reader.ReadString()
		if err != nil {
			return nil, err
		}
		bPos, err := reader.ReadPosition()
		if err != nil {
			return nil, err
		}
		eid, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		height, err := reader.ReadFloat()
		if err != nil {
			return nil, err
		}
		ticks, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		data = ParticleVibration{
			PositionSourceType: posSrcType,
			BlockPosition:      bPos,
			EntityID:           eid,
			EntityEyeHeight:    height,
			Ticks:              ticks,
		}
	case 95:
		value, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		data = ParticleShriek(value)
	}
	return Particle{
		ParticleID:    pid,
		LongDistance:  longDist,
		X:             x,
		Y:             y,
		Z:             z,
		OffsetX:       ox,
		OffsetY:       oy,
		OffsetZ:       oz,
		MaxSpeed:      speed,
		ParticleCount: count,
		Data:          data,
	}, nil
}

func (p Particle) IsValid() (reason error) {
	// TODO: validate data
	return nil
}
