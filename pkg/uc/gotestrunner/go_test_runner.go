package gotestrunner

import (
	"encoding/json"
	"fmt"
	"github.com/Adhara-Tech/itrunner/pkg/uc/dependencymanager"
)

type TestRunner interface {
	RunTest(test GoTest) (*GoTestResult, error)
}

var _ TestRunner = (*DefaultTestRunner)(nil)

const testRunnerConfEnvVarName = "ITEST_RUNNER_CONF_DEFAULT"

type DefaultTestRunner struct {
}

func (d DefaultTestRunner) RunTest(test GoTest) (*GoTestResult, error) {

	testEnvData := testExecutionData{
		EnvConfigFormat: "",
		EnvData:         "",
	}

	envData, err := json.Marshal(testEnvData)
	if err != nil {
		return nil, err
	}

	envVar := fmt.Sprintf("%s=%s", testRunnerConfEnvVarName, string(envData))

	args := make([]string, 0)
	// TODO gotestsum must be an option
	args = append(args, "test")
	args = append(args, test.Packages...)
	exitCode, err := Command("go", []string{envVar}, args...).ExecuteWithLog()
	testResult := TestSuccess
	if err != nil {
		return nil, err
	} else {
		if exitCode != 0 {
			testResult = TestFailure
		}
	}

	return &GoTestResult{
		Result: testResult,
	}, nil
}

type InfraProvider interface {
	SpinUpContainer(id string) (*dependencymanager.Container, error)
}

func NewDefaultTestRunner() DefaultTestRunner {
	return DefaultTestRunner{}
}
