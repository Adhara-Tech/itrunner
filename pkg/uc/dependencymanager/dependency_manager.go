package dependencymanager

import (
	"fmt"
)

type DependencyManager interface {
	GetDependencyInfo(dependencyID string) (*DependencyInfo, error)
	ShutDownDependencies()
}

type DependencyInfo struct {
	Host string
	//TODO we may need multiport in the future
	Port int
}

var _ DependencyManager = (*DefaultDependencyManager)(nil)

type DefaultDependencyManager struct {
	containerProvider containerProvider
	dependencies      map[string]innerDependency
}

func NewDefaultDependencyManager(dependenciesFilePath string) (*DefaultDependencyManager, error) {
	// Reads the config file and store the info in a private field
	// Inits the container provider
	panic("not implemented")
}

type innerDependency struct {
	container      *Container
	dependencyInfo DependencyInfo
}

func (d DefaultDependencyManager) ShutDownDependencies() {
	for depID, dep := range d.dependencies {
		err := dep.container.Purge()
		if err != nil {
			fmt.Println("Dependency (container) with id " + depID + " count not be stopped")
			fmt.Println(err)
		}
	}
}

func (d DefaultDependencyManager) GetDependencyInfo(dependencyID string) (*DependencyInfo, error) {
	// request infra:

	dep, ok := d.dependencies[dependencyID]

	if ok {
		return &dep.dependencyInfo, nil
	}

	fmt.Println("starting container " + dependencyID)
	container, err := d.containerProvider.SpinUpContainer(dependencyID)
	if err != nil {
		return nil, err
	}

	port, err := container.GetPortAsInt("TODO it comes from config")
	if err != nil {
		return nil, err
	}

	dependencyInfo := DependencyInfo{
		Host: container.GetIP("TODO it comes from config"),
		Port: port,
	}

	d.dependencies[dependencyID] = innerDependency{
		container:      container,
		dependencyInfo: dependencyInfo,
	}

	return &dependencyInfo, nil
}
