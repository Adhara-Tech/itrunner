package resultrender

import (
	"io"

	"github.com/Adhara-Tech/itrunner/pkg/itrunner"

	"github.com/olekukonko/tablewriter"
)

type CommandLineRender struct {
}

func (r CommandLineRender) Render(result itrunner.SuiteExecutionResult, writer io.Writer) error {
	data := make([][]string, 0)

	for _, currentTestExecutionResult := range result.AllTestResults {
		for _, currentVersionExecutionResult := range currentTestExecutionResult.VersionExecutionResults {
			resultStr := "Failure"
			if currentVersionExecutionResult.Result == itrunner.TestSuccess {
				resultStr = "Success"
			}
			data = append(data, []string{currentTestExecutionResult.Name, currentVersionExecutionResult.ID, resultStr})
		}
	}

	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"Group", "Version", "Result"})
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	//table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()

	return nil
}
