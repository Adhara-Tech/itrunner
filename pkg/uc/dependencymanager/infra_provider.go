package dependencymanager

import (
	"github.com/pkg/errors"
)

type containerProvider struct {
	Dependencies   DependencyCollection
	containersPool ContainersPool
}

func NewInfraProvider(deps DependencyCollection) *containerProvider {
	return &containerProvider{
		Dependencies:   deps,
		containersPool: *InitContainerPool(),
	}

}

func (d containerProvider) SpinUpContainer(id string, inDocker bool) (*Container, error) {
	containerSpec := d.findContainerSpec(id)
	if containerSpec == nil {
		return nil, errors.Errorf("Container with ID %s not found in config", id)
	}
	config := ContainerRunConfig{
		Repository:   containerSpec.Container.Repository,
		Tag:          containerSpec.Container.Tag,
		Env:          containerSpec.Container.Env,
		Name:         containerSpec.ID,
		PortBindings: containerSpec.Container.PortBindings,
		ExposedPorts: containerSpec.Container.ExposedPorts,
		Mounts:       containerSpec.Container.Mounts,
		Entrypoint:   containerSpec.Container.Entrypoint,
		Cmd:          containerSpec.Container.Cmd,
		ExtraHosts:   containerSpec.Container.ExtraHosts,
		Labels:       containerSpec.Container.Labels,
	}
	container, err := d.containersPool.Run(config, inDocker, func(container *Container) error {
		return nil
	})
	return container, err
}

func (d containerProvider) findContainerSpec(id string) *Dependency {
	for _, dependency := range d.Dependencies.Dependencies {
		if dependency.Container != nil && dependency.ID == id {
			return &dependency
		}
	}
	return nil
}
