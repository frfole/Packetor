package packetor

import (
	"Packetor/packetor/decode"
	"Packetor/packetor/decode/cs_handshake"
	"Packetor/packetor/decode/cs_login"
	"Packetor/packetor/decode/cs_status"
	"Packetor/packetor/decode/sc_login"
	"Packetor/packetor/decode/sc_play"
	"Packetor/packetor/decode/sc_status"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
	"net"
	"time"
)

// Route handles traffic through Packetor
// TODO: check if we are processing packets in correct order?
type Route struct {
	fronter  *Fronter
	fConn    net.Conn
	bConn    net.Conn
	errCh    chan error
	state    byte
	cReg     decode.PacketRegistry
	sReg     decode.PacketRegistry
	compress bool
}

func (r *Route) Start() {
	r.cReg = decode.PacketRegistry{
		Packets: map[byte]map[int32]decode.PacketEntry{
			0: {
				0x00: decode.PacketEntry{Decode: cs_handshake.Handshake{}.Read, Handle: r.handleHandshakeC},
			},
			1: {
				0x00: decode.PacketEntry{Decode: cs_status.StatusRequest{}.Read},
				0x01: decode.PacketEntry{Decode: cs_status.PingRequest{}.Read},
			},
			2: {
				0x00: decode.PacketEntry{Decode: cs_login.LoginStart{}.Read},
				0x01: decode.PacketEntry{Decode: cs_login.EncryptionResponse{}.Read},
				0x02: decode.PacketEntry{Decode: cs_login.LoginPluginResponse{}.Read},
			},
			3: {},
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
				0x02: decode.PacketEntry{Decode: sc_login.LoginSuccess{}.Read, Handle: func(packet decode.Packet) (err error) {
					r.state = 3
					return nil
				}},
				0x03: decode.PacketEntry{Decode: sc_login.SetCompression(0).Read, Handle: func(packet decode.Packet) (err error) {
					r.compress = packet.(sc_login.SetCompression).Compression()
					return nil
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
				0x0a: decode.PacketEntry{Decode: sc_play.BlockUpdate{}.Read},
				0x0b: decode.PacketEntry{Decode: sc_play.BossBar{}.Read},
				0x0c: decode.PacketEntry{Decode: sc_play.ChangeDifficulty{}.Read},
				0x0d: decode.PacketEntry{Decode: sc_play.ChunkBiomes{}.Read},
				0x0e: decode.PacketEntry{Decode: sc_play.ClearTitles{}.Read},
				0x0f: decode.PacketEntry{Decode: sc_play.CommandSuggestionsResponse{}.Read},
				0x10: decode.PacketEntry{Decode: sc_play.Commands{}.Read},
				0x11: decode.PacketEntry{Decode: sc_play.CloseContainer{}.Read},
				0x12: decode.PacketEntry{Decode: sc_play.SetContainerContent{}.Read},
				0x13: decode.PacketEntry{Decode: sc_play.SetContainerProperty{}.Read},
				0x14: decode.PacketEntry{Decode: sc_play.SetContainerSlot{}.Read},
				0x15: decode.PacketEntry{Decode: sc_play.SetCooldown{}.Read},
				0x16: decode.PacketEntry{Decode: sc_play.ChatSuggestions{}.Read},
				0x17: decode.PacketEntry{Decode: sc_play.PluginMessage{}.Read},
				0x18: decode.PacketEntry{Decode: sc_play.DamageEvent{}.Read},
				0x19: decode.PacketEntry{Decode: sc_play.DeleteMessage{}.Read},
				0x1a: decode.PacketEntry{Decode: sc_play.Disconnect{}.Read},
				0x1b: decode.PacketEntry{Decode: sc_play.DisguisedChatMessage{}.Read},
				0x1c: decode.PacketEntry{Decode: sc_play.EntityEvent{}.Read},
				0x1d: decode.PacketEntry{Decode: sc_play.Explosion{}.Read},
				0x1e: decode.PacketEntry{Decode: sc_play.UnloadChunk{}.Read},
				0x1f: decode.PacketEntry{Decode: sc_play.GameEvent{}.Read},
				0x20: decode.PacketEntry{Decode: sc_play.OpenHorseScreen{}.Read},
				0x21: decode.PacketEntry{Decode: sc_play.HurtAnimation{}.Read},
				0x22: decode.PacketEntry{Decode: sc_play.InitializeWorldBorder{}.Read},
				0x23: decode.PacketEntry{Decode: sc_play.KeepAlive{}.Read},
				0x24: decode.PacketEntry{Decode: sc_play.ChunkData{}.Read},
				0x25: decode.PacketEntry{Decode: sc_play.WorldEvent{}.Read},
				0x26: decode.PacketEntry{Decode: sc_play.Particle{}.Read},
				0x27: decode.PacketEntry{Decode: sc_play.UpdateLight{}.Read},
				0x28: decode.PacketEntry{Decode: sc_play.Login{}.Read},
				0x29: decode.PacketEntry{Decode: sc_play.MapData{}.Read},
				0x2a: decode.PacketEntry{Decode: sc_play.MerchantOffers{}.Read},
				0x2b: decode.PacketEntry{Decode: sc_play.UpdateEntityPosition{}.Read},
				0x2c: decode.PacketEntry{Decode: sc_play.UpdateEntityPositionRotation{}.Read},
				0x2d: decode.PacketEntry{Decode: sc_play.UpdateEntityRotation{}.Read},
				0x2e: decode.PacketEntry{Decode: sc_play.MoveVehicle{}.Read},
				0x2f: decode.PacketEntry{Decode: sc_play.OpenBook{}.Read},
				0x30: decode.PacketEntry{Decode: sc_play.OpenScreen{}.Read},
				0x31: decode.PacketEntry{Decode: sc_play.OpenSignEditor{}.Read},
				0x32: decode.PacketEntry{Decode: sc_play.Ping{}.Read},
				0x33: decode.PacketEntry{Decode: sc_play.PlaceGhostRecipe{}.Read},
				0x34: decode.PacketEntry{Decode: sc_play.PlayerAbilities{}.Read},
				0x35: decode.PacketEntry{Decode: sc_play.PlayerChatMessage{}.Read},
				0x36: decode.PacketEntry{Decode: sc_play.EndCombat{}.Read},
				0x37: decode.PacketEntry{Decode: sc_play.EnterCombat{}.Read},
				0x38: decode.PacketEntry{Decode: sc_play.CombatDeath{}.Read},
				0x39: decode.PacketEntry{Decode: sc_play.PlayerInfoRemove{}.Read},
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
		raw, err := reader.ReadPacket()
		if err != nil {
			r.errCh <- fmt.Errorf("client->server read: %w", err)
			return
		}
		if _, err = r.bConn.Write(raw); err != nil {
			r.errCh <- fmt.Errorf("client->server write: %w", err)
			return
		}
		if err = r.cReg.HandleNewPacket(r.state, reader); err != nil {
			if errors.Is(err, error2.ErrSoft) && errors.Is(err, error2.ErrUnknownPacket) {
				continue
			}
			if errors.Is(err, error2.ErrSoft) {
				fmt.Printf("client->server handle: %v\n", err)
			} else {
				r.errCh <- fmt.Errorf("client->server handle: %w", err)
			}
		}
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
		raw, err := reader.ReadPacket()
		if err != nil {
			r.errCh <- fmt.Errorf("server->client read: %w", err)
			return
		}
		if _, err = r.fConn.Write(raw); err != nil {
			r.errCh <- fmt.Errorf("server->client write: %w", err)
			return
		}
		if err = r.sReg.HandleNewPacket(r.state, reader); err != nil {
			if errors.Is(err, error2.ErrSoft) && errors.Is(err, error2.ErrUnknownPacket) {
				continue
			}
			if errors.Is(err, error2.ErrSoft) {
				fmt.Printf("server->client handle: %v\n", err)
			} else {
				r.errCh <- fmt.Errorf("server->client handle: %w", err)
			}
		}
		reader.HasComp = r.compress
	}
}
