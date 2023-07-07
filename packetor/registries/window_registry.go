package registries

type (
	Window struct {
		// name of the Window
		name string
		// size of the Window excluding player slots
		size int32
	}
	// WindowRegistry is registry containing list of registry ids mapped to Window
	WindowRegistry struct {
		// base maps registry id to Window
		base []Window
	}
)

func (reg *WindowRegistry) load(_ string) error {
	reg.base = []Window{
		{name: "generic_9x1", size: 9},
		{name: "generic_9x2", size: 18},
		{name: "generic_9x3", size: 27},
		{name: "generic_9x4", size: 36},
		{name: "generic_9x5", size: 45},
		{name: "generic_9x6", size: 54},
		{name: "generic_3x3", size: 9},
		{name: "anvil", size: 3},
		{name: "beacon", size: 1},
		{name: "blast_furnace", size: 3},
		{name: "brewing_stand", size: 5},
		{name: "crafting", size: 10},
		{name: "enchantment", size: 2},
		{name: "furnace", size: 3},
		{name: "grindstone", size: 3},
		{name: "hopper", size: 5},
		{name: "lectern", size: 1},
		{name: "loom", size: 4},
		{name: "merchant", size: 3},
		{name: "shulker_box", size: 27},
		{name: "smithing", size: 4},
		{name: "smoker", size: 3},
		{name: "cartography_table", size: 3},
		{name: "stonecutter", size: 2},
	}
	return nil
}

// GetByRegistryID returns Window corresponding to given registry id
func (reg *WindowRegistry) GetByRegistryID(registryId int32) *Window {
	if registryId < 0 || len(reg.base) < int(registryId) {
		return nil
	}
	return &reg.base[registryId]
}
