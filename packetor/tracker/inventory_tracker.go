package tracker

import (
	"Packetor/packetor/decode"
	"Packetor/packetor/decode/cs_play"
	"Packetor/packetor/decode/sc_play"
	"Packetor/packetor/registries"
	"fmt"
)

// InventoryTracker keeps track inventory usage by the player
type InventoryTracker struct {
	HasOpenInventory   bool
	WindowID           int32
	IsSecondaryOpen    bool
	IsSecondaryHorse   bool
	SecondaryWindow    *registries.Window
	SecondarySlotCount int32
}

// UpdateInventory handle packets related to inventory
func (receiver InventoryTracker) UpdateInventory(packet decode.Packet, _ decode.PacketContext) error {
	if openScreen, ok := packet.(sc_play.OpenScreen); ok {
		receiver.HasOpenInventory = false
		receiver.IsSecondaryOpen = true
		receiver.IsSecondaryHorse = false
		receiver.SecondaryWindow = registries.GetRegistry().Windows().GetByRegistryID(openScreen.WindowType)
		receiver.WindowID = openScreen.WindowID
		receiver.SecondarySlotCount = receiver.SecondaryWindow.GetSize()
	} else if openScreen, ok := packet.(sc_play.OpenHorseScreen); ok {
		receiver.HasOpenInventory = false
		receiver.IsSecondaryOpen = true
		receiver.IsSecondaryHorse = true
		receiver.SecondaryWindow = nil
		receiver.WindowID = int32(openScreen.WindowID)
		receiver.SecondarySlotCount = openScreen.SlotCount
	} else if _, ok := packet.(sc_play.CloseContainer); ok {
		receiver.HasOpenInventory = false
		receiver.IsSecondaryOpen = false
		receiver.IsSecondaryHorse = false
		receiver.SecondaryWindow = nil
		receiver.WindowID = -1
		receiver.SecondarySlotCount = 0
	} else if _, ok := packet.(cs_play.CloseContainer); ok {
		receiver.HasOpenInventory = false
		receiver.IsSecondaryOpen = false
		receiver.IsSecondaryHorse = false
		receiver.SecondaryWindow = nil
		receiver.WindowID = -1
		receiver.SecondarySlotCount = 0
	} else if clickPacket, ok := packet.(cs_play.ClickContainer); ok {
		receiver.HasOpenInventory = true
		if receiver.WindowID != int32(clickPacket.WindowID) {
			return fmt.Errorf("client tried to edit wrong window id")
		}
	}
	return nil
}
