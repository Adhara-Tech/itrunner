package foostorage_test

import (
	"testing"

	"github.com/Adhara-Tech/itrunner/cmd/exportedtypes"
	"github.com/Adhara-Tech/project1-example/pkg/storage/foostorage"
	"github.com/stretchr/testify/assert"
)

type ExampleTestConfig struct {
	ApplicationHost string `mapstructure:"applicationHost" yaml:"applicationHost"`
	ApplicationPort string `mapstructure:"applicationPort" yaml:"applicationPort"`
}

func TestDoFoo(t *testing.T) {
	foostorage.DoFoo()
	var config ExampleTestConfig
	err := exportedtypes.ReadTestEnvExecutionData(&config)

	assert.NoError(t, err)
	assert.NotZero(t, config.ApplicationHost)
	assert.NotZero(t, config.ApplicationPort)
}
