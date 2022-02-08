package resultrender

import (
	"io"

	"github.com/Adhara-Tech/itrunner/pkg/itrunner"
)

type Render interface {
	Render(itrunner.SuiteExecutionResult, io.Writer) error
}
