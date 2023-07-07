package registries

import "fmt"

const (
	RegistryBlock RegistryType = iota
	RegistryWindow
)

type (
	// RegistryType identifies type of registry
	RegistryType uint
	// dataRegistry provides basic methods for registry
	dataRegistry interface {
		// load loads registry with data for given version
		load(version string) error
	}
	// Registry provides basic registries
	Registry struct {
		registries map[RegistryType]dataRegistry
	}
)

var registryInstance = Registry{
	registries: map[RegistryType]dataRegistry{},
}

// GetRegistry returns instance of Registry
func GetRegistry() *Registry {
	return &registryInstance
}

// Load loads all registries with data vor given version
func (reg *Registry) Load(version string) error {
	reg.registries[RegistryBlock] = &BlockRegistry{}
	reg.registries[RegistryWindow] = &WindowRegistry{}
	for registryType, registry := range reg.registries {
		err := registry.load(version)
		if err != nil {
			return fmt.Errorf("failed to load registry %v: %w", registryType, err)
		}
	}
	return nil
}

// Blocks returns shared instance of BlockRegistry
func (reg *Registry) Blocks() *BlockRegistry {
	return reg.registries[RegistryBlock].(*BlockRegistry)
}

// Windows returns shared instance of WindowRegistry
func (reg *Registry) Windows() *WindowRegistry {
	return reg.registries[RegistryWindow].(*WindowRegistry)
}
