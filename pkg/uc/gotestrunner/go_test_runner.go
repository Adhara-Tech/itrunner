package gotestrunner

type TestRunner interface {
	RunTests(set Suite) (*SuiteExecutionResult, error)
}

var _ TestRunner = (*DefaultTestRunner)(nil)

type DefaultTestRunner struct {
}

func (d DefaultTestRunner) RunTests(testSet Suite) (*SuiteExecutionResult, error) {

	allResults := make([]TestGroupExecutionResult, 0)

	for _, testDefinition := range testSet.AllTests {
		testGroupExecutionResult, err := d.doExecuteTestGroup(testDefinition)
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, *testGroupExecutionResult)
	}

	return &SuiteExecutionResult{AllTestResults: allResults}, nil
}

func (d DefaultTestRunner) doExecuteTestGroup(group TestGroup) (*TestGroupExecutionResult, error) {

	results := make([]VersionExecutionResult, 0)

	args := make([]string, 0)
	// TODO gotestsum must be an option
	args = append(args, "test")
	args = append(args, group.Packages...)
	for _, version := range group.Versions {

		exitCode, err := Command("go", version.Env, args...).ExecuteWithLog()
		testResult := TestSuccess
		if err != nil {
			return nil, err
		} else {
			if exitCode != 0 {
				testResult = TestFailure
			}
		}

		results = append(results, VersionExecutionResult{
			ID:     version.ID,
			Result: testResult,
		})
	}

	return &TestGroupExecutionResult{
		Name: group.Name,
		VersionExecutionResults: results,
	}, nil
}
