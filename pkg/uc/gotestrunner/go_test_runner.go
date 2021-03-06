package gotestrunner

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Adhara-Tech/itrunner/cmd/exportedtypes"
	"github.com/Adhara-Tech/itrunner/pkg/uc/dependencymanager"
	"io/ioutil"
)

type TestRunner interface {
	RunTest(test GoTest) (*GoTestResult, error)
}

var _ TestRunner = (*DefaultTestRunner)(nil)

type DefaultTestRunner struct {
}

func (d DefaultTestRunner) RunTest(test GoTest) (*GoTestResult, error) {

	fileData, err := ioutil.ReadFile(test.EnvConfigFilePath)
	if err != nil {
		return nil, err
	}

	testEnvData := exportedtypes.TestEnvExecutionData{
		EnvConfigFormat: test.EnvConfigFormat,
		EnvData:         string(fileData),
	}

	envData, err := json.Marshal(testEnvData)
	if err != nil {
		return nil, err
	}

	base64EnvData := base64.StdEncoding.EncodeToString(envData)
	envVar := fmt.Sprintf("%s=%s", exportedtypes.TestRunnerConfEnvVarName, base64EnvData)

	fmt.Println(envVar)

	args := make([]string, 0)
	// TODO gotestsum must be an option
	args = append(args, "test")
	args = append(args, test.ExtraArgs...)
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
