package integrationtestrunner

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/AdharaProjects/compatibility-matrix-test-executor/pkg/uc/gotestrunner"

	"gopkg.in/yaml.v3"

	"github.com/AdharaProjects/compatibility-matrix-test-executor/pkg/uc/resultrender"
)

type OutputFormat string

const (
	OutputFormatJson  OutputFormat = "json"
	OutputFormatTable OutputFormat = "table"
)

type RunnerOptions struct {
	CompatibilityMatrixConfigFilePath string
	OutputFile                        string
	OutputFormat                      OutputFormat
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
				ID:  currentVersion.Name,
				Env: currentVersion.EnvVarList,
			})
		}

		testSet.AllTests = append(testSet.AllTests, currentTestGroup)

	}

	result, err := testRunner.RunTests(testSet)
	if err != nil {
		fmt.Println(err)
	}

	var render resultrender.Render
	switch opts.OutputFormat {
	case OutputFormatJson:
		render = resultrender.JsonRender{}
	case OutputFormatTable:
		render = resultrender.CommandLineRender{}
	default:
		render = resultrender.CommandLineRender{}
	}

	var writer io.Writer
	if opts.OutputFile != "" {
		writer, err = os.Create(opts.OutputFile)
		if err != nil {
			return err
		}
	} else {
		writer = os.Stdout
	}
	err = render.Render(*result, writer)
	if err != nil {
		return err
	}

	return nil
}
