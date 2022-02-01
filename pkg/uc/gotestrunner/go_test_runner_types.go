package gotestrunner

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
	ID  string //Ex:1.9
	Env []string
	//VersionDependencyList []ContainerReference
	//TestConfig            TestConfig
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
