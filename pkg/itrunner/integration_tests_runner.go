package itrunner

import (
	"github.com/Adhara-Tech/itrunner/pkg/uc/configmaker"
	"github.com/Adhara-Tech/itrunner/pkg/uc/dependencymanager"
	"github.com/Adhara-Tech/itrunner/pkg/uc/gotestrunner"
)

type IntegrationTestsRunner interface {
}

type DefaultIntegrationTestsRunner struct {
	dependencyManager dependencymanager.DependencyManager
	testRunner        gotestrunner.TestRunner
}

func NewDefaultIntegrationTestsRunner(dependencyManager dependencymanager.DependencyManager) (*DefaultIntegrationTestsRunner, error) {
	testRunner := gotestrunner.NewDefaultTestRunner()
	return &DefaultIntegrationTestsRunner{
		dependencyManager: dependencyManager,
		testRunner:        testRunner,
	}, nil
}

func (d DefaultIntegrationTestsRunner) RunSuite(testSet Suite) (*SuiteExecutionResult, error) {

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

	for _, version := range group.Versions {

		// Obtain info about dependencies. Dependency manager starts services when possible or retrieve the info
		// to access the dependency
		configOptions := configmaker.GenerateConfigOptions{
			OutputPath:   version.TestConfig.OutputPath,
			TemplatePath: version.TestConfig.TemplatePath,
			TemplateData: make(map[string]interface{}),
		}

		for _, dependency := range version.TestConfig.InputDataFrom.Dependencies {
			depInfo, err := d.dependencyManager.GetDependencyInfo(dependency.ID)
			if err != nil {
				return nil, err
			}
			configOptions.TemplateData[dependency.TemplateVar] = depInfo
		}

		// Request all dependencies
		for _, dependency := range version.DependsOn {
			_, err := d.dependencyManager.GetDependencyInfo(dependency.ID)
			if err != nil {
				return nil, err
			}
		}

		//Generate configuration
		generatedConfigOutput, err := configmaker.GenerateConfig(configOptions)
		if err != nil {
			return nil, err
		}

		// Execute tests
		testExecutionResult, err := d.testRunner.RunTest(gotestrunner.GoTest{
			Packages:          group.Packages,
			EnvConfigFormat:   "YAML", //TODO need to be added to config... or we can try to deduce it
			EnvConfigFilePath: generatedConfigOutput.OutputFilePath,
			ExtraArgs:         version.TestConfig.GoTestExtraArgs,
		})

		if err != nil {
			return nil, err
		}

		testResult := TestSuccess
		if testExecutionResult.Result != gotestrunner.TestSuccess {
			testResult = TestFailure
		}

		results = append(results, VersionExecutionResult{
			ID:     version.ID,
			Result: testResult,
		})
	}

	// Dependencies are shutdown after all tests are executed
	d.dependencyManager.ShutDownDependencies()

	return &TestGroupExecutionResult{
		Name:                    group.Name,
		VersionExecutionResults: results,
	}, nil
}
