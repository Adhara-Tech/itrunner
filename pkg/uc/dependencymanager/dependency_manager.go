package dependencymanager

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DependencyManagerOptions struct {
	DependenciesFilePath string
	InDocker             bool
}

type DependencyManager interface {
	GetDependencyInfo(dependencyID string) (*DependencyInfo, error)
	ShutDownDependencies()
}

type DependencyInfo struct {
	Host  string
	Ports map[string]int
}

var _ DependencyManager = (*DefaultDependencyManager)(nil)

type DefaultDependencyManager struct {
	containerProvider containerProvider
	dependencies      map[string]innerDependency
	inDocker          bool
}

func NewDefaultDependencyManager(opts DependencyManagerOptions) (*DefaultDependencyManager, error) {
	// Reads the config file and store the info in a private field
	// Inits the container provider

	configDataBytes, err := os.ReadFile(opts.DependenciesFilePath)
	if err != nil {
		return nil, err
	}
	replacedConfigDataBytes := os.ExpandEnv(string(configDataBytes))

	var dependencies DependencyCollection

	err = yaml.Unmarshal([]byte(replacedConfigDataBytes), &dependencies)
	if err != nil {
		return nil, err
	}

	return &DefaultDependencyManager{
		containerProvider: *NewInfraProvider(dependencies),
		dependencies:      make(map[string]innerDependency),
		inDocker:          opts.InDocker,
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
	container, err := d.containerProvider.SpinUpContainer(dependencyID, d.inDocker)
	if err != nil {
		return nil, fmt.Errorf("Failed to request dependency %s: %w", dependencyID, err)
	}

	// TODO simple man implementation: it assumes there's always a port to be exposed, picking the first one
	portNames := container.GetPortNames()
	if len(portNames) == 0 {
		return nil, fmt.Errorf("Container of dependency %s does not expose any port", dependencyID)
	}
	firstPort := portNames[0]

	ports := make(map[string]int)

	for _, portName := range portNames {
		port, err := container.GetPortAsInt(portName)
		if err != nil {
			return nil, err
		}
		ports[portName] = port
	}

	dependencyInfo := DependencyInfo{
		Host:  container.GetIP(firstPort),
		Ports: ports,
	}

	fmt.Printf("%+v\n", dependencyInfo)

	d.dependencies[dependencyID] = innerDependency{
		container:      container,
		dependencyInfo: dependencyInfo,
	}

	return &dependencyInfo, nil
}
