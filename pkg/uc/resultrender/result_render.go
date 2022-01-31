package resultrender

import (
	"io"

	"github.com/AdharaProjects/compatibility-matrix-test-executor/pkg/uc/gotestrunner"
)

type Render interface {
	Render(gotestrunner.SuiteExecutionResult, io.Writer) error
}
