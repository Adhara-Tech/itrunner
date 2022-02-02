package resultrender

import (
	"io"

	"github.com/Adhara-Tech/itrunner/pkg/uc/gotestrunner"

	"github.com/olekukonko/tablewriter"
)

type CommandLineRender struct {
}

func (r CommandLineRender) Render(result gotestrunner.SuiteExecutionResult, writer io.Writer) error {
	data := make([][]string, 0)
	//{
	//	[]string{"1/1/2014", "Domain name", "1234", "$10.98"},
	//	[]string{"1/1/2014", "January Hosting", "2345", "$54.95"},
	//	[]string{"1/4/2014", "February Hosting", "3456", "$51.00"},
	//	[]string{"1/4/2014", "February Extra Bandwidth", "4567", "$30.00"},
	//}

	for _, currentTestExecutionResult := range result.AllTestResults {
		for _, currentVersionExecutionResult := range currentTestExecutionResult.VersionExecutionResults {
			resultStr := "Failure"
			if currentVersionExecutionResult.Result == gotestrunner.TestSuccess {
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
