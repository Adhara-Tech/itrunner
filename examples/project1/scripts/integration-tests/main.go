package main

import (
	"fmt"

	"github.com/Adhara-Tech/itrunner/cmd/integrationtestrunner"
)

func main() {
	opt := integrationtestrunner.RunnerOptions{
		CompatibilityMatrixConfigFilePath:       "./compatibility-matrix-tests.yaml",
		CompatibilityMatrixDependenciesFilePath: "./dependencies.yaml",
		OutputFormat:                            integrationtestrunner.OutputFormatJson,
		OutputFile:                              "results.json",
	}
	_, err := integrationtestrunner.Run(opt)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DONE")
}
