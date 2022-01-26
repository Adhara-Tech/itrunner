package containertesthelper

import "github.com/ory/dockertest/docker"

type ContainerInfo struct {
	ContainerRunConfig ContainerRunConfig `mapstructure:"container"`
}

// ContainerRunConfig
type ContainerRunConfig struct {
	Repository   string                               `mapstructure:"repository"`
	Tag          string                               `mapstructure:"tag"`
	Env          []string                             `mapstructure:"env"`
	Name         string                               `mapstructure:"name"`
	Entrypoint   []string                             `mapstructure:"entrypoint"`
	Cmd          []string                             `mapstructure:"cmd"`
	Mounts       []string                             `mapstructure:"mounts"`
	Links        []string                             `mapstructure:"links"`
	ExposedPorts []string                             `mapstructure:"exposed_ports"`
	ExtraHosts   []string                             `mapstructure:"extra_hosts"`
	Labels       map[string]string                    `mapstructure:"labels"`
	PortBindings map[docker.Port][]docker.PortBinding `mapstructure:"port_bindings"`
	InDocker     bool                                 `mapstructure:"in_docker"`
}
