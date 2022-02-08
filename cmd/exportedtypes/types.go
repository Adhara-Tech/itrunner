package exportedtypes

const TestRunnerConfEnvVarName = "ITEST_RUNNER_CONF_DEFAULT"

type TestConfigFormat string

const (
	TestConfigFormatJson TestConfigFormat = "JSON"
	TestConfigFormatYaml TestConfigFormat = "YAML"
	TestConfigFormatRaw  TestConfigFormat = "RAW"
)

type TestEnvExecutionData struct {
	EnvConfigFormat TestConfigFormat `json:"configFormat"`
	EnvData         string           `json:"data"`
}
