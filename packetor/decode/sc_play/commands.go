package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

const (
	CommandParserBBool int32 = iota
	CommandParserBFloat
	CommandParserBDouble
	CommandParserBInteger
	CommandParserBLong
	CommandParserBString
	CommandParserMEntity
	CommandParserMGameProfile
	CommandParserMBlockPos
	CommandParserMColumnPos
	CommandParserMVec3
	CommandParserMVec2
	CommandParserMBlockState
	CommandParserMBlockPredicate
	CommandParserMItemStack
	CommandParserMItemPredicate
	CommandParserMColor
	CommandParserMComponent
	CommandParserMMessage
	CommandParserMNbt
	CommandParserMNbtTag
	CommandParserMNbtPath
	CommandParserMObjective
	CommandParserMObjectiveCriteria
	CommandParserMOperation
	CommandParserMParticle
	CommandParserMAngle
	CommandParserMRotation
	CommandParserMScoreboardSlot
	CommandParserMScoreHolder
	CommandParserMSwizzle
	CommandParserMTeam
	CommandParserMItemSlot
	CommandParserMResourceLocation
	CommandParserMFunction
	CommandParserMEntityAnchor
	CommandParserMIntRange
	CommandParserMFloatRange
	CommandParserMDimension
	CommandParserMGamemode
	CommandParserMTime
	CommandParserMResourceOrTag
	CommandParserMResourceOrTagKey
	CommandParserMResource
	CommandParserMResourceKey
	CommandParserMTemplateMirror
	CommandParserMTemplateRotation
	CommandParserMHeightmap
	CommandParserMUuid
)

type CommandNode struct {
	Flags           uint8
	Children        []int32
	RedirectNode    int32
	Name            string
	ParserID        int32
	Properties      []byte
	SuggestionsType string
}

type Commands struct {
	Nodes   []*CommandNode
	RootIdx int32
}

func (p Commands) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	nodes := make([]*CommandNode, count)
	for i := int32(0); i < count; i++ {
		node := new(CommandNode)
		node.Flags, err = reader.ReadUByte()
		if err != nil {
			return nil, err
		}
		childrenC, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		} else if childrenC < 0 {
			return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", childrenC), error2.ErrDecodeTooSmall)
		}
		node.Children = make([]int32, childrenC)
		for j := int32(0); j < childrenC; j++ {
			childrenIdx, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			node.Children[j] = childrenIdx
		}
		if (node.Flags & 0x08) == 0x08 {
			node.RedirectNode, err = reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
		}
		if (node.Flags&0x03) == 1 || (node.Flags&0x03) == 2 {
			node.Name, err = reader.ReadString0(32767)
			if err != nil {
				return nil, err
			}
		}
		if (node.Flags & 0x03) == 2 {
			node.ParserID, err = reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			switch node.ParserID {
			case CommandParserBFloat:
				flags, err := reader.ReadUByte()
				if err != nil {
					return nil, err
				}
				data, err := reader.ReadBytesExact(int((flags&0b1)<<2 + (flags&0b10)<<1))
				if err != nil {
					return nil, err
				}
				node.Properties = append([]byte{flags}, data...)
			case CommandParserBDouble:
				flags, err := reader.ReadUByte()
				if err != nil {
					return nil, err
				}
				data, err := reader.ReadBytesExact(int((flags&0b1)<<3 + (flags&0b10)<<2))
				if err != nil {
					return nil, err
				}
				node.Properties = append([]byte{flags}, data...)
			case CommandParserBInteger:
				flags, err := reader.ReadUByte()
				if err != nil {
					return nil, err
				}
				data, err := reader.ReadBytesExact(int((flags&0b1)<<2 + (flags&0b10)<<1))
				if err != nil {
					return nil, err
				}
				node.Properties = append([]byte{flags}, data...)
			case CommandParserBLong:
				flags, err := reader.ReadUByte()
				if err != nil {
					return nil, err
				}
				data, err := reader.ReadBytesExact(int((flags&0b1)<<3 + (flags&0b10)<<2))
				if err != nil {
					return nil, err
				}
				node.Properties = append([]byte{flags}, data...)
			case CommandParserBString:
				_, node.Properties, err = reader.ReadVarIntRaw()
				if err != nil {
					return nil, err
				}
			case CommandParserMEntity:
				node.Properties, err = reader.ReadBytesExact(1)
				if err != nil {
					return nil, err
				}
			case CommandParserMScoreHolder:
				node.Properties, err = reader.ReadBytesExact(1)
				if err != nil {
					return nil, err
				}
				// TODO: ranges?
			case CommandParserMTime:
				node.Properties, err = reader.ReadBytesExact(4)
				if err != nil {
					return nil, err
				}
			case CommandParserMResourceOrTag, CommandParserMResourceOrTagKey, CommandParserMResource, CommandParserMResourceKey:
				value, err := reader.ReadIdentifier()
				if err != nil {
					return nil, err
				}
				node.Properties = []byte(value)
			}
		}
		if (node.Flags & 0x10) == 0x10 {
			node.SuggestionsType, err = reader.ReadIdentifier()
			if err != nil {
				return nil, err
			}
		}
		nodes[i] = node
	}
	rootIdx, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return Commands{
		Nodes:   nodes,
		RootIdx: rootIdx,
	}, nil
}

func (p Commands) IsValid() (reason error) {
	// TODO: validate
	return nil
}
