package dependencymanager

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"github.com/ory/dockertest"
)

//var TestConfigPath string

type ContainersPool struct {
	pool *dockertest.Pool
}

type Container struct {
	myPool           *ContainersPool
	resource         *dockertest.Resource
	isDockerInDocker bool
}

func (container *Container) GetPortNames() []string {
	result := make([]string, 0)
	for portName := range container.resource.Container.NetworkSettings.Ports {
		result = append(result, string(portName))
	}
	return result
}

func (container *Container) GetPort(portName string) string {
	return container.resource.GetPort(portName)
}

func (container *Container) GetPortAsInt(portName string) (int, error) {
	portStr := container.resource.GetPort(portName)
	return strconv.Atoi(portStr)
}

func (container *Container) getBoundIP(portName string) string {
	return container.resource.GetBoundIP(portName)
}

func (container *Container) getGatewayIP(portName string) string {
	return container.resource.GetBoundIP(portName)
}

func (container *Container) GetIP(portName string) string {
	if container.isDockerInDocker {
		return container.getGatewayIP(portName)
	}
	return container.getBoundIP(portName)
}

func InitContainerPool() *ContainersPool {
	return InitContainerPoolWithEndpoint("")
}

func InitContainerPoolWithEndpoint(endpoint string) *ContainersPool {
	pool, err := dockertest.NewPool(endpoint)
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	c := &ContainersPool{}
	c.pool = pool

	return c
}

func (container *Container) Purge() error {
	return container.myPool.pool.Purge(container.resource)
}

func (config *ContainerRunConfig) buildAbsoluteMountPath() {
	for i, relativePath := range config.Mounts {
		absPath, err := filepath.Abs(relativePath)
		// TODO: better way to handle this error
		if err != nil {
			panic(err)
		}
		config.Mounts[i] = absPath
	}
}

func (containersPool *ContainersPool) Run(config ContainerRunConfig, dockerInDocker bool, waitFunction func(container *Container) error) (*Container, error) {
	config.buildAbsoluteMountPath()

	runOptions := &dockertest.RunOptions{
		Name:         config.Name,
		Repository:   config.Repository,
		Tag:          config.Tag,
		Env:          config.Env,
		Entrypoint:   config.Entrypoint,
		Cmd:          config.Cmd,
		Mounts:       config.Mounts,
		Links:        config.Links,
		ExposedPorts: config.ExposedPorts,
		ExtraHosts:   config.ExtraHosts,
		Labels:       config.Labels,
		PortBindings: config.PortBindings,
	}

	resource, err := containersPool.pool.RunWithOptions(runOptions)
	if err != nil {
		return nil, fmt.Errorf("Running container with options: %w", err)
	}

	container := &Container{
		myPool:           containersPool,
		resource:         resource,
		isDockerInDocker: dockerInDocker,
	}

	err = containersPool.pool.Retry(func() error {
		return waitFunction(container)
	})

	if err != nil {
		return nil, fmt.Errorf("Waiting container to be ready: %w", err)
	}

	return container, nil
}
