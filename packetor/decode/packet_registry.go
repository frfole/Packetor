package decode

import (
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

type PacketEntry struct {
	Decode func(reader PacketReader) (packet Packet, err error)
	Handle []func(packet Packet) (err error)
}

type PacketRegistry struct {
	Packets map[byte]map[int32]PacketEntry
}

func (r PacketRegistry) HandleNewPacket(state byte, reader PacketReader) (err error) {
	packetId, err := reader.ReadVarInt()
	if err != nil {
		return err
	}
	packets, ok := r.Packets[state]
	if !ok {
		return errors.Join(fmt.Errorf("unknown net state %d", state), error2.ErrSoft, error2.ErrUnknownNetState)
	}
	entry, ok := packets[packetId]
	if !ok {
		return errors.Join(fmt.Errorf("unknown packet id %x/%d", packetId, state), error2.ErrSoft, error2.ErrUnknownPacket)
	}
	packet, err := entry.Decode(reader)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to decode packet %x/%d", packetId, state), error2.ErrPacketDecode, err)
	} else if reason := packet.IsValid(); reason != nil {
		return errors.Join(fmt.Errorf("invalid packet %x/%d: %w", packetId, state, reason), error2.ErrSoft, error2.ErrPacketInvalid)
	}
	if entry.Handle != nil {
		for i := range entry.Handle {
			err = entry.Handle[i](packet)
			if err != nil && errors.Is(err, error2.ErrSoft) {
				logrus.Errorf("soft failed to handle packet %x/%d using handle index %d", packetId, state, i)
			} else if err != nil {
				return errors.Join(fmt.Errorf("failed to handle packet %x/%d using handle index %d", packetId, state, i), error2.ErrPacketHandle, err)
			}
		}
	}
	return nil
}
