package packetor

import (
	"Packetor/packetor/decode"
	"Packetor/packetor/decode/cs_handshake"
	"Packetor/packetor/decode/cs_login"
	"Packetor/packetor/decode/cs_play"
	"Packetor/packetor/decode/cs_status"
	"Packetor/packetor/decode/sc_login"
	"Packetor/packetor/decode/sc_play"
	"Packetor/packetor/decode/sc_status"
	error2 "Packetor/packetor/error"
	"Packetor/packetor/tracker"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"reflect"
	"runtime/trace"
	"time"
)

// Route handles traffic through Packetor
type Route struct {
	fronter  *Fronter
	fConn    net.Conn
	bConn    net.Conn
	errCh    chan error
	state    byte
	cReg     decode.PacketRegistry
	sReg     decode.PacketRegistry
	compress bool
	tracker  tracker.Tracker
}

func (r *Route) Start() {
	r.tracker = tracker.NewTracker()
	r.cReg = decode.PacketRegistry{
		Packets: map[byte]map[int32]decode.PacketEntry{
			0: {
				0x00: decode.PacketEntry{Decode: cs_handshake.Handshake{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.handleHandshakeC,
				}},
			},
			1: {
				0x00: decode.PacketEntry{Decode: cs_status.StatusRequest{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.PacketTracker.LimitCount(1),
				}},
				0x01: decode.PacketEntry{Decode: cs_status.PingRequest{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.PacketTracker.LimitCount(1),
				}},
			},
			2: {
				0x00: decode.PacketEntry{Decode: cs_login.LoginStart{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.PacketTracker.LimitCount(1),
					r.tracker.PacketTracker.EnsureOrder([]reflect.Type{}, []reflect.Type{reflect.TypeOf(cs_login.EncryptionResponse{})}),
				}},
				0x01: decode.PacketEntry{Decode: cs_login.EncryptionResponse{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.PacketTracker.LimitCount(1),
					r.tracker.PacketTracker.EnsureOrder([]reflect.Type{}, []reflect.Type{reflect.TypeOf(cs_login.LoginStart{})}),
				}},
				0x02: decode.PacketEntry{Decode: cs_login.LoginPluginResponse{}.Read},
			},
			3: {
				0x00: decode.PacketEntry{Decode: cs_play.ConfirmTeleportation{}.Read},
				0x01: decode.PacketEntry{Decode: cs_play.QueryBlockEntityTag{}.Read},
				0x02: decode.PacketEntry{Decode: cs_play.ChangeDifficulty{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.printPacket,
				}},
				0x03: decode.PacketEntry{Decode: cs_play.MessageAck{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.printPacket,
				}},
				0x04: decode.PacketEntry{Decode: cs_play.ChatCommand{}.Read},
				0x05: decode.PacketEntry{Decode: cs_play.ChatMessage{}.Read},
				0x06: decode.PacketEntry{Decode: cs_play.PlayerSession{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.printPacket,
				}},
				0x07: decode.PacketEntry{Decode: cs_play.ClientCommand{}.Read},
				0x08: decode.PacketEntry{Decode: cs_play.ClientInformation{}.Read},
				0x09: decode.PacketEntry{Decode: cs_play.CommandSuggestionsRequest{}.Read},
				0x0a: decode.PacketEntry{Decode: cs_play.ClickContainerButton{}.Read},
				0x0b: decode.PacketEntry{Decode: cs_play.ClickContainer{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.InventoryTracker.UpdateInventory,
				}},
				0x0c: decode.PacketEntry{Decode: cs_play.CloseContainer{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.InventoryTracker.UpdateInventory,
				}},
				0x0d: decode.PacketEntry{Decode: cs_play.PluginMessage{}.Read},
				0x0e: decode.PacketEntry{Decode: cs_play.EditBook{}.Read},
				0x0f: decode.PacketEntry{Decode: cs_play.QueryEntityTag{}.Read},
				0x10: decode.PacketEntry{Decode: cs_play.Interact{}.Read},
				0x11: decode.PacketEntry{Decode: cs_play.JigsawGenerate{}.Read},
				0x12: decode.PacketEntry{Decode: cs_play.KeepAlive{}.Read},
				0x13: decode.PacketEntry{Decode: cs_play.LockDifficulty{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.printPacket,
				}},
				0x14: decode.PacketEntry{Decode: cs_play.SetPlayerPosition{}.Read},
				0x15: decode.PacketEntry{Decode: cs_play.SetPlayerPositionRotation{}.Read},
				0x16: decode.PacketEntry{Decode: cs_play.SetPlayerRotation{}.Read},
				0x17: decode.PacketEntry{Decode: cs_play.SetPlayerOnGround{}.Read},
				0x18: decode.PacketEntry{Decode: cs_play.MoveVehicle{}.Read},
				0x19: decode.PacketEntry{Decode: cs_play.PaddleBoat{}.Read},
				0x1a: decode.PacketEntry{Decode: cs_play.PickItem{}.Read},
				0x1b: decode.PacketEntry{Decode: cs_play.PlaceRecipe{}.Read},
				0x1c: decode.PacketEntry{Decode: cs_play.PlayerAbilities{}.Read},
				0x1d: decode.PacketEntry{Decode: cs_play.PlayerAction{}.Read},
				0x1e: decode.PacketEntry{Decode: cs_play.PlayerCommand{}.Read},
				0x1f: decode.PacketEntry{Decode: cs_play.PlayerInput{}.Read},
				0x20: decode.PacketEntry{Decode: cs_play.Pong{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.printPacket,
				}},
				0x21: decode.PacketEntry{Decode: cs_play.ChangeRecipeBookSettings{}.Read},
				0x22: decode.PacketEntry{Decode: cs_play.SetSeenRecipe{}.Read},
				0x23: decode.PacketEntry{Decode: cs_play.RenameItem{}.Read},
				0x24: decode.PacketEntry{Decode: cs_play.ResourcePack{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.printPacket,
				}},
				0x25: decode.PacketEntry{Decode: cs_play.SeenAdvancements{}.Read},
				0x26: decode.PacketEntry{Decode: cs_play.SelectTrade{}.Read},
				0x27: decode.PacketEntry{Decode: cs_play.SetBeaconEffect{}.Read},
				0x28: decode.PacketEntry{Decode: cs_play.SetHeldItem{}.Read},
				0x29: decode.PacketEntry{Decode: cs_play.ProgramCommandBlock{}.Read},
				0x2a: decode.PacketEntry{Decode: cs_play.ProgramCommandBlockMinecart{}.Read},
				0x2b: decode.PacketEntry{Decode: cs_play.SetCreativeModeSlot{}.Read},
				0x2c: decode.PacketEntry{Decode: cs_play.ProgramJigsawBlock{}.Read},
				0x2d: decode.PacketEntry{Decode: cs_play.ProgramStructureBlock{}.Read},
				0x2e: decode.PacketEntry{Decode: cs_play.UpdateSign{}.Read},
				0x2f: decode.PacketEntry{Decode: cs_play.ArmSwing{}.Read},
				0x30: decode.PacketEntry{Decode: cs_play.TeleportToEntity{}.Read},
				0x31: decode.PacketEntry{Decode: cs_play.UseItemOn{}.Read},
				0x32: decode.PacketEntry{Decode: cs_play.UseItem{}.Read},
			},
		},
	}
	r.sReg = decode.PacketRegistry{
		Packets: map[byte]map[int32]decode.PacketEntry{
			0: {},
			1: {
				0x00: decode.PacketEntry{Decode: sc_status.StatusResponse{}.Read},
				0x01: decode.PacketEntry{Decode: sc_status.PingResponse{}.Read},
			},
			2: {
				0x00: decode.PacketEntry{Decode: sc_login.Disconnect{}.Read},
				0x01: decode.PacketEntry{Decode: sc_login.EncryptionRequest{}.Read},
				0x02: decode.PacketEntry{Decode: sc_login.LoginSuccess{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.handleLoginSuccess,
				}},
				0x03: decode.PacketEntry{Decode: sc_login.SetCompression(0).Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.handleSetCompression,
				}},
				0x04: decode.PacketEntry{Decode: sc_login.LoginPluginRequest{}.Read},
			},
			3: {
				0x00: decode.PacketEntry{Decode: sc_play.BundleDelimiter{}.Read},
				0x01: decode.PacketEntry{Decode: sc_play.SpawnEntity{}.Read},
				0x02: decode.PacketEntry{Decode: sc_play.SpawnExperienceOrb{}.Read},
				0x03: decode.PacketEntry{Decode: sc_play.SpawnPlayer{}.Read},
				0x04: decode.PacketEntry{Decode: sc_play.EntityAnimation{}.Read},
				0x05: decode.PacketEntry{Decode: sc_play.AwardStatistics{}.Read},
				0x06: decode.PacketEntry{Decode: sc_play.AcknowledgeBlockChange{}.Read},
				0x07: decode.PacketEntry{Decode: sc_play.SetBlockDestroyStage{}.Read},
				0x08: decode.PacketEntry{Decode: sc_play.BlockEntityData{}.Read},
				0x09: decode.PacketEntry{Decode: sc_play.BlockAction{}.Read},
				0x0a: decode.PacketEntry{Decode: sc_play.BlockUpdate{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.WorldTracker.UpdateChunk,
				}},
				0x0b: decode.PacketEntry{Decode: sc_play.BossBar{}.Read},
				0x0c: decode.PacketEntry{Decode: sc_play.ChangeDifficulty{}.Read},
				0x0d: decode.PacketEntry{Decode: sc_play.ChunkBiomes{}.Read},
				0x0e: decode.PacketEntry{Decode: sc_play.ClearTitles{}.Read},
				0x0f: decode.PacketEntry{Decode: sc_play.CommandSuggestionsResponse{}.Read},
				0x10: decode.PacketEntry{Decode: sc_play.Commands{}.Read},
				0x11: decode.PacketEntry{Decode: sc_play.CloseContainer{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.InventoryTracker.UpdateInventory,
				}},
				0x12: decode.PacketEntry{Decode: sc_play.SetContainerContent{}.Read},
				0x13: decode.PacketEntry{Decode: sc_play.SetContainerProperty{}.Read},
				0x14: decode.PacketEntry{Decode: sc_play.SetContainerSlot{}.Read},
				0x15: decode.PacketEntry{Decode: sc_play.SetCooldown{}.Read},
				0x16: decode.PacketEntry{Decode: sc_play.ChatSuggestions{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.printPacket,
				}},
				0x17: decode.PacketEntry{Decode: sc_play.PluginMessage{}.Read},
				0x18: decode.PacketEntry{Decode: sc_play.DamageEvent{}.Read},
				0x19: decode.PacketEntry{Decode: sc_play.DeleteMessage{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.printPacket,
				}},
				0x1a: decode.PacketEntry{Decode: sc_play.Disconnect{}.Read},
				0x1b: decode.PacketEntry{Decode: sc_play.DisguisedChatMessage{}.Read},
				0x1c: decode.PacketEntry{Decode: sc_play.EntityEvent{}.Read},
				0x1d: decode.PacketEntry{Decode: sc_play.Explosion{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.WorldTracker.UpdateChunk,
				}},
				0x1e: decode.PacketEntry{Decode: sc_play.UnloadChunk{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.WorldTracker.UpdateChunk,
				}},
				0x1f: decode.PacketEntry{Decode: sc_play.GameEvent{}.Read},
				0x20: decode.PacketEntry{Decode: sc_play.OpenHorseScreen{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.InventoryTracker.UpdateInventory,
				}},
				0x21: decode.PacketEntry{Decode: sc_play.HurtAnimation{}.Read},
				0x22: decode.PacketEntry{Decode: sc_play.InitializeWorldBorder{}.Read},
				0x23: decode.PacketEntry{Decode: sc_play.KeepAlive{}.Read},
				0x24: decode.PacketEntry{Decode: sc_play.ChunkData{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.WorldTracker.UpdateChunk,
				}},
				0x25: decode.PacketEntry{Decode: sc_play.WorldEvent{}.Read},
				0x26: decode.PacketEntry{Decode: sc_play.Particle{}.Read},
				0x27: decode.PacketEntry{Decode: sc_play.UpdateLight{}.Read},
				0x28: decode.PacketEntry{Decode: sc_play.Login{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					func(packet decode.Packet, _ decode.PacketContext) (err error) {
						return r.tracker.OnLogin(packet.(sc_play.Login))
					},
				}},
				0x29: decode.PacketEntry{Decode: sc_play.MapData{}.Read},
				0x2a: decode.PacketEntry{Decode: sc_play.MerchantOffers{}.Read},
				0x2b: decode.PacketEntry{Decode: sc_play.UpdateEntityPosition{}.Read},
				0x2c: decode.PacketEntry{Decode: sc_play.UpdateEntityPositionRotation{}.Read},
				0x2d: decode.PacketEntry{Decode: sc_play.UpdateEntityRotation{}.Read},
				0x2e: decode.PacketEntry{Decode: sc_play.MoveVehicle{}.Read},
				0x2f: decode.PacketEntry{Decode: sc_play.OpenBook{}.Read},
				0x30: decode.PacketEntry{Decode: sc_play.OpenScreen{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.InventoryTracker.UpdateInventory,
				}},
				0x31: decode.PacketEntry{Decode: sc_play.OpenSignEditor{}.Read},
				0x32: decode.PacketEntry{Decode: sc_play.Ping{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.printPacket,
				}},
				0x33: decode.PacketEntry{Decode: sc_play.PlaceGhostRecipe{}.Read},
				0x34: decode.PacketEntry{Decode: sc_play.PlayerAbilities{}.Read},
				0x35: decode.PacketEntry{Decode: sc_play.PlayerChatMessage{}.Read},
				0x36: decode.PacketEntry{Decode: sc_play.EndCombat{}.Read},
				0x37: decode.PacketEntry{Decode: sc_play.EnterCombat{}.Read},
				0x38: decode.PacketEntry{Decode: sc_play.CombatDeath{}.Read},
				0x39: decode.PacketEntry{Decode: sc_play.PlayerInfoRemove{}.Read},
				0x3a: decode.PacketEntry{Decode: sc_play.PlayerInfoUpdate{}.Read},
				0x3b: decode.PacketEntry{Decode: sc_play.LookAt{}.Read},
				0x3c: decode.PacketEntry{Decode: sc_play.SynchronizePlayerPosition{}.Read},
				0x3d: decode.PacketEntry{Decode: sc_play.UpdateRecipeBook{}.Read},
				0x3e: decode.PacketEntry{Decode: sc_play.RemoveEntities{}.Read},
				0x3f: decode.PacketEntry{Decode: sc_play.RemoveEntityEffect{}.Read},
				0x40: decode.PacketEntry{Decode: sc_play.ResourcePack{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.printPacket,
				}},
				0x41: decode.PacketEntry{Decode: sc_play.Respawn{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.OnRespawn,
				}},
				0x42: decode.PacketEntry{Decode: sc_play.SetHeadRotation{}.Read},
				0x43: decode.PacketEntry{Decode: sc_play.UpdateSectionBlocks{}.Read, Handle: []func(packet decode.Packet, context decode.PacketContext) (err error){
					r.tracker.WorldTracker.UpdateChunk,
				}},
				0x44: decode.PacketEntry{Decode: sc_play.SelectAdvancementsTab{}.Read},
				0x45: decode.PacketEntry{Decode: sc_play.ServerData{}.Read},
				0x46: decode.PacketEntry{Decode: sc_play.SetActionBarText{}.Read},
				0x47: decode.PacketEntry{Decode: sc_play.SetBorderCenter{}.Read},
				0x48: decode.PacketEntry{Decode: sc_play.SetBorderLerpSize{}.Read},
				0x49: decode.PacketEntry{Decode: sc_play.SetBorderSize{}.Read},
				0x4a: decode.PacketEntry{Decode: sc_play.SetBorderWarningDelay{}.Read},
				0x4b: decode.PacketEntry{Decode: sc_play.SetBorderWarningDistance{}.Read},
				0x4c: decode.PacketEntry{Decode: sc_play.SetCamera{}.Read},
				0x4d: decode.PacketEntry{Decode: sc_play.SetHeldItem{}.Read},
				0x4e: decode.PacketEntry{Decode: sc_play.SetCenterChunk{}.Read},
				0x4f: decode.PacketEntry{Decode: sc_play.SetRenderDistance{}.Read},
				0x50: decode.PacketEntry{Decode: sc_play.SetDefaultSpawnPosition{}.Read},
				0x51: decode.PacketEntry{Decode: sc_play.DisplayObjective{}.Read},
				0x52: decode.PacketEntry{Decode: sc_play.SetEntityMetadata{}.Read},
				0x53: decode.PacketEntry{Decode: sc_play.LinkEntities{}.Read},
				0x54: decode.PacketEntry{Decode: sc_play.SetEntityVelocity{}.Read},
				0x55: decode.PacketEntry{Decode: sc_play.SetEquipment{}.Read},
				0x56: decode.PacketEntry{Decode: sc_play.SetExperience{}.Read},
				0x57: decode.PacketEntry{Decode: sc_play.SetHealth{}.Read},
				0x58: decode.PacketEntry{Decode: sc_play.UpdateObjectives{}.Read},
				0x59: decode.PacketEntry{Decode: sc_play.SetPassengers{}.Read},
				0x5a: decode.PacketEntry{Decode: sc_play.UpdateTeams{}.Read},
				0x5b: decode.PacketEntry{Decode: sc_play.UpdateScore{}.Read},
				0x5c: decode.PacketEntry{Decode: sc_play.SetSimulationDistance{}.Read},
				0x5d: decode.PacketEntry{Decode: sc_play.SetSubtitleText{}.Read},
				0x5e: decode.PacketEntry{Decode: sc_play.UpdateTime{}.Read},
				0x5f: decode.PacketEntry{Decode: sc_play.SetTitleText{}.Read},
				0x60: decode.PacketEntry{Decode: sc_play.SetTitleAnimationTimes{}.Read},
				0x61: decode.PacketEntry{Decode: sc_play.EntitySoundEffect{}.Read},
				0x62: decode.PacketEntry{Decode: sc_play.SoundEffect{}.Read},
				0x63: decode.PacketEntry{Decode: sc_play.StopSound{}.Read},
				0x64: decode.PacketEntry{Decode: sc_play.SystemChatMessage{}.Read},
				0x65: decode.PacketEntry{Decode: sc_play.SetTabListHeaderFooter{}.Read},
				0x66: decode.PacketEntry{Decode: sc_play.TagQueryResponse{}.Read},
				0x67: decode.PacketEntry{Decode: sc_play.PickupItem{}.Read},
				0x68: decode.PacketEntry{Decode: sc_play.TeleportEntity{}.Read},
				0x69: decode.PacketEntry{Decode: sc_play.UpdateAdvancements{}.Read},
				0x6a: decode.PacketEntry{Decode: sc_play.UpdateAttributes{}.Read},
				0x6b: decode.PacketEntry{Decode: sc_play.FeatureFlags{}.Read},
				0x6c: decode.PacketEntry{Decode: sc_play.EntityEffect{}.Read},
				0x6d: decode.PacketEntry{Decode: sc_play.UpdateRecipes{}.Read},
				0x6e: decode.PacketEntry{Decode: sc_play.UpdateTags{}.Read},
			},
		},
	}
	go r.handleFBTraffic()
	go r.handleBFTraffic()
}

