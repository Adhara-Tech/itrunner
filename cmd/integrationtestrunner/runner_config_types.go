package integrationtestrunner

type CompatibilityMatrixTestConfig struct {
	Suite Suite `yaml:"suite"`
}

type Suite struct {
	TestGroupList []TestGroup `yaml:"testGroups"`
}

type TestGroup struct {
	Name        string        `yaml:"groupName"`
	PackageList []string      `yaml:"packages"`
	VersionList []TestVersion `yaml:"versions"`
}

type TestVersion struct {
	Name       string     `yaml:"versionName"`
	EnvVarList []string   `yaml:"env"`
	TestConfig TestConfig `yaml:"testConfig"`
}

type TestConfig struct {
	TemplatePath  string                  `yaml:"templatePath"`
	InputDataFrom TestConfigInputDataFrom `yaml:"inputDataFrom"`
	OutputPath    string                  `yaml:"outputPath"`
}

type TestConfigInputDataFrom struct {
	ContainerTestConfigList []ContainerTestConfig `yaml:"containers"`
}

type ContainerTestConfig struct {
	ContainerID string `yaml:"containerId"`
	TemplateVar string `yaml:"templateVar"`
}

/*
suite:
  - groupName: postgres
    packages:
      - ./test/integration/db/...
    versions:
      - versionName: 10.9
        env:
          - CONFIG=./tmp-configs/databases/postgres_config.yml
        dependsOn:
          - containers:
              - id: postgres_10.9
              #- id: prometheus
        testConfig:
          templatePath: ./config-templates/databases/postgres_config.tpl.yml
          input:
            containers:
              - containerId: postgres_10.9
                templateVar: db
            files:
              - path: some_file.yaml

            as: fileConfig
          output: ./tmp-configs/databases/postgres_config.yml
      - version: 10.11
        env:
          config: path_to_config
*/
