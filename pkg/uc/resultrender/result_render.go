package resultrender

import (
	"io"

	"github.com/Adhara-Tech/itrunner/pkg/uc/gotestrunner"
)

type Render interface {
	Render(gotestrunner.SuiteExecutionResult, io.Writer) error
}
