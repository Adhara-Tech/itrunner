package exportedtypes

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadTestEnvExecutionData(config interface{}) error {
	base64ConfigStr, ok := os.LookupEnv(TestRunnerConfEnvVarName)

	if !ok {
		return fmt.Errorf("env variable %s not found or empty content", TestRunnerConfEnvVarName)
	}

	configStr, err := base64.StdEncoding.DecodeString(base64ConfigStr)
	if err != nil {
		return err
	}

	var envExecutionData TestEnvExecutionData
	err = json.Unmarshal([]byte(configStr), &envExecutionData)
	if err != nil {
		return err
	}

	switch envExecutionData.EnvConfigFormat {
	case TestConfigFormatYaml:
		err := yaml.Unmarshal([]byte(envExecutionData.EnvData), config)
		if err != nil {
			return err
		}
	case TestConfigFormatJson:
		err := json.Unmarshal([]byte(envExecutionData.EnvData), config)
		if err != nil {
			return err
		}
	case TestConfigFormatRaw:
		return fmt.Errorf("raw format not implemented")
	default:
		return errors.New("unknown test config format")
	}

	return nil
}
