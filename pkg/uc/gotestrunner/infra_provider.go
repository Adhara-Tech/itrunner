package gotestrunner

import (
	"github.com/AdharaProjects/compatibility-matrix-test-executor/pkg/containertesthelper"
	"github.com/pkg/errors"
)

var _ InfraProvider = (*DefaultInfraProvider)(nil)

type DefaultInfraProvider struct {
	Dependencies   DependenciesList
	containersPool containertesthelper.ContainersPool
}

func NewInfraProvider(deps DependenciesList) *DefaultInfraProvider {
	return &DefaultInfraProvider{
		Dependencies:   deps,
		containersPool: *containertesthelper.InitContainerPool(),
	}

}

func (d *DefaultInfraProvider) SpinUpContainer(id string) (*containertesthelper.Container, error) {
	containerSpec := d.findContainerSpec(id)
	if containerSpec == nil {
		return nil, errors.Errorf("Container with ID %s not found in config", id)
	}
	config := containertesthelper.ContainerRunConfig{
		Repository: containerSpec.Repository,
		Tag:        containerSpec.Tag,
		Env:        containerSpec.Env,
		Name:       containerSpec.ID,
	}
	container, err := d.containersPool.Run(config, false, func(container *containertesthelper.Container) error {
		return nil
	})
	return container, err
}

func (infraProvider DefaultInfraProvider) findContainerSpec(id string) *ContainerSpec {
	for _, container := range infraProvider.Dependencies.Containers {
		if container.ID == id {
			return &container
		}
	}
	return nil
}