func (r *Route) handleFBTraffic() {
	reader := decode.NewPacketReader(r.fConn)
	for {
		if err := r.fConn.SetReadDeadline(time.Now().Add(r.fronter.timeout)); err != nil {
			r.errCh <- fmt.Errorf("client->server set read deadline: %w", err)
			return
		}
		if err := r.bConn.SetWriteDeadline(time.Now().Add(r.fronter.timeout)); err != nil {
			r.errCh <- fmt.Errorf("client->server set write deadline: %w", err)
			return
		}
		region := trace.StartRegion(context.Background(), "FB read")
		raw, err := reader.ReadPacket()
		if err != nil {
			r.errCh <- fmt.Errorf("client->server read: %w", err)
			return
		}
		if _, err = r.bConn.Write(raw); err != nil {
			r.errCh <- fmt.Errorf("client->server write: %w", err)
			return
		}
		region.End()
		region = trace.StartRegion(context.Background(), "FB handle")
		packet, packetContext, err := r.cReg.ReadNewPacket(decode.ServerBound, r.state, reader)
		if err != nil {
			if errors.Is(err, error2.ErrSoft) && errors.Is(err, error2.ErrUnknownPacket) {
				continue
			}
			if errors.Is(err, error2.ErrSoft) {
				logrus.Errorf("client->server decode: %v", err)
			} else {
				r.errCh <- fmt.Errorf("client->server decode: %w", err)
			}
		}
		r.tracker.PacketTracker.UpdateCount(packetContext, packet)
		if err = r.cReg.HandlePacket(packet, packetContext); err != nil {
			if errors.Is(err, error2.ErrSoft) && errors.Is(err, error2.ErrUnknownPacket) {
				continue
			}
			if errors.Is(err, error2.ErrSoft) {
				logrus.Errorf("client->server handle: %v", err)
			} else {
				r.errCh <- fmt.Errorf("client->server handle: %w", err)
			}
		}
		region.End()
		reader.HasComp = r.compress
	}
}

