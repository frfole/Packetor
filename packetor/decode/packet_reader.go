package decode

import (
	error2 "Packetor/packetor/error"
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"io"
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

func (r *PacketReader) ReadString(maxLength int) (value string, err error) {
	length, err := r.ReadVarInt()
	if err != nil {
		return "", errors.Join(error2.ErrDecodeReadFail, err)
	} else if length < 0 {
		return "", errors.Join(fmt.Errorf("string length must be atleast 0, was %d", length), error2.ErrDecodeFormat)
	} else if int(length) > (maxLength * 4) {
		return "", errors.Join(fmt.Errorf("encoded string length must be atmost %d was %d", maxLength*4, length), error2.ErrDecodeLength)
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

func (r *PacketReader) ReadUuid() (value uuid.UUID, err error) {
	data := make([]byte, 16)
	n, err := r.rd.Read(data)
	if err != nil {
		return uuid.Nil, errors.Join(error2.ErrDecodeReadFail, err)
	} else if n != 16 {
		return uuid.Nil, errors.Join(fmt.Errorf("uuid length mismatch"), error2.ErrDecodeLength)
	}
	value, err = uuid.FromBytes(data)
	if err != nil {
		return uuid.Nil, errors.Join(fmt.Errorf("uuid parse"), err)
	} else {
		return value, nil
	}
}
