package dependencymanager

type DependencyCollection struct {
	Dependencies []Dependency `yaml:"dependencies"`
}

type Dependency struct {
	ID   string      `yaml:"id"`
	Type string      `yaml:"type"`
	Data interface{} `yaml:"-"`
}

type ServiceDependency struct {
	Host string
	Port int
}

/*

dependencies: # moved to 1st level (2 yamls)
  - id: postgres_10.9
    type: service
    data:
      host: my.cool.domain
      port: 8080
  - id: postgres_10.9
    type: container
    data:
      repository:
      tag:
      env:
        - "foo=bar"
      name:
      entrypoint:
        - ""
      cmd:
        - ""
      mounts:
        - "/config"
      links:
        - ""
      exposedPorts:
        - ""
      extraHosts:
        - ""
      labels:
        key: value
      inDocker: #TODO move it to a cmd flag
*/
