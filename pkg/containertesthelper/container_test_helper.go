package containertesthelper

import (
	"log"
	"strconv"

	"shared.mod/go-commons/pkg/testhelpers/testconfig"

	"shared.mod/go-commons/pkg/exerrors"

	"github.com/ory/dockertest"
)

var TestConfigPath string

type ContainersPool struct {
	pool *dockertest.Pool
}

type Container struct {
	myPool           *ContainersPool
	resource         *dockertest.Resource
	isDockerInDocker bool
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
		config.Mounts[i] = testconfig.BuildAbsolutePathFromTestConfigPath(relativePath)
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
		return nil, exerrors.WrapUnknownWithMsg(err, "Running container with options")
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
		return nil, exerrors.WrapUnknownWithMsg(err, "Waiting container to be ready")
	}

	return container, nil
}
