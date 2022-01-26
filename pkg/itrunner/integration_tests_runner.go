package itrunner

type IntegrationTestsRunner interface {
}

type TestSuite struct {
}

type TestSuiteExecutionResult struct {
}

type DefaultIntegrationTestsRunner struct {
}

func (runner *DefaultIntegrationTestsRunner) RunSuite() error {
	// loop test groups
	// -- current test group
	// --- start dependencies
	// --- generate config
	// --- execute
	// end test groups
	return nil
}
