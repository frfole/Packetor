package decode

import (
	"Packetor/packetor/nbt"
	"fmt"
)

type (
	Angle    byte
	Position int64
	Slot     struct {
		Present   bool
		ItemID    int32
		ItemCount uint8
		ItemNbt   nbt.Compound
	}
	BitSet []uint64
)

func (a Angle) asFloat32() float32 {
	return float32(a) * (360.0 / 256.0)
}

func (a Angle) asFloat64() float64 {
	return float64(a) * (360.0 / 256.0)
}

func (p Position) X() int32 {
	return int32(p >> 38)
}

func (p Position) Y() int32 {
	return int32(p << 52 >> 52)
}

func (p Position) Z() int32 {
	return int32(p << 26 >> 38)
}

func (p Position) String() string {
	return fmt.Sprintf("{%v %v %v}", p.X(), p.Y(), p.Z())
}

func (b BitSet) IsSet(index int) bool {
	if index>>6 >= len(b) {
		return false
	}
	return (b[index>>6] >> (index & 63)) == 1
}
