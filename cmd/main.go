package main

import "github.com/AdharaProjects/compatibility-matrix-test-executor/cmd/integrationtestrunner"

func main() {
	integrationtestrunner.Run(integrationtestrunner.RunnerOptions{
		CompatibilityMatrixConfigFilePath: "",
	})
}
