package resultrender

import (
	"fmt"
	"os"

	"github.com/AdharaProjects/compatibility-matrix-test-executor/pkg/uc/gotestrunner"

	"github.com/olekukonko/tablewriter"
)

type CommandLineRender struct {
}

func (r CommandLineRender) Render(result []gotestrunner.SuiteExecutionResult) {
	data := make([][]string, len(result))
	//{
	//	[]string{"1/1/2014", "Domain name", "1234", "$10.98"},
	//	[]string{"1/1/2014", "January Hosting", "2345", "$54.95"},
	//	[]string{"1/4/2014", "February Hosting", "3456", "$51.00"},
	//	[]string{"1/4/2014", "February Extra Bandwidth", "4567", "$30.00"},
	//}

	for index, currentResult := range result {
		for _,currentTestExecutionResult := range currentResult.AllTestResults {
			for _, currentVersionExecutionResult := range currentTestExecutionResult.VersionExecutionResults{
				data[index] = []string{currentTestExecutionResult.Name, currentVersionExecutionResult.ID, fmt.Sprint(currentVersionExecutionResult.Result)}
			}
		}

	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Group", "Version", "Result"})
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	//table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}
