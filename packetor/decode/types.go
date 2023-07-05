package decode

import (
	"Packetor/packetor/nbt"
	"fmt"
	"github.com/sirupsen/logrus"
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
		GetIndex(state int32) int32
	}
	PaletteSingle int32
	PaletteArray  []int32
	PaletteDirect struct {
	}
	PaletteContainer struct {
		Dimension    uint8
		BitsPerEntry uint8
		MaxBits      uint8
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

func (receiver PaletteSingle) GetIndex(state int32) int32 {
	if int32(receiver) != state {
		return -1
	}
	return 0
}

func (receiver PaletteArray) Get(index int32) int32 {
	return receiver[index]
}

func (receiver PaletteArray) GetIndex(state int32) int32 {
	for i := range receiver {
		if receiver[i] == state {
			return int32(i)
		}
	}
	return -1
}

func (receiver PaletteDirect) Get(index int32) int32 {
	return index
}

func (receiver PaletteDirect) GetIndex(state int32) int32 {
	return state
}

func (p *PaletteContainer) getState0(x int, y int, z int) int32 {
	if _, ok := p.Palette.(PaletteSingle); ok {
		return p.Palette.Get(0)
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
	return int32(value)
}

func (p *PaletteContainer) GetState(x int, y int, z int) int32 {
	return p.Palette.Get(p.getState0(x, y, z))
}

func (p *PaletteContainer) SetState(x int, y int, z int, state int32) {
	stateIdx := p.Palette.GetIndex(state)
	if stateIdx == -1 {
		if palette, ok := p.Palette.(PaletteArray); ok {
			if len(palette) >= 1<<p.BitsPerEntry {
				p.resize(p.BitsPerEntry + 1)
			}
			stateIdx = int32(len(palette))
			p.Palette = append(palette, state)
		} else if palette, ok := p.Palette.(PaletteSingle); ok {
			dimension := int(p.Dimension)
			p.Data = make([]uint64, dimension*dimension*dimension)
			p.resize(1)
			p.Palette = PaletteArray{int32(palette), state}
			stateIdx = 1
		} else {
			logrus.Fatal("not in palette")
			return
		}
	}

	p.setState0(x, y, z, stateIdx)
}

func (p *PaletteContainer) setState0(x int, y int, z int, stateIdx int32) {
	var sectionIndex int
	if p.Dimension == 16 {
		sectionIndex = ((y & 0xf) << 8) | ((z & 0xf) << 4) | (x & 0xf)
	} else {
		sectionIndex = ((y & 0x3) << 4) | ((z & 0x3) << 2) | (x & 0x3)
	}
	valuesPerLong := 64 / int(p.BitsPerEntry)
	index := sectionIndex / valuesPerLong
	bitIndex := (sectionIndex - index*valuesPerLong) * int(p.BitsPerEntry)

	clear := uint64((1 << p.BitsPerEntry) - 1)
	p.Data[index] = p.Data[index]&^(clear<<bitIndex) | (uint64(stateIdx) << bitIndex)
}

func (p *PaletteContainer) resize(newBits uint8) {
	if newBits > p.MaxBits {
		newBits = 15
	}
	dimension := int(p.Dimension)
	oldStates := make([]int32, dimension*dimension*dimension)
	for rx := 0; rx < dimension; rx++ {
		for ry := 0; ry < dimension; ry++ {
			for rz := 0; rz < dimension; rz++ {
				oldStates[(rx*dimension+ry)*dimension+rz] = p.getState0(rx, ry, rz)
			}
		}
	}
	p.BitsPerEntry = newBits
	for rx := 0; rx < dimension; rx++ {
		for ry := 0; ry < dimension; ry++ {
			for rz := 0; rz < dimension; rz++ {
				p.setState0(rx, ry, rz, oldStates[(rx*dimension+ry)*dimension+rz])
			}
		}
	}
}
