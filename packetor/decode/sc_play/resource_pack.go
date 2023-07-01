package sc_play

import "Packetor/packetor/decode"

type ResourcePack struct {
	URL              string
	Hash             string
	Forced           bool
	HasPromptMessage bool
	PromptMessage    string
}

func (p ResourcePack) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	url, err := reader.ReadString0(32767)
	if err != nil {
		return nil, err
	}
	hash, err := reader.ReadString0(40)
	if err != nil {
		return nil, err
	}
	forced, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	hasPrompt, err := reader.ReadBoolean()
	if err != nil {
		return nil, err
	}
	prompt := ""
	if hasPrompt {
		prompt, err = reader.ReadChat()
		if err != nil {
			return nil, err
		}
	}
	return ResourcePack{
		URL:              url,
		Hash:             hash,
		Forced:           forced,
		HasPromptMessage: hasPrompt,
		PromptMessage:    prompt,
	}, nil
}

func (p ResourcePack) IsValid() (reason error) {
	return nil
}
