package dependencymanager

import "github.com/ory/dockertest/docker"

type Dependency struct {
	ID        string              `yaml:"id"`
	Container *ContainerRunConfig `yaml:"container"`
}

// ContainerRunConfig
type ContainerRunConfig struct {
	Repository   string                               `yaml:"repository"`
	Tag          string                               `yaml:"tag"`
	Env          []string                             `yaml:"env"`
	Name         string                               `yaml:"name"`
	Entrypoint   []string                             `yaml:"entrypoint"`
	Cmd          []string                             `yaml:"cmd"`
	Mounts       []string                             `yaml:"mounts"`
	Links        []string                             `yaml:"links"`
	ExposedPorts []string                             `yaml:"exposed_ports"`
	ExtraHosts   []string                             `yaml:"extra_hosts"`
	Labels       map[string]string                    `yaml:"labels"`
	PortBindings map[docker.Port][]docker.PortBinding `yaml:"port_bindings"`
	Auth         docker.AuthConfiguration             `yaml:"dockerauth"`
}
