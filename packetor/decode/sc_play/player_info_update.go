package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
)

type PlayerInfoUpdateAddProperty struct {
	Name      string
	Value     string
	IsSigned  bool
	Signature string
}

type PlayerInfoUpdateAdd struct {
	Name       string
	Properties []PlayerInfoUpdateAddProperty
}

type PlayerInfoUpdateInit struct {
	HasSignatureData bool
	ChatSessionID    uuid.UUID
	PubKeyExpiryTime int64
	EncodedPubKey    []byte
	PubKeySig        []byte
}

type PlayerInfoUpdateGameMode struct {
	GameMode int32
}

type PlayerInfoUpdateListed struct {
	Listed bool
}

type PlayerInfoUpdatePing struct {
	Ping int32
}

type PlayerInfoUpdateName struct {
	HasDisplayName bool
	DisplayName    string
}

type PlayerInfoUpdateActions struct {
	Uuid        uuid.UUID
	AddPlayer   PlayerInfoUpdateAdd
	InitChat    PlayerInfoUpdateInit
	GameMode    PlayerInfoUpdateGameMode
	Listed      PlayerInfoUpdateListed
	Ping        PlayerInfoUpdatePing
	DisplayName PlayerInfoUpdateName
}

type PlayerInfoUpdate struct {
	Action  int8
	Actions []PlayerInfoUpdateActions
}

func (p PlayerInfoUpdate) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	actionTypes, err := reader.ReadSByte()
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	actions := make([]PlayerInfoUpdateActions, count)
	for i := int32(0); i < count; i++ {
		actions[i] = PlayerInfoUpdateActions{}
		actions[i].Uuid, err = reader.ReadUuid()
		if err != nil {
			return nil, err
		}
		if (actionTypes & 0x01) == 0x01 {
			name, err := reader.ReadString0(16)
			if err != nil {
				return nil, err
			}
			length, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			} else if length < 0 {
				return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", length), error2.ErrDecodeTooSmall)
			}
			properties := make([]PlayerInfoUpdateAddProperty, length)
			for j := int32(0); j < length; j++ {
				propName, err := reader.ReadString0(32767)
				if err != nil {
					return nil, err
				}
				propValue, err := reader.ReadString0(32767)
				if err != nil {
					return nil, err
				}
				hasSig, err := reader.ReadBoolean()
				if err != nil {
					return nil, err
				}
				propSig := ""
				if hasSig {
					propSig, err = reader.ReadString0(32767)
					if err != nil {
						return nil, err
					}
				}
				properties[j] = PlayerInfoUpdateAddProperty{
					Name:      propName,
					Value:     propValue,
					IsSigned:  hasSig,
					Signature: propSig,
				}
			}
			actions[i].AddPlayer = PlayerInfoUpdateAdd{
				Name:       name,
				Properties: properties,
			}
		}
		if (actionTypes & 0x02) == 0x02 {
			hasSig, err := reader.ReadBoolean()
			if err != nil {
				return nil, err
			}
			if hasSig {
				csid, err := reader.ReadUuid()
				if err != nil {
					return nil, err
				}
				pkExpiry, err := reader.ReadLong()
				if err != nil {
					return nil, err
				}
				pkSize, err := reader.ReadVarInt()
				if err != nil {
					return nil, err
				}
				pk, err := reader.ReadBytesExact(int(pkSize))
				if err != nil {
					return nil, err
				}
				sigSize, err := reader.ReadVarInt()
				if err != nil {
					return nil, err
				}
				pkSig, err := reader.ReadBytesExact(int(sigSize))
				if err != nil {
					return nil, err
				}
				actions[i].InitChat = PlayerInfoUpdateInit{
					HasSignatureData: hasSig,
					ChatSessionID:    csid,
					PubKeyExpiryTime: pkExpiry,
					EncodedPubKey:    pk,
					PubKeySig:        pkSig,
				}
			} else {
				actions[i].InitChat = PlayerInfoUpdateInit{HasSignatureData: hasSig}
			}
		}
		if (actionTypes & 0x04) == 0x04 {
			gm, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			actions[i].GameMode = PlayerInfoUpdateGameMode{GameMode: gm}
		}
		if (actionTypes & 0x08) == 0x08 {
			listed, err := reader.ReadBoolean()
			if err != nil {
				return nil, err
			}
			actions[i].Listed = PlayerInfoUpdateListed{Listed: listed}
		}
		if (actionTypes & 0x10) == 0x10 {
			ping, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			actions[i].Ping = PlayerInfoUpdatePing{Ping: ping}
		}
		if (actionTypes & 0x20) == 0x20 {
			hasDN, err := reader.ReadBoolean()
			if err != nil {
				return nil, err
			}
			if hasDN {
				dn, err := reader.ReadChat()
				if err != nil {
					return nil, err
				}
				actions[i].DisplayName = PlayerInfoUpdateName{
					HasDisplayName: hasDN,
					DisplayName:    dn,
				}
			} else {
				actions[i].DisplayName = PlayerInfoUpdateName{
					HasDisplayName: hasDN,
					DisplayName:    "",
				}
			}
		}
	}
	return PlayerInfoUpdate{
		Action:  actionTypes,
		Actions: actions,
	}, nil
}

func (p PlayerInfoUpdate) IsValid() (reason error) {
	for i, action := range p.Actions {
		if (p.Action & 0x04) == 0x04 {
			if action.GameMode.GameMode < 0 || 3 < action.GameMode.GameMode {
				return fmt.Errorf("gamemode must be in <0; 3> was %d at index %d", action.GameMode.GameMode, i)
			}
		}
	}
	return nil
}
