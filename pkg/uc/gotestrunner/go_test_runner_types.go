package gotestrunner

import (
	"encoding/json"

	"github.com/Adhara-Tech/itrunner/cmd/exportedtypes"
)

type TestResult int

const (
	TestSuccess TestResult = iota
	TestFailure
)

type GoTest struct {
	Packages          []string
	EnvConfigFormat   exportedtypes.TestConfigFormat
	EnvConfigFilePath string
	EnvVarName        string
	ExtraArgs         []string
}

type GoTestResult struct {
	Result TestResult
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
