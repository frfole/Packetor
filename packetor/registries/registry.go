package registries

type (
	dataRegistry interface {
		load(version string) error
	}
	Registry struct {
		blocks *BlockRegistry
	}
)

var registryInstance = Registry{
	blocks: &BlockRegistry{},
}

func GetRegistry() *Registry {
	return &registryInstance
}

func (reg *Registry) Load(version string) error {
	err := reg.blocks.load(version)
	if err != nil {
		return err
	}
	return nil
}

func (reg *Registry) Blocks() *BlockRegistry {
	return reg.blocks
}
