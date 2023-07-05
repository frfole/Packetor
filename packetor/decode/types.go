package decode

import (
	"Packetor/packetor/nbt"
	"fmt"
	"reflect"
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
	BitSet  []uint64
	Palette interface {
		Get(index int32) int32
	}
	PaletteSingle int32
	PaletteArray  []int32
	PaletteDirect struct {
	}
	PaletteContainer struct {
		Dimension    uint8
		BitsPerEntry uint8
		Palette      Palette
		Data         []uint64
	}
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

func (receiver PaletteSingle) Get(_ int32) int32 {
	return int32(receiver)
}

func (receiver PaletteArray) Get(index int32) int32 {
	return receiver[index]
}

func (receiver PaletteDirect) Get(index int32) int32 {
	return index
}

func (p PaletteContainer) GetState(x int, y int, z int) int32 {
	if reflect.TypeOf(p.Palette) == reflect.TypeOf(PaletteSingle(0)) {
		return 0
	}
	var sectionIndex int
	if p.Dimension == 16 {
		sectionIndex = ((y & 0xf) << 8) | ((z & 0xf) << 4) | (x & 0xf)
	} else {
		sectionIndex = ((y & 0x3) << 4) | ((z & 0x3) << 2) | (x & 0x3)
	}
	valuesPerLong := 64 / int(p.BitsPerEntry)
	index := sectionIndex / valuesPerLong
	bitIndex := (sectionIndex - index*valuesPerLong) * int(p.BitsPerEntry)
	value := (p.Data[index] >> bitIndex) & ((1 << p.BitsPerEntry) - 1)
	return p.getState0(int32(value))
}

func (p PaletteContainer) getState0(data int32) int32 {
	return p.Palette.Get(data)
}
