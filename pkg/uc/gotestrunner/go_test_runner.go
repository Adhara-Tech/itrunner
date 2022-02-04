package gotestrunner

import (
	"fmt"

	"github.com/Adhara-Tech/itrunner/pkg/containertesthelper"
)

type TestRunner interface {
	RunTests(set Suite) (*SuiteExecutionResult, error)
}

var _ TestRunner = (*DefaultTestRunner)(nil)

type DefaultTestRunner struct {
	infraProvider InfraProvider
}

type InfraProvider interface {
	SpinUpContainer(id string) (*containertesthelper.Container, error)
}

func NewTestRunner(infraProvider InfraProvider) DefaultTestRunner {
	return DefaultTestRunner{
		infraProvider: infraProvider,
	}
}

func (d DefaultTestRunner) RunTests(testSet Suite) (*SuiteExecutionResult, error) {

	allResults := make([]TestGroupExecutionResult, 0)

	for _, testDefinition := range testSet.AllTests {
		testGroupExecutionResult, err := d.doExecuteTestGroup(testDefinition)
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, *testGroupExecutionResult)
	}

	return &SuiteExecutionResult{AllTestResults: allResults}, nil
}

func (d DefaultTestRunner) doExecuteTestGroup(group TestGroup) (*TestGroupExecutionResult, error) {

	results := make([]VersionExecutionResult, 0)

	args := make([]string, 0)
	// TODO gotestsum must be an option
	args = append(args, "test")
	args = append(args, group.Packages...)
	for _, version := range group.Versions {

		fmt.Println("starting version " + version.ID)
		// request infra:
		var containers []*containertesthelper.Container = make([]*containertesthelper.Container, 0)
		for _, dependency := range version.DependsOn {
			fmt.Println("starting container " + dependency.ID)
			container, err := d.infraProvider.SpinUpContainer(dependency.ID)
			if err != nil {
				return nil, err
			}
			containers = append(containers, container)
		}

		exitCode, err := Command("go", version.Env, args...).ExecuteWithLog()
		testResult := TestSuccess
		if err != nil {
			return nil, err
		} else {
			if exitCode != 0 {
				testResult = TestFailure
			}
		}

		// shutdown infra
		for _, container := range containers {
			if err = container.Purge(); err != nil {
				return nil, err
			}
		}

		results = append(results, VersionExecutionResult{
			ID:     version.ID,
			Result: testResult,
		})
	}

	return &TestGroupExecutionResult{
		Name:                    group.Name,
		VersionExecutionResults: results,
	}, nil
}
