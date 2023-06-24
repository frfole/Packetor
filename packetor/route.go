package packetor

import (
	"Packetor/packetor/decode"
	"Packetor/packetor/decode/cs_handshake"
	"Packetor/packetor/decode/sc_login"
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
				0: decode.PacketEntry{Decode: cs_handshake.Handshake{}.Read, Handle: r.handleHandshakeC},
			},
			1: {},
			2: {},
			3: {},
		},
	}
	r.sReg = decode.PacketRegistry{
		Packets: map[byte]map[int32]decode.PacketEntry{
			0: {},
			1: {},
			2: {
				0x02: decode.PacketEntry{Decode: sc_login.LoginSuccess{}.Read},
				0x03: decode.PacketEntry{Decode: sc_login.SetCompression(0).Read, Handle: func(packet decode.Packet) (err error) {
					r.compress = packet.(sc_login.SetCompression).Compression()
					return nil
				}},
			},
			3: {},
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
			if errors.Is(err, error2.ErrSoft) {
				fmt.Printf("server->client handle: %v\n", err)
			} else {
				r.errCh <- fmt.Errorf("server->client handle: %w", err)
			}
		}
		reader.HasComp = r.compress
	}
}
