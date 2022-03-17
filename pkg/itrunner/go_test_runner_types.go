package itrunner

import "encoding/json"

type TestResult int

const (
	TestSuccess TestResult = iota
	TestFailure
)

type Suite struct {
	AllTests []TestGroup
}

type SuiteExecutionResult struct {
	AllTestResults []TestGroupExecutionResult `json:"groups"`
}

type TestGroupExecutionResult struct {
	Name                    string                   `json:"name"`
	VersionExecutionResults []VersionExecutionResult `json:"versions"`
}

type TestDependency struct {
	ID string
}

type VersionExecutionResult struct {
	ID     string     `json:"version"`
	Result TestResult `json:"result"`
}

type TestGroup struct {
	Name     string //Ex: Postgres
	Packages []string
	Versions []Version
}

type Version struct {
	ID        string //Ex:1.9
	Env       []string
	DependsOn []TestDependency
	//VersionDependencyList []ContainerReference
	TestConfig VersionTestConfig
}

type VersionTestConfig struct {
	TemplatePath    string              `yaml:"templatePath"`
	InputDataFrom   ConfigInputDataFrom `yaml:"inputDataFrom"`
	OutputPath      string              `yaml:"outputPath"`
	GoTestExtraArgs []string            `yaml:"extraArgs"`
}

type ConfigInputDataFrom struct {
	Dependencies []ConfigInputDataFromDependency `yaml:"dependencies"`
}

type ConfigInputDataFromDependency struct {
	ID          string `yaml:"id"`
	TemplateVar string `yaml:"templateVar"`
}

func (testResult TestResult) MarshalJSON() ([]byte, error) {
	switch testResult {
	case TestSuccess:
		return json.Marshal("SUCCESS")
	case TestFailure:
		return json.Marshal("FAILURE")
	default:
		return json.Marshal("UNKNOWN")
	}
}
