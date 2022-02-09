package dependencymanager

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
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

	configDataBytes, err := ioutil.ReadFile(dependenciesFilePath)
	if err != nil {
		return nil, err
	}

	var dependencies DependencyCollection

	err = yaml.Unmarshal(configDataBytes, &dependencies)
	if err != nil {
		return nil, err
	}

	return &DefaultDependencyManager{
		containerProvider: *NewInfraProvider(dependencies),
		dependencies:      make(map[string]innerDependency),
	}, nil
}

type innerDependency struct {
	container      *Container
	dependencyInfo DependencyInfo
}

func (d *DefaultDependencyManager) ShutDownDependencies() {
	for depID, dep := range d.dependencies {
		err := dep.container.Purge()
		if err != nil {
			fmt.Println("Dependency (container) with id " + depID + " count not be stopped")
			fmt.Println(err)
		}
	}

	d.dependencies = make(map[string]innerDependency)
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
		return nil, fmt.Errorf("Failed to request dependency %s: %w", dependencyID, err)
	}

	// TODO simple man implementation: it assumes there's always a port to be exposed, picking the first one
	portNames := container.GetPortNames()
	if len(portNames) == 0 {
		return nil, fmt.Errorf("Container of dependency %s does not expose any port", dependencyID)
	}
	firstPort := portNames[0]

	port, err := container.GetPortAsInt(firstPort)
	if err != nil {
		return nil, err
	}

	dependencyInfo := DependencyInfo{
		Host: container.GetIP(firstPort),
		Port: port,
	}

	fmt.Printf("%+v\n", dependencyInfo)

	d.dependencies[dependencyID] = innerDependency{
		container:      container,
		dependencyInfo: dependencyInfo,
	}

	return &dependencyInfo, nil
}
