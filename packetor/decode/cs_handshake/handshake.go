package cs_handshake

import (
	"Packetor/packetor/decode"
	"fmt"
)

type Handshake struct {
	ProtocolVersion int32
	ServerAddr      string
	ServerPort      uint16
	NextState       int32
}

func (h Handshake) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	protVer, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	srvAddr, err := reader.ReadString0(255)
	if err != nil {
		return nil, err
	}
	srvPort, err := reader.ReadUShort()
	if err != nil {
		return nil, err
	}
	nextState, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return Handshake{
		ProtocolVersion: protVer,
		ServerAddr:      srvAddr,
		ServerPort:      srvPort,
		NextState:       nextState,
	}, err
}

func (h Handshake) IsValid() (reason error) {
	// TODO: maybe validate correct protocol version
	if h.NextState < 1 || 2 < h.NextState {
		return fmt.Errorf("next state must be in <1; 2> was %d", h.NextState)
	}
	return nil
}
