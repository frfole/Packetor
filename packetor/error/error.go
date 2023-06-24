package error

import "errors"

var (
	ErrSoft            = errors.New("soft error")
	ErrUnknownNetState = errors.New("unknown net state")
	ErrUnknownPacket   = errors.New("unknown packet")
	ErrPacketDecode    = errors.New("failed to decode packet")
	ErrPacketHandle    = errors.New("failed to handle packet")
	ErrPacketInvalid   = errors.New("packet format is invalid")
	ErrDecodeTooBig    = errors.New("too big")
	ErrDecodeFormat    = errors.New("invalid format")
	ErrDecodeLength    = errors.New("length mismatch")
	ErrDecodeReadFail  = errors.New("failed to read")
	ErrNoCompression   = errors.New("no compression")
)
