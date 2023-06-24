package decode

import (
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type PacketEntry struct {
	Decode func(reader PacketReader) (packet Packet, err error)
	Handle func(packet Packet) (err error)
}

type PacketRegistry struct {
	Packets map[byte]map[int32]PacketEntry
}

func (r PacketRegistry) HandleNewPacket(state byte, reader PacketReader) (err error) {
	packetId, err := reader.ReadVarInt()
	if err != nil {
		return err
	}
	println("packet", packetId)
	packets, ok := r.Packets[state]
	if !ok {
		return errors.Join(fmt.Errorf("unknown net state %d", state), error2.ErrSoft, error2.ErrUnknownNetState)
	}
	entry, ok := packets[packetId]
	if !ok {
		return errors.Join(fmt.Errorf("unknown packet id %d/%d", packetId, state), error2.ErrSoft, error2.ErrUnknownPacket)
	}
	packet, err := entry.Decode(reader)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to decode packet %d/%d", packetId, state), error2.ErrPacketDecode, err)
	} else if reason := packet.IsValid(); reason != nil {
		return errors.Join(fmt.Errorf("invalid packet %d/%d: %w", packetId, state, reason), error2.ErrSoft, error2.ErrPacketInvalid)
	}
	if entry.Handle != nil {
		if err = entry.Handle(packet); err != nil {
			return errors.Join(fmt.Errorf("failed to handle packet %d/%d", packetId, state), error2.ErrPacketHandle, err)
		}
	}
	return nil
}
