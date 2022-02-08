package configmaker_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Adhara-Tech/itrunner/pkg/uc/configmaker"
)

var configRenderResult = `containers:
  ip: 192.168.10.5
  port: 8080`

func TestGenerateConfig(t *testing.T) {
	data := make(map[string]interface{})

	data["ExampleContainer"] = struct {
		Ip   string
		Port string
	}{
		Ip:   "192.168.10.5",
		Port: "8080",
	}

	opts := configmaker.GenerateConfigOptions{
		OutputPath:   "example-template",
		TemplatePath: "testdata/example.tpl.yaml",
		TemplateData: data,
	}

	configOutput, err := configmaker.GenerateConfig(opts)

	require.NoError(t, err)

	fileBytes, err := ioutil.ReadFile(configOutput.OutputFilePath)
	require.NoError(t, err)
	require.Equal(t, configRenderResult, string(fileBytes))
}
