package decode

import (
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	// ServerBound means that packet is going from the client to the server
	ServerBound PacketDirection = true
	// ClientBound means that packet is going from the server to the client
	ClientBound PacketDirection = false
)

// PacketDirection indicates which way is the packet going
type PacketDirection bool

// PacketContext provides basic for context for packet
type PacketContext struct {
	direction PacketDirection
	state     byte
	id        int32
}

// PacketEntry provides way to read and handle read packet
type PacketEntry struct {
	// Decode is a function used to decode packet
	Decode func(reader PacketReader) (packet Packet, err error)
	// Handle is a list of functions to run with decoded packet
	Handle []func(packet Packet, context PacketContext) (err error)
}

// PacketRegistry contains mapping of state and packet id to PacketEntry used to decode new packet
type PacketRegistry struct {
	// Packet is mapping state and packet id to PacketEntry used to decode new packet
	Packets map[byte]map[int32]PacketEntry
}

// State returns the state in which packet was received
func (receiver PacketContext) State() byte {
	return receiver.state
}

// Direction returns the direction of the packet
func (receiver PacketContext) Direction() PacketDirection {
	return receiver.direction
}

// ReadNewPacket reads new packet and validates it
func (r PacketRegistry) ReadNewPacket(direction PacketDirection, state byte, reader PacketReader) (packet Packet, context PacketContext, err error) {
	// read packet id
	packetId, err := reader.ReadVarInt()
	if err != nil {
		return nil, PacketContext{direction, state, packetId}, err
	}

	// find corresponding PacketEntry, decode the packet and validate it
	packets, ok := r.Packets[state]
	if !ok {
		return nil, PacketContext{direction, state, packetId}, errors.Join(fmt.Errorf("unknown net state %d", state), error2.ErrSoft, error2.ErrUnknownNetState)
	}
	entry, ok := packets[packetId]
	if !ok {
		return nil, PacketContext{direction, state, packetId}, errors.Join(fmt.Errorf("unknown packet id %x/%d", packetId, state), error2.ErrSoft, error2.ErrUnknownPacket)
	}
	packet, err = entry.Decode(reader)
	if err != nil {
		return nil, PacketContext{direction, state, packetId}, errors.Join(fmt.Errorf("failed to decode packet %x/%d", packetId, state), error2.ErrPacketDecode, err)
	} else if reason := packet.IsValid(); reason != nil {
		return nil, PacketContext{direction, state, packetId}, errors.Join(fmt.Errorf("invalid packet %x/%d: %w", packetId, state, reason), error2.ErrSoft, error2.ErrPacketInvalid)
	}
	return packet, PacketContext{direction, state, packetId}, nil
}

// HandlePacket passes given packet to every handler associated with the packet
func (r PacketRegistry) HandlePacket(packet Packet, context PacketContext) (err error) {
	packets, ok := r.Packets[context.state]
	if !ok {
		return errors.Join(fmt.Errorf("unknown net state %d", context.state), error2.ErrSoft, error2.ErrUnknownNetState)
	}
	entry, ok := packets[context.id]
	if !ok {
		return errors.Join(fmt.Errorf("unknown packet id %x/%d", context.id, context.state), error2.ErrSoft, error2.ErrUnknownPacket)
	}
	if entry.Handle != nil {
		for i := range entry.Handle {
			err = entry.Handle[i](packet, context)
			if err != nil && errors.Is(err, error2.ErrSoft) {
				logrus.Errorf("soft failed to handle packet %x/%d using handle index %d", context.id, context.state, i)
			} else if err != nil {
				return errors.Join(fmt.Errorf("failed to handle packet %x/%d using handle index %d", context.id, context.state, i), error2.ErrPacketHandle, err)
			}
		}
	}
	return nil
}
