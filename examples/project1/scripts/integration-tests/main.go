package main

import (
	"fmt"

	"github.com/AdharaProjects/compatibility-matrix-test-executor/cmd/integrationtestrunner"
	//"github.com/AdharaProjects/compatibility-matrix-test-executor/cmd/runner"
)

func main() {
	opt := integrationtestrunner.RunnerOptions{
		CompatibilityMatrixConfigFilePath: "./compatibility-matrix-tests.yaml",
		OutputFormat:                      integrationtestrunner.OutputFormatJson,
		OutputFile:                        "results.json",
	}
	err := integrationtestrunner.Run(opt)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DONE")
}
