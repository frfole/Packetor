package decode

import (
	error2 "Packetor/packetor/error"
	"Packetor/packetor/nbt"
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"io"
	"math"
	"net"
	"unicode/utf8"
)

type PacketReader struct {
	zrd     io.ReadCloser
	parent  *bytes.Reader
	rd      io.Reader
	conn    net.Conn
	data    []byte
	HasComp bool
}

func NewPacketReader(conn net.Conn) PacketReader {
	return PacketReader{
		zrd:     nil,
		parent:  nil,
		rd:      nil,
		conn:    conn,
		HasComp: false,
	}
}

func NewPacketReaderBytes(raw []byte) PacketReader {
	return PacketReader{
		zrd:     nil,
		parent:  nil,
		rd:      bytes.NewReader(raw),
		conn:    nil,
		data:    raw,
		HasComp: false,
	}
}

func (r *PacketReader) ReadPacket() (raw []byte, err error) {
	length, raw, err := r.readVarIntC()
	if err != nil {
		return nil, err
	}
	data := make([]byte, length)
	_, err = io.ReadFull(r.conn, data)
	if err != nil {
		return nil, err
	}
	raw = append(raw, data...)
	r.parent = bytes.NewReader(data)
	r.rd = r.parent
	if r.HasComp {
		err = r.setupDecompressor()
		if err != nil && !errors.Is(err, error2.ErrNoCompression) {
			return nil, fmt.Errorf("%x %w", data, err)
		} else if err != nil && errors.Is(err, error2.ErrNoCompression) {
			err = nil // avoid throwing error, when packet not compressed
		} else {
			r.rd = r.zrd
		}
	}
	return raw, err
}

func (r *PacketReader) setupDecompressor() (err error) {
	n, err := r.ReadVarInt()
	if err != nil {
		return err
	} else if n == 0 {
		return error2.ErrNoCompression // avoid decompressing uncompressed packet
	}
	if r.zrd == nil {
		r.zrd, err = zlib.NewReader(r.parent)
	} else {
		err = r.zrd.(zlib.Resetter).Reset(r.parent, nil)
	}
	if err != nil {
		return fmt.Errorf("setting up zlib reader: %d %w", n, err)
	} else {
		return nil
	}
}

func (r *PacketReader) readVarIntC() (value int32, raw []byte, err error) {
	value = 0
	raw = []byte{}
	pos := 0
	curByte := make([]byte, 1)

	for {
		n, err := r.conn.Read(curByte)
		if n != 1 {
			return value, raw, errors.Join(error2.ErrDecodeReadFail, err)
		}
		raw = append(raw, curByte[0])
		value |= int32(curByte[0]&0x7f) << pos
		if (curByte[0] & 0x80) == 0 {
			break
		}
		pos += 7
		if pos >= 32 {
			return value, raw, fmt.Errorf("VarInt[%x] %w", raw, error2.ErrDecodeTooBig)
		}
	}
	return value, raw, nil
}

func (r *PacketReader) ReadVarIntRaw() (value int32, raw []byte, err error) {
	value = 0
	raw = []byte{}
	pos := 0
	curByte := make([]byte, 1)

	for {
		n, err := r.rd.Read(curByte)
		if n != 1 {
			return value, raw, errors.Join(error2.ErrDecodeReadFail, err)
		}
		raw = append(raw, curByte[0])
		value |= int32(curByte[0]&0x7f) << pos
		if (curByte[0] & 0x80) == 0 {
			break
		}
		pos += 7
		if pos >= 32 {
			return value, raw, fmt.Errorf("VarInt %w", error2.ErrDecodeTooBig)
		}
	}
	return value, raw, nil
}

func (r *PacketReader) ReadVarInt() (value int32, err error) {
	value = 0
	pos := 0
	curByte := make([]byte, 1)

	for {
		n, err := r.rd.Read(curByte)
		if n != 1 {
			return value, errors.Join(error2.ErrDecodeReadFail, err)
		}
		value |= int32(curByte[0]&0x7f) << pos
		if (curByte[0] & 0x80) == 0 {
			break
		}
		pos += 7
		if pos >= 32 {
			return value, fmt.Errorf("VarInt %w", error2.ErrDecodeTooBig)
		}
	}
	return value, nil
}

