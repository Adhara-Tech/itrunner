package main

import "github.com/Adhara-Tech/itrunner/cmd/integrationtestrunner"

func main() {
	integrationtestrunner.Run(integrationtestrunner.RunnerOptions{
		CompatibilityMatrixConfigFilePath: "",
	})
}
