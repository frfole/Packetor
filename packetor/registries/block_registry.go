package registries

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Block struct {
	// name of the Block
	name string
}

// BlockRegistry is registry containing list of registry ids and block state ids mapped to Block
type BlockRegistry struct {
	// base maps registry id to Block
	base []Block
	// states maps state id to registry id
	states []int32
}

func (reg *BlockRegistry) load(version string) error {
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
	reg.base = make([]Block, len(data))
	for i := range data {
		if i != int(data[i].Id) {
			return fmt.Errorf("id mismatch for %v", data[i].Name)
		}
		reg.base[i] = Block{name: data[i].Name}
		if maxState < data[i].MaxState {
			maxState = data[i].MaxState
		}
	}

	reg.states = make([]int32, maxState+1)
	for i := range data {
		for j := data[i].MinState; j <= data[i].MaxState; j++ {
			reg.states[j] = int32(i)
		}
	}
	return nil
}

// GetByRegistryID returns Block corresponding to given registry id
func (reg *BlockRegistry) GetByRegistryID(registryId int32) *Block {
	if registryId < 0 || len(reg.base) < int(registryId) {
		return nil
	}
	return &reg.base[registryId]
}

// GetByState returns Block corresponding to given block state id
func (reg *BlockRegistry) GetByState(stateId int32) *Block {
	if stateId < 0 || len(reg.states) < int(stateId) {
		return nil
	}
	return reg.GetByRegistryID(reg.states[stateId])
}

// GetName returns the name of the Block
func (b *Block) GetName() string {
	return b.name
}
