package dependencymanager

import (
	"github.com/Adhara-Tech/itrunner/pkg/itrunner"
	"github.com/pkg/errors"
)

type containerProvider struct {
	Dependencies   itrunner.DependenciesList
	containersPool ContainersPool
}

func NewInfraProvider(deps itrunner.DependenciesList) *containerProvider {
	return &containerProvider{
		Dependencies:   deps,
		containersPool: *InitContainerPool(),
	}

}

func (d containerProvider) SpinUpContainer(id string) (*Container, error) {
	containerSpec := d.findContainerSpec(id)
	if containerSpec == nil {
		return nil, errors.Errorf("Container with ID %s not found in config", id)
	}
	config := ContainerRunConfig{
		Repository: containerSpec.Repository,
		Tag:        containerSpec.Tag,
		Env:        containerSpec.Env,
		Name:       containerSpec.ID,
	}
	container, err := d.containersPool.Run(config, false, func(container *Container) error {
		return nil
	})
	return container, err
}

func (d containerProvider) findContainerSpec(id string) *itrunner.ContainerSpec {
	for _, container := range d.Dependencies.Containers {
		if container.ID == id {
			return &container
		}
	}
	return nil
}
