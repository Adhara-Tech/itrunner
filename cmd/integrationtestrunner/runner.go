package integrationtestrunner

import (
	"encoding/json"
	"fmt"
	"github.com/AdharaProjects/compatibility-matrix-test-executor/pkg/uc/gotestrunner"
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"github.com/AdharaProjects/compatibility-matrix-test-executor/pkg/uc/resultrender"
)

type RunnerOptions struct {
	CompatibilityMatrixConfigFilePath string
}

func Run(opts RunnerOptions) error {

	configDataBytes, err := ioutil.ReadFile(opts.CompatibilityMatrixConfigFilePath)
	if err != nil {
		return err
	}

	var config CompatibilityMatrixTestConfig

	err = yaml.Unmarshal(configDataBytes, &config)
	if err != nil {
		return err
	}

	dataBytes, err := json.MarshalIndent(config, "", "   ")
	if err != nil {
		return err
	}

	fmt.Println(string(dataBytes))

	testRunner := gotestrunner.DefaultTestRunner{}
	testSet := gotestrunner.Suite{}
	testSet.AllTests = make([]gotestrunner.TestGroup, 0)




	for _, testGroup := range config.Suite.TestGroupList {
		currentTestGroup := gotestrunner.TestGroup{
			Name:     testGroup.Name,
			Packages: testGroup.PackageList,
			Versions: make([]gotestrunner.Version, 0),
		}

		for _, currentVersion := range testGroup.VersionList {
			currentTestGroup.Versions = append(currentTestGroup.Versions, gotestrunner.Version{
				ID:                    currentVersion.Name,
				Env:                   currentVersion.EnvVarList,
				//VersionDependencyList: currentVersion.,

				//TestConfig:            gotestrunner.TestConfig{
				//	TemplatePath:
				//},
			})
		}

		testSet.AllTests = append(testSet.AllTests, currentTestGroup)

	}

	testSet.AllTests = append(testSet.AllTests, gotestrunner.TestGroup{
		Name:     "Postgresql demo",
		Packages: []string{"./test/integration/db/..."},
		Versions: []gotestrunner.Version{
			{
				ID:  "10.9",
				Env: []string{"CUSTOM_KEY=value"},
			},
			{
				ID:  "10.10",
				Env: []string{"CUSTOM_KEY=value"},
			},
			{
				ID:  "10.11",
				Env: []string{"CUSTOM_KEY=value"},
			},
		},
	})
	testSet.AllTests = append(testSet.AllTests, gotestrunner.TestGroup{
		Name:     "Rabbit demo",
		Packages: []string{"./test/integration/db/..."},
		Versions: []gotestrunner.Version{
			{
				ID:  "1.4",
				Env: []string{"CUSTOM_KEY=value"},
			},
			{
				ID:  "2.4",
				Env: []string{"CUSTOM_KEY=value"},
			},
			{
				ID:  "3.7",
				Env: []string{"CUSTOM_KEY=value"},
			},
		},
	})
	result, err := testRunner.RunTests(testSet)
	if err != nil {
		fmt.Println(err)
	}

	render := resultrender.CommandLineRender{}
	render.Render([]gotestrunner.SuiteExecutionResult{*result})

	return nil
}
