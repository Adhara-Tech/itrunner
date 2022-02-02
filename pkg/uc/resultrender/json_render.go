package resultrender

import (
	"encoding/json"
	"io"

	"github.com/Adhara-Tech/itrunner/pkg/uc/gotestrunner"
)

type JsonRender struct {
}

func (r JsonRender) Render(result gotestrunner.SuiteExecutionResult, writer io.Writer) error {
	content, err := json.Marshal(map[string]gotestrunner.SuiteExecutionResult{
		"suite": result,
	})
	if err != nil {
		return err
	}

	_, err = writer.Write(content)
	if err != nil {
		return err
	}
	return nil
}
