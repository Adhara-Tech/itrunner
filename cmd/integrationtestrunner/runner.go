package integrationtestrunner

import (
	"encoding/json"
	"fmt"
	"github.com/Adhara-Tech/itrunner/pkg/uc/dependencymanager"
	"io"
	"io/ioutil"
	"os"

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

	var dependencieslist itrunner.DependenciesList
	for _, dependency := range config.Suite.Dependencies.Containers {
		container := itrunner.ContainerSpec{
			ID:         dependency.ID,
			Repository: dependency.Repository,
			Tag:        dependency.Tag,
			Env:        dependency.Env,
		}
		dependencieslist.Containers = append(dependencieslist.Containers, container)

	}

	// TODO path to the dependencies file must be provided
	dependencyManager, err := dependencymanager.NewDefaultDependencyManager("TODO")
	if err != nil {
		return err
	}
	testRunner, err := itrunner.NewDefaultIntegrationTestsRunner(dependencyManager)
	if err != nil {
		return err
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
			currentTestGroup.Versions = append(currentTestGroup.Versions, version)
		}

		testSuite.AllTests = append(testSuite.AllTests, currentTestGroup)

	}

	result, err := testRunner.RunSuite(testSuite)
	if err != nil {
		return err
	}

	render := newRender(opts.OutputFormat)

	// TODO extract to factory function?
	var resultsWriter io.Writer
	if opts.OutputFile != "" {
		resultsWriter, err = os.Create(opts.OutputFile)
		if err != nil {
			return err
		}
	} else {
		resultsWriter = os.Stdout
	}

	err = render.Render(*result, resultsWriter)
	if err != nil {
		return err
	}

	return nil
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
