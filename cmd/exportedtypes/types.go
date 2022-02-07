package exportedtypes

type TestConfigFormat string

const (
	TestConfigFormatJson TestConfigFormat = "JSON"
	TestConfigFormatYaml TestConfigFormat = "YAML"
)

type TestEnvExecutionData struct {
	EnvConfigFormat TestConfigFormat `json:"configFormat"`
	EnvData         string           `json:"data"`
}
