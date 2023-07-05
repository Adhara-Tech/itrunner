package integrationtestrunner

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Adhara-Tech/itrunner/pkg/uc/dependencymanager"

	"github.com/Adhara-Tech/itrunner/pkg/itrunner"

	"gopkg.in/yaml.v3"

	"github.com/Adhara-Tech/itrunner/pkg/uc/resultrender"
)

type OutputFormat string

const (
	OutputFormatJson  OutputFormat = "json"
	OutputFormatTable OutputFormat = "table"
)

type RunnerOptions struct {
	CompatibilityMatrixConfigFilePath       string
	CompatibilityMatrixDependenciesFilePath string
	OutputFile                              string
	OutputFormat                            OutputFormat
	InDocker                                bool
}

func Run(opts RunnerOptions) (*itrunner.SuiteExecutionResult, error) {

	configDataBytes, err := os.ReadFile(opts.CompatibilityMatrixConfigFilePath)
	if err != nil {
		return nil, err
	}
	replacedConfigDataBytes := os.ExpandEnv(string(configDataBytes))

	var config CompatibilityMatrixTestConfig
	err = yaml.Unmarshal([]byte(replacedConfigDataBytes), &config)
	if err != nil {
		return nil, err
	}

	dataBytes, err := json.MarshalIndent(config, "", "   ")
	if err != nil {
		return nil, err
	}

	fmt.Println(string(dataBytes))

	dependencyManager, err := dependencymanager.NewDefaultDependencyManager(dependencymanager.DependencyManagerOptions{
		DependenciesFilePath: opts.CompatibilityMatrixDependenciesFilePath,
		InDocker:             opts.InDocker,
	})
	if err != nil {
		return nil, err
	}
	testRunner, err := itrunner.NewDefaultIntegrationTestsRunner(dependencyManager)
	if err != nil {
		return nil, err
	}
	testSuite := itrunner.Suite{}
	testSuite.AllTests = make([]itrunner.TestGroup, 0)

	for _, testGroup := range config.Suite.TestGroupList {
		currentTestGroup := itrunner.TestGroup{
			Name:     testGroup.Name,
			Packages: testGroup.PackageList,
			Versions: make([]itrunner.Version, 0),
		}

		for _, currentVersion := range testGroup.VersionList {
			version := itrunner.Version{
				ID:  currentVersion.Name,
				Env: currentVersion.EnvVarList,
			}

			for _, dependency := range currentVersion.DependsOn {
				version.DependsOn = append(version.DependsOn, itrunner.TestDependency{
					ID: dependency.ID,
				})

			}
			version.TestConfig = itrunner.VersionTestConfig{
				TemplatePath:    currentVersion.TestConfig.TemplatePath,
				InputDataFrom:   itrunner.ConfigInputDataFrom{Dependencies: make([]itrunner.ConfigInputDataFromDependency, 0)},
				OutputPath:      currentVersion.TestConfig.OutputPath,
				GoTestExtraArgs: currentVersion.TestConfig.GoTestExtraArgs,
			}

			for _, currentDependency := range currentVersion.TestConfig.InputDataFrom.ContainerTestConfigList {
				version.TestConfig.InputDataFrom.Dependencies = append(version.TestConfig.InputDataFrom.Dependencies, itrunner.ConfigInputDataFromDependency{
					ID:          currentDependency.ContainerID,
					TemplateVar: currentDependency.TemplateVar,
				})
			}
			currentTestGroup.Versions = append(currentTestGroup.Versions, version)
		}

		testSuite.AllTests = append(testSuite.AllTests, currentTestGroup)

	}

	result, err := testRunner.RunSuite(testSuite)
	if err != nil {
		return result, err
	}

	render := newRender(opts.OutputFormat)

	// TODO extract to factory function?
	var resultsWriter io.Writer
	if opts.OutputFile != "" {
		directory := filepath.Dir(opts.OutputFile)

		if err = os.MkdirAll(directory, 0777); err != nil {
			return result, err
		}

		resultsWriter, err = os.Create(opts.OutputFile)
		if err != nil {
			return result, err
		}
	} else {
		resultsWriter = os.Stdout
	}

	err = render.Render(*result, resultsWriter)
	if err != nil {
		return result, err
	}

	return result, nil
}

func newRender(format OutputFormat) resultrender.Render {
	switch format {
	case OutputFormatJson:
		return resultrender.JsonRender{}
	case OutputFormatTable:
		return resultrender.CommandLineRender{}
	default:
		return resultrender.CommandLineRender{}
	}
}