func (r *Route) handleBFTraffic() {
	reader := decode.NewPacketReader(r.bConn)
	for {
		if err := r.bConn.SetReadDeadline(time.Now().Add(r.fronter.timeout)); err != nil {
			r.errCh <- fmt.Errorf("server->client set read deadline: %w", err)
			return
		}
		if err := r.fConn.SetWriteDeadline(time.Now().Add(r.fronter.timeout)); err != nil {
			r.errCh <- fmt.Errorf("server->client set write deadline: %w", err)
			return
		}
		region := trace.StartRegion(context.Background(), "BF read")
		raw, err := reader.ReadPacket()
		if err != nil {
			r.errCh <- fmt.Errorf("server->client read: %w", err)
			return
		}
		if _, err = r.fConn.Write(raw); err != nil {
			r.errCh <- fmt.Errorf("server->client write: %w", err)
			return
		}
		region.End()
		region = trace.StartRegion(context.Background(), "BF handle")
		packet, packetContext, err := r.sReg.ReadNewPacket(decode.ClientBound, r.state, reader)
		if err != nil {
			if errors.Is(err, error2.ErrSoft) && errors.Is(err, error2.ErrUnknownPacket) {
				continue
			}
			if errors.Is(err, error2.ErrSoft) {
				logrus.Errorf("server->client decode: %v", err)
			} else {
				r.errCh <- fmt.Errorf("server->client decode: %w", err)
			}
		}
		r.tracker.PacketTracker.UpdateCount(packetContext, packet)
		if err = r.sReg.HandlePacket(packet, packetContext); err != nil {
			if errors.Is(err, error2.ErrSoft) && errors.Is(err, error2.ErrUnknownPacket) {
				continue
			}
			if errors.Is(err, error2.ErrSoft) {
				logrus.Errorf("server->client handle: %v", err)
			} else {
				r.errCh <- fmt.Errorf("server->client handle: %w", err)
			}
		}
		region.End()
		reader.HasComp = r.compress
		reader.Context.WorldHeight = r.tracker.WorldTracker.DimensionType.Height / 16
	}
}
