package cs_play

import "Packetor/packetor/decode"

type RenameItem struct {
	Name string
}

func (p RenameItem) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	name, err := reader.ReadString()
	if err != nil {
		return nil, err
	}
	return RenameItem{Name: name}, nil
}

func (p RenameItem) IsValid() (reason error) {
	return nil
}
