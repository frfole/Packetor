package data

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Block struct {
	name string
}

type BlockRegistry struct {
	// protos maps protocol id to block
	protos []Block
	// states maps state id to protocol id
	states []int32
}

var blockRegistryInst = BlockRegistry{}

func GetBlockRegistry() *BlockRegistry {
	return &blockRegistryInst
}

func (r *BlockRegistry) Load(version string) error {
	resp, err := http.Get(fmt.Sprintf("https://github.com/PrismarineJS/minecraft-data/raw/master/data/pc/%v/blocks.json", version))
	if err != nil {
		return fmt.Errorf("failed to get blocks.json: %w", err)
	}

	var data []struct {
		Id       int32  `json:"id"`
		Name     string `json:"name"`
		MinState int32  `json:"minStateId"`
		MaxState int32  `json:"maxStateId"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return fmt.Errorf("failed to decode blocks.json: %w", err)
	}

	maxState := int32(-1)
	r.protos = make([]Block, len(data))
	for i := range data {
		if i != int(data[i].Id) {
			return fmt.Errorf("id mismatch for %v", data[i].Name)
		}
		r.protos[i] = Block{name: data[i].Name}
		if maxState < data[i].MaxState {
			maxState = data[i].MaxState
		}
	}

	r.states = make([]int32, maxState+1)
	for i := range data {
		for j := data[i].MinState; j <= data[i].MaxState; j++ {
			r.states[j] = int32(i)
		}
	}
	return nil
}

func (r *BlockRegistry) GetByProtoID(protoId int32) *Block {
	return &r.protos[protoId]
}

func (r *BlockRegistry) GetByState(stateId int32) *Block {
	return r.GetByProtoID(r.states[stateId])
}
