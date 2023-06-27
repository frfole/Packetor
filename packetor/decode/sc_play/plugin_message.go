package sc_play

import "Packetor/packetor/decode"

type PluginMessage struct {
	Channel string
	Data    []byte
}

func (p PluginMessage) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	channel, err := reader.ReadIdentifier()
	if err != nil {
		return nil, err
	}
	data, err := reader.ReadBytes(1048576)
	if err != nil {
		return nil, err
	}
	return PluginMessage{
		Channel: channel,
		Data:    data,
	}, nil
}

func (p PluginMessage) IsValid() (reason error) {
	return nil
}
