package sc_play

import "Packetor/packetor/decode"

type SetContainerProperty struct {
	WindowID uint8
	Property int16
	Value    int16
}

func (p SetContainerProperty) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	wid, err := reader.ReadUByte()
	if err != nil {
		return nil, err
	}
	property, err := reader.ReadShort()
	if err != nil {
		return nil, err
	}
	value, err := reader.ReadShort()
	if err != nil {
		return nil, err
	}
	return SetContainerProperty{
		WindowID: wid,
		Property: property,
		Value:    value,
	}, nil
}

func (p SetContainerProperty) IsValid() (reason error) {
	return nil
}
