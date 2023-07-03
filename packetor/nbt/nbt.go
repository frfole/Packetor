package nbt

import (
	"errors"
	"io"
	"math"
	"unsafe"
)

const (
	TagEnd byte = iota
	TagByte
	TagShort
	TagInt
	TagLong
	TagFloat
	TagDouble
	TagByteArr
	TagString
	TagList
	TagCompound
	TagIntArr
	TagLongArr
)

var (
	ErrUnknownTag = errors.New("unknown tag")
)

type Tag interface {
	Type() byte
}

type Byte int8
type Short int16
type Int int32
type Long int64
type Float float32
type Double float64
type ByteArray []int8
type String string
type List []Tag
type Compound map[string]Tag
type IntArray []int32
type LongArray []int64

// ReadNbt returns compound inside reader parsed and wrapped in another
// compound a:{} would return {a:{}}
func ReadNbt(reader io.Reader) (value Compound, err error) {
	return readNbt(reader, true)
}

func readNbt(reader io.Reader, isTop bool) (value Compound, err error) {
	tagType := make([]byte, 1)
	nameLenB := make([]byte, 2)
	compound := Compound(map[string]Tag{})
	for {
		_, err = io.ReadFull(reader, tagType)
		if isTop && errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return Compound{}, err
		}
		if tagType[0] == TagEnd {
			break
		}
		_, err = io.ReadFull(reader, nameLenB)
		if err != nil {
			return Compound{}, err
		}
		nameLen := uint16(nameLenB[0])<<8 | uint16(nameLenB[1])
		nameB := make([]byte, nameLen)
		_, err = io.ReadFull(reader, nameB)
		if err != nil {
			return Compound{}, err
		}
		name := string(nameB)
		tag, err := readTag(reader, tagType[0])
		if err != nil {
			return Compound{}, err
		}
		compound[name] = tag
		if isTop {
			break
		}
	}

	return compound, nil
}

func readTag(reader io.Reader, tagType byte) (tag Tag, err error) {
	switch tagType {
	case TagByte:
		data := make([]byte, 1)
		_, err = io.ReadFull(reader, data)
		if err != nil {
			return nil, err
		}
		return Byte(data[0]), nil
	case TagShort:
		data := make([]byte, 2)
		_, err = io.ReadFull(reader, data)
		if err != nil {
			return nil, err
		}
		return Short(int16(data[0])<<8 | int16(data[1])), nil
	case TagInt:
		data := make([]byte, 4)
		_, err = io.ReadFull(reader, data)
		if err != nil {
			return nil, err
		}
		return Int(int32(data[0])<<24 | int32(data[1])<<16 | int32(data[2])<<8 | int32(data[3])), nil
	case TagLong:
		data := make([]byte, 8)
		_, err = io.ReadFull(reader, data)
		if err != nil {
			return nil, err
		}
		return Long(int64(data[0])<<56 | int64(data[1])<<48 | int64(data[2])<<40 | int64(data[3])<<32 | int64(data[4])<<24 | int64(data[5])<<16 | int64(data[6])<<8 | int64(data[7])), nil
	case TagFloat:
		data := make([]byte, 4)
		_, err = io.ReadFull(reader, data)
		if err != nil {
			return nil, err
		}
		return Float(math.Float32frombits(uint32(data[0])<<24 | uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3]))), nil
	case TagDouble:
		data := make([]byte, 8)
		_, err = io.ReadFull(reader, data)
		if err != nil {
			return nil, err
		}
		return Double(math.Float64frombits(uint64(data[0])<<56 | uint64(data[1])<<48 | uint64(data[2])<<40 | uint64(data[3])<<32 | uint64(data[4])<<24 | uint64(data[5])<<16 | uint64(data[6])<<8 | uint64(data[7]))), nil
	case TagByteArr:
		dataLenB := make([]byte, 4)
		_, err = io.ReadFull(reader, dataLenB)
		if err != nil {
			return nil, err
		}
		dataLen := int32(dataLenB[0])<<24 | int32(dataLenB[1])<<16 | int32(dataLenB[2])<<8 | int32(dataLenB[3])
		data := make([]byte, dataLen)
		_, err = io.ReadFull(reader, data)
		if err != nil {
			return nil, err
		}
		return ByteArray(unsafe.Slice((*int8)(unsafe.Pointer(&data[0])), dataLen)), nil
	case TagString:
		dataLenB := make([]byte, 2)
		_, err = io.ReadFull(reader, dataLenB)
		if err != nil {
			return nil, err
		}
		data := make([]byte, uint16(dataLenB[0])<<8|uint16(dataLenB[1]))
		_, err = io.ReadFull(reader, data)
		if err != nil {
			return nil, err
		}
		return String(data), nil
	case TagList:
		header := make([]byte, 5)
		_, err = io.ReadFull(reader, header)
		if err != nil {
			return nil, err
		}
		length := int(int32(header[1])<<24 | int32(header[2])<<16 | int32(header[3])<<8 | int32(header[4]))
		if length <= 0 {
			return List([]Tag{}), nil
		}
		data := make([]Tag, length)
		for i := 0; i < length; i++ {
			t, err := readTag(reader, header[0])
			if err != nil {
				return nil, err
			}
			data[i] = t
		}
		return List(data), nil
	case TagCompound:
		nbt, err := readNbt(reader, false)
		if err != nil {
			return nil, err
		}
		return nbt, nil
	case TagIntArr:
		dataLenB := make([]byte, 4)
		_, err = io.ReadFull(reader, dataLenB)
		if err != nil {
			return nil, err
		}
		dataLen := int32(dataLenB[0])<<24 | int32(dataLenB[1])<<16 | int32(dataLenB[2])<<8 | int32(dataLenB[3])
		data := make([]byte, dataLen*4)
		_, err = io.ReadFull(reader, data)
		if err != nil {
			return nil, err
		}
		return IntArray(unsafe.Slice((*int32)(unsafe.Pointer(&data[0])), dataLen)), nil
	case TagLongArr:
		dataLenB := make([]byte, 4)
		_, err = io.ReadFull(reader, dataLenB)
		if err != nil {
			return nil, err
		}
		dataLen := int32(dataLenB[0])<<24 | int32(dataLenB[1])<<16 | int32(dataLenB[2])<<8 | int32(dataLenB[3])
		data := make([]byte, dataLen*8)
		_, err = io.ReadFull(reader, data)
		if err != nil {
			return nil, err
		}
		return LongArray(unsafe.Slice((*int64)(unsafe.Pointer(&data[0])), dataLen)), nil
	}
	return nil, ErrUnknownTag
}

func (receiver Byte) Type() byte {
	return TagByte
}
func (receiver Int) Type() byte {
	return TagInt
}
func (receiver Short) Type() byte {
	return TagShort
}
func (receiver Long) Type() byte {
	return TagLong
}
func (receiver Float) Type() byte {
	return TagFloat
}
func (receiver Double) Type() byte {
	return TagDouble
}
func (receiver ByteArray) Type() byte {
	return TagByteArr
}
func (receiver String) Type() byte {
	return TagString
}
func (receiver List) Type() byte {
	return TagList
}
func (receiver Compound) Type() byte {
	return TagCompound
}
func (receiver IntArray) Type() byte {
	return TagIntArr
}
func (receiver LongArray) Type() byte {
	return TagLongArr
}
