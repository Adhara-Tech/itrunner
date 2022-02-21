# IT Runner - Integration Tests executor

_IT Runner_ is a tool to automate the execution of integration tests across multiple versions of dependencies. It helps to build compatibiliy matrix by executing the same suite of tests against different setups of components.

> **DISCLAIMER**: _IT Runner_ is still in its early beta stages. Please mind that breaking changes in its APIs and configuration files format can be made until a first 1.X version is stabilised.

- [IT Runner - Integration Tests executor](#it-runner---integration-tests-executor)
  * [Usage](#usage)
    + [Suite configuration](#suite-configuration)
    + [Templating](#templating)
    + [Dependencies configuration](#dependencies-configuration)
    + [Invoking the tool from GO code](#invoking-the-tool-from-go-code)
    + [Using the config library in go tests](#using-the-config-library-in-go-tests)
  * [Roadmap](#roadmap)

## Usage

_ITRunner_ relies on two configuration files to bootstrap the components to be used in the integration tests (dependencies configuration file) and to execute the tests themselves (suite configuration file).

### Suite configuration

```yaml
suite:
  testGroups:
    - groupName: postgres                        # descriptive name of the group of tests
      packages:
        - ./tests/persistence/...                # GO packages containing test files

      versions:

        - versionName: 9.6.16                    # descriptive name of the version of the dependency(ies) being tested
          dependsOn:
            - id: postgres_9_6_16                # list of dependency ids that will be requested before the test execution
          testConfig:
            templatePath: ./path/to/config.tpl   # golang template of custom config used by the tests. It is rendered with data from dependencies
            inputDataFrom:
              dependencies:
                - id: postgres_9_6_16            # Id of the dependency
                  templateVar: Postgres          # Variable to host the dependency data when templating
            outputPath: ./configs/rendered.yml      # output of the result config file after rendering the template

        - versionName: 10.12
          dependsOn:
              - id: postgres_10_12
          testConfig:
            templatePath: ./paht/to/another/config.tpl
            inputDataFrom:
              dependencies:
                - id: postgres_10_12
                  templateVar: Postgres
            outputPath: ./configs/another-render.yml
```

### Templating

When using templates to render configuration for tests, IT Runner will inject the following information for each dependency specified in `testConfig`:
* Host
* Port

### Dependencies configuration

See the [dependencymanager.ContainerRunConfig struct](pkg/uc/dependencymanager/container_test_types.go#L11-L24) for the allowed fields of the `container` config section.

```yaml
dependencies:

  - id: postgres_9_6_16             # Id of the dependency, used in the suite config file
    container:                      # Specs to run a Docker container
      repository: "postgres"
      tag: "9.6.16"
      env:
        - "POSTGRES_PASSWORD="
        - "POSTGRES_DB="

  - id: postgres_10_12
    container:
      repository: "postgres"
      tag: "10.12"
      env:
        - "POSTGRES_PASSWORD="
      exposedPorts:
        - "5432:5432"
```

### Invoking the tool from GO code

Please check this example on how to invoke IT Runner from a GO script:

```go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Adhara-Tech/itrunner/pkg/itrunner"

	"github.com/Adhara-Tech/itrunner/cmd/integrationtestrunner"
)

func main() {
	opt := integrationtestrunner.RunnerOptions{
		CompatibilityMatrixConfigFilePath:       "path/to/itrunner-suite.yaml",
		CompatibilityMatrixDependenciesFilePath: "path/to/itrunner-dependencies.yaml",
		InDocker:                                false,   # set to true when you are executing the test from inside a docker container

	}

	result, err := integrationtestrunner.Run(opt)

	if err != nil {
		os.Exit(1)  // an error happened during the test execution (for instance, docker daemon down or similar non-test related issues)
	}

	for _, testResult := range result.AllTestResults {
		for _, versionResult := range testResult.VersionExecutionResults {
			if versionResult.Result != itrunner.TestSuccess {
				os.Exit(1)    // in this example: return non-zero exit code if one of the test failed
			}
		}
	}
}
```

### Using the config library in go tests

IT Runner provides a `github.com/Adhara-Tech/itrunner/cmd/exportedtypes` package with a `ReadTestEnvExecutionData(data interface{})` function that allows you to load the configuration rendered from `testConfig` settings into the `data` structure.


```
package foostorage_test

import (
	"testing"

	"github.com/Adhara-Tech/itrunner/cmd/exportedtypes"
	"github.com/your/coolproject/foostorage"
	"github.com/stretchr/testify/assert"
)

type ExampleTestConfig struct {
	ApplicationHost string `mapstructure:"applicationHost" yaml:"applicationHost"`
	ApplicationPort string `mapstructure:"applicationPort" yaml:"applicationPort"`
}

func TestDoFoo(t *testing.T) {
	var config ExampleTestConfig
	err := exportedtypes.ReadTestEnvExecutionData(&config)
	assert.NoError(t, err)

	err = foostorage.DoFoo(config)
	assert.NoError(t, err)
}
```


## Roadmap

* Multiple ports exposed as part of a dependency (containers and pre provisioned services)
* Support for pre provisioned services
* Fine-grained control to shut down dependencies (now only after a group). Options can be after a version or after a suite
* Redefine test executable (support for gotestsum)
* Labeling to discriminate tests to be executed
* Redefine env variable used to inject data to tests
* Support to inject several variables
    ```
      testVars:
	- name: configPath
	  value: ./tmp-configs/databases/postgres_config.yml
	- name: configValue
	  type: [raw,yaml,json]
	  fromFile: ./tmp-configs/databases/postgres_config.yml #reads the file
    ```
* Support more config formats other than `YAML` in `exportedtypes.ReadTestEnvExecutionData(data interface{})`

