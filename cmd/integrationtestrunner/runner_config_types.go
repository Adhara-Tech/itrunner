package integrationtestrunner

type CompatibilityMatrixTestConfig struct {
	Suite Suite `yaml:"suite"`
}

type Suite struct {
	TestGroupList []TestGroup      `yaml:"testGroups"`
	Dependencies  DependenciesList `yaml:"dependencies"`
}

type TestGroup struct {
	Name                 string        `yaml:"groupName"`
	PackageList          []string      `yaml:"packages"`
	VersionList          []TestVersion `yaml:"versions"`
	CoverProfileFilePath string        `yaml:"coverProfileFilePath"`
	CoverPackages        []string      `yaml:"coverPackages"`
}

type TestVersion struct {
	Name       string           `yaml:"versionName"`
	EnvVarList []string         `yaml:"env"`
	TestConfig TestConfig       `yaml:"testConfig"`
	DependsOn  []TestDependency `yaml:"dependsOn"`
}

type TestDependency struct {
	ID string
}

type TestConfig struct {
	TemplatePath  string                  `yaml:"templatePath"`
	InputDataFrom TestConfigInputDataFrom `yaml:"inputDataFrom"`
	OutputPath    string                  `yaml:"outputPath"`
}

type TestConfigInputDataFrom struct {
	ContainerTestConfigList []ContainerTestConfig `yaml:"dependencies"`
}

type ContainerTestConfig struct {
	ContainerID string `yaml:"id"`
	TemplateVar string `yaml:"templateVar"`
}

type DependenciesList struct {
	Containers []ContainerSpec `yaml:"containers"`
}

type ContainerSpec struct {
	ID         string   `yaml:"id"`
	Repository string   `yaml:"repository"`
	Tag        string   `yaml:"tag"`
	Env        []string `yaml:"env"`
	Ports      []string `yaml:"ports"`
}
