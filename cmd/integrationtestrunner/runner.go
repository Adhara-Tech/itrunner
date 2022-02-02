package integrationtestrunner

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/Adhara-Tech/itrunner/pkg/uc/gotestrunner"

	"gopkg.in/yaml.v3"

	"github.com/Adhara-Tech/itrunner/pkg/uc/resultrender"
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
