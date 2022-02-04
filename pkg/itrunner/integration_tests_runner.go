package itrunner

import (
	"fmt"
)

type IntegrationTestsRunner interface {
}

type TestSuite struct {
}

type TestSuiteExecutionResult struct {
}

type DefaultIntegrationTestsRunner struct {
	//goTestRunner GoTestRunner
}

//func (runner *DefaultIntegrationTestsRunner) RunSuite() error {
//	// loop test groups
//	// -- current test group
//	// --- start dependencies (if not already started)
//	// --- generate config
//	// --- execute
//	// end test groups
//	return nil
//}


func (d DefaultIntegrationTestsRunner)  RunSuite(testSet Suite) (*SuiteExecutionResult, error) {

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

func (d DefaultIntegrationTestsRunner) doExecuteTestGroup(group TestGroup) (*TestGroupExecutionResult, error) {

	results := make([]VersionExecutionResult, 0)

	args := make([]string, 0)
	// TODO gotestsum must be an option
	args = append(args, "test")
	args = append(args, group.Packages...)
	for _, version := range group.Versions {

		fmt.Println("starting version " + version.ID)
		// request infra:
		//TODO call dependency manager to retrieve infra info and spinup if possible
		//var containers []*containertesthelper.Container = make([]*containertesthelper.Container, 0)
		//for _, dependency := range version.DependsOn {
		//	fmt.Println("starting container " + dependency.ID)
		//	container, err := d.infraProvider.SpinUpContainer(dependency.ID)
		//	if err != nil {
		//		return nil, err
		//	}
		//	containers = append(containers, container)
		//}

		//TODO call config templater using dependnecies info

		//TODO prepare it test runner config data that will be added as part of env variables when running the tests

		// TODO call gotestrunner
		//exitCode, err := Command("go", version.Env, args...).ExecuteWithLog()
		//testResult := TestSuccess
		//if err != nil {
		//	return nil, err
		//} else {
		//	if exitCode != 0 {
		//		testResult = TestFailure
		//	}
		//}


		// TODO call dependency manager to release dependencies
		// shutdown infra
		//for _, container := range containers {
		//	if err = container.Purge(); err != nil {
		//		return nil, err
		//	}
		//}

		//results = append(results, VersionExecutionResult{
		//	ID:     version.ID,
		//	Result: testResult,
		//})
	}

	return &TestGroupExecutionResult{
		Name:                    group.Name,
		VersionExecutionResults: results,
	}, nil
}