func (r *PacketReader) ReadVarLong() (value int64, err error) {
	value = 0
	pos := 0
	curByte := make([]byte, 1)

	for {
		n, err := r.rd.Read(curByte)
		if n != 1 {
			return value, errors.Join(error2.ErrDecodeReadFail, err)
		}
		value |= int64(curByte[0]&0x7f) << pos
		if (curByte[0] & 0x80) == 0 {
			break
		}
		pos += 7
		if pos >= 64 {
			return value, fmt.Errorf("VarLong %w", error2.ErrDecodeTooBig)
		}
	}
	return value, nil
}

func (r *PacketReader) ReadBoolean() (value bool, err error) {
	data := make([]byte, 1)
	n, err := r.rd.Read(data)
	if err != nil {
		return false, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != 1 {
		return false, errors.Join(fmt.Errorf("boolean length mismatch (excepted %d was %d)", 1, n), error2.ErrDecodeLength)
	} else {
		return data[0] == 1, nil
	}
}

func (r *PacketReader) ReadSByte() (value int8, err error) {
	data := make([]byte, 1)
	n, err := r.rd.Read(data)
	if err != nil {
		return 0, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != 1 {
		return 0, errors.Join(fmt.Errorf("byte length mismatch (excepted %d was %d)", 1, n), error2.ErrDecodeLength)
	} else {
		return int8(data[0]), nil
	}
}

func (r *PacketReader) ReadUByte() (value uint8, err error) {
	data := make([]byte, 1)
	n, err := r.rd.Read(data)
	if err != nil {
		return 0, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != 1 {
		return 0, errors.Join(fmt.Errorf("ubyte length mismatch (excepted %d was %d)", 1, n), error2.ErrDecodeLength)
	} else {
		return data[0], nil
	}
}

func (r *PacketReader) ReadShort() (value int16, err error) {
	data := make([]byte, 2)
	n, err := r.rd.Read(data)
	if err != nil {
		return 0, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != 2 {
		return 0, errors.Join(fmt.Errorf("short length mismatch (excepted %d was %d)", 2, n), error2.ErrDecodeLength)
	} else {
		return (int16(data[0]) << 8) | int16(data[1]), nil
	}
}

func (r *PacketReader) ReadUShort() (value uint16, err error) {
	data := make([]byte, 2)
	n, err := r.rd.Read(data)
	if err != nil {
		return 0, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != 2 {
		return 0, errors.Join(fmt.Errorf("ushort length mismatch (excepted %d was %d)", 2, n), error2.ErrDecodeLength)
	} else {
		return (uint16(data[0]) << 8) | uint16(data[1]), nil
	}
}

func (r *PacketReader) ReadInt() (value int32, err error) {
	b := make([]byte, 4)
	n, err := r.rd.Read(b)
	if err != nil {
		return 0, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != 4 {
		return 0, errors.Join(fmt.Errorf("int length mismatch (excepted %d was %d)", 4, n), error2.ErrDecodeLength)
	} else {
		return int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3]), nil
	}
}

func (r *PacketReader) ReadLong() (value int64, err error) {
	b := make([]byte, 8)
	n, err := r.rd.Read(b)
	if err != nil {
		return 0, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != 8 {
		return 0, errors.Join(fmt.Errorf("long length mismatch (excepted %d was %d)", 8, n), error2.ErrDecodeLength)
	} else {
		return int64(b[7]) | int64(b[6])<<8 | int64(b[5])<<16 | int64(b[4])<<24 |
			int64(b[3])<<32 | int64(b[2])<<40 | int64(b[1])<<48 | int64(b[0])<<56, nil
	}
}

func (r *PacketReader) ReadULong() (value uint64, err error) {
	b := make([]byte, 8)
	n, err := r.rd.Read(b)
	if err != nil {
		return 0, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != 8 {
		return 0, errors.Join(fmt.Errorf("ulong length mismatch (excepted %d was %d)", 8, n), error2.ErrDecodeLength)
	} else {
		return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
			uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56, nil
	}
}

func (r *PacketReader) ReadFloat() (value float32, err error) {
	b := make([]byte, 4)
	n, err := r.rd.Read(b)
	if err != nil {
		return 0, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != 4 {
		return 0, errors.Join(fmt.Errorf("encoded float length mismatch (excepted %d was %d)", 4, n), error2.ErrDecodeLength)
	} else {
		return math.Float32frombits(uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[0])), nil
	}
}

func (r *PacketReader) ReadDouble() (value float64, err error) {
	b := make([]byte, 8)
	n, err := r.rd.Read(b)
	if err != nil {
		return 0, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != 8 {
		return 0, errors.Join(fmt.Errorf("encoded float length mismatch (excepted %d was %d)", 4, n), error2.ErrDecodeLength)
	} else {
		return math.Float64frombits(
			uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 | uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7]),
		), nil
	}
}

func (r *PacketReader) ReadString0(maxLength int) (value string, err error) {
	length, err := r.ReadVarInt()
	if err != nil {
		return "", errors.Join(error2.ErrDecodeReadFail, err)
	} else if length < 0 {
		return "", errors.Join(fmt.Errorf("string length must be atleast 0, was %d", length), error2.ErrDecodeFormat)
	} else if int(length) > (maxLength * 4) {
		return "", errors.Join(fmt.Errorf("encoded string length must be atmost %d was %d", maxLength*4, length), error2.ErrDecodeLength)
	}
	if length == 0 {
		return "", nil
	}
	data := make([]byte, length)
	n, err := r.rd.Read(data)
	if err != nil {
		return "", errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != int(length) {
		return "", errors.Join(fmt.Errorf("encoded string length mismatch (excepted %d was %d)", length, n), error2.ErrDecodeFormat)
	}
	str := string(data)
	if utf8.RuneCountInString(str) > maxLength {
		return "", errors.Join(fmt.Errorf("string length must be atmost %d was %d", maxLength, length), error2.ErrDecodeLength)
	}
	return str, nil
}

func (r *PacketReader) ReadString() (value string, err error) {
	return r.ReadString0(32767)
}

func (r *PacketReader) ReadChat() (value string, err error) {
	return r.ReadString0(262144)
}

func (r *PacketReader) ReadIdentifier() (value string, err error) {
	return r.ReadString0(32767)
}

func (r *PacketReader) ReadSlot() (value Slot, err error) {
	value = Slot{}
	value.Present, err = r.ReadBoolean()
	if err != nil {
		return value, fmt.Errorf("slot (present) decode failed: %w", err)
	} else if !value.Present {
		return value, nil
	} else {
		value.ItemID, err = r.ReadVarInt()
		if err != nil {
			return value, fmt.Errorf("slot (item id) decode failed: %w", err)
		}
		value.ItemCount, err = r.ReadUByte()
		if err != nil {
			return value, fmt.Errorf("slot (item count) decode failed: %w", err)
		}
		value.ItemNbt, err = r.ReadNbt()
		if err != nil {
			return value, fmt.Errorf("slot (item nbt) decode failed: %w", err)
		}
		return value, nil
	}
}

func (r *PacketReader) ReadNbt() (value nbt.Compound, err error) {
	compound, err := nbt.ReadNbt(r.rd)
	if err != nil {
		return nil, fmt.Errorf("NBT decode failed: %w", err)
	} else {
		return compound, nil
	}
}

func (r *PacketReader) ReadPosition() (value Position, err error) {
	long, err := r.ReadLong()
	if err != nil {
		return 0, errors.Join(fmt.Errorf("decode position"), err)
	}
	return Position(long), nil
}

func (r *PacketReader) ReadAngle() (value Angle, err error) {
	data := make([]byte, 1)
	n, err := r.rd.Read(data)
	if err != nil {
		return 0, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != 1 {
		return 0, errors.Join(fmt.Errorf("angle length mismatch (excepted %d was %d)", 1, n), error2.ErrDecodeLength)
	} else {
		return Angle(data[0]), nil
	}
}

func (r *PacketReader) ReadUuid() (value uuid.UUID, err error) {
	data := make([]byte, 16)
	n, err := r.rd.Read(data)
	if err != nil {
		return uuid.Nil, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != 16 {
		return uuid.Nil, errors.Join(fmt.Errorf("uuid length mismatch (excepted %d was %d)", 16, n), error2.ErrDecodeLength)
	}
	value, err = uuid.FromBytes(data)
	if err != nil {
		return uuid.Nil, errors.Join(fmt.Errorf("uuid parse"), err)
	} else {
		return value, nil
	}
}

func (r *PacketReader) ReadBytesExact(length int) (value []byte, err error) {
	data := make([]byte, length)
	n, err := r.rd.Read(data)
	if err != nil {
		return nil, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != length {
		return nil, errors.Join(fmt.Errorf("byte array length mismatch (excepted %d was %d)", length, n), error2.ErrDecodeLength)
	} else {
		return data, nil
	}
}

func (r *PacketReader) ReadBytes(maxLen int) (value []byte, err error) {
	data := make([]byte, maxLen)
	n, err := r.rd.Read(data)
	if err != nil {
		return nil, errors.Join(error2.ErrDecodeReadFail, err)
	} else {
		return data[:n], nil
	}
}

func (r *PacketReader) ReadBitSet0(bitCount int) (value BitSet, err error) {
	if bitCount < 0 {
		return nil, errors.Join(fmt.Errorf("bit count must be atleast 0 was %d", bitCount), error2.ErrDecodeLength)
	}
	bytesCount := bitCount >> 3
	if (bitCount & 0b111) != 0 {
		bytesCount += 1
	}
	raw, err := r.ReadBytesExact(bytesCount)
	if err != nil {
		return nil, errors.Join(error2.ErrDecodeReadFail, err)
	}
	longCount := bitCount >> 6
	if (bytesCount & 0b111) != 0 {
		padding := make([]byte, 8-(bytesCount&0b111))
		raw = append(padding, raw...)
		longCount += 1
	}
	value = make([]uint64, longCount)
	for i := 0; i < longCount; i++ {
		value[i] = uint64(raw[i<<3+7])<<56 | uint64(raw[i<<3+6])<<48 | uint64(raw[i<<3+5])<<40 | uint64(raw[i<<3+4])<<32 |
			uint64(raw[i<<3+3])<<24 | uint64(raw[i<<3+2])<<16 | uint64(raw[i<<3+1])<<8 | uint64(raw[i<<3])
	}
	return value, nil
}

func (r *PacketReader) ReadBitSet() (value BitSet, err error) {
	count, err := r.ReadVarInt()
	if err != nil {
		return nil, errors.Join(error2.ErrDecodeReadFail, err)
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 was %d", count), error2.ErrDecodeLength)
	}
	value = make([]uint64, count)
	for i := int32(0); i < count; i++ {
		value[i], err = r.ReadULong()
		if err != nil {
			return nil, errors.Join(error2.ErrDecodeReadFail, err)
		}
	}
	return value, nil
}

func (r *PacketReader) ReadPaletteContainer(dimension uint8, maxBits uint8) (value PaletteContainer, err error) {
	bits, err := r.ReadUByte()
	if err != nil {
		return PaletteContainer{}, err
	}

	var palette Palette
	if bits == 0 {
		stateId, err := r.ReadVarInt()
		if err != nil {
			return PaletteContainer{}, err
		}
		palette = PaletteSingle(stateId)
	} else if bits <= maxBits {
		count, err := r.ReadVarInt()
		if err != nil {
			return PaletteContainer{}, err
		} else if count < 0 {
			return PaletteContainer{}, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
		}
		stateIds := make([]int32, count)
		for i := int32(0); i < count; i++ {
			stateIds[i], err = r.ReadVarInt()
			if err != nil {
				return PaletteContainer{}, err
			}
		}
		palette = PaletteArray(stateIds)
	} else {
		palette = PaletteDirect{}
	}

	count, err := r.ReadVarInt()
	if err != nil {
		return PaletteContainer{}, err
	} else if count < 0 {
		return PaletteContainer{}, errors.Join(fmt.Errorf("count must be atleast 0 was %d", count), error2.ErrDecodeTooSmall)
	}
	data := make([]uint64, count)
	for i := int32(0); i < count; i++ {
		data[i], err = r.ReadULong()
		if err != nil {
			return PaletteContainer{}, err
		}
	}
	return PaletteContainer{
		Dimension:    dimension,
		BitsPerEntry: bits,
		Palette:      palette,
		Data:         data,
	}, nil
}
