package tracker

import (
	"Packetor/packetor/decode"
	"fmt"
	"reflect"
)

// PacketTracker keeps track of received packet counts
type PacketTracker struct {
	// packetCounts holds information on how many times was given packet received
	packetCounts map[decode.PacketDirection]map[byte]map[reflect.Type]uint32
}

// newPacketTracker creates a new PacketTracker and initialize it partially
func newPacketTracker() PacketTracker {
	tracker := PacketTracker{
		packetCounts: map[decode.PacketDirection]map[byte]map[reflect.Type]uint32{},
	}
	tracker.packetCounts[decode.ServerBound] = map[byte]map[reflect.Type]uint32{}
	tracker.packetCounts[decode.ClientBound] = map[byte]map[reflect.Type]uint32{}
	return tracker
}

// UpdateCount increases counter which indicates how many times was the given packet received
func (tracker PacketTracker) UpdateCount(context decode.PacketContext, packet decode.Packet) {
	_, ok := tracker.packetCounts[context.Direction()][context.State()]
	if !ok {
		tracker.packetCounts[context.Direction()][context.State()] = map[reflect.Type]uint32{}
	}
	count, ok := tracker.packetCounts[context.Direction()][context.State()][reflect.TypeOf(packet)]
	if !ok {
		count = 0
	}
	count++
	tracker.packetCounts[context.Direction()][context.State()][reflect.TypeOf(packet)] = count
}

// LimitCount creates a new function limiting amount of times a certain packet can be received
func (tracker PacketTracker) LimitCount(limit uint32) func(decode.Packet, decode.PacketContext) error {
	return func(packet decode.Packet, context decode.PacketContext) (err error) {
		count, ok := tracker.packetCounts[context.Direction()][context.State()][reflect.TypeOf(packet)]
		if ok && count == limit+1 {
			return fmt.Errorf("packet %v was received more than allowed (%d)", reflect.TypeOf(packet), limit)
		}
		return nil
	}
}

// EnsureOrder creates a new function that ensures that the packet was received between given packets
func (tracker PacketTracker) EnsureOrder(after []reflect.Type, before []reflect.Type) func(decode.Packet, decode.PacketContext) error {
	return func(packet decode.Packet, context decode.PacketContext) (err error) {
		counts := tracker.packetCounts[context.Direction()][context.State()]
		for _, afterPacket := range after {
			if count, ok := counts[afterPacket]; !ok || count == 0 {
				return fmt.Errorf("packet %v was received before %v", reflect.TypeOf(packet), afterPacket)
			}
		}
		for _, beforePacket := range before {
			if count, ok := counts[beforePacket]; ok && count != 0 {
				return fmt.Errorf("packet %v was received after %v", reflect.TypeOf(packet), beforePacket)
			}
		}
		return nil
	}
}
